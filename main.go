package main
import (
	"github.com/Syfaro/telegram-bot-api"
	"log"
	"os"
	"encoding/json"
	"fmt"
)

type Config struct {
	Applications []Application
}

type Application struct {
	Name string
	Db   Db
	BotAPIToken string
}

type Db struct {
	Host   string
	User   string
	Pass   string
	Dbname string
}


func main() {

	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	config := new(Config)
	err := decoder.Decode(&config)
	if err != nil {
		// handle it
	}

	fmt.Println(config.Applications[0].BotAPIToken)

	// подключаемся к боту с помощью токена
	bot, err := tgbotapi.NewBotAPI("token")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// инициализируем канал, куда будут прилетать обновления от API
	var ucfg tgbotapi.UpdateConfig = tgbotapi.NewUpdate(0)
	ucfg.Timeout = 60
	upd, _ := bot.GetUpdatesChan(ucfg)
	// читаем обновления из канала
	for {
		select {
		case update := <-upd:
			// Пользователь, который написал боту
			UserName := update.Message.From.UserName

			// ID чата/диалога.
			// Может быть идентификатором как чата с пользователем
			// (тогда он равен UserID) так и публичного чата/канала
			ChatID := update.Message.Chat.ID

			// Текст сообщения
			Text := update.Message.Text

			log.Printf("[%s] %d %s", UserName, ChatID, Text)

			// Ответим пользователю его же сообщением
			reply := Text
			// Созадаем сообщение
			msg := tgbotapi.NewMessage(ChatID, reply)
			// и отправляем его
	var ucfg tgbotapi.UpdateConfig = tgbotapi.NewUpdate(0)
	ucfg.Timeout = 60
	upd, _ := bot.GetUpdatesChan(ucfg)
	for {
		select {
		case update := <-upd:
			UserName := update.Message.From.UserName
			ChatID := update.Message.Chat.ID
			Text := update.Message.Text
			log.Printf("[%s] %d %s", UserName, ChatID, Text)
			reply := Text
			msg := tgbotapi.NewMessage(ChatID, reply)
			bot.Send(msg)
		}

	}
	}
	}
}
