package main

import (
	"gopkg.in/telegram-bot-api.v4"
	//"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"encoding/json"
	"fmt"
	"os"
)



type BotTelegram struct {
	telegramBOT  	string `json:"TelegramBotToken"`
}


type Config struct {
	TelegramBotToken string
}

var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Заказать обрытный звонок"),
		tgbotapi.NewKeyboardButton("Оставить заявку в тех.поддержку"),
		tgbotapi.NewKeyboardButton("Новый проект"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("FAQ"),
		tgbotapi.NewKeyboardButton("Оставить контакты для обатной связи"),
		tgbotapi.NewKeyboardButton("Перейти на сайт компании"),
	),
)
func LoadConfiguration(file string) string {
	var config Config
	fmt.Println(file)
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config.TelegramBotToken
}


func main() {
	// используя токен создаем новый инстанс бота
	telegramBotToken := LoadConfiguration("./config.json")
	bot, err := tgbotapi.NewBotAPI(telegramBotToken)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)
	// u - структура с конфигом для получения апдейтов
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	// используя конфиг u создаем канал в который будут прилетать новые сообщения
	updates, err := bot.GetUpdatesChan(u)
	// в канал updates прилетают структуры типа Update
	// вычитываем их и обрабатываем
	for update := range updates {
		// универсальный ответ на любое сообщение
		reply := ""
		if update.Message == nil {
			continue
			}
		// логируем от кого какое сообщение пришло
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)


		///Анализируем Меню
		switch update.Message.Text {
		case "FAQ" :
			reply = "IP-телефония - это телефонная связь через интернет, по протоколу IP. Под IP-телефонией подразумевается набор коммуникационных протоколов, VoIP оборудования, программного обеспечения, технологий и методов, обеспечивающих традиционные для телефонной связи функции: набор номера, дозвон и двустороннее голосовое общение, а также видеообщение по сети Интернет или любым другим IP-сетям. "
		case "Заказать обрытный звонок":
			// Если нет -> запрашиваем его


		}


		// комманда - сообщение, начинающееся с "/"
		key := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Command())
		switch update.Message.Command() {
		case "start":
			reply = "Привет! Я телеграм-бот компании HuskyTech"
			key.ReplyMarkup = numericKeyboard
			bot.Send(key)
		case "close":
			key.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			bot.Send(key)
		}
		// создаем ответное сообщение
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
		// отправляем
		bot.Send(msg)
	}
}
