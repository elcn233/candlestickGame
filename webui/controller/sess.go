package controller

import (
	"github.com/kataras/iris/v12/sessions"
)

var (
	cookieNameForSessionID = "session"
	Sess                   = sessions.New(sessions.Config{Cookie: cookieNameForSessionID})
)
