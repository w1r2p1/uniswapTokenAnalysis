package main

import (
	"encoding/json"
	"bytes"
	"net/http"
	"io/ioutil"
	"fmt"
)

type Payload struct {
	OperationName interface{} `json:"operationName"`
	Variables     Variables   `json:"variables"`
	Query         string      `json:"query"`
}

type Variables struct {
	AllPairs []string `json:"allPairs"`
}

func main()  {
	data := Payload{
		// fill struct
		OperationName:nil,
		Variables: Variables{
			AllPairs:[]string{"0xc5be99a02c6857f9eac67bbce58df5572498f40c"},
		},
		Query:"query ($allPairs: [Bytes]!) {\n  mints(first: 30, where: {pair_in: $allPairs}, orderBy: timestamp, orderDirection: desc) {\n    transaction {\n      id\n      timestamp\n      __typename\n    }\n    pair {\n      token0 {\n        id\n        symbol\n        __typename\n      }\n      token1 {\n        id\n        symbol\n        __typename\n      }\n      __typename\n    }\n    to\n    liquidity\n    amount0\n    amount1\n    amountUSD\n    __typename\n  }\n  burns(first: 30, where: {pair_in: $allPairs}, orderBy: timestamp, orderDirection: desc) {\n    transaction {\n      id\n      timestamp\n      __typename\n    }\n    pair {\n      token0 {\n        id\n        symbol\n        __typename\n      }\n      token1 {\n        id\n        symbol\n        __typename\n      }\n      __typename\n    }\n    sender\n    liquidity\n    amount0\n    amount1\n    amountUSD\n    __typename\n  }\n  swaps(first: 30, where: {pair_in: $allPairs}, orderBy: timestamp, orderDirection: desc) {\n    transaction {\n      id\n      timestamp\n      __typename\n    }\n    id\n    pair {\n      token0 {\n        id\n        symbol\n        __typename\n      }\n      token1 {\n        id\n        symbol\n        __typename\n      }\n      __typename\n    }\n    amount0In\n    amount0Out\n    amount1In\n    amount1Out\n    amountUSD\n    to\n    __typename\n  }\n}\n",
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	reqBody := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "https://api.thegraph.com/subgraphs/name/ianlapham/unsiwap3", reqBody)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Origin", "https://uniswap.info")
	req.Header.Set("Referer", "https://uniswap.info/pair/0xc5be99a02c6857f9eac67bbce58df5572498f40c")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "cross-site")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(respBody))

	var respData ResponseData
	err = json.Unmarshal(respBody, &respData)
	if err != nil {
		panic(err)
	}

	fmt.Println(respData)
}

type ResponseData struct {
	Data Data `json:"data"`
}

type Data struct {
	Burns []Burns `json:"burns"`
	Mints []Mints `json:"mints"`
	Swaps []Swaps `json:"swaps"`
}

type Burns struct {
	Typename  string `json:"__typename"`
	Amount0   string `json:"amount0"`
	Amount1   string `json:"amount1"`
	AmountUSD string `json:"amountUSD"`
	Liquidity string `json:"liquidity"`
	Pair Pair `json:"pair"`
	Sender string `json:"sender"`
	Transaction struct {
		Typename  string `json:"__typename"`
		ID        string `json:"id"`
		Timestamp string `json:"timestamp"`
	} `json:"transaction"`
}

type Pair struct {
	Typename string `json:"__typename"`
	Token0 Token `json:"token0"`
	Token1 Token `json:"token1"`
}

type Token struct {
	Typename string `json:"__typename"`
	ID       string `json:"id"`
	Symbol   string `json:"symbol"`
}

type Mints struct {
	Typename  string `json:"__typename"`
	Amount0   string `json:"amount0"`
	Amount1   string `json:"amount1"`
	AmountUSD string `json:"amountUSD"`
	Liquidity string `json:"liquidity"`
	Pair Pair `json:"pair"`
	To string `json:"to"`
	Transaction Transaction `json:"transaction"`
}

type Transaction struct {
	Typename  string `json:"__typename"`
	ID        string `json:"id"`
	Timestamp string `json:"timestamp"`
}

type Swaps struct {
	Typename   string `json:"__typename"`
	Amount0In  string `json:"amount0In"`
	Amount0Out string `json:"amount0Out"`
	Amount1In  string `json:"amount1In"`
	Amount1Out string `json:"amount1Out"`
	AmountUSD  string `json:"amountUSD"`
	ID         string `json:"id"`
	Pair Pair `json:"pair"`
	To string `json:"to"`
	Transaction Transaction `json:"transaction"`
}
