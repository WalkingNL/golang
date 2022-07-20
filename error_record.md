## error record

#### go build vs go run

    package main

    import (
        "fmt"
        "os"
        "path/filepath"
    )

    func main() {
        dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
        if err != nil {
            panic(err)
        }

        fmt.Println("args0: ", os.Args[0], "\nargs1: ", os.Args[1], "\ndirs: ", dir)
    }
    

##### go run

##### go build


