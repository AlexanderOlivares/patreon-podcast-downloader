package cmd

import (
	"ppd/download"
	"ppd/metadata"
	"ppd/util"

	"github.com/spf13/cobra"
)

var downloadLatestCommand = &cobra.Command{
	Use:   "download-latest",
	Short: "Download the n-latest episodes from a feed",
	Long: `Download the n-latest episodes from a feed
	
		This command allows you to download a specific number of the latest episodes from a feed. Omit the latest-episodes flag to download all episodes from the feed

		For example:
			podcasts download-latest --feed-url https://www.patreon.com/rss/xxx?auth=xxx latest-episodes 3 --output /output/path/directory --date-prefix true
			or 
			podcasts download-one -f=https://www.patreon.com/rss/xxx?auth=xxx -l=3 -o=/output/path/directory -d=true

			Replace "xxx" with the feed name and your auth token 
			Replace "3" with the number of episodes you want to download (omit to download all)
			Replace "/output/path/directory" with the path to an existing directory you wish to download files to
			If date-prefix is true, episode file names will begin with the publication date in format YYYY-MM-DD
	`,
	Run: downloadLatest,
}

func init() {
	rootCmd.AddCommand((downloadLatestCommand))
	defaultOutputDir := util.GetDefaultDownloadDirectory()
	downloadLatestCommand.Flags().StringP("feed-url", "f", "12345", "Feed URL")
	downloadLatestCommand.Flags().IntP("latest-episodes", "l", -1, "Download the latest n episodes, or omit to download all")
	downloadLatestCommand.Flags().StringP("output", "o", defaultOutputDir, "Path to directory to download files to")
	downloadLatestCommand.Flags().BoolP("date-prefix", "d", false, "Path to directory to download files to")
}

func downloadLatest(cmd *cobra.Command, args []string) {
	feedUrl, fError := cmd.Flags().GetString("feed-url")
	nLatestEpisodes, nError := cmd.Flags().GetInt("latest-episodes")
	path, oError := cmd.Flags().GetString("output")
	prefixWithDate, pError := cmd.Flags().GetBool("date-prefix")

	util.CheckAndPrintErrors(fError, nError, oError, pError)

	commandInput := download.DownloadOptions{
		NLatestEpisodes:     nLatestEpisodes,
		OutputPath:          path,
		PrefixWithPubDate:   prefixWithDate,
		TargetEpisodeNumber: "",
	}

	episodes := metadata.GetMetadata(feedUrl)
	download.DownloadEpisodes(episodes, commandInput)
}
