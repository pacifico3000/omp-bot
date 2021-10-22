package announcement

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

func (c *StreamingAnnouncementCommander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		c.SendBotErrorMessage(inputMessage, "Usage: /get__streaming__announcement {announcement index}", "Get")
		return
	}

	a, err := c.announcementService.Describe(uint64(idx))
	if err != nil || a == nil {
		log.Printf("fail to get announcement with idx %d: %v", idx, err)
		c.SendBotErrorMessage(inputMessage, "Failed to get announcement with id: " + strconv.Itoa(idx), "Get")
		return
	}
	if a == nil {
		log.Printf("fail to get announcement with idx %d: %v", idx, err)
		c.SendBotErrorMessage(inputMessage, "No such announcement with id: " + strconv.Itoa(idx), "Get")
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		a.String(),
	)

	c.SendBotMessage(msg, "Get")
}
