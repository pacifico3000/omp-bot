package announcement

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"log"
)

func (c *StreamingAnnouncementCommander) List(inputMessage *tgbotapi.Message) {
	var outputMsgText string
	announcements, err := c.announcementService.List(0, pageLimit)
	if err != nil {
		log.Printf("StreamingAnnouncement.List: error getting announcements list - %v", err)
		return
	}
	for _, p := range announcements {
		outputMsgText += p.String()
		outputMsgText += "\n----------------------\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	serializedData, err := json.Marshal(CallbackListData{
		Offset: 1,
	})
	if err != nil {
		log.Printf("StreamingAnnouncement.List: error serializing callback data - %v", err)
		return
	}

	callbackPath := path.CallbackPath{
		Domain:       "streaming",
		Subdomain:    "announcement",
		CallbackName: "list",
		CallbackData: string(serializedData),
	}
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()),
		),
	)

	c.SendBotMessage(msg, "List")
}