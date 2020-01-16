// https://play.golang.org/p/1U5ImsDAd0n

/*
    TODO:
       Загрузка и сохранения файла
       Кнопки в диалоге окна 
*/

package main

import (
	"gopkg.in/telegram-bot-api.v4"
	//"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const (
	// TestToken               = "153667468:AAHlSHlMqSt1f_uFmVRJbm5gntu2HI4WW8I"
	// ChatID                  = 76918703
	// SupergroupChatID        = -1001120141283
	// ReplyToMessageID        = 35
	ExistingPhotoFileID     = "AgADAgADw6cxG4zHKAkr42N7RwEN3IFShCoABHQwXEtVks4EH2wBAAEC"
	ExistingDocumentFileID  = "BQADAgADOQADjMcoCcioX1GrDvp3Ag"
	ExistingAudioFileID     = "BQADAgADRgADjMcoCdXg3lSIN49lAg"
	ExistingVoiceFileID     = "AwADAgADWQADjMcoCeul6r_q52IyAg"
	ExistingVideoFileID     = "BAADAgADZgADjMcoCav432kYe0FRAg"
	ExistingVideoNoteFileID = "DQADAgADdQAD70cQSUK41dLsRMqfAg"
	ExistingStickerFileID   = "BQADAgADcwADjMcoCbdl-6eB--YPAg"
)


type BotTelegram struct {
	telegramBOT  	string           `json:"TelegramBotToken"`
}

type Config struct {
	TelegramBotToken string
}



// https://tlgrm.ru/docs/bots/api#inlinekeyboardbutton
// Button description
var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Заказать обратный звонок"),
		tgbotapi.NewKeyboardButton("Оставить заявку в тех.поддержку"),
		tgbotapi.NewKeyboardButton("Координаты компании"),
		tgbotapi.NewKeyboardButton("Docs"),
	),

	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("FAQ"),
		tgbotapi.NewKeyboardButton("Оставить контакты для обатной связи"),
		tgbotapi.NewKeyboardButton("Перейти на сайт компании"),
	),
    
    tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("☀️ Фото дня"),
		tgbotapi.NewKeyboardButton("Музыка"),
		tgbotapi.NewKeyboardButton("Речь"),
	),

     tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Новости"),
		tgbotapi.NewKeyboardButton("Голос"),
		tgbotapi.NewKeyboardButton("Контакт"),
		tgbotapi.NewKeyboardButton("Венуе"),

		
		
	),


tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Test"),
	),
)



// Main
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

		ChatID:=update.Message.Chat.ID

		// универсальный ответ на любое сообщение
		reply := ""

		if update.Message == nil {
			continue
		}

		// логируем от кого какое сообщение пришло
		log.Printf("[%s] %s %v", update.Message.From.UserName,               // Пользователь
			                     update.Message.Text,                        // Сообщение 
			                     update.Message.MessageID ,                  // Номер сообщения
			                 )

        MessageID:=update.Message.MessageID  


        // Ответы на ввод с клавиатуры
        switch update.Message.Text {
		       case "Артур":		reply = "Привет  друг."
		       case "Анна":			reply = "Я заню тебя ты возглавляешь отдел разработки."
		       case "Роман":		reply = "Я заню тебя ты администратор на Барсе."
		       case "Рома":	   	    reply = "Я уже занкомился с тобой !! Ты администратор на Барсе. Рад тебя видеть снова."
		       case "Релиз":		reply = "На сегодняшнюю дату 68 релиз"
		       case "Счет":			reply = "Ваш счет в банке 246777 открыт 01.02.2018"
		       case "Остаток":		reply = "На вашем счету : 1000 грн. и $200"
		       case "Дата":			reply = "17.01.2018"
		       case "11111":		reply = "Заявка в работе и ждет ответа от разработчика"
		       case "00000":		reply = "Заявка не обработана и ждет утверждение от банка"
		       case "22222":		reply = "Ожидание утверждения от отдела администрации"
		       case "Ярослав":		reply = "Привет Ярик ! Рад видеть службу поддержки у себя в гостях."
		       case "Ярик":			reply = "Привет Ярик ! Повторно рад тебя видеть!! Привет службе поддержки!"
		       case "Евгений":		reply = "Привет Женя ! Рад видеть руководителя поддержки у себя в гостях!"
		       case "Андрей":		reply = "Привет  ! Фамилия твоя не ? "
		       case "Документ":		reply = `Документы   "Оновлення від  встановлено на тестові полігони. Проводиться тестування. Також отримано оновлення в Release_2.0.38, для тестування в рамках релізу Станом на 18.01.2018 проведено попереднє тестування на тестовому полігоні BARSPROJ. Виявлені помилки зі сторони АБС: 1) після відкриття карткового рахунку 2625, з АБС до Електронного архіву предається три однакових повідомлення по угоді ДКБО та два однакових повідомлення по рахунку БПК 2) Після резервування рахунку 2600 по угоді ДБО, рахунок не відображається в меню ""Зарезервовані рахунки"" та по зарезервованому рахунку не передається інформація з АБС в Електронний архів 3) Є можливість акцептувати зарезервований рахунок під роллю ""фронт-офіс"". Помилки передані розробнику. 4) Помилка при реєстрації угоди зарплатного проекту Очікуємо встановлення релізу 2.0.38.1 на тестовий полігон Барсу для проведення комплексного тестування."`
		       case "Оля": 			reply = `Привет Оля ! Я тебя знаю!  !!!! `
		       
		}



	// Анализ Меню кнопочек
	switch update.Message.Text {
		
        case "Test":
        	reply = "Привет всем участникам соревнований!."

		case "Test dont work":     

	          markup := tgbotapi.InlineKeyboardMarkup{
		            InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
			                                   []tgbotapi.InlineKeyboardButton{
				                                        tgbotapi.InlineKeyboardButton{Text: "Семен Семенович"},
			             },
		            },
	                  }

	                   edit := tgbotapi.NewEditMessageReplyMarkup(ChatID, MessageID, markup)         
		               //edit.BaseEdit.MessageID != edit.ReplyToMessageID 
	                   bot.Send(edit)	


               


		case "FAQ" :
			  reply = "IP-телефония - это телефонная связь через интернет, по протоколу IP. Под IP-телефонией подразумевается набор коммуникационных протоколов, VoIP оборудования, программного обеспечения, технологий и методов, обеспечивающих традиционные для телефонной связи функции: набор номера, дозвон и двустороннее голосовое общение, а также видеообщение по сети Интернет или любым другим IP-сетям. "
		

		case "Заказать обратный звонок":
			  reply = "Ждите вам перезвонят в течении 5 мин."

		case "Координаты компании" 	  :
             bot.Send(tgbotapi.NewLocation(ChatID, 50.384616, 30.625453)) 

		case "Новости":
			reply = "<h2>News</h2>"

        case "Венуе" :
        	venue := tgbotapi.NewVenue(ChatID, "A Test Location", "123 Test Street", 40, 40)
            bot.Send(venue)	

		case "Контакт":
           сontact := tgbotapi.NewContact(ChatID, "+38 097 09 98", "Геда Абрамовна")
           bot.Send(сontact)


        case "Голос":
        	  msg := tgbotapi.NewVoiceUpload(ChatID, "tests/voice.ogg")
	          msg.Duration = 10
	          bot.Send(msg)


        case "Фото недели":
               bot.Send(tgbotapi.NewDocumentUpload(ChatID, "tests/image.jpg"))


		case "☀️ Фото дня":
             // msg := tgbotapi.NewDocumentUpload(ChatID, "tests/image.jpg")
             // msg.ReplyToMessageID = ReplyToMessageID


             f, _ := os.Open("tests/image.jpg")
	         reader := tgbotapi.FileReader{Name: "image.jpg", Reader: f, Size: -1}

	         msg := tgbotapi.NewPhotoUpload(ChatID, reader)
	         msg.Caption          = "Test"
	         msg.ReplyToMessageID = update.Message.MessageID

	         _,err:=bot.Send(msg)
             if err != nil {
	              fmt.Println(err)
	           }


        // Отправка в телеграмм зараннее сохранненого стикера 
        case "Test fish": 
               bot.Send(tgbotapi.NewMessage(ChatID,        "Здрастуйте я рыба !!! "))
               bot.Send(tgbotapi.NewDocumentUpload(ChatID, "tests/fish.tgs"))

        // Не понятно зачем нужен
        case "New Webhook":        
				time.Sleep(time.Second * 2)
				bot.RemoveWebhook()

				// wh := tgbotapi.NewWebhook("https://example.com/tgbotapi-test/rtest-1111111111")
					wh := tgbotapi.NewWebhook("http://arttech.inf.ua")
				bot.SetWebhook(wh)
				
				info,_ := bot.GetWebhookInfo()
				
				if info.LastErrorDate != 0 {
					fmt.Printf("[Telegram callback failed] %s", info.LastErrorMessage)
				}
				bot.RemoveWebhook()

        // Просомтр токена 
        case "Get Bot token":        
                  fmt.Println(bot.Token, bot.Self.UserName)
        
        // Просомтр ИД сообщения 
        case "Get Message ID":        
             msg, _:= bot.Send(tgbotapi.NewMessage(ChatID, "Testing editing."))
             fmt.Println(msg.MessageID)

        // Очистка конопок с экрана     
        case "Clear all button":        
                      msg            := tgbotapi.NewMessage(ChatID, "By keyboard")
                      msg.ReplyMarkup = tgbotapi.ReplyKeyboardRemove{RemoveKeyboard:true,	Selective:false}
                      bot.Send(msg) 	

        // Простое сообщение     
        case "Test soobs":        
             msg := tgbotapi.NewMessage(ChatID, "Test soob")
             bot.Send(msg) 	

        // Загрузка видео 
        case "Video note":        
             msg := tgbotapi.NewVideoNoteUpload(ChatID, 240, "tests/videonote.mp4")
	         msg.Duration = 10
	         bot.Send(msg)

        // Загрузка текстового документа в телеграмм
        case "Load txt file":        
              msg := tgbotapi.NewDocumentUpload(ChatID, "tests/news.txt")
	          bot.Send(msg)

        // Очистка экрана от кнопок
        case "Clear button":        
             u := tgbotapi.NewUpdate(0)
	         bot.GetUpdates(u)


        case "Test replay":
             msg := tgbotapi.NewMessage(ChatID, "Test soob")
             msg.ReplyToMessageID = update.Message.MessageID             	
             bot.Send(msg) 	

        // Линк на сайт
        case "Link to site":
        	msg := tgbotapi.NewMessage(ChatID, "http://arttech.inf.ua")
	        bot.Send(msg)

        // Форматирование теста сообщения при помощи маркдауна
        case "Test markdown":
        	msg := tgbotapi.NewMessage(ChatID, "A test *message* from the test library in telegram-bot-api ")
            msg.ParseMode = "Markdown"
	        bot.Send(msg)

        // Сообщение в формате маркдаун
        case "Test notify test":
        	msg          := tgbotapi.NewMessage(ChatID, "A test message from the test library in telegram-bot-api")
	        msg.ParseMode = "Markdown"
	        message, _   := bot.Send(msg)

	        pinChatMessageConfig := tgbotapi.PinChatMessageConfig{
		                ChatID:              message.Chat.ID,
		                MessageID:           message.MessageID,
		                DisableNotification: false,
	        }
	        bot.PinChatMessage(pinChatMessageConfig)


        // Удаление сообщения по его ИД 
        case  "Delete message" :
			msg := tgbotapi.NewMessage(ChatID, "A test message from the test library in telegram-bot-api")
			msg.ParseMode = "markdown"
			message, _ := bot.Send(msg)

			deleteMessageConfig := tgbotapi.DeleteMessageConfig{
					ChatID:    message.Chat.ID,
					MessageID: message.MessageID,
			}
			 bot.DeleteMessage(deleteMessageConfig)
	        


        // Pdf загрузка но когда большой =65.5 Мб файл программа вываливается !
        case  "Docs" :
              msg := tgbotapi.NewDocumentUpload(ChatID, "tests/shild.pdf")
	          bot.Send(msg)


        // Загрузка фото с сайта
        case  "Arttech foto" :
                // Показывает фото с сайта
				cfg := tgbotapi.NewMediaGroup(ChatID, []interface{}{
					tgbotapi.NewInputMediaPhoto("http://arttech.inf.ua/pic/idx/sport.jpeg"),
				})
                bot.Send(cfg)

        
        // Показывает несколько фото с сайта
        case  "Test same foto" :
                
				cfg := tgbotapi.NewMediaGroup(ChatID, []interface{}{
					tgbotapi.NewInputMediaPhoto("https://i.imgur.com/unQLJIb.jpg"),
					tgbotapi.NewInputMediaPhoto("https://i.imgur.com/J5qweNZ.jpg"),
					tgbotapi.NewInputMediaVideo("https://i.imgur.com/F6RmI24.mp4"),
				})
                bot.Send(cfg)

       
        case  "Recall" :
            msg, _ := bot.Send(tgbotapi.NewMessage(ChatID, "Testing editing."))
	        edit := tgbotapi.EditMessageTextConfig{
	      	            BaseEdit:  tgbotapi.BaseEdit {
			            ChatID:    ChatID,
			            MessageID: msg.MessageID,
		     },
		         Text: "Обратный звонок заказан.",
	        }

	        bot.Send(edit)


              //    bot.Send(tgbotapi.NewChatAction(ChatID, tgbotapi.ChatTyping))
         	    // // агрузка фото как стикера без рамок и внизу подписи
         	    // msg = tgbotapi.NewStickerUpload(ChatID, "tests/image.jpg")
         	    // bot.Send(msg)  

        // Показывает на карте точку
        case  "Локация" :
                venue := tgbotapi.NewVenue(ChatID, "аше место назначения", "123 улю Шевченко ", 40, 40)
	            bot.Send(venue)

        // Прокрутка музыки
	    case "Музыка":
	          msg := tgbotapi.NewAudioUpload(ChatID, "tests/audio.mp3")
	          msg.Title     = "TEST"
	          msg.Duration  = 10
	          msg.Performer = "TEST"
	          msg.MimeType  = "audio/mpeg"
	          msg.FileSize  = 688
	           _, err := bot.Send(msg)    
	           
	           if err != nil {
	              fmt.Println(err)
	           }
   
		}


        // /start   - начало работы бота
        // /close   - очистка кнопок
		// комманда - сообщение, начинающееся с "/"
		key := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Command())

		switch update.Message.Command() {
				case "start":
					  reply = "Привет! Я телеграм-бот Барс."
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

// ***************************************************************
// Чтение Ключа из json.config
// ***************************************************************
func LoadConfiguration(file string) string {
	var config Config
	fmt.Println("Config file is was reading. File name: " + file)
	configFile, err := os.Open(file)
	defer configFile.Close()
	
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config.TelegramBotToken
}
