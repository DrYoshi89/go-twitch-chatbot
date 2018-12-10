package App

// ircApp
import (
	"net"
	"net/textproto"
	"../tools"
)

type App struct {
	ConnInfo *IrcBase
	Running chan bool
	Chans []chan string
	Conn net.Conn
	Buffer *textproto.Reader

	CN struct {
		PingLoop chan bool
	}
}

func NewApp(server string, port string, nick string, channel string, pass string) (app *App) {
	app = &App{
		ConnInfo: &IrcBase {
			Server: server,
			Port: port,
			Nick: nick,
			Channel: channel,
			Pass: pass,
		},
		Chans: []chan string{
			make(chan string, 100),
		},
	}
	app.CN.PingLoop = make(chan bool)
	app.Running = make(chan bool)

	return
}

func (app *App) Run() (err error) {
	tools.Defer()

	err = app.Connect()
	app.PingLoop()
	return
}
