package configs

type Configs struct {
	Host string
	Port string
	Nick string
	Channel string
	Auth string
}

func Get() *Configs {
	return &Configs{
		Host: "irc.chat.twitch.tv", // Host (default)
		Port: "6667", // Porta (default)
		Nick: "nick-here", // Nick do Bot/ContaNormal
		Channel: "channel-here", // Canal para se auto conectar
		Auth: "oauth:token-here", // Token do Bot/ContaNormal
	}
}