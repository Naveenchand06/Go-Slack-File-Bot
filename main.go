package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	// * You can add list of channel IDs in channels slice. But in this example we are only sending files to one file
	channelsArr := []string{os.Getenv("CHANNEL_ID")}
	// * You can add the name of files you want to upload to a channel (In this we are testing by uploading pdf file)
	filesArr := []string{"dummy.pdf"}

	for i:=0; i<len(filesArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelsArr,
			File: filesArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n", err)
		}
		fmt.Printf("Name: %s, URL:%s\n", file.Name, file.URL)
	}

}