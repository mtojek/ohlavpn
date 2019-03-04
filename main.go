package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/mtojek/ohlavpn/ipapi"
	"github.com/mtojek/ohlavpn/vpn"
)

func main() {
	var countryCode string
	var limit int
	var geoIP bool

	flag.StringVar(&countryCode, "c", "us", "country code")
	flag.IntVar(&limit, "l", 5, "proxy server limit")
	flag.BoolVar(&geoIP, "g", false, "check GeoIP data")
	flag.Parse()

	ipAPIClient := ipapi.NewClient()
	vpnClient := vpn.NewClient()
	err := vpnClient.Initialize()
	if err != nil {
		log.Fatalf("Error occurred while initializing VPN API vpnClient: %v", err)
	}

	tunnels, err := vpnClient.FindTunnels(countryCode, limit)
	if err != nil {
		log.Fatalf("Error occurred while finding VPN tunnels: %v", err)
	}

	if len(tunnels.Servers) == 0 {
		log.Fatal("No proxy servers found")
	}

	fmt.Printf("Login: %s, Password: %s\n\n", tunnels.Login, tunnels.Password)

	for _, tunnel := range tunnels.Servers {
		fmt.Print(tunnel.String())

		if geoIP {
			geoIPData, err := ipAPIClient.GeoIP(tunnel.Host)
			if err != nil {
				fmt.Printf("\terror checking GeoIP data: %v", err)
			} else {
				fmt.Printf("\t%v", geoIPData.String())
			}
		}

		fmt.Println()
	}
}
