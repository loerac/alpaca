package main

import (
    "bytes"
    "encoding/json"
    "fmt"
)

func checkErr(e error) {
    if nil != e {
        panic(e)
    }
}

type argError struct {
    arg  int
    prob string
}

func (e *argError) Error() string {
    return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func PrintPrettyJson(str string) string {
    var prettyJSON bytes.Buffer
    err := json.Indent(&prettyJSON, []byte(str), "", "\t")
    checkErr(err)
    return string(prettyJSON.Bytes())
}
