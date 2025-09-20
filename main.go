package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/imrany/spindle/internal/scrape"
	"github.com/imrany/spindle/server"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Profile struct {
	Addr string
	Port int
}

var rootCmd = &cobra.Command{
	Use:   "crawler",
	Short: "An open source, lightweight web crawler and scraper",
	Long:  `An open-source, lightweight web crawler and scraper. It can discover links on the web (crawler) and extract structured data from webpages (scraper).`,
	Run:   runServer,
}

func runServer(_ *cobra.Command, _ []string) {
	profile := Profile{
		Addr: viper.GetString("addr"),
		Port: viper.GetInt("port"),
	}

	fmt.Printf("Starting server on %s:%d...\n", profile.Addr, profile.Port)
	if err := server.StartServer(profile.Addr, profile.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func init() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [server|url]")
		os.Exit(1)
	}

	firstArg := os.Args[1]

	// If argument is "server", configure cobra/viper for server mode
	if firstArg == "server" {
		// Load .env file if available
		if err := godotenv.Load(); err != nil {
			slog.Warn("No .env file found or failed to load", "error", err)
		}

		viper.AutomaticEnv()

		// Bind CLI flags and ENV variables
		rootCmd.PersistentFlags().String("addr", "0.0.0.0", "Bind address")
		rootCmd.PersistentFlags().IntP("port", "p", 5020, "Set the port for the server")

		_ = viper.BindPFlag("addr", rootCmd.PersistentFlags().Lookup("addr"))
		_ = viper.BindPFlag("port", rootCmd.PersistentFlags().Lookup("port"))

	} else if len(firstArg) > 7 &&
		(firstArg[:7] == "http://" || firstArg[:8] == "https://") {

		// CLI mode → directly extract info from URL
		url := firstArg
		pageInfo, err := scrape.ExtractInfo(url)
		if err != nil {
			log.Fatalf("Error scraping URL: %v", err)
		}

		fmt.Printf(
			"Title: %s\nDescription: %s\nFavicon: %s\nVideo: %s\nLinks: %v\nImages: %v\n",
			pageInfo.Title, pageInfo.Description, pageInfo.Favicon,
			pageInfo.Video, pageInfo.Links, pageInfo.Images,
		)

		os.Exit(0) // exit after CLI scrape
	} else {
		fmt.Println("Usage: go run main.go [server|url]")
		os.Exit(1)
	}
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		slog.Error("failed to run command", "error", err)
		os.Exit(1)
	}
}
