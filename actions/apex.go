package actions

import (
	"encoding/json"
	"github.com/vicanso/go-axios"
	"khl-meobot/model/kook/post"
	"khl-meobot/utils/api"
	KookCard "khl-meobot/utils/kook-card"
	"log"
)

func GetApexMapRotate(targetID, quote string, updateFlags bool) (*axios.Response, error) {
	info, _ := api.ApexMapRotate()
	if info.Ranked.Next.Map == "Broken Moon" {
		info.Ranked.Next.Asset = "https://media.contentapi.ea.com/content/dam/apex-legends/common/articles/broken-moon/s-15-00-moon.png.adapt.1920w.png"
	}
	apexMessage := []KookCard.CardMessage{
		{
			Type:  "card",
			Theme: "warning",
			Size:  "lg",
			Modules: []KookCard.CardMessageModule{
				{
					Type: "section",
					Text: &KookCard.TextModule{
						Type:    "kmarkdown",
						Content: "**Apex 地图轮换**",
					},
					Mode: "right",
					Accessory: &KookCard.ElementsModule{
						Type:  "button",
						Theme: "success",
						Click: "return-val",
						Value: "refresh-apexmap",
						Text: &KookCard.TextModule{
							Type:    "plain-text",
							Content: "刷新",
						},
					},
				},
				{
					Type: "divider",
				},
				{
					Type: "section",
					Text: &KookCard.TextModule{
						Type:    "kmarkdown",
						Content: "**🗺 大逃杀**",
					},
				},
				{
					Type: "section",
					Text: &KookCard.TextModule{
						Type: "paragraph",
						Cols: 2,
						Fields: &[]KookCard.FieldsModule{
							{
								Type:    "kmarkdown",
								Content: "**当前地图**：[" + api.MaptoZh(info.BattleRoyale.Current.Map) + "](" + info.BattleRoyale.Current.Asset + ")",
							},
							{
								Type:    "kmarkdown",
								Content: " **下一轮换**：[" + api.MaptoZh(info.BattleRoyale.Next.Map) + "](" + info.BattleRoyale.Next.Asset + ")",
							},
						},
					},
				},
				{
					Type:    "countdown",
					Mode:    "day",
					EndTime: info.BattleRoyale.Current.End * 1000,
				},
				{
					Type: "divider",
				},
				{
					Type: "section",
					Text: &KookCard.TextModule{
						Type:    "kmarkdown",
						Content: "**🗺 竞技场**",
					},
				},
				{
					Type: "section",
					Text: &KookCard.TextModule{
						Type: "paragraph",
						Cols: 2,
						Fields: &[]KookCard.FieldsModule{
							{
								Type:    "kmarkdown",
								Content: "**当前地图**：[" + api.MaptoZh(info.Arenas.Current.Map) + "](" + info.Arenas.Current.Asset + ")",
							},
							{
								Type:    "kmarkdown",
								Content: " **下一轮换**：[" + api.MaptoZh(info.Arenas.Next.Map) + "](" + info.Arenas.Next.Asset + ")",
							},
						},
					},
				},
				{
					Type:    "countdown",
					Mode:    "day",
					EndTime: info.Arenas.Current.End * 1000,
				},
				{
					Type: "divider",
				},
				{
					Type: "section",
					Text: &KookCard.TextModule{
						Type:    "kmarkdown",
						Content: "**🗺 排位赛**",
					},
				},
				{
					Type: "section",
					Text: &KookCard.TextModule{
						Type: "paragraph",
						Cols: 2,
						Fields: &[]KookCard.FieldsModule{
							{
								Type:    "kmarkdown",
								Content: "**当前地图**：[" + api.MaptoZh(info.Ranked.Current.Map) + "](" + info.Ranked.Current.Asset + ")",
							},
							{
								Type:    "kmarkdown",
								Content: " **下一轮换**：[" + api.MaptoZh(info.Ranked.Next.Map) + "](" + info.Ranked.Next.Asset + ")",
							},
						},
					},
				},
				{
					Type:    "countdown",
					Mode:    "day",
					EndTime: info.Ranked.Current.End * 1000,
				},
				{
					Type: "divider",
				},
				{
					Type: "section",
					Text: &KookCard.TextModule{
						Type:    "kmarkdown",
						Content: "**🗺 排位竞技场**",
					},
				},
				{
					Type: "section",
					Text: &KookCard.TextModule{
						Type: "paragraph",
						Cols: 2,
						Fields: &[]KookCard.FieldsModule{
							{
								Type:    "kmarkdown",
								Content: "**当前地图**：[" + api.MaptoZh(info.ArenasRanked.Current.Map) + "](" + info.ArenasRanked.Current.Asset + ")",
							},
							{
								Type:    "kmarkdown",
								Content: " **下一轮换**：[" + api.MaptoZh(info.ArenasRanked.Next.Map) + "](" + info.ArenasRanked.Next.Asset + ")",
							},
						},
					},
				},
				{
					Type:    "countdown",
					Mode:    "day",
					EndTime: info.ArenasRanked.Current.End * 1000,
				},
			},
		},
	}
	apexMessageBytes, _ := json.Marshal(apexMessage)
	log.Println(string(apexMessageBytes))
	if updateFlags {
		return api.UpdateChannelMessage(post.ChannelMessageUpdate{
			MsgID:   targetID,
			Content: string(apexMessageBytes),
		})
	} else {
		return api.SendChannelMessage(post.ChannelMessageCreate{
			Type:     10,
			TargetID: targetID,
			Content:  string(apexMessageBytes),
			Quote:    quote,
		})
	}
}

func SendMessage(targetID, content, quote string, msgType int) (*axios.Response, error) {
	return api.SendChannelMessage(post.ChannelMessageCreate{
		Type:     msgType,
		TargetID: targetID,
		Content:  content,
		Quote:    quote,
	})
}

func SendTempMessage(targetID, content, quote, tempTargetID string, msgType int) (*axios.Response, error) {
	return api.SendChannelMessage(post.ChannelMessageCreate{
		Type:         msgType,
		TargetID:     targetID,
		Content:      content,
		Quote:        quote,
		TempTargetID: tempTargetID,
	})
}

func SendDirectMessage(targetID, content, quote string) (*axios.Response, error) {
	return api.SendDirectMessage(post.ChannelMessageCreate{
		Type:     9,
		TargetID: targetID,
		Content:  content,
		Quote:    quote,
	})
}
