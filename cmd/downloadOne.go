package cmd

import (
	"ppd/download"
	"ppd/metadata"
	"ppd/util"

	"github.com/spf13/cobra"
)

var downloadOneCommand = &cobra.Command{
	Use:   "download-one",
	Short: "Download a single episode from a feed",
	Long: `Download a single episode from a feed
	
		This command allows you to download 1 episode from a feed. 

		For example:
			podcasts download-one --episode-GUID 123456789 --feed-url https://www.patreon.com/rss/xxx?auth=xxx --output /output/path/directory --date-prefix true
			or 
			podcasts download-one -e=123456789 -f=https://www.patreon.com/rss/xxx?auth=xxx -o=/output/path/directory -d=true

			Replace "xxx" with the feed name and your auth token 
			Replace "123456789" with the number of episodes you want to download (omit to download all)
			Replace "/output/path/directory" with the path to an existing directory you wish to download files to
			If date-prefix is true, episode file names will begin with the publication date in format YYYY-MM-DD
	`,
	Run: downloadOne,
}

func init() {
	rootCmd.AddCommand((downloadOneCommand))
	defaultOutputDir := util.GetDefaultDownloadDirectory()
	downloadOneCommand.Flags().StringP("feed-url", "f", "12345", "Feed URL")
	downloadOneCommand.Flags().StringP("episode-GUID", "e", "", "GUID of the episode you wish to download")
	downloadOneCommand.Flags().StringP("output", "o", defaultOutputDir, "Path to directory to download files to")
	downloadOneCommand.Flags().BoolP("date-prefix", "d", false, "Prefix the downloaded file name with the date of publication")
}

func downloadOne(cmd *cobra.Command, args []string) {
	feedUrl, fError := cmd.Flags().GetString("feed-url")
	episodeGUID, nError := cmd.Flags().GetString("episode-GUID")
	path, oError := cmd.Flags().GetString("output")
	prefixWithDate, pError := cmd.Flags().GetBool("date-prefix")

	util.CheckAndPrintErrors(fError, nError, oError, pError)

	commandInput := download.DownloadOptions{
		NLatestEpisodes:     1,
		OutputPath:          path,
		PrefixWithPubDate:   prefixWithDate,
		TargetEpisodeNumber: episodeGUID,
	}

	episodes := metadata.GetMetadata(feedUrl, episodeGUID)
	download.DownloadEpisodes(episodes, commandInput)
}
