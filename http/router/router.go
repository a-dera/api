package router

import (
	"github.com/duniapay/api/http/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func New() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		//handler = auth.JwtMiddleware.Handler(handler)
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

var routes = Routes{
	Route{
		"HealthCheck",
		"GET",
		"/_ah/health",
		handlers.HealthCheck,
	},
	Route{
		"PostHealthCheck",
		"POST",
		"/v1/health/check",
		handlers.PostHealthCheck,
	},
	Route{
		"InitizalizeTransaction",
		"POST",
		"/v1/transactions/initialize",
		handlers.PostHealthCheck,
	},
	Route{
		"VerifyTransaction",
		strings.ToUpper("GET"),
		"/v1/transactions/verify/{transactionId}",
		handlers.HealthCheck,
	},
	Route{
		"TransactionTotals",
		strings.ToUpper("GET"),
		"/v1/transactions/totals",
		handlers.HealthCheck,
	},
	Route{
		"VerifyTransaction",
		strings.ToUpper("GET"),
		"/v1/transactions/totals",
		handlers.HealthCheck,
	},
	Route{
		"ExportTransactions",
		strings.ToUpper("GET"),
		"/v1/transactions/export",
		handlers.HealthCheck,
	},
	Route{
		"ViewTransactionTimeline",
		strings.ToUpper("GET"),
		"/v1/transactions/timeline/{transactionId}",
		handlers.HealthCheck,
	},
	Route{
		"ListTransactions",
		strings.ToUpper("GET"),
		"/v1/transactions",
		handlers.HealthCheck,
	},
	Route{
		"FetchTransaction",
		strings.ToUpper("GET"),
		"/v1/transactions/{transactionId}",
		handlers.HealthCheck,
	},
	Route{
		"ChargeAuthorization",
		strings.ToUpper("POST"),
		"/v1/transactions/charge_authorization",
		handlers.HealthCheck,
	},
	Route{
		"CheckAuthorization",
		strings.ToUpper("POST"),
		"/v1/transactions/check_authorization",
		handlers.PostHealthCheck,
	},
}
