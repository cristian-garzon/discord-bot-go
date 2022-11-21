package utils

const PrefixPath = "C! "

type Command struct {
	Prefix      []string
	Title       string
	Description string
}

var Commands = []Command{
	{
		Title:       "ping",
		Prefix:      Ping,
		Description: "tu haces ping, yo hago pong :3",
	}, {
		Title:       "avatar",
		Prefix:      Avatar,
		Description: "puedo mostrar tu avatar",
	},
	{
		Title:       "help",
		Prefix:      Help,
		Description: "te muestro todo lo que puedo hacer!!",
	},
}

var (
	Avatar = []string{"avatar"}
	Ping   = []string{"ping"}
	Help   = []string{"help", "h"}
)
