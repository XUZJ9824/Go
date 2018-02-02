package main

import (
        "bufio"
        "fmt"
        "io"
        "io/ioutil"
        "os"        
)

func check( e error){
    if (e != nil )
    {
        panic(e)
    }
}

func test_file_io(){
    dat, err := ioutil.ReadFile("c:\\temp\test.log");
    check(err)
    fmt.Print(string(dat))
    
    f, err := os.Open("c:\\temp\\aa.log")
    check(err)
    
    
    
}

