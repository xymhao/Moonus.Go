package Proxy

import "fmt"

type Seller interface {
	sell(name string)
}

type Station struct {
	stock int
}

func (station Station) sell(name string) {
	if station.stock > 0 {
		station.stock--
	} else {
		fmt.Println("库存为空")
	}
}

type StationProxy struct {
	station Station
}

func (proxy StationProxy) sell(name string) {
	//before do something
	proxy.station.sell(name)
	//after do something
}
