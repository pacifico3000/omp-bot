package announcement

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *StreamingAnnouncementCommander) Default(inputMessage *tgbotapi.Message) {
	c.SendBotMessage(tgbotapi.NewMessage(inputMessage.Chat.ID, "You wrote: "+inputMessage.Text), "Default")
}