package websocket

import (
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// 全局连接池（用于存储每个直播间的所有连接）
var roomConnMap = make(map[string]map[*websocket.Conn]bool)

var roomConnMapLock = &sync.Mutex{}

func WebsHandler(c *gin.Context) {
	roomID := c.Param("roomID")
	//1.升级为webSocket协议
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true }, // 跨域
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("Upgrade error: %v", err)
		return
	}
	defer conn.Close()

	//2.加入连接池
	roomConnMapLock.Lock()
	if _, ok := roomConnMap[roomID]; !ok {
		roomConnMap[roomID] = make(map[*websocket.Conn]bool)
	}
	roomConnMap[roomID][conn] = true
	roomConnMapLock.Unlock()

	defer func() {
		roomConnMapLock.Lock()
		delete(roomConnMap[roomID], conn)
		roomConnMapLock.Unlock()
	}()

	for {
		// 读消息
		mt, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("ReadMessage error: %v", err)
			break
		}
		// 广播给房间内所有连接
		roomConnMapLock.Lock()
		for client := range roomConnMap[roomID] {
			_ = client.WriteMessage(mt, message)
		}
		roomConnMapLock.Unlock()
	}
}
