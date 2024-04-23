package server

import (
	"context"
	"os"
	"team-manager/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BotServer struct {
	log        *log.Helper
	bot        *tgbotapi.BotAPI
	botService *service.BotService
}

func NewBotServer(
	logger log.Logger,
	botService *service.BotService,
) *BotServer {
	bs := &BotServer{
		log:        log.NewHelper(log.With(logger, "module", "server/bot")),
		botService: botService,
	}

	tgBot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))
	if err != nil {
		log.Errorf("failed on create bot: %v", err)
	} else {
		bs.bot = tgBot
	}

	return bs
}

func (bs *BotServer) Start(ctx context.Context) error {
	bs.telegramBot()
	bs.log.Info("bot started")

	return nil
}

func (bs *BotServer) Stop(ctx context.Context) error {
	bs.bot.StopReceivingUpdates()
	bs.log.Info("bot stopped")

	return nil
}

func (bs *BotServer) telegramBot() {

	//Устанавливаем время обновления
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	//Получаем обновления от бота
	updates := bs.bot.GetUpdatesChan(u)

	for update := range updates {

		// handle request
		reply, err := bs.botService.BotRequest(context.Background(), update)
		if err != nil {
			bs.log.Errorf("BotRequest error: %v", err)
			bs.bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Something went wrong. Please, try again later."))
		}

		// send reply
		if reply != nil {
			bs.bot.Send(reply)
		} else {
			bs.log.Errorf("Reply not found")
			if update.Message != nil {
				bs.bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Something went wrong. Please, try again later."))
			}
		}
	}
}
