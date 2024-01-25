package main

import (
	"fmt"
	"github.com/gammazero/deque"
	"github.com/shopspring/decimal"
)

type Trade struct {
	Id        string
	Type      string
	TotalGold decimal.Decimal
	TotalUSD  decimal.Decimal
	OurTag    string
	TheirTag  string
}

const (
	BUY  = "BUY"
	SELL = "SELL"
)

func NetPositionForTag(trades []*Trade, tag string) (decimal.Decimal, decimal.Decimal) {
	//first create a map[string][]string of every tag and its neighbors
	//for the tag we're given, traverse the neighbors to grab other trades in that group
	// calculate net position for trades
	// would use set as value here
	tagToNeighbors := make(map[string][]string)
	for _, trade := range trades {
		tagToNeighbors[trade.OurTag] = append(tagToNeighbors[trade.OurTag], trade.TheirTag)
		tagToNeighbors[trade.TheirTag] = append(tagToNeighbors[trade.TheirTag], trade.OurTag)
	}

	tagsQueue := deque.New[string]()
	tagsQueue.PushFront(tag)

	taggedTrades := make(map[*Trade]bool)
	visitedTags := map[string]bool{}

	for tagsQueue.Len() > 0 {
		curTag := tagsQueue.PopFront()
		visitedTags[curTag] = true
		for _, trade := range trades {
			if trade.OurTag == curTag || trade.TheirTag == curTag {
				taggedTrades[trade] = true
			}
		}
		neighbors := tagToNeighbors[curTag]
		for _, tag := range neighbors {
			if _, found := visitedTags[tag]; !found {
				tagsQueue.PushFront(tag)
			}
		}
	}

	taggedTradesSlice := make([]*Trade, 0)
	for trade, _ := range taggedTrades {
		taggedTradesSlice = append(taggedTradesSlice, trade)
	}

	return NetPosition(taggedTradesSlice)
}

func NetPosition(trades []*Trade) (decimal.Decimal, decimal.Decimal) {
	fmt.Println("Calculating Net trades for....")
	var (
		totalGold decimal.Decimal
		totalUSD  decimal.Decimal
	)
	for _, trade := range trades {
		fmt.Printf("Trade: %+v\n", *trade)
		if trade.Type == BUY {
			totalGold = totalGold.Add(trade.TotalGold)
			totalUSD = totalUSD.Sub(trade.TotalUSD)
		} else {
			totalGold = totalGold.Sub(trade.TotalGold)
			totalUSD = totalUSD.Add(trade.TotalUSD)
		}
	}
	return totalGold, totalUSD
}

func main() {
	trades := []*Trade{
		{
			Id:        "A",
			Type:      "BUY",
			TotalGold: decimal.NewFromFloat(23.55),
			TotalUSD:  decimal.NewFromFloat(10.00),
			OurTag:    "X",
			TheirTag:  "1",
		},
		{
			Id:        "B",
			Type:      "BUY",
			TotalGold: decimal.NewFromFloat(23.55),
			TotalUSD:  decimal.NewFromFloat(10.00),
			OurTag:    "X",
			TheirTag:  "0",
		},
		{
			Id:        "C",
			Type:      BUY,
			TotalGold: decimal.NewFromFloat(23.55),
			TotalUSD:  decimal.NewFromFloat(10.00),
			OurTag:    "Y",
			TheirTag:  "1",
		},
		{
			Id:        "D",
			Type:      SELL,
			TotalGold: decimal.NewFromFloat(23.55),
			TotalUSD:  decimal.NewFromFloat(10.00),
			OurTag:    "Z",
			TheirTag:  "2",
		},
		{
			Id:        "E",
			Type:      SELL,
			TotalGold: decimal.NewFromFloat(23.55),
			TotalUSD:  decimal.NewFromFloat(10.00),
			OurTag:    "Y",
			TheirTag:  "3",
		},
		{
			Id:        "F",
			Type:      BUY,
			TotalGold: decimal.NewFromFloat(23.55),
			TotalUSD:  decimal.NewFromFloat(10.00),
			OurTag:    "W",
			TheirTag:  "3",
		},
	}

	NetPositionForTag(trades, "Z")
}
