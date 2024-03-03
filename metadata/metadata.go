package metadata

import (
	"fmt"
	"log"

	"github.com/mmcdole/gofeed"
)

type Episode struct {
	Title string
	Date  string
	URL   string
	GUID  string
}

func GetMetadata(feedUrl string) []Episode {

	fp := gofeed.NewParser()

	feed, err := fp.ParseURL(feedUrl)
	if err != nil {
		log.Fatal("Error parsing RSS feed:", err)
	}

	fmt.Println("Feed Title:", feed.Title)
	fmt.Println("Feed Description:", feed.Description)

	var episodes []Episode
	for _, item := range feed.Items {
		episode := Episode{
			Title: item.Title,
			Date:  item.Published,
			URL:   item.Enclosures[0].URL,
			GUID:  item.GUID,
		}
		episodes = append(episodes, episode)
	}
	return episodes
}
