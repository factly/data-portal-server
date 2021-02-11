package main

import (
	"log"
	"net/http"

	"github.com/factly/data-portal-server/util/razorpay"
	"github.com/factly/x/meilisearchx"
	"github.com/spf13/viper"

	"github.com/factly/data-portal-server/action"
	"github.com/factly/data-portal-server/config"
	"github.com/factly/data-portal-server/model"
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
	model.SetupDB(viper.GetString("postgres.dsn"))

	model.Migration()

	meilisearchx.SetupMeiliSearch("data-portal", []string{"name", "slug", "description", "title", "contact_name", "contact_email", "license", "caption", "alt_text", "plan_name", "plan_info"})

	razorpay.SetupClient()

	err := config.CreateSuperOrganisation()
	if err != nil {
		log.Println(err)
	}

	// register routes
	userRouter := action.RegisterUserRoutes()
	adminRouter := action.RegisterAdminRoutes()
	webhookRouter := action.RegisterWebHookRoutes()

	go func() {
		log.Fatal(http.ListenAndServe(":7720", userRouter))
	}()

	go func() {
		log.Fatal(http.ListenAndServe(":7722", webhookRouter))
	}()

	log.Fatal(http.ListenAndServe(":7721", adminRouter))

}
