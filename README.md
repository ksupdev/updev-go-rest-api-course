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
    
    ```

 
 