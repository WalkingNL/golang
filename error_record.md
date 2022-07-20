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

        fmt.Println("args: ", os.Args[0], "\ndirs: ", dir)
    }

run this code in two ways: "go build" and "go run", what's the difference of the results.

##### go build

    # commands
    cd ./local_files/
    go build // generate binary file local_files
    ./local_files

    # result
    args:  ./local_files 
    dirs:  /Users/nlgong/Desktop/Coding/go/colly/_examples/local_files

##### go run
    # commands
    cd ./local_files/

    # result
    args:  /var/folders/c6/rnkpvx4140516mv516vhwdy40000gn/T/go-build3035963001/b001/exe/local_files 
    dirs:  /var/folders/c6/rnkpvx4140516mv516vhwdy40000gn/T/go-build3035963001/b001/exe


the results is completely different by these two ways. the latter is in a tmp dir so that we got the tmp path
