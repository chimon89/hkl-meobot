package actions

import (
	"encoding/base64"
	"encoding/json"
	kookEvent "khl-meobot/model/kook/event"
	"khl-meobot/utils/api"
	KookCard "khl-meobot/utils/kook-card"
	"log"
	"net/url"
	"strings"
)

func ReqLoginSpotify(event kookEvent.Event) {
	resp, _ := SendMessage(event.D.TargetID, "请点击链接登录：[LOGIN](https://kookbot.dailypics.cn/spotify/req)", event.D.MsgID, 9)
	log.Println(string(resp.Data))
}

func GetSpotifySearchResultToKook(event kookEvent.Event, q string) {
	res, err := GetSpotifySearchResult(q)
	if err != nil {
		SendMessage(event.D.TargetID, "请求失败！", event.D.MsgID, 0)
	}
	modules := make([]KookCard.CardMessageModule, 0)
	for _, track := range res.Tracks.Items {
		artists := track.Artists[0].Name
		if len(track.Artists) == 1 {
			for _, artist := range track.Artists {
				if artist.Name != artists {
					artists += "/" + artist.Name
				}
			}
		}
		trackerReVal := TrackReturnVal{
			URI:       track.URI,
			TrackName: track.Name,
			Artists:   artists,
		}
		tr, _ := json.Marshal(trackerReVal)
		returnValRaw := base64.StdEncoding.EncodeToString(tr)
		modules = append(modules, KookCard.CardMessageModule{
			Type: "section",
			Text: &KookCard.TextModule{
				Type:    "kmarkdown",
				Content: "[" + track.Name + " - " + artists + "](" + track.ExternalUrls.Spotify + ") *" + track.Album.Name + "*",
			},
			Mode: "right",
			Accessory: &KookCard.ElementsModule{
				Type:  "button",
				Theme: "success",
				Click: "return-val",
				Value: "spotify " + returnValRaw,
				Text: &KookCard.TextModule{
					Type:    "plain-text",
					Content: "添加",
				},
			},
		})
	}
	searchResultMessage := []KookCard.CardMessage{
		{
			Type:    "card",
			Theme:   "secondary",
			Size:    "lg",
			Modules: modules,
		},
	}

	if len(modules) != 0 {
		msg, _ := json.Marshal(searchResultMessage)
		log.Println(string(msg))

		SendMessage(event.D.TargetID, string(msg), event.D.MsgID, 10)
	} else {
		SendMessage(event.D.TargetID, "未找到结果", event.D.MsgID, 0)
	}

}

func filter[T any](slice []T, f func(T) bool) []T {
	var n []T
	for _, e := range slice {
		if f(e) {
			n = append(n, e)
		}
	}
	return n
}

func AddItemToQueue(event kookEvent.Event) {
	dl, _ := GetSpotifyDevicesList()
	activeDevicesList := filter(dl.Devices, func(s struct {
		ID               string `json:"id"`
		IsActive         bool   `json:"is_active"`
		IsPrivateSession bool   `json:"is_private_session"`
		IsRestricted     bool   `json:"is_restricted"`
		Name             string `json:"name"`
		Type             string `json:"type"`
		VolumePercent    int    `json:"volume_percent"`
	}) bool {
		return s.IsActive == true
	})
	if len(activeDevicesList) != 0 {
		returnValBase64 := strings.Fields(event.D.Extra.Body.Value)[1]
		returnValBytes, _ := base64.StdEncoding.DecodeString(returnValBase64)
		var returnVal TrackReturnVal
		json.Unmarshal(returnValBytes, &returnVal)
		values := make(url.Values)
		values.Set("uri", returnVal.URI)
		resp, _ := api.SpotifyInstance().Post("/me/player/queue", nil, values)
		if resp.Status != 204 {
			resp, _ = SendMessage(event.D.Extra.Body.TargetID, string(resp.Data), "", 0)
			log.Println(string(resp.Data))
			return
		}
		resp, _ = SendMessage(event.D.Extra.Body.TargetID, returnVal.TrackName+" - "+returnVal.Artists+" 添加成功!", "", 0)
		log.Println(string(resp.Data))
	} else {
		resp, _ := SendTempMessage(event.D.Extra.Body.TargetID, "没有正在播放的设备", "", event.D.Extra.Body.UserID, 0)
		log.Println(string(resp.Data))
	}
}

func SpotifySkipToNext(event kookEvent.Event) {
	qu, _ := GetSpotifyQueue()
	api.SpotifyInstance().Post("/me/player/next", nil)
	if len(qu.Queue) > 1 {
		SendMessage(event.D.TargetID, "操作完毕\n> **当前播放：**"+qu.Queue[0].Name+" - "+qu.Queue[0].Artists[0].Name+"\n**下一首: **"+qu.Queue[1].Name+" - "+qu.Queue[1].Artists[0].Name, "", 9)
		return
	}
	SendMessage(event.D.TargetID, "操作完毕", "", 0)
}

func SpotifySkipToPrevious(event kookEvent.Event) {
	api.SpotifyInstance().Post("/me/player/previous", nil)
	SendMessage(event.D.TargetID, "操作完毕", "", 0)
}
