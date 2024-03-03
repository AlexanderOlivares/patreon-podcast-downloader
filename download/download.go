package download

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"ppd/metadata"
	"ppd/util"
	"sync"
	"time"
)

type DownloadOptions struct {
	OutputPath          string
	NLatestEpisodes     int
	TargetEpisodeNumber string
	PrefixWithPubDate   bool
}

func DownloadEpisodes(episodesMetadata []metadata.Episode, options DownloadOptions) {
	fmt.Printf("Downloading to %s", options.OutputPath)

	absOutputDir, err := filepath.Abs(options.OutputPath)
	if err != nil {
		fmt.Printf("Error getting absolute path for output directory: %s\n", err)
	}

	episodes := func() []metadata.Episode {
		// Download all episodes
		if options.NLatestEpisodes == -1 {
			return episodesMetadata
		}
		return episodesMetadata[:options.NLatestEpisodes]
	}()

	var wg sync.WaitGroup
	errChan := make(chan error, len(episodes))

	for _, episode := range episodes {

		title := util.SanitizeFileName(episode.Title)
		var datePrefix = ""

		parsedTime, err := time.Parse(time.RFC1123, episode.Date)
		if err != nil {
			fmt.Println("Error parsing date:", err)
			return
		}

		datePrefix = parsedTime.Format("2006-01-02")

		fileName := func() string {
			if options.PrefixWithPubDate {
				return datePrefix + "_" + title
			}
			return title
		}()

		path := filepath.Join(absOutputDir, fileName)
		file := fmt.Sprintf("%s.mp3", path)

		wg.Add(1)
		fmt.Printf("Downloading episode:  %s\n", episode.Title)
		go downloadFile(episode.URL, file, &wg, errChan)
	}

	wg.Wait()

	close(errChan)
	for err := range errChan {
		if err != nil {
			fmt.Printf("Error downloading file: %s\n", err)
		}
	}

}

func downloadFile(url, outputPath string, wg *sync.WaitGroup, errChan chan error) {
	defer wg.Done()

	out, err := os.Create(outputPath)
	if err != nil {
		errChan <- err
		return
	}
	defer out.Close()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		errChan <- err
		return
	}
	req.Header.Set("Accept-Encoding", "gzip, deflate")

	client := &http.Client{
		Transport: &http.Transport{
			TLSNextProto: make(map[string]func(authority string, c *tls.Conn) http.RoundTripper),
		},
	}

	resp, err := client.Do(req)

	if err != nil {
		errChan <- err
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errChan <- fmt.Errorf("failed to download %s: %s", url, resp.Status)
		return
	}

	writer := bufio.NewWriter(out)
	_, err = io.Copy(writer, resp.Body)

	if err != nil {
		errChan <- err
		return
	}

	fmt.Printf("File downloaded: %s\n", outputPath)
}
