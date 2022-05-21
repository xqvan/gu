package gu

import (
	"fmt"
	"log"
	"runtime"
)

const (
	alertUrl       = "https://api.day.app/push"
	deviceKeyI11Go = "kzNGcnEmJW8kHFzWoYxv2C"
)

var (
	deviceKey = deviceKeyI11Go
)

func SetDeviceKey(key string) {
	deviceKey = key
}

func TitleContent(title, content string) {
	reqBody := map[string]any{
		"device_key": deviceKey,
		"title":      title,
		"body":       content,
	}
	get, err := CommonPostJson(alertUrl, nil, reqBody)
	log.Printf("\n[alert]: \ntitle: %s \ncontent:%s \nrsp:%s \nerr:%v\n\n", title, content, string(get), err)
}

func FatalTitleContent(title, content string) {
	TitleContent(title, content)
	log.Fatalln(title, content)
}

func FatalError(err error) {
	title := fmt.Sprintln(getFileAndLine(2))
	content := err.Error()
	FatalTitleContent(title, content)
}

func getFileAndLine(level int) (string, int) {
	_, file, line, _ := runtime.Caller(level)
	return file, line
}
