package pronote

import (
	telegram "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/theovidal/bacbot/lib"

	"github.com/theovidal/bacbot/pronote/api"
)

func HomeworkCommand() lib.Command {
	return lib.Command{
		Execute: func(bot *telegram.BotAPI, update *telegram.Update, _ []string, _ map[string]interface{}) error {
			response, err := api.GetHomework()
			if err != nil {
				return err
			}

			if len(response.Homeworks) == 0 {
				msg := telegram.NewMessage(update.Message.Chat.ID, "🍃 Aucun devoir n'a été rédigé pour le moment.")
				_, err = bot.Send(msg)
				return err
			}

			content := ""
			for _, homework := range response.Homeworks {
				content += homework.String()
			}

			msg := telegram.NewMessage(update.Message.Chat.ID, content)
			msg.ParseMode = "MarkdownV2"
			msg.DisableWebPagePreview = true
			_, err = bot.Send(msg)

			return err
		},
	}
}
