package announcement

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/service/streaming/announcement"
	"log"
	"strconv"
	"strings"
)

func (c *StreamingAnnouncementCommander) Edit(inputMessage *tgbotapi.Message) {
	args := strings.SplitN(inputMessage.CommandArguments(), " ", 2)
	if len(args) != 2 {
		c.sendEditFormatMessage(inputMessage)
		return
	}

	idx, err := strconv.ParseUint(args[0], 10, 64)
	if err != nil {
		log.Println("wrong args", args)
		c.sendEditFormatMessage(inputMessage)
		return
	}

	var parsedData AnouncementData
	err = json.Unmarshal([]byte(args[1]), &parsedData)
	if err != nil {
		log.Printf("StreamingAnnouncementCommander.Edit: "+
			"error reading json data for type AnouncementData from "+
			"input string %v - %v", args, err)
		c.sendEditFormatMessage(inputMessage)
		return
	}

	a := announcement.Announcement{
		Author:       parsedData.Author,
		TimePlanned:  parsedData.TimePlanned,
		Title:        parsedData.Title,
		Description:  parsedData.Description,
		ThumbnailUrl: parsedData.ThumbnailUrl,
	}
	if c.announcementService.Update(idx, a) != nil {
		log.Printf("fail to update announcement with idx %d: %v", idx, err)
		c.SendBotErrorMessage(inputMessage, "Failed to update announcement with id: " + strconv.FormatUint(idx, 10), "Edit")
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"Announcement with id: " + strconv.FormatUint(idx, 10) + " was updated successfully",
	)

	c.SendBotMessage(msg, "Edit")
}

func (c *StreamingAnnouncementCommander) sendEditFormatMessage(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"Usage: /edit__streaming__announcement {announcement index} {announcement json}\n" +
			"JSON fields are:\nauthor(string),\n" +
			"title(string),\n" +
			"description(string),\n" +
			"time_planned(timestamp),\n" +
			"thumbnail_url(string)",
	)

	c.SendBotMessage(msg, "Edit")
}