package announcement

import (
	"strconv"
	"time"
)

var allEntities = []Announcement{
	{
		ID: 1,
		Author: "John Doe",
		TimePlanned: time.Now(),
		Title: "Sample 1",
		Description: "Sample description",
		ThumbnailUrl: "example.com",
	},
	{
		ID: 2,
		Author: "Jane Doe",
		TimePlanned: time.Now(),
		Title: "Sample 2",
		Description: "Sample description",
		ThumbnailUrl: "example.com",
	},
	{
		ID: 3,
		Author: "John Doe",
		TimePlanned: time.Now(),
		Title: "Sample 3",
		Description: "Sample description",
		ThumbnailUrl: "example.com",
	},
	{
		ID: 4,
		Author: "Jane Doe",
		TimePlanned: time.Now(),
		Title: "Sample 4",
		Description: "Sample description",
		ThumbnailUrl: "example.com",
	},
	{
		ID: 5,
		Author: "John Doe",
		TimePlanned: time.Now(),
		Title: "Sample 5",
		Description: "Sample description",
		ThumbnailUrl: "example.com",
	},
	{
		ID: 6,
		Author: "Jane Doe",
		TimePlanned: time.Now(),
		Title: "Sample 6",
		Description: "Sample description",
		ThumbnailUrl: "example.com",
	},
	{
		ID: 7,
		Author: "John Doe",
		TimePlanned: time.Now(),
		Title: "Sample 7",
		Description: "Sample description",
		ThumbnailUrl: "example.com",
	},
	{
		ID: 8,
		Author: "Jane Doe",
		TimePlanned: time.Now(),
		Title: "Sample 8",
		Description: "Sample description",
		ThumbnailUrl: "example.com",
	},
	{
		ID: 9,
		Author: "John Doe",
		TimePlanned: time.Now(),
		Title: "Sample 9",
		Description: "Sample description",
		ThumbnailUrl: "example.com",
	},
	{
		ID: 10,
		Author: "Jane Doe",
		TimePlanned: time.Now(),
		Title: "Sample 10",
		Description: "Sample description",
		ThumbnailUrl: "example.com",
	},
}

type Announcement struct {
	ID uint64
	Author string
	TimePlanned time.Time
	Title string
	Description string
	ThumbnailUrl string
}

func (a *Announcement) FormattedTime() string {
	return a.TimePlanned.UTC().Format(time.UnixDate)
}

func (a *Announcement) String() string {
	return "ID: " + strconv.FormatUint(a.ID, 10) + "\n" +
		"Author: " + a.Author + "\n" +
		"TimePlanned: " + a.FormattedTime() + "\n" +
		"Title: " + a.Title + "\n" +
		"Description: " + a.Description + "\n" +
		"Thumbnail: " + a.ThumbnailUrl
}