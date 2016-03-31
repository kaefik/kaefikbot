package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"go-board-money/parsebank"

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

var (
	numtoken string
	reply    string
	msg      tgbotapi.MessageConfig
)

func GetKursValuta() string {
	resstr := ""
	_, res := parsebank.RunGoBoardValutaHtml()
	//	pick.Savestrtofile(todir+"board-money.html", str)

	usdkurspokupka := res[0]
	usdkursprodaja := res[1]
	eurkurspokupka := res[2]
	eurkursprodaja := res[3]
	//	fmt.Println(res)

	resstr += " Лучшая покупка USD: " + usdkurspokupka.Namebank + " - " + parsebank.FloatToString(usdkurspokupka.Pokupka) + " \n Лучшая продажа USD: " + usdkursprodaja.Namebank + " - " + parsebank.FloatToString(usdkursprodaja.Prodaja) + "\n Лучшая покупка EUR: " + eurkurspokupka.Namebank + " - " + parsebank.FloatToString(eurkurspokupka.Pokupka) + "\n Лучшая продажа EUR: " + eurkursprodaja.Namebank + " - " + parsebank.FloatToString(eurkursprodaja.Prodaja)
	//

	//	usdkursprodaja
	//	eurkurspokupka
	//	eurkursprodaja

	return resstr
}

func main() {

	fmt.Println("Start kaefik bot...")

	//	linkbanks := parsebank.Initlinksbank() // ссылки на банки

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
		log.Printf("update.Message.Text = %s", update.Message.Text)
		strcmd := strings.ToLower(strings.TrimSpace(update.Message.Text))
		switch strcmd {
		case "/start":
			reply = fmt.Sprintf(`Привет @%s! Я тут слежу за порядком. Веди себя хорошо.`, update.Message.From.UserName)
		case "kurs":
			reply = fmt.Sprintf(GetKursValuta())
		case "курс":
			reply = fmt.Sprintf(GetKursValuta())
		default:
			reply = update.Message.Text
		}

		msg = tgbotapi.NewMessage(update.Message.Chat.ID, reply)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
	fmt.Println("Stop kaefik bot...")
}
