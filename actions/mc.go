package actions

import (
	"encoding/json"
	"github.com/PassTheMayo/mcstatus/v4"
	"github.com/vicanso/go-axios"
	"khl-meobot/model/kook/post"
	"khl-meobot/utils/api"
	KookCard "khl-meobot/utils/kook-card"
	"strconv"
)

func GetMcServerStatus(targetID, quote, server, port string) (*axios.Response, error) {
	if server == "" {
		server = "mcn.cooo.cool"
	}
	if port == "" {
		port = "25565"
	}
	portNum, err := strconv.ParseUint(port, 10, 16)
	if err != nil {
		return SendMessage(targetID, "端口号非法", quote, 0)
	}
	response, err := mcstatus.Status(server, uint16(portNum))
	if err != nil {
		return SendMessage(targetID, "服务器超时或不在线", quote, 0)
	}
	playerList := "> 没有玩家数据"
	if len(response.Players.Sample) != 0 {
		playerList = "> "
		for _, v := range response.Players.Sample {
			playerList += v.Clean + "\n"
		}
	}
	cardMsg := []KookCard.CardMessage{
		{
			Type:  "card",
			Theme: "secondary",
			Size:  "lg",
			Modules: []KookCard.CardMessageModule{
				{
					Type: "header",
					Text: &KookCard.TextModule{
						Type:    "plain-text",
						Content: "Minecraft Server Status",
					},
				},
				{
					Type: "divider",
				},
				{
					Type: "section",
					Text: &KookCard.TextModule{
						Type:    "kmarkdown",
						Content: "**服务器：**" + server + ":" + port,
					},
				},
				{
					Type: "section",
					Text: &KookCard.TextModule{
						Type:    "kmarkdown",
						Content: "**消息：**" + response.MOTD.Clean,
					},
				},
				{
					Type: "section",
					Text: &KookCard.TextModule{
						Type:    "kmarkdown",
						Content: "**版本：**" + response.Version.Name,
					},
				},
				{
					Type: "section",
					Text: &KookCard.TextModule{
						Type:    "kmarkdown",
						Content: "**玩家：**" + strconv.Itoa(response.Players.Online) + "/" + strconv.Itoa(response.Players.Max),
					},
				},
				{
					Type: "section",
					Text: &KookCard.TextModule{
						Type:    "kmarkdown",
						Content: playerList,
					},
				},
			},
		},
	}
	msg, _ := json.Marshal(cardMsg)
	return api.SendChannelMessage(post.ChannelMessageCreate{
		Type:     10,
		TargetID: targetID,
		Content:  string(msg),
		Quote:    quote,
	})
}
