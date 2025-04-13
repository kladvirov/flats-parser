package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"os"
	"strconv"
)

type BotConfig struct {
	Token  string
	ChatID int
}

type Bot struct {
	Tg  *tgbotapi.BotAPI
	Cfg *BotConfig
}

func New() *Bot {
	cfg := newConfig()

	tgBot, err := tgbotapi.NewBotAPI(cfg.Token)
	if err != nil {
		return nil
	}

	return &Bot{
		Tg:  tgBot,
		Cfg: cfg,
	}
}

func (b *Bot) SendMediaWithText(photoURLs []string, caption string) error {
	var mediaGroup []interface{}

	for i, url := range photoURLs {
		photo := tgbotapi.NewInputMediaPhoto(tgbotapi.FileURL(url))
		if i == 0 {
			photo.Caption = caption
			photo.ParseMode = tgbotapi.ModeMarkdown
		}
		mediaGroup = append(mediaGroup, photo)
	}

	if len(mediaGroup) == 0 {
		return nil
	}

	msg := tgbotapi.MediaGroupConfig{
		ChatID: int64(b.Cfg.ChatID),
		Media:  mediaGroup,
	}

	_, err := b.Tg.SendMediaGroup(msg)
	if err != nil {
		return nil
	}

	return nil
}

func newConfig() *BotConfig {
	token := os.Getenv("TG_BOT_TOKEN")
	if token == "" {
		return nil
	}

	chatIDStr := os.Getenv("TG_CHAT_ID")
	if chatIDStr == "" {
		return nil
	}
	chatID, err := strconv.Atoi(chatIDStr)
	if err != nil {
		return nil
	}

	return &BotConfig{
		Token:  token,
		ChatID: chatID,
	}
}
