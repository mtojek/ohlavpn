# ohlavpn
**Find working VPN proxy in the given county | Hola VPN**

Use Hola VPN API to grab a list of fast VPN proxy servers all over the world. Use the IP-API endpoint to check their geolocation and owners.

## Disclaimer

**This code is only for educational purposes. Please do not overuse the Hola VPN infrastructure or avoid Netflix limits.**

## Getting started

```
$ go get github.com/mtojek/ohlavpn
$ ohlavpn -h
Usage of ohlavpn:
  -c string
    	country code (default "us")
  -g	check GeoIP data
  -l int
    	proxy server limit (default 5)
 ```

### First steps

Fetch active French proxy servers:
 
```
$ ohlavpn -c fr -l 5 -g
Login: 339e68ea036644a7987307a4ae965713, Password: 97a5ae965312

https	158.255.215.125:22222	AS39326 HighSpeed Office Limited	Paris	France	Edis France		EDIS GmbH	Île-de-France	75001
https	178.32.172.248:22222	AS16276 OVH SAS	Paris	France	OVH ISP		CO LTD MAXSERVER	Île-de-France	75000
https	193.70.63.13:22222	AS16276 OVH SAS	Gravelines	France	OVH		OVH SAS	Hauts-de-France	59820
https	51.255.175.61:22222	AS16276 OVH SAS	Strasbourg	France	OVH SAS		OVH	Grand Est	67000
https	178.32.227.14:22222	AS16276 OVH SAS	Roubaix	France	OVH SAS		CO LTD MAXSERVER	Hauts-de-France	59100
```
