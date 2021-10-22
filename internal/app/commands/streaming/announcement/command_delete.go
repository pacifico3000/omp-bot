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
		c.SendBotErrorMessage(inputMessage, "Usage: /delete__streaming__announcement {announcement index}", "Delete")
		return
	}

	a, err := c.announcementService.Remove(uint64(idx))
	if err != nil {
		log.Printf("fail to remove announcement with idx %d: %v", idx, err)
		c.SendBotErrorMessage(inputMessage, "Failed to remove announcement with id: " + strconv.Itoa(idx), "Delete")
		return
	}
	if !a {
		log.Printf("fail to remove announcement with idx %d: %v", idx, err)
		c.SendBotErrorMessage(inputMessage, "No such announcement with id: " + strconv.Itoa(idx), "Delete")
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"Announcement was removed successfully",
	)

	c.SendBotMessage(msg, "Delete")
}