package api

import (
	"github.com/vicanso/go-axios"
	"khl-meobot/model/kook/post"
)

func SendChannelMessage(data post.ChannelMessageCreate) (resp *axios.Response, err error) {
	return KookInstance().Post("/message/create", data)
}

func UpdateChannelMessage(data post.ChannelMessageUpdate) (resp *axios.Response, err error) {
	return KookInstance().Post("/message/update", data)
}

func SendDirectMessage(data post.ChannelMessageCreate) (resp *axios.Response, err error) {
	return KookInstance().Post("/direct-message/create", data)
}
