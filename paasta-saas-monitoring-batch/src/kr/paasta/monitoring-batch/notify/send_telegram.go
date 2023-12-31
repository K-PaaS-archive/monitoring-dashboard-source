package notify

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os"
	"saas-monitoring-batch/util"
)

var config map[string]string

func init() {
	var err error
	config, err = util.ReadConfig("config.ini")
	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}
}

func SendChatBot(serviceType string, receivers []int64, message string, token string) {
	if len(receivers) > 0 {
		//token := config["telegram." + serviceType +  ".token"]
		bot, err := tgbotapi.NewBotAPI(token)
		if err != nil {
			fmt.Println("Failed to get telegram client connection! :", err)
			return
		}
		bot.Debug = true

		for _, receiver := range receivers {
			msg := tgbotapi.NewMessage(receiver, message)
			botMsg, botErr := bot.Send(msg)
			fmt.Printf(">>>>> botMsg=[%v], botErr[%v]\n", botMsg, botErr)
		}

	}

}
