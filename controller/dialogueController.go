package controller

import (
	"PushSystem/resp"
	"PushSystem/service"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"time"
)

type Group struct {
	Users map[uint]User
}

type User struct {
	Conn    *websocket.Conn
	Context string
}

type Dia struct {
	ID      uint
	Context string
}

var dialogueGroups = make(map[uint]map[uint]User)

func WebSocketConn(ctx *gin.Context) {
	g, e := strconv.Atoi(ctx.Query("GroupID"))
	if e != nil {
		zap.L().Debug(fmt.Sprint("test"))
		ctx.JSON(resp.SUCCESS, resp.NewInvalidResp())
		return
	}
	u, e := strconv.Atoi(ctx.Query("UserID"))
	userID := uint(u)
	groupID := uint(g)
	zap.L().Debug(fmt.Sprint(userID, groupID))
	if DialogueIdentify(userID, groupID) {
		upgrade := websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		}
		conn, e := upgrade.Upgrade(ctx.Writer, ctx.Request, nil)
		if e != nil {
			ctx.JSON(resp.SUCCESS, resp.NewErrorResp())
			return
		}
		defer conn.Close()
		if dialogueGroups[groupID] == nil {
			dialogueGroups[groupID] = make(map[uint]User)
		}
		dialogueGroups[groupID][userID] = User{
			Conn: conn,
		}
		for {
			time.Sleep(time.Second)
			mt, message, err := conn.ReadMessage()
			zap.L().Debug(fmt.Sprint(message))
			if err != nil {
				zap.L().Debug(fmt.Sprint(err.Error()))
				break
			}
			d := Dia{}
			_ = json.Unmarshal(message, &d)
			zap.L().Debug(fmt.Sprint(d))
			err = conn.WriteMessage(mt, message)
			if err != nil {
				zap.L().Debug(fmt.Sprint(err.Error()))
				break
			}
		}
		conn.SetCloseHandler(func(code int, text string) error {
			e := conn.Close()
			if e != nil {
				zap.L().Debug(fmt.Sprintln(e.Error()))
				return e
			}
			zap.L().Debug(fmt.Sprintln("delete user", dialogueGroups[groupID][userID]))
			delete(dialogueGroups[groupID], userID)
			return nil
		})

	}
}

func getAllDialogue() {

}

func DialogueIdentify(userID, groupID uint) bool {
	messageService := new(service.MessageService)
	belongToUser, e := messageService.IsBelongToUser(userID, groupID)
	if e != nil {
		zap.L().Debug("response1")
		return false
	} else if !belongToUser {
		zap.L().Debug("response2")
		return false
	} else {
		return true
	}
}
