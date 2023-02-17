package tgbot

import (
	"CandlestickGame/image"
	"fmt"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	log "github.com/sirupsen/logrus"
)

// /start 命令的处理程序
func start(b *gotgbot.Bot, ctx *ext.Context) error {
	log.Infof("用户开始，姓名： [%s] 用户名： [%s] ", ctx.Message.From.FirstName, ctx.Message.From.Username)
	/*_, err := ctx.EffectiveMessage.Reply(b, fmt.Sprintf("<b>你好</b>, 我是 @%s", b.User.Username), &gotgbot.SendMessageOpts{
		ParseMode: "html",
		ReplyMarkup: gotgbot.InlineKeyboardMarkup{
			InlineKeyboard: [][]gotgbot.InlineKeyboardButton{{
				{Text: "按钮", CallbackData: "start_callback"},
			}},
		},
	})*/

	// 获取文件
	f, err := image.Assets.Open("模板.png")
	if err != nil {
		return fmt.Errorf("failed to open source: %w", err)
	}

	b.SendPhoto(ctx.Message.From.Id, f, &gotgbot.SendPhotoOpts{
		Caption: "你好！",
	})
	if err != nil {
		return fmt.Errorf("failed to send start message: %w", err)
	}
	return nil
}
