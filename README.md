# updev-go-rest-api-course
Example project https://tutorialedge.net/courses/go-rest-api-course/02-project-setup/

### command
 - ทำการกำหนด module
    ```powershell
    D:\GO\go mod init github.com/ksupdev/updev-go-rest-api-course
    ```

 - Create directory cmd/service และทำการสร้าง ``main.go`` ภายใต้ folder นั้นด้วย
    ```powershell
    mkdir cmd/server
    ```
    ```GO
    [filename : main.go]
    package main

    import "fmt"

    func main(){
        fmt.Println("GO REST API Course")
    }
    ```
    ทำการทดสอบ 

    ``` powershell
    PS D:\GO\updev-go-rest-api-course> go run .\cmd\server\main.go
    GO REST API Course
    ```

- ทำการสร้าง ส่วนของการ Connect database โดยการใช้ struct ที่มีการระบุการทำงานให้เป็นแบบ pointers ** ใน GO เราสามารถเลือกได้ว่า Struct ที่เราสร้างขึ้นนั้นจะมมีการทำงานแบบ pointers หรือจะให้เป็นแต่ value ธรรมดา
    ```GO
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
    ```
    
    ```GO
    func (app *App) Run() error {
        fmt.Println("Setting Up Our App")
        return nil
    }
    ```
    เป็นการประกาศ Method Run สำหรับ App struct{}

- Implementing the Transport Package (health check)
    ```powershell
    PS D:\GO\updev-go-rest-api-course>mkdir internal
    PS D:\GO\updev-go-rest-api-course\internal>mkdir transport
    PS D:\GO\updev-go-rest-api-course\internal\transport>mkdir http
    ```
    ทำการสร้าง ``handler.go``

    ```GO

    ```

    ทำการ implement "github.com/gorilla/mux"

    ```powershell
    PS D:\git-myself\GO\go-projects\labs\updev-go-rest-api-course> go get github.com/gorilla/mux                                                                               
    PS D:\git-myself\GO\go-projects\labs\updev-g
    go: downloading github.com/gorilla/mux v1.8.0
    go get: added github.com/gorilla/mux v1.8.0

    ```

    ทำการ implement net/http เพื่อใช้สำหรับกับรับ Http Request

    ทำการทดสอบ

    ```powershell
    PS D:\git-myself\GO\go-projects\labs\updev-go-rest-api-course> go run .\cmd\server\main.go
    GO REST API Course
    Setting Up Our App
    Setting Up Routes
    
    ..... New terminal

    PS D:\git-myself\GO\go-projects\labs\updev-go-rest-api-course> curl http://localhost:8080/api/health
    StatusCode        : 200
    StatusDescription : OK
    Content           : I'm alive!
    ```
 
 