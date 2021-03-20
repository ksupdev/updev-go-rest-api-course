package main

import (
	"net/http"

	"github.com/ksupdev/updev-go-rest-api-course/internal/comment"
	"github.com/ksupdev/updev-go-rest-api-course/internal/database"
	transportHTTP "github.com/ksupdev/updev-go-rest-api-course/internal/transport/http"
	log "github.com/sirupsen/logrus"
)

// App - contain application informaion
// to database connections
type App struct {
	Name    string
	Version string
}

// Run setup our application
func (app *App) Run() error {
	// fmt.Println("Setting Up Our App")
	log.SetFormatter(&log.JSONFormatter{})
	log.WithFields(
		log.Fields{
			"AppName":    app.Name,
			"AppVersion": app.Version,
		}).Info("Setting up application")

	var err error
	db, err := database.NewDatabase()
	if err != nil {
		return err
	}

	err = database.MigrateDB(db)
	if err != nil {
		return err
	}

	commentService := comment.NewService(db)

	handler := transportHTTP.NewHandler(commentService)
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8088", handler.Router); err != nil {

		// fmt.Println("Failed to set up server")
		log.Error("Failed to set up server")
		return err
	}

	return nil
}

func main() {
	// fmt.Println("GO REST API Course")
	app := App{
		Name:    "Commenting Service",
		Version: "1.0.0",
	}
	if err := app.Run(); err != nil {
		// fmt.Println("Error starting up our REST API")
		log.Error("Error starting up our REST API")
		log.Fatal(err)
		// fmt.Println(err)
	}

}
