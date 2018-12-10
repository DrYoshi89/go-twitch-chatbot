package App

import (
	"net"
)

/**
 * [Name]			Connection info
 */
type IrcBase struct {
	Server string
	Port string
	Nick string
	User string
	Channel string
	Pass string
	Conn net.Conn
}

/**
 * [Name] 			ROOMSTATE
 * [Reference] 		https://dev.twitch.tv/docs/irc/tags/#roomstate-twitch-tags
 * [Example]		> @emote-only=<emote only>;room-id=<channel ID>;r9k=<r9k>;slow=<slow>;:tmi.twitch.tv ROOMSTATE #chatrooms:<channel ID>:<room UUID>:<message>
 * [Comment] 		When a user joins a channel or a room setting is changed.
 * 					For a join, the message contains all chat-room settings. For changes, only the relevant tag is sent.
 */
type Room struct {
	Name string
	R9k int
	Slow int
	BroadcasterLang string
	FollowersOnly int
	EmoteOnly int
	SubsOnly int
}

/**
 * [Name] 			PRIVMSG
 * [Reference]		https://dev.twitch.tv/docs/irc/chat-rooms/#privmsg-twitch-chat-rooms
 * [Example]		> :<user>!<user>@<user>.tmi.twitch.tv PRIVMSG #chatrooms:<channel ID>:<room UUID> :This is a sample message
 * [Comment] 		Send a message to a chat room.
 */
type Message struct {
	Raw []chan string
	MsgID string

	UserID string
	UserType string
	User string
	Mod string
	Subscriber string
	Turbo string

	Room string
	Command string
	Text string
	badge string
	bits string
}

/**
 * [Name] 			NOTICE
 * [Reference]		hhttps://dev.twitch.tv/docs/irc/chat-rooms/#notice-twitch-chat-rooms
 * [Example]		> @msg-id=<msg id>:tmi.twitch.tv NOTICE #chatrooms:<channel ID>:<room UUID> :<message>
 * [Comment] 		General notices from the server.
 */
type UserNotice struct {
	Raw []chan string
	MsgID string
	Tmi string

	MsgType string
	SustemMsg string
	Room string

	Command string
	Text string

	UserID string
	UserType string
	User string
	Mod string
	Subscriber string
	Turbo string
	badge string
	bits string
}
