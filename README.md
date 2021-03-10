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
    และทำการ Implement ``handler.go`` เพื่อทำการจัดการ Routes ตามที่มีการ Request เข้ามา ซึ่งก็คือจะทำการรับ request ``/api/health``และทำการ response ``I'm alive!``


    และทำการ implement net/http เพื่อใช้สำหรับรับ Http Request และส่วไปให้ handler.goเป็นคนจัดการกับ Request นั้นๆ

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

- Running Postgres Locally with Docker
    Getting a docker instance up and running locally using docker
    ```docker
    docker run --name some-postgres -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgres

    PS D:\git-myself\GO\go-projects\labs\updev-go-rest-api-course> docker ps
    CONTAINER ID   IMAGE      COMMAND                  CREATED          STATUS          PORTS                    NAMES
    f7cf9e4973b6   postgres   "docker-entrypoint.s…"   22 seconds ago   Up 21 seconds   0.0.0.0:5432->5432/tcp   some-postgres

    ```
- Implementing the Database Package
    
    Create a new directory within `internal` directory called ``database``
    ```
    mkdir .\internal\database 
    ```

    Create a new file called database.go และทำการ import GORM และ postgress driver ของ สำหรับ gorm
    ```GO
        import (
            "fmt"

            "github.com/jinzhu/gorm"
            _ "github.com/jinzhu/gorm/dialects/postgres"
        )
    ```
    จาก Code จะเห็นว่ามีการ ใช้ ``_`` อยู่หน้า ``github.com/jinzhu/gorm/dialects/postgres`` นั้นเป็นการบอกว่าเป็นการ Import ในส่วนของ module ที่มีการเกี่ยวข้อง หรือมีผลกับ module ด้านบน

    ทำการ run command เพื่อ Download modulec และติดตั้งใน project

    ```powershell
    PS D:\git-myself\GO\go-projects\labs\updev-go-rest-api-course> go get github.com/jinzhu/gorm
    go: downloading github.com/jinzhu/gorm v1.9.16
    go: downloading github.com/jinzhu/inflection v1.0.0
    go get: added github.com/jinzhu/gorm v1.9.16
    PS D:\git-myself\GO\go-projects\labs\updev-go-rest-api-course> go get github.com/jinzhu/gorm/dialects/postgres
    go: downloading github.com/lib/pq v1.1.1
    ```

    ทำการ Implement ในส่วน defind dtabase properties และส่วนของการ Connect database ที่ ``file database.go``


    ทำการ run เพื่อ Test 
    ``` powershell
    PS D:\git-myself\GO\go-projects\labs\updev-go-rest-api-course> go run .\cmd\server\main.go
    GO REST API Course
    Setting Up Our App
    Setting up new database connection
    Setting Up Routes

    ```


 
 