package main

import (
	"fmt"
	"net/http"

	"github.com/factly/data-portal-server/action"
	"github.com/factly/data-portal-server/config"
	"github.com/factly/data-portal-server/model"
	"github.com/factly/data-portal-server/util/meili"
)

// @title Data portal API
// @version 1.0
// @description This is a sample server.

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:7721
// @BasePath /

func main() {
	config.SetupVars()

	// db setup
	model.SetupDB(config.DSN)

	model.Migration()

	meili.SetupMeiliSearch()

	// register routes
	userRouter := action.RegisterUserRoutes()
	adminRouter := action.RegisterAdminRoutes()

	fmt.Println("swagger-ui http://localhost:7720/swagger/index.html")

	go http.ListenAndServe(":7720", userRouter)
	http.ListenAndServe(":7721", adminRouter)

}
