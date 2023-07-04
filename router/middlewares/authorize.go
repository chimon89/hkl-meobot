package middlewares

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"khl-meobot/model/errors"
	kookEvent "khl-meobot/model/kook/event"
	"log"
)

func CheckVerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var event kookEvent.Event
		if err := c.BindJSON(&event); err != nil {
			log.Println(err.Error())
			c.Abort()
			c.JSON(400, errors.ErrorResponseDev{
				Ok:           false,
				ErrorMessage: err.Error(),
				Data:         event,
			})
			return
		}
		bytes, _ := json.Marshal(event)
		log.Println(string(bytes))
		if event.D.VerifyToken != "u8iu3nvGcYlr-uy6" {
			c.Abort()
			c.JSON(401, errors.ErrorResponse{
				Ok:           false,
				ErrorMessage: "Verify Token Error",
			})
			return
		}
		c.Set("Event", event)
		c.Next()
	}
}
