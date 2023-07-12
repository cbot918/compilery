package main

import (
	"encoding/json"
	"fmt"
	"log"
)

var (
	lg = fmt.Println
	lf = fmt.Printf
)

var lv = func(v any) {
	fmt.Printf("%+v", v)
}

var lj = func(v any) {
	res, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Println("json 轉換失敗, 搜尋 lj")
		log.Fatal(err)
	}
	fmt.Println(res)
}
