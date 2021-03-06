package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Syfaro/telegram-bot-api" 
)

var (
	numtoken  string
	reply     string
	replyfile string
	msg       tgbotapi.MessageConfig

	testfilename = "board-money.html" // для теста функции GetFile
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

	//	var ff bool
	//	ff = true
	//	ExistingPhotoFileID := "AgADAgADw6cxG4zHKAkr42N7RwEN3IFShCoABHQwXEtVks4EH2wBAAEC"

	var msgi tgbotapi.PhotoConfig
	var msgf tgbotapi.DocumentConfig

	for update := range updates {
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		log.Printf("update.Message.Text = %s", update.Message.Text)
		strcmd := strings.ToLower(strings.TrimSpace(update.Message.Text))
		switch strcmd {
		case "/start":
			reply = fmt.Sprintf(`Привет @%s! Я тут слежу за порядком. Веди себя хорошо.`, update.Message.From.UserName)
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, reply)
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		case "kurs":
			//			reply, replyfile = GetKursValuta(update.Message.From.UserName + "-")
			reply = fmt.Sprintf(reply)
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, reply)
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
			msgf = tgbotapi.NewDocumentUpload(update.Message.Chat.ID, replyfile)
			bot.Send(msgf)
		case "курс":
			//			reply, replyfile = GetKursValuta(update.Message.From.UserName + "-")
			reply = fmt.Sprintf(reply)
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, reply)
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
			msgf = tgbotapi.NewDocumentUpload(update.Message.Chat.ID, replyfile)
			bot.Send(msgf)
		case "file":
			msgf = tgbotapi.NewDocumentUpload(update.Message.Chat.ID, testfilename)
			bot.Send(msgf)
		case "img": // возвращает картинку в текущий чат
			//			ff = false
			msgi = tgbotapi.NewPhotoUpload(update.Message.Chat.ID, "image.jpg")
			msgi.Caption = "TestImg"
			bot.Send(msgi)
		case "!quit!":
			// сделать
			tgbotapi.NewMessage(update.Message.Chat.ID, reply)
			msg.ReplyToMessageID = update.Message.MessageID
			msg.Text = "Bye-Bye!!!"
			bot.Send(msg)
			return
		default:
			reply = update.Message.Text
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, reply)
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		}

	}
	fmt.Println("Stop kaefik bot...")
}

//------------------
// возвращает текст сообщения лучших предложений валют и возвращает путь и имя файла доски валют
//func GetKursValuta(todir string) (string, string) {
//	resstr := ""
//	//	if todir != "" {
//	//		if _, err := os.Stat(todir); os.IsNotExist(err) {
//	//			os.Mkdir(todir, 0666)
//	//		}
//	//	}
//	//	resfile := todir + string(os.PathSeparator) + "board-money.html"
//	resfile := todir + "board-money.html"
//	str, res := parsebank.RunGoBoardValutaHtml()

//	pick.Savestrtofile(resfile, str) // не сохраняет в папке

//	usdkurspokupka := res[0]
//	usdkursprodaja := res[1]
//	eurkurspokupka := res[2]
//	eurkursprodaja := res[3]

//	resstr += " Лучшая покупка USD: " + usdkurspokupka.Namebank + " - " + parsebank.FloatToString(usdkurspokupka.Pokupka) + " \n Лучшая продажа USD: " + usdkursprodaja.Namebank + " - " + parsebank.FloatToString(usdkursprodaja.Prodaja) + "\n Лучшая покупка EUR: " + eurkurspokupka.Namebank + " - " + parsebank.FloatToString(eurkurspokupka.Pokupka) + "\n Лучшая продажа EUR: " + eurkursprodaja.Namebank + " - " + parsebank.FloatToString(eurkursprodaja.Prodaja)

//	return resstr, resfile
//}
