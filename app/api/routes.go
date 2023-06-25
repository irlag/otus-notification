package api

import "errors"

var (
	RouteNotExistError = errors.New("route not found")
)

type Route struct {
	Name   string
	Method string
	Path   string
	Secure bool
}

type Routes map[string]Route

func (r *Routes) Find(name string) (Route, error) {
	if appRoute, ok := AppRoutes[name]; ok {
		return appRoute, nil
	}
	return Route{}, RouteNotExistError
}

var AppRoutes = Routes{
	"metrics": Route{
		Name:   "metrics",
		Method: "GET",
		Path:   "/metrics",
		Secure: false,
	},

	"healthcheck": Route{
		Name:   "healthcheck",
		Method: "GET",
		Path:   "/health",
		Secure: false,
	},

	"event_send": Route{
		Name:   "event_send",
		Method: "POST",
		Path:   "/event",
		Secure: false,
	},
}
