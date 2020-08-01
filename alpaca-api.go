package main

type Order struct {
    Symbol  string  `json:"symbol"`
    Qty     int     `json:"qty"`
    Side    string  `json:"side"`
    Type    string  `json:"type"`
    TimeInForce string `json:"time_in_force"`
}

const (
    API_KEY     string = "INVALID-KEY-ID"
    API_SECRET  string = "INVALID-SECRET-KEY"
    API_KEY_HEADER string = "APCA-API-KEY-ID"
    API_SECRET_HEADER string = "APCA-API-SECRET-KEY"
    BASE_URL    string = "https://paper-api.alpaca.markets/"
    ACCT_URL    string = BASE_URL + "v2/account"
    ORDER_URL   string = BASE_URL + "v2/orders"
)

/**
 * @brief:  Authenticate the account
 *
 * @return: Details of the given account is valid,
 *          Else error
 **/
func AcctAuth() string {
    return HttpGet(ACCT_URL)
}

/**
 * @brief:  List all the open orders on the account
 *
 * @return: Details on the open orders on success,
 *          Else error
 **/
func ListOrders() string {
    return HttpGet(ORDER_URL)
}

/**
 * @brief:  Request an order
 *
 * @arg:    _symbol - The ticker of the company
 * @arg:    _side - Buy or sell
 * @arg:    _type - market, limit, stop, or stop_loss
 * @arg:    _qty - Number of shares to buy
 *
 * @return: Details on the request order on success,
 *          Else error
 **/
func RequestOrder(_symbol, _side, _type, _time_in_force string, _qty int) string {
    order := Order {
        Symbol: _symbol,
        Qty: _qty,
        Side: _side,
        Type: _type,
        TimeInForce: _time_in_force,
    }

    return HttpPost(ORDER_URL, order)
}
