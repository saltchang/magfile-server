package router

import "github.com/saltchang/magfile-server/handler"

var h = handler.NewHandler(database)

// Routes is the list of routes
var Routes = map[string]interface{}{
	"^/$":             h.HomeHandler,
	"^/users(/)?$":    h.UsersHandler,
	"^/users/(\\d)+$": h.UsersHandler,
}
