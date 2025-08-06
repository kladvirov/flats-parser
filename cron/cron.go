package cron

import (
	"bytes"
	"flats-parser/adapter"
	"flats-parser/constants"
	"flats-parser/helper"
	"flats-parser/parser"
	"flats-parser/parser/kufar"
	"flats-parser/parser/onliner"
	"flats-parser/parser/options"
	"flats-parser/parser/realt"
	flats "flats-parser/repositories"
	"flats-parser/telegram"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-co-op/gocron/v2"
)

type AdProcessor[T any, R any] struct {
	ParseURL  string
	AdType    int8
	BuildOpts func() []options.Option
	TgBot     *telegram.Bot
	BuildURL  func([]string) []string
	MakeDesc  func(T) string
	ExtractID func(T) int
	GetAds    func(R) []T
	GetImages func(T) []string
}

type Job interface {
	Execute() error
}

type RealtSendJob struct {
	processor *AdProcessor[realt.Flat, realt.GraphQLResponse]
}

func NewRealtSendJob(bot *telegram.Bot) *RealtSendJob {
	return &RealtSendJob{
		processor: &AdProcessor[realt.Flat, realt.GraphQLResponse]{
			ParseURL: constants.REALT_PARSE_URL,
			AdType:   constants.T_REALT,
			TgBot:    bot,
			BuildOpts: func() []options.Option {
				body := realt.BuildFlatsBody()
				return []options.Option{
					options.WithMethod(http.MethodPost, bytes.NewReader(body)),
					options.WithHeader("Content-Type", "application/json"),
					options.WithHeader("Accept", "*/*"),
					options.WithHeader("Accept-Language", "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7"),
					options.WithHeader("Origin", "https://realt.by"),
					options.WithHeader("Referer", "https://realt.by/rent/flat-for-long/?sortType=createdAt&page=1"),
					options.WithHeader("x-realt-client", "www@4.24.5"),
					options.WithHeader("Cookie", `consent={"analytics":true,"advertising":true,"functionality":true}`),
				}
			},
			BuildURL: helper.BuildRealtURL,
			MakeDesc: func(flat realt.Flat) string {
				ad := helper.RealtAd{Flat: flat}
				return helper.MakeDesc(ad)
			},
			ExtractID: func(f realt.Flat) int { return f.Code },
			GetAds: func(r realt.GraphQLResponse) []realt.Flat {
				var ads []realt.Flat

				flatAds := r.Data.SearchObjects.Body.Results

				for _, flat := range flatAds {
					if flat.Price > constants.HIGHER_SEARCH_PRICE || flat.Price < constants.LOWER_SEARCH_PRICE {
						continue
					}
					ads = append(ads, flat)
				}
				return ads
			},
			GetImages: func(f realt.Flat) []string { return f.Images },
		},
	}
}

func (r *RealtSendJob) Execute() error {
	return r.processor.Execute()
}

type KufarSendJob struct {
	processor *AdProcessor[kufar.Ad, kufar.Response]
}

func NewKufarSendJob(bot *telegram.Bot) *KufarSendJob {
	return &KufarSendJob{
		processor: &AdProcessor[kufar.Ad, kufar.Response]{
			ParseURL: constants.KUFAR_PARSE_URL,
			AdType:   constants.T_KUFAR,
			TgBot:    bot,
			BuildURL: helper.BuildKufarURL,
			MakeDesc: func(ad kufar.Ad) string {
				f := helper.KufarAd{Ad: ad}
				return helper.MakeDesc(f)
			},
			ExtractID: func(a kufar.Ad) int { return a.AdID },
			GetAds:    func(r kufar.Response) []kufar.Ad { return r.Ads },
			GetImages: func(a kufar.Ad) []string {
				paths := make([]string, len(a.Images))
				for i, img := range a.Images {
					paths[i] = img.Path
				}
				return paths
			},
		},
	}
}

func (k *KufarSendJob) Execute() error {
	return k.processor.Execute()
}

type OnlinerSendJob struct {
	processor *AdProcessor[onliner.Apartment, onliner.Response]
}

func NewOnlinerSendJob(bot *telegram.Bot) *OnlinerSendJob {
	return &OnlinerSendJob{
		processor: &AdProcessor[onliner.Apartment, onliner.Response]{
			ParseURL: constants.ONLINER_PARSE_URL,
			AdType:   constants.T_ONLINER,
			TgBot:    bot,
			BuildURL: helper.BuildOnlinerURL,
			MakeDesc: func(apartment onliner.Apartment) string {
				ad := helper.OnlinerAd{Apartment: apartment}
				return helper.MakeDesc(ad)
			},
			ExtractID: func(a onliner.Apartment) int { return a.ID },
			GetAds: func(r onliner.Response) []onliner.Apartment {
				var ads []onliner.Apartment
				for _, apartment := range r.Apartments {
					price, err := strconv.ParseFloat(apartment.Price.Amount, 64)
					if err != nil || price > constants.HIGHER_SEARCH_PRICE || price < constants.LOWER_SEARCH_PRICE {
						continue
					}
					ads = append(ads, apartment)
				}
				return ads
			},
			GetImages: func(a onliner.Apartment) []string { return []string{a.Photo} },
		},
	}
}

func (o *OnlinerSendJob) Execute() error {
	return o.processor.Execute()
}

func (p *AdProcessor[T, R]) Execute() error {
	var (
		res R
		err error
	)

	if p.BuildOpts != nil {
		res, err = parser.Parse[R](p.ParseURL, p.BuildOpts()...)
	} else {
		res, err = parser.Parse[R](p.ParseURL)
	}
	if err != nil {
		return fmt.Errorf("ошибка парсинга: %w", err)
	}

	ads := p.GetAds(res)

	oldFlats := flats.Get(p.AdType, helper.ExtractIDs(ads, p.ExtractID))
	oldIDs := make(map[int]struct{}, len(oldFlats))
	for _, plot := range oldFlats {
		oldIDs[plot.RemoteID] = struct{}{}
	}

	log.Println(oldFlats)

	var newAds []T
	for _, ad := range ads {
		if _, exists := oldIDs[p.ExtractID(ad)]; !exists {
			newAds = append(newAds, ad)
		}
	}

	if len(newAds) < 1 {
		log.Println("Length is zero")
		return nil
	}

	adaptFlats := adapter.AdsToFlats(p.AdType, newAds, p.ExtractID)

	for _, newAd := range newAds {
		if (len(newAds)) > 10 {
			log.Println("Too many new ads, skipping...")
			break
		}

		images := p.BuildURL(p.GetImages(newAd))
		err := p.TgBot.SendMediaWithText(images, p.MakeDesc(newAd))
		if err != nil {
			log.Printf("ошибка отправки медиа для объявления %d: %v", p.ExtractID(newAd), err)
			continue
		}
		time.Sleep(4 * time.Second)
		log.Println("Отправил сообщение")
	}

	if err := flats.Insert(adaptFlats); err != nil {
		return fmt.Errorf("ошибка вставки в базу данных: %w", err)
	}

	log.Println("Успешное выполнение")

	return nil
}

func runJob(j Job) func() error {
	return func() error {
		return j.Execute()
	}
}

func RunScheduler(j Job) error {
	s, err := gocron.NewScheduler()
	if err != nil {
		panic(err)
	}

	task := gocron.NewTask(runJob(j))

	_, err = s.NewJob(gocron.DurationJob(1*time.Minute), task)
	if err != nil {
		return err
	}

	s.Start()

	select {}
}
