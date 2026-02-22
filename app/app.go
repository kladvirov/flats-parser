package app

import (
	"flats-parser/cron"
	"flats-parser/db"
	"flats-parser/telegram"
	"github.com/go-co-op/gocron/v2"
	"log/slog"
	"math/rand"
	"time"
)

func Run() {
	db.Init()
	tgBot := telegram.New()
	kufarJob := cron.NewKufarSendJob(tgBot)
	realtJob := cron.NewRealtSendJob(tgBot)
	onlinerJob := cron.NewOnlinerSendJob(tgBot)

	slog.Info("Jobs have been created")

	s, err := gocron.NewScheduler()
	if err != nil {
		panic(err)
	}

	jobs := []cron.Job{realtJob, onlinerJob, kufarJob}

	for _, job := range jobs {
		task := gocron.NewTask(func(j cron.Job) func() {
			return func() {
				if err := j.Execute(); err != nil {
					slog.Error("Job execution error", "error", err)
				}
			}
		}(job))

		randV := time.Duration(rand.Intn(15)) * time.Second
		cronJob := gocron.DurationJob(45*time.Second + randV)
		_, err = s.NewJob(
			cronJob,
			task,
		)
		if err != nil {
			slog.Error("Failed to create job", "error", err)
		}
	}

	s.Start()
	slog.Info("Scheduler started")

	select {}
}
