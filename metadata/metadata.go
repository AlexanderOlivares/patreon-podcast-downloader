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

func GetMetadata(feedUrl string, targetEpisodeGUID ...string) []Episode {

	fp := gofeed.NewParser()

	feed, err := fp.ParseURL(feedUrl)
	if err != nil {
		log.Fatal("Error parsing RSS feed:", err)
	}

	fmt.Println("Feed Title:", feed.Title)
	fmt.Println("Feed Description:", feed.Description)

	episodes := make([]Episode, 0, len(feed.Items))
	for _, item := range feed.Items {
		// Only return the targeted episode
		if len(targetEpisodeGUID) > 0 && item.GUID == targetEpisodeGUID[0] {
			return []Episode{
				{
					Title: item.Title,
					Date:  item.Published,
					URL:   item.Enclosures[0].URL,
					GUID:  item.GUID,
				},
			}
		}
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
