package messenger

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Messenger struct {
	bot     *tg.BotAPI
	updates tg.UpdatesChannel
	chats   []int64
}

func New(token string, chats []int64) (*Messenger, error) {
	bot, err := tg.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	bot.Debug = false

	return &Messenger{
		bot:     bot,
		updates: nil,
		chats:   chats,
	}, nil
}

func (m *Messenger) Send(msg string) error {
	for _, ci := range m.chats {
		msgTg := tg.NewMessage(ci, msg)
		msgTg.ParseMode = tg.ModeHTML
		_, err := m.bot.Send(msgTg)
		if err != nil {
			return err
		}
	}

	return nil
}
