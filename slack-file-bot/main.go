package main

import (
"fmt"
"os"
"github.com/slack-go/slack"
)

func main(){

os.Setenv("SLACK_BOT_TOKEN","xoxb-6134226757906-6131379879589-zGXSOHlGTefELwAhmksHW0AO")
os.Setenv("CHANNEL_ID","C064ARPLH8R")
api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
channelArr := []string{os.Getenv("CHANNEL_ID")}
fileArr := []string{"Social_link.txt"} //slice

for i :=0; i<len(fileArr); i++{

	params := slack.FileUploadParameters{
		Channels: channelArr,
		File: fileArr[i],
	}
	file, err := api.UploadFile(params)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("Name: %s, URL: %s\n", file.Name, file.URL)
}




}