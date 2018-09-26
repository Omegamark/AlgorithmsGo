package main

import "fmt"

func maxStockProfit(slice []int) int {
    maxProfit := -1
    buyPrice := 0
    sellPrice := 0

    changeBuyPrice := true

    for i := 0; i < len(slice) - 1; i++ {
       if changeBuyPrice == true {buyPrice = slice[i]}
       
        sellPrice = slice[i+1]
        if buyPrice > sellPrice {
            changeBuyPrice = true
        } else {
            fmt.Println("I'm buyPrice", buyPrice)
            fmt.Println("I'm sellPrice", sellPrice)

            tempProfit := sellPrice - buyPrice
            fmt.Println("I'm temp", tempProfit)
            if maxProfit < tempProfit { maxProfit = tempProfit }
            changeBuyPrice = false
        }
    }
    return maxProfit
}

func main() {
stocks := []int{32, 46, 26, 38, 40, 48, 42, 12}
    fmt.Println(maxStockProfit(stocks))
}