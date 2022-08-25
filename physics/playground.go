package physics

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type PhysicsPlayground struct {
	r     *gin.Engine
	world *World
}

func NewPhysicsPlayground(world *World) *PhysicsPlayground {
	p := PhysicsPlayground{
		r:     gin.Default(),
		world: world,
	}

	p.r.StaticFile("/", "./physics/playground.html")

	var upgrader = websocket.Upgrader{}

	p.r.GET("/websocket", func(ctx *gin.Context) {
		// Upgrade to websocket protocol
		ws, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer ws.Close()

		for {
			msgType, msg, err := ws.ReadMessage()
			if err != nil {
				fmt.Println(err)
				return
			}

			packet := SocketPacket{}
			if err := json.Unmarshal(msg, &packet); err != nil {
				fmt.Println(err)
				return
			}

			if err := msgHandler(&p, msgType, packet, ws); err != nil {
				fmt.Println(err)
				return
			}
		}
	})

	return &p
}

// Starts the physics playground at the given port
func (p *PhysicsPlayground) Run(port int) {
	p.r.Run(":" + fmt.Sprint(port))
}

// A response sent by the msg handler
type SocketPacket struct {
	Name string `json:"name"`
	Data any    `json:"data"`
}

func sendState(world *World, msgType int, ws *websocket.Conn) error {
	bodies := world.Bodies()
	r := SocketPacket{
		Name: "state",
		Data: bodies,
	}
	jsonStr, err := json.Marshal(r)
	if err != nil {
		return fmt.Errorf("msg handler: %s", err)
	}
	return ws.WriteMessage(msgType, jsonStr)
}

func msgHandler(
	playground *PhysicsPlayground, msgType int, packet SocketPacket, ws *websocket.Conn,
) error {
	switch packet.Name {
	case "state":
		return sendState(playground.world, msgType, ws)
	case "step":
		playground.world.Step(300)
		return sendState(playground.world, msgType, ws)
	}
	return errors.New("msg handler: invalid msg: " + packet.Name)
}
