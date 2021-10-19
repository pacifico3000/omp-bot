package announcement

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"log"
)

type CallbackListData struct {
	Offset int `json:"offset"`
}

func (c *StreamingAnnouncementCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	parsedData := CallbackListData{}
	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	if err != nil {
		log.Printf("StreamingAnnouncementCommander.CallbackList: " +
			"error reading json data for type CallbackListData from " +
			"input string %v - %v", callbackPath.CallbackData, err)
		return
	}

	var outputMsgText string
	announcements, err := c.announcementService.List(uint64(parsedData.Offset), pageLimit)
	if err != nil {
		log.Printf("StreamingAnnouncement.CallbackList: error getting announcements list - %v", err)
		return
	}

	for _, p := range announcements {
		outputMsgText += p.String()
		outputMsgText += "\n----------------------\n"
	}
	if len(announcements) <= 0 {
		outputMsgText = "No more items :("
	}
	msg := tgbotapi.NewMessage(callback.Message.Chat.ID, outputMsgText)

	var buttons []tgbotapi.InlineKeyboardButton

	if parsedData.Offset > 0 {
		offsetData, _ := json.Marshal(CallbackListData{
			Offset: parsedData.Offset - 1,
		})
		callbackPath = path.CallbackPath{
			Domain:       "streaming",
			Subdomain:    "announcement",
			CallbackName: "list",
			CallbackData: string(offsetData),
		}
		buttons = append(buttons,
			tgbotapi.NewInlineKeyboardButtonData("Previous page", callbackPath.String()),
		)
	}

	if len(announcements) == pageLimit {
		offsetData, err := json.Marshal(CallbackListData{
			Offset: parsedData.Offset + 1,
		})
		if err != nil {
			log.Printf("StreamingAnnouncementCommander.CallbackList: " +
				"error building json data for type CallbackListData")
			return
		}

		callbackPath = path.CallbackPath{
			Domain:       "streaming",
			Subdomain:    "announcement",
			CallbackName: "list",
			CallbackData: string(offsetData),
		}
		buttons = append(buttons,
			tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()),
		)
	}

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(buttons...),
	)

	c.SendBotMessage(msg, "CallbackList")
}
