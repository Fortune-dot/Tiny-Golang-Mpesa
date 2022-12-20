package main

import (
    "bytes"
    "fmt"
    "net/http"
    "net/url"
)

const baseURL = "https://tinypesa.com/api/v1/"

func initTransaction() {
    formData := url.Values{
        "amount": {"10"},
        "msisdn": {"0740408496"},
    }

    client := &http.Client{}
    req, err := http.NewRequest("POST", baseURL+"express/initialize", bytes.NewReader([]byte(formData.Encode())))
    if err != nil {
        panic(err.Error())
    }

    req.Header.Add("Accept", "application/json")
    req.Header.Add("Apikey", "UX0bRhL3c20")
    req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

    res, err := client.Do(req)
    if err != nil {
        panic(err.Error())
    }

    if res.StatusCode != http.StatusOK {
        fmt.Println("Error:", res.Status)
    } else {
        fmt.Println("Check Phone for STK Push")
    }
}

func main() {
    initTransaction()
}
