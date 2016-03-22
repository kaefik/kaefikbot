package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Syfaro/telegram-bot-api"
)

//// чтение файла с именем namefи возвращение содержимое файла, иначе текст ошибки
func readfiletxt(namef string) string {
	file, err := os.Open(namef)
	if err != nil {
		return "handle the error here"
	}
	defer file.Close()
	// get the file size
	stat, err := file.Stat()
	if err != nil {
		return "error here"
	}
	// read the file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return "error here"
	}
	return string(bs)
}

//получение токена из конфиг файла
func GetToken(namef string) string {
	fstr := readfiletxt(namef)
	res := (strings.Split(fstr, ";"))[0]
	return res
}
func main() {
	var numtoken string
	fmt.Println("Start kaefik bot...")

	numtoken = GetToken("cfg.token")

	bot, err := tgbotapi.NewBotAPI(numtoken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
	fmt.Println("Stop kaefik bot...")
}
