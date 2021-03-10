package main

import "fmt"

// App - the struct which contains thinks like pointers
// to database connections
type App struct{}

// Run setup our application
func (app *App) Run() error {
	fmt.Println("Setting Up Our App")
	return nil
}


func main(){
	fmt.Println("GO REST API Course")
	app := App{}
	if err := app.Run(); err != nil{
		fmt.Println("Error starting up our REST API")
		fmt.Println(err)
	}

}