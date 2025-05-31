package helper

import (
	"flats-parser/constants"
	"flats-parser/parser/kufar"
	"flats-parser/parser/onliner"
	"flats-parser/parser/realt"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Advertisement interface {
	GetType() string
	GetLink() string
	GetDescription() string
	GetAddress() string
	GetPrice() int
	GetCurrency() string
	GetCommonSquare() float64
	GetLivingSquare() float64
	GetFloor() int
}

type KufarAd struct {
	kufar.Ad
}

func (a KufarAd) GetType() string {
	return "KUFAR"
}

func (a KufarAd) GetLink() string {
	return a.AdLink
}

func (a KufarAd) GetDescription() string {
	return a.BodyShort
}

func (a KufarAd) GetAddress() string {
	return a.AccountParameters[1].V.(string)
}

func (a KufarAd) GetPrice() int {
	price, err := strconv.Atoi(a.PriceUSD)
	if err != nil {
		return 0
	}
	return price / 100
}

func (a KufarAd) GetCurrency() string {
	return "USD"
}

func (a KufarAd) GetCommonSquare() float64 {
	return squareHelper(a, regexp.MustCompile(`(?i)общая площадь`))
}

func (a KufarAd) GetLivingSquare() float64 {
	return squareHelper(a, regexp.MustCompile(`(?i)жилая площадь`))
}

func (a KufarAd) GetFloor() int {
	for _, param := range a.AdParameters {
		if param.PL != "Этаж" {
			continue
		}

		if vSlice, ok := param.V.([]interface{}); ok {
			if len(vSlice) == 0 {
				return 0
			}

			switch v := vSlice[0].(type) {
			case int:
				return v
			case float64:
				return int(v)
			default:
				return 0
			}
		}
	}

	return 0
}

func squareHelper(a KufarAd, r *regexp.Regexp) float64 {
	for _, param := range a.AdParameters {
		if !r.MatchString(param.PL) {
			continue
		}
		return param.V.(float64)
	}

	return 0.00
}

type RealtAd struct {
	realt.Flat
}

func (a RealtAd) GetType() string {
	return "REALT"
}

func (a RealtAd) GetLink() string {
	return fmt.Sprintf("https://realt.by/rent-flat-for-long/object/%d", a.Code)
}

func (a RealtAd) GetDescription() string {
	return a.Headline
}

func (a RealtAd) GetAddress() string {
	return fmt.Sprintf("%s\nБлижайшая станция метро: %s", a.Address, a.Metro)
}

func (a RealtAd) GetPrice() int {
	return a.Price
}

func (a RealtAd) GetCurrency() string {
	return "USD"
}

func (a RealtAd) GetCommonSquare() float64 {
	return a.AreaTotal
}

func (a RealtAd) GetLivingSquare() float64 {
	return a.AreaLiving
}

func (a RealtAd) GetFloor() int {
	return a.Floor
}

type OnlinerAd struct {
	onliner.Apartment
}

func (a OnlinerAd) GetType() string {
	return "ONLINER"
}

func (a OnlinerAd) GetLink() string {
	return a.URL
}

func (a OnlinerAd) GetDescription() string {
	var roomType string
	switch a.RentType {
	case "1_room":
		roomType = "1-комнатная"
	case "2_rooms":
		roomType = "2-комнатная"
	case "3_rooms":
		roomType = "3-комнатная"
	case "4_rooms":
		roomType = "4-комнатная"
	case "5_rooms":
		roomType = "5-комнатная"
	default:
		roomType = "Квартира"
	}
	return fmt.Sprintf("%s квартира", roomType)
}

func (a OnlinerAd) GetAddress() string {
	return a.Location.Address
}

func (a OnlinerAd) GetPrice() int {
	price, err := strconv.ParseFloat(a.Price.Amount, 64)
	if err != nil {
		return 0
	}
	return int(price)
}

func (a OnlinerAd) GetCurrency() string {
	return a.Price.Currency
}

func (a OnlinerAd) GetCommonSquare() float64 {
	return 0 // Not available in API
}

func (a OnlinerAd) GetLivingSquare() float64 {
	return 0 // Not available in API
}

func (a OnlinerAd) GetFloor() int {
	return 0 // Not available in API
}

func MakeDesc(ad Advertisement) string {
	var desc strings.Builder

	desc.WriteString(fmt.Sprintf("#%s\n\n", ad.GetType()))
	desc.WriteString(fmt.Sprintf("Ссылка: %s\n\n", ad.GetLink()))

	if description := ad.GetDescription(); description != "" {
		desc.WriteString(fmt.Sprintf("Описание: %s\n\n", description))
	}

	desc.WriteString(fmt.Sprintf("Адрес: %s\n\n", ad.GetAddress()))
	desc.WriteString(fmt.Sprintf("Этаж: %d\n\n", ad.GetFloor()))
	desc.WriteString(fmt.Sprintf("Общая площадь: %.1f м²\nЖилая площадь: %.1f м²\n\n", ad.GetCommonSquare(), ad.GetLivingSquare()))
	desc.WriteString(fmt.Sprintf("Цена: %d %s", ad.GetPrice(), ad.GetCurrency()))

	return desc.String()
}

func ExtractIDs[T any](items []T, getID func(T) int) []int {
	ids := make([]int, len(items))
	for i, item := range items {
		ids[i] = getID(item)
	}
	return ids
}

func BuildKufarURL(images []string) []string {
	images = limitSlice(images, 10)
	urls := make([]string, len(images))
	for i, image := range images {
		urls[i] = fmt.Sprintf(constants.KUFAR_GALLERY_PATH, image)
	}
	return urls
}

func BuildRealtURL(images []string) []string {
	return limitSlice(images, 10)
}

func BuildOnlinerURL(photos []string) []string {
	if len(photos) == 0 {
		return []string{}
	}
	return []string{photos[0]}
}

func limitSlice[T any](slice []T, maxLen int) []T {
	if len(slice) > maxLen {
		return slice[:maxLen]
	}
	return slice
}
