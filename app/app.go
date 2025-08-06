package app

import (
	"flats-parser/cron"
	"flats-parser/telegram"
	"log"
)

func Run() {
	//db.Init()
	tgBot := telegram.New()
	kufarJob := cron.NewKufarSendJob(tgBot)
	realtJob := cron.NewRealtSendJob(tgBot)
	onlinerJob := cron.NewOnlinerSendJob(tgBot)

	log.Println("Jobs have been created")

	go func() {
		err := cron.RunScheduler(kufarJob)
		if err != nil {
			log.Printf("Ошибка в kufarJob: %v", err)
		}
	}()

	go func() {
		err := cron.RunScheduler(realtJob)
		if err != nil {
			log.Printf("Ошибка в realtJob: %v", err)
		}
	}()

	go func() {
		err := cron.RunScheduler(onlinerJob)
		if err != nil {
			log.Printf("Ошибка в onlinerJob: %v", err)
		}
	}()

	select {}
}
