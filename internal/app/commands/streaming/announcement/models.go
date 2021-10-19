package announcement

import "time"

const pageLimit = 5

type AnouncementData struct {
	Author       string 	`json:"author"`
	Title        string 	`json:"title"`
	Description  string 	`json:"description"`
	TimePlanned  time.Time `json:"time_planned"`
	ThumbnailUrl string 	`json:"thumbnail_url"`
}