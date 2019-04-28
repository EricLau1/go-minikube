package utils

type Social struct {
	Email    string `json:"email"`
	Github   string `json:"github"`
	Youtube  string `json:"youtube"`
	LinkedIn string `json:"linkedIn"`
}

type Author struct {
	Name   string `json:"name"`
	Social Social `json:"social"`
}

type Api struct {
	Version     string `json:"version"`
	Description string `json:"description"`
	Language    string `json:"language"`
	Author      Author `json:"author"`
	HelpUrl     string `json:"help_url"`
}

var INFO = Api{
	"1.0", 
	"Test Backend Developer: Avenue Securities",
	"go version go1.11.5 linux/amd64",
	Author{
		"Eric Lau de Oliveira",
		Social{
			"ericlau.oliveira@gmail.com",
			"https://github.com/EricLau1",
			"www.youtube.com/channel/UCr_3nxsd5v6g980xNese71w",
			"https://www.linkedin.com/in/ericlau2",
		},
	},
	"localhost:9000/help",
}

type Help struct {
	Uri         string `json:"uri"`
	Method      string `json:"method"`
	Description string `json:"description"`
}

var HELPS = []Help{
	Help{
		"/owners",
		"GET",
		"Paged list of owners. Add the parameter on the route: ?page=1",
	},
	Help{
		"/owners/1",
		"GET",
		"Return Owner by ID",
	},
	Help{
		"/owners",
		"POST",
		"Add Owner. Send a JSON with: { first_name, last_name, email, password, status, gender }. Return the new Owner's Wallet",
	},
	Help{
		"/owners/1",
		"PUT",
		`Update Owner by ID.
		 Option #1: Send a JSON with: { first_name, last_name, email, gender }. 		 
		 Option #2: Add parameter to disable account: ?disable=true.
		 Return Affrected Rows`,
	},
	Help{
		"/wallets",
		"GET",
		"paged list of wallets. Add the parameter on the route: ?page=1",
	},
	Help{
		"/wallets/1",
		"GET",
		"Return Wallet by ID",
	},
	Help{
		"/wallets",
		"PUT",
		"Add money in Wallet. Send a JSON: { public_key, cash }. Return Affected Rows",
	},
	Help{
		"/wallets/{public_key}",
		"POST",
		`Transfer money to another account.
		 Inform the public key of the Target Wallet at the URL.
		 Send a JSON with Origin Wallet: { public_key, cash } in the request body.
		 Returns the transfer Log.
		`,
	},
	Help{
		"/logs",
		"GET",
		"Paged list of logs. Add the parameter on the route: ?page=1",
	},
}

