package main

import (
    "fmt"
)

func main() {
    /* Get details of the account */
    acct := AcctAuth()

    /* Request an order */
    //reqOrder := RequestOrder("AAPL", "buy", "market", "day", 10)

    /* List all open orders */
    allOrders := ListOrders()

    fmt.Println("Account Information:")
    fmt.Println("====================")
    fmt.Println(PrintPrettyJson(acct))

    fmt.Println("Request Order:")
    fmt.Println("====================")
    //fmt.Println(json.MarshalIndent(reqOrder, "", "    "))

    fmt.Println("List All Order:")
    fmt.Println("====================")
    fmt.Println(PrintPrettyJson(allOrders))
}
