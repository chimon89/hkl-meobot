package actions

import (
	"github.com/gin-gonic/gin"
	kookEvent "khl-meobot/model/kook/event"
	"khl-meobot/model/kook/response"
	"reflect"
	"strings"
)

func KookEntry(c *gin.Context) {
	res, _ := c.Get("Event")
	event := res.(kookEvent.Event)

	if event.D.Challenge != "" {
		c.JSON(200, response.ChallengeResponse{Challenge: event.D.Challenge})
		return
	}

	c.JSON(200, gin.H{"ok": true})

	if kc := CommandParse(event.D.Content); !event.D.Extra.Author.Bot && kc.IsCommand {
		switch kc.Args[0] {
		case "/mie":
			go SendMessage(event.D.TargetID, "mie~", "", 0)
			break
		case "/apexmap":
			go GetApexMapRotate(event.D.TargetID, event.D.MsgID, false)
			break
		case "/mc":
			switch len(kc.Args) {
			case 1:
				go GetMcServerStatus(event.D.TargetID, event.D.MsgID, "", "")
				break
			case 2:
				go GetMcServerStatus(event.D.TargetID, event.D.MsgID, kc.Args[1], "")
				break
			default:
				go GetMcServerStatus(event.D.TargetID, event.D.MsgID, kc.Args[1], kc.Args[2])
				break
			}
		case "/spm_login":
			go ReqLoginSpotify(event)
			break
		case "/sps":
			go GetSpotifySearchResultToKook(event, kc.SingleArg)
			break
		case "/spnext":
			go SpotifySkipToNext(event)
			break
		case "/spprev":
			go SpotifySkipToPrevious(event)
			break
		default:
			go SendMessage(event.D.TargetID, "没有这条指令", event.D.MsgID, 0)
			break
		}
	}

	if reflect.TypeOf(event.D.Extra.Type).String() == "string" && event.D.Extra.Type.(string) == "message_btn_click" && event.D.Extra.Body.Value == "refresh-apexmap" {
		go GetApexMapRotate(event.D.Extra.Body.MsgID, "", true)
	}

	if reflect.TypeOf(event.D.Extra.Type).String() == "string" && event.D.Extra.Type.(string) == "message_btn_click" && strings.Contains(event.D.Extra.Body.Value, "spotify") {
		go AddItemToQueue(event)
	}

}
