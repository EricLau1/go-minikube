package routes

import (
	"go-minikube/api/controllers"
	"net/http"
)

type Route struct {
	Pattern string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
}

var Routes = []Route{
	Route{"/owners", "GET", controllers.GetOwners},
	Route{"/owners/{id}", "GET", controllers.GetOwner},
	Route{"/owners", "POST", controllers.PostOwner},
	Route{"/owners/{id}", "PUT", controllers.PutOwner},
	Route{"/wallets", "GET", controllers.GetWallets},
	Route{"/wallets/{id}", "GET", controllers.GetWallet},
	Route{"/wallets", "PUT", controllers.PutWallet},
	Route{"/wallets/{public_key}", "POST", controllers.PostWallet},
	Route{"/logs", "GET", controllers.GetLogs},
	Route{"/", "GET", controllers.GetHome},
	Route{"/help", "GET", controllers.GetHelp},
}
