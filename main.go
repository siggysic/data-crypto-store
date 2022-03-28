package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/markcheno/go-quote"
)

const (
	DATE_FORMAT = "2006-01-02"
)

func main() {

	startTime := time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC)
	endTime := time.Date(2021, 12, 31, 0, 0, 0, 0, time.UTC)

	runningTime := startTime
	breakTime := endTime.Format(DATE_FORMAT)

	f, err := os.Create("btc-data-4h.csv")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	isBreak := false
	for {
		if isBreak {
			break
		}

		rTime := runningTime.Add(24 * time.Hour)
		rTimeF := rTime.Format(DATE_FORMAT)
		if rTimeF == breakTime {
			isBreak = true
		}
		sTimeF := runningTime.Format(DATE_FORMAT)
		fmt.Printf("%s - %s Start..\n", sTimeF, rTimeF)

		btcData, err := quote.NewQuoteFromBinance("BTCUSDT", sTimeF, rTimeF, quote.Hour4)
		if err != nil {
			panic(err)
		}

		priceStrs := []string{}
		for _, price := range btcData.Close {
			priceStrs = append(priceStrs, fmt.Sprintf("%f", price))
		}

		prices := strings.Join(priceStrs, ",")

		_, err = f.WriteString(sTimeF + "," + prices + "\n")
		if err != nil {
			panic(err)
		}
		runningTime = rTime

		fmt.Printf("%s - %s DONE..\n", sTimeF, rTimeF)
	}

}
