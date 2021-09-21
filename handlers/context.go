package handlers

import "github.com/PayAbbhiRS/appkit"

var appCtx *appkit.AppContext

//SetAppContext sets the application context in the handlers
func SetAppContext(ac *appkit.AppContext) {
	appCtx = ac
}
