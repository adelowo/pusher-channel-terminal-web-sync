package main

import (
	"bufio"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	pusher "github.com/pusher/pusher-http-go"
)

const (
	channelName = "realtime-terminal"
	eventName   = "logs"
)

func main() {

	info, err := os.Stdin.Stat()
	if err != nil {
		log.Fatal(err)
	}

	if info.Mode()&os.ModeCharDevice != 0 || info.Size() <= 0 {
		log.Println("This command is intended to be used as a pipe such as yourprogram | thisprogram")
		os.Exit(0)
	}

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	appID := os.Getenv("PUSHER_APP_ID")
	appKey := os.Getenv("PUSHER_APP_KEY")
	appSecret := os.Getenv("PUSHER_APP_SECRET")
	appCluster := os.Getenv("PUSHER_APP_CLUSTER")
	appIsSecure := os.Getenv("PUSHER_APP_SECURE")

	var isSecure bool
	if appIsSecure == "1" {
		isSecure = true
	}

	client := &pusher.Client{
		AppId:   appID,
		Key:     appKey,
		Secret:  appSecret,
		Cluster: appCluster,
		Secure:  isSecure,
		HttpClient: &http.Client{
			Timeout: time.Minute * 2,
		},
	}

	reader := bufio.NewReader(os.Stdin)

	// If you happen to be outputting a lot of data in the terminal
	// and you want to "slow down" a bit
	// Just remember to call writer.Flush after the for loop
	// writer := bufio.NewWriter(pusherChannelWriter{client:client})
	writer := pusherChannelWriter{client: client}

	for {
		in, _, err := reader.ReadLine()
		if err != nil && err == io.EOF {
			break
		}

		if _, err := writer.Write(in); err != nil {
			log.Fatalln(err)
		}
	}

}

type pusherChannelWriter struct {
	client *pusher.Client
}

func (pusher pusherChannelWriter) Write(p []byte) (int, error) {
	s := string(p)
	_, err := pusher.client.Trigger(channelName, eventName, s)
	return len(p), err
}
