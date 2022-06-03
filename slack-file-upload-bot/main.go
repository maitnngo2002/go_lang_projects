package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "COPY AND PASTE THE SLACK BOT TOKEN HERE")
	os.Setenv("CHANNEL_ID", "COPY AND PASTE THE CHANNEL ID HERE") // get the channel id in the slack workspace
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"backend.pdf"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File: fileArr[i],
		}
		file, err := api.UploadFile(params)

		if err != nil {
			fmt.Printf("%s \n", err)
			return
		}
		fmt.Printf("Name: %s, URL: %s\n", file.Name, file.URL)
	}
}