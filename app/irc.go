package App

import (
	"net"
	"net/textproto"
	"bufio"
	"log"
	"fmt"
	"time"
)

func (app *App) Connect() (err error) {
	app.Conn, err = net.Dial("tcp", app.ConnInfo.Server + ":" + app.ConnInfo.Port)
	if err != nil{
		log.Fatal("Unable to connect to IRC server ", err)
		app.Conn.Close()
		return
	}

	reader := bufio.NewReader(app.Conn)
	app.Buffer = textproto.NewReader(reader)

	log.Printf("Connected to IRC server %s (%s) \n", app.ConnInfo.Server, app.Conn.RemoteAddr())
	fmt.Fprintf(app.Conn, "PASS %s\r\n", app.ConnInfo.Pass)
	fmt.Fprintf(app.Conn, "NICK %s\r\n", app.ConnInfo.Nick)

	// fmt.Fprintf(app.Conn, "CAP REQ :twitch.tv/membership CAP REQ :twitch.tv/tags CAP REQ :twitch.tv/commands\r\n")

	// fmt.Fprintf(app.Conn, "JOIN #%s\r\n", app.ConnInfo.Channel)

	go func(app *App) {
		app.Send("CAP REQ :twitch.tv/membership")
		app.Send("CAP REQ :twitch.tv/tags")
		app.Send("CAP REQ :twitch.tv/commands")
		time.Sleep(1 * time.Second)
		app.Send(fmt.Sprintf("JOIN #%s\r\n", app.ConnInfo.Channel))
	}(app)

	return
}

func (app *App) Maker() {
}

func (app *App) Send(command string) {
	fmt.Fprintf(app.Conn, "%s\r\n", command)
}

func (app *App) Msg(channel string, msg string) {
	fmt.Fprintf(app.Conn, "PRIVMSG #%s :%s\r\n", channel, msg)
}

func (app *App) Join(channel string) {
	fmt.Fprintf(app.Conn, "JOIN #%s\r\n", channel)
}

func (app *App) Leave(channel string) {
	fmt.Fprintf(app.Conn, "PART #%s\r\n", channel)
}

func (app *App) PingLoop() {
	fmt.Println("PingLoop started")
	go func() {
		for {
			// time.Sleep(time.Second * 120)
			time.Sleep(time.Second * 1)
			select {
				case <-app.CN.PingLoop:
					fmt.Println("Ping stoped")
					return
				default:
					// app.Send("PING :tmi.twitch.tv")
			}
		}
	}()
}