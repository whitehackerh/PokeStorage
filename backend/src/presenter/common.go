package presenter

import (
	"fmt"
	"time"
)

type CommonPresenter struct {
	Time string      `json:"time"`
	Url  string      `json:"url"`
	Data interface{} `json:"data"`
}

func NewCommonPresenter(url string, data interface{}) CommonPresenter {
	t := time.Now()
	return CommonPresenter{
		Time: fmt.Sprintf(
			"%d-%02d-%02d %02d:%02d:%02d.%03d",
			t.Year(), t.Month(), t.Day(),
			t.Hour(), t.Minute(), t.Second(), t.Nanosecond()/1000000,
		),
		Url:  url,
		Data: data,
	}
}
