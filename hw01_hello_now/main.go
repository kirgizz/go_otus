package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	ntpServer := "0.beevik-ntp.pool.ntp.org"
	serverTime := time.Now()
	ntpTime, err := ntp.Time(ntpServer)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("current time: %s\n", serverTime.Round(0))
	fmt.Printf("exact time: %s\n", ntpTime.Round(0))
}
