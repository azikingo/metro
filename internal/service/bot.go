package service

import (
	"context"
	"fmt"
	"golang.org/x/text/width"
	"math"
	"strings"

	"github.com/go-kratos/kratos/v2/log"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gocolly/colly/v2"
	v1 "team-manager/api/helloworld/v1"
	"team-manager/internal/biz"
)

type BotService struct {
	log *log.Helper

	v1.UnimplementedGreeterServer
	userBiz *biz.UsersUsecase
}

// NewBotService new a greeter service.
func NewBotService(
	userBiz *biz.UsersUsecase,
	logger log.Logger,
) *BotService {
	return &BotService{
		userBiz: userBiz,
		log:     log.NewHelper(log.With(logger, "module", "service/bot")),
	}
}

func (bs *BotService) BotRequest(ctx context.Context, update tgbotapi.Update) (tgbotapi.Chattable, error) {
	message := update.Message
	reply := tgbotapi.Chattable(nil)

	// Check if we've gotten a message update.
	if message != nil {
		chatId := update.Message.Chat.ID
		tgUser := update.Message.From

		if message.Command() != "" {
			switch message.Command() {
			case "start":
				bs.RegisterUser(ctx, tgUser)
				reply = tgbotapi.NewMessage(chatId, "Welcome to the Metro bot! You can track metro schedule here!")

			case "metro":
				reply = tgbotapi.NewMessage(update.Message.Chat.ID, getScheduleContent())

			case "help":
				reply = tgbotapi.NewMessage(chatId, "I did a list to you here for commands that I can do:\n"+
					"/start - Start the bot\n"+
					"/metro - To get metro schedule\n"+
					"/help - Show this message\n")

			default:
				reply = tgbotapi.NewMessage(chatId, "Unknown command, use one of the commands in the menu.")
			}

			return reply, nil
		} else if message.Text != "" {
			reply = tgbotapi.NewMessage(chatId, "Welcome to the Almaty Metro bot! Please, use the commands in the menu.")
			return reply, nil
		}
	}

	return reply, nil
}

func (bs *BotService) RegisterUser(ctx context.Context, tgUser *tgbotapi.User) {
	_, err := bs.userBiz.CreateUser(ctx, tgUser)
	if err != nil {
		bs.log.Errorf("CreateUser error: %v", err)
		return
	}
}

type Metro struct {
	Schedules          []MetroSchedule
	IsFirstCellForward bool
	LongNameLength     int
}

type MetroSchedule struct {
	Station    string
	FirstCell  string
	SecondCell string
}

func parseTest() (Metro, error) {
	c := colly.NewCollector()

	metro := Metro{}
	isFirstCellForward := false
	longNameLength := 0
	c.OnHTML(".schedule__table", func(e *colly.HTMLElement) {
		e.ForEach("div.schedule__table-body", func(i int, el *colly.HTMLElement) {
			station := MetroSchedule{}
			el.ForEach("div", func(j int, el2 *colly.HTMLElement) {
				text := strings.Trim(el2.Text, " ")
				if j == 0 {
					station.Station = text
					longNameLength = int(math.Max(float64(longNameLength), float64(len(width.Narrow.String(text)))))
				} else if j == 1 {
					station.FirstCell = text
				} else {
					station.SecondCell = text
					if i == 0 && j == 2 && text == "" {
						isFirstCellForward = true
					}
				}
			})
			metro.Schedules = append(metro.Schedules, station)
		})
	})

	err := c.Visit("https://metroalmaty.kz/ru/schedule")

	metro.IsFirstCellForward = isFirstCellForward
	metro.LongNameLength = longNameLength

	return metro, err
}

func getScheduleContent() string {
	result := ""

	metro, err := parseTest()
	if err != nil {
		fmt.Println(err)
		return ""
	}

	for i, s := range metro.Schedules {
		if i == 0 || i == len(metro.Schedules)-1 {
			if s.FirstCell == "" {
				result += fmt.Sprint("       🏁       | ")
				//if i == 0 && metro.IsFirstCellForward {
				//	result += fmt.Sprint("      ⬆️      | ")
				//} else {
				//	result += fmt.Sprint("      ⬇️      | ")
				//}
			} else {
				result += fmt.Sprint(s.FirstCell + " ↓ | ")
			}
			if s.SecondCell == "" {
				result += fmt.Sprint("       🏁       ")
				//if i == 0 && metro.IsFirstCellForward {
				//	result += fmt.Sprint("      ⬆️     ")
				//} else {
				//	result += fmt.Sprint("      ⬇️     ")
				//}
			} else {
				result += fmt.Sprint(" ↑ " + s.SecondCell + " ")
			}
		} else {
			result += fmt.Sprint(s.FirstCell+" ↓ | ↑ ") + fmt.Sprint(s.SecondCell)
		}
		result += fmt.Sprintln(" | ", s.Station)
	}

	return result
}
