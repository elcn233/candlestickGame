package tgbot

import (
	"fmt"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	log "github.com/sirupsen/logrus"
)

// 点击按钮后的应答程序
func startCB(b *gotgbot.Bot, ctx *ext.Context) error {
	// 这里有点坑，不能直接从 ctx 里获取 username、firstname，要从 ctx.Update.CallbackQuery 里获取
	cb := ctx.Update.CallbackQuery

	log.Infof("用户点击按钮，姓名： [%s] 用户名： [%s] ", cb.From.FirstName, cb.From.Username)
	_, err := cb.Answer(b, &gotgbot.AnswerCallbackQueryOpts{
		Text: "你按下了一个按钮！",
	})
	if err != nil {
		return fmt.Errorf("failed to answer start callback query: %w", err)
	}

	_, _, err = cb.Message.EditText(b, "您编辑了这条消息。", nil)
	if err != nil {
		return fmt.Errorf("failed to edit start message text: %w", err)
	}
	return nil
}
