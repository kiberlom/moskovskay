package telegram

import "github.com/kiberlom/moskovskay/internal/config"

type BotWorker interface {
	Send(string) error
}

func NewBot(c *config.Config) BotWorker {

	return &botWork{
		config: c,
	}

	// in := make(chan string, 2)
	// go bot(in)

}

// func bot(in <-chan string) {
// 	for {
// 		select {
// 		case v := <-in:
// 			fmt.Println(v)
// 		}
// 	}
// }
