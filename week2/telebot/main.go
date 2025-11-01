package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	telebot "gopkg.in/telebot.v4"
)

// var TeleToken = os.Getenv("TELE_TOKEN")
var appVersion = "0.1.0"

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	TeleToken := os.Getenv("TELE_TOKEN")
	// fmt.Print(TeleToken)
	if TeleToken == "" {
		log.Fatal("BOT_TOKEN is not set in .env")
	}

	var kbotCmd = &cobra.Command{
		Use:     "kbot",
		Aliases: []string{"start"},
		Short:   "Telegram bot for controlling traffic light signals",
		Long: `A Telegram bot that allows controlling traffic light signals through GPIO pins.
	The bot accepts commands to toggle red, amber, and green lights on/off.`,

		Run: func(cmd *cobra.Command, args []string) {
			if TeleToken == "" {
				log.Fatal("TELE_TOKEN environment variable is not set")
			}

			fmt.Printf("kbot %s started\n", appVersion)

			kbot, err := telebot.NewBot(telebot.Settings{
				URL:    "",
				Token:  TeleToken,
				Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
			})

			if err != nil {
				log.Fatalf("Please check TELE_TOKEN env variable. %s", err)
			}

			kbot.Handle("/start", func(c telebot.Context) error {
				return c.Send("Привіт! Я бот для курсу DevOps.")
			})

			kbot.Handle(telebot.OnText, func(m telebot.Context) error {
				log.Printf("Ти написав боту: %s %s\n", m.Message().Payload, m.Text())
				return m.Send("Я отримав: " + m.Text())
			})

			kbot.Start()
		},
	}

	if err := kbotCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
