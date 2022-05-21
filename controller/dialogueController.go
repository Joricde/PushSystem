package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"net/http"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WebsocketConnect(c *gin.Context) {
	zap.L().Debug("T")
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		zap.L().Debug(err.Error())
		return
	}
	defer ws.Close()
	for {
		mt, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		err = ws.WriteMessage(mt, message)
		if err != nil {
			break
		}
	}
}
