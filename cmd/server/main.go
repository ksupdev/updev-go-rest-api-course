package main

import (
	"fmt"
	"net/http"

	"github.com/ksupdev/updev-go-rest-api-course/internal/comment"
	"github.com/ksupdev/updev-go-rest-api-course/internal/database"
	transportHTTP "github.com/ksupdev/updev-go-rest-api-course/internal/transport/http"
)

// App - the struct which contains thinks like pointers
// to database connections
type App struct{}

// Run setup our application
func (app *App) Run() error {
	fmt.Println("Setting Up Our App")

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

		fmt.Println("Failed to set up server")
		return err
	}

	return nil
}

func main() {
	fmt.Println("GO REST API Course")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up our REST API")
		fmt.Println(err)
	}

}
