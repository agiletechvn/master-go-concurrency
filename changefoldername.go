package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
    "regexp"
)

func main() {
    pwd, err := os.Getwd()
    folderPattern := regexp.MustCompile("(\\d+_Codes)$")

    files, err := ioutil.ReadDir(pwd)
    if err != nil {
        log.Fatal(err)
    }

    for _, f := range files {
        res := folderPattern.FindStringSubmatch(f.Name())
        if len(res) > 0 {
            path := filepath.Join(pwd, f.Name())
            newpath := filepath.Join(pwd, res[0])
            fmt.Printf("mv %q %q\n", path, newpath)
            os.Rename(path, newpath)
        }
    }
}
