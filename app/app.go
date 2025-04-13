package app

import (
	"flats-parser/cron"
	"flats-parser/db"
	"flats-parser/telegram"
	"log"
)

func Run() {
	db.Init()
	tgBot := telegram.New()
	kufarJob := cron.NewKufarSendJob(tgBot)
	realtJob := cron.NewRealtSendJob(tgBot)

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

	select {}
}
