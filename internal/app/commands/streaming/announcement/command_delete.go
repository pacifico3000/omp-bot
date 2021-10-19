package announcement

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

func (c *StreamingAnnouncementCommander) Delete(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			"Usage: /delete__streaming__announcement {announcement index}",
		)
		c.SendBotMessage(msg, "Delete")
		return
	}

	a, err := c.announcementService.Remove(uint64(idx))
	if err != nil {
		log.Printf("fail to remove announcement with idx %d: %v", idx, err)
		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			"Failed to remove announcement with id: " + strconv.Itoa(idx),
		)
		c.SendBotMessage(msg, "Delete")
		return
	}
	if !a {
		log.Printf("fail to remove announcement with idx %d: %v", idx, err)
		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			"No such announcement with id: " + strconv.Itoa(idx),
		)
		c.SendBotMessage(msg, "Delete")
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"Announcement was removed successfully",
	)

	c.SendBotMessage(msg, "Delete")
}