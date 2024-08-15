package presenter

import (
	"fmt"
	"time"
)

type CommonPresenter struct {
	Time string      `json:"time"`
	Api  string      `json:"api"`
	Data interface{} `json:"data"`
}

func NewCommonPresenter(api string, data interface{}) CommonPresenter {
	t := time.Now()
	return CommonPresenter{
		Time: fmt.Sprintf(
			"%d-%02d-%02d %02d:%02d:%02d.%03d",
			t.Year(), t.Month(), t.Day(),
			t.Hour(), t.Minute(), t.Second(), t.Nanosecond()/1000000,
		),
		Api:  api,
		Data: data,
	}
}
