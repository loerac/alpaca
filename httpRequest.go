package main

import (
    "bytes"
    "encoding/json"
    "io/ioutil"
    "net/http"
)

/**
 * @brief:  A HTTP GET request
 *
 * @arg:    url - URL of the API
 *
 * @return: Details on GET request on success,
 *          Else error
 **/
func HttpGet(url string) string {
    req, err := http.NewRequest("GET", url, nil)
    checkErr(err)

    req.Header.Set(API_KEY_HEADER, API_KEY)
    req.Header.Set(API_SECRET_HEADER, API_SECRET)

    res, err := http.DefaultClient.Do(req)
    checkErr(err)
    defer res.Body.Close()

    body, err := ioutil.ReadAll(res.Body)
    checkErr(err)

    return string(body)
}

/**
 * @brief:  A HTTP POST request
 *
 * @arg:    url - URL of the API
 *
 * @return: Details on POST request on success,
 *          Else error
 **/
func HttpPost(url string, order Order) string {
    orderBytes := new(bytes.Buffer)
    json.NewEncoder(orderBytes).Encode(order)
    req, err := http.NewRequest("POST", url, orderBytes)
    checkErr(err)

    req.Header.Set(API_KEY_HEADER, API_KEY)
    req.Header.Set(API_SECRET_HEADER, API_SECRET)

    res, err := http.DefaultClient.Do(req)
    checkErr(err)
    defer res.Body.Close()

    //fmt.Println("response Status:", res.Status)
    //fmt.Println("response Headers:", res.Header)
    body, err := ioutil.ReadAll(res.Body)
    checkErr(err)

    return string(body)
}
