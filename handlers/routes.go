package handlers

import "github.com/PayAbbhiRS/bridge-app-svc/models"

var routes []models.Route

// GetRoutes returns a list of routes defined in the http router
func GetRoutes() []models.Route {
	return routes
}

// SetRoutes sets a list of routes to be used in the http router
func SetRoutes(rts []models.Route) {
	routes = rts
}
