package main

import (
	"github.com/deifyed/statusmsg/pkg/battery"
	"log"
)

func main() {
	batteryStatus, err := battery.GetBatteryStatus()
	if err != nil {
		log.Fatal(err)
	}
	
	println(batteryStatus.String())
}
