package main

import (
	"context"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"slices"
	"strings"
	"time"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/spf13/viper"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	cfg, err := load_config()
	if err != nil {
		log.Fatal(err)
	}

	opts := []bot.Option{
		bot.WithDefaultHandler(execShell),
		bot.WithCheckInitTimeout(5 * time.Minute),
	}

	b, err := bot.New(cfg.TELEGRAM_BOT_TOKEN, opts...)
	if err != nil {
		panic(err)
	}

	b.Start(ctx)
}

type Config struct {
	TELEGRAM_BOT_TOKEN string `mapstructure:"TELEGRAM_BOT_TOKEN"`
	ALLOWED_USERS      string `mapstructure:"ALLOWED_USERS"`
}

func load_config() (cfg Config, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&cfg)
	return
}

func execShell(ctx context.Context, b *bot.Bot, update *models.Update) {
	msg := update.Message.Text

	if msg == "/start" {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "HI 👋",
		})
		return
	}

	username := update.Message.From.Username
	if grantUser(username) == false {
		log.Printf("%s trying to access...\n", username)
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "403 - NOT AUTHORIZE",
		})
		return
	}

	var text string
	var cmd = strings.Split(msg, " ")

	log.Printf("%s - %s\n", username, msg)

	stdout, err := exec.Command(cmd[0], cmd[1:]...).Output()

	text = string(stdout)
	if err != nil {
		text = err.Error()
	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   text,
	})
}

func grantUser(username string) bool {
	cfg, _ := load_config()

	users := strings.Split(cfg.ALLOWED_USERS, ",")

	exist := slices.Contains(users, username)

	return exist
}
