package main

type mcxBhav struct {
	Date            string
	Sessionid       int32
	Markettype      string
	Instrumentid    int32
	Instrumentname  string
	Symbol          string
	Expirydate      string
	Reserved1       string
	Strike          float64
	Opttype         string
	Prevclose       float64
	Open            float64
	High            float64
	Low             float64
	Close           float64
	Volume          int64
	Totalvalue      float64
	Lifetimehigh    float64
	Lifetimelow     float64
	Unit            string
	Settlementprice float64
	Numberoftrades  float64
	Oi              float64
	Avgtradeprice   float64
	Reserved2       string
	Currencycode    string
}
