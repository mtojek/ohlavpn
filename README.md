# ohlavpn

[![Build Status](https://travis-ci.org/mtojek/ohlavpn.svg?branch=master)](https://travis-ci.org/mtojek/ohlavpn)

Status: **Done**

**Find working VPN proxy all over the world**

Use Hola VPN API to grab a list of fast VPN proxy servers all over the world. Use the IP-API endpoint to check their geolocation and owners (mind the service limits).

## Disclaimer

**This project is only for educational purposes. Please do not overuse the Hola VPN infrastructure or avoid Netflix limits.**

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
https://user-uuid-2188652440aa480b9b789bc8ba5a2882:5db3017a5b01@159.203.90.106:22222	AS14061 DigitalOcean, LLC	Clifton	United States	DigitalOcean, LLC		Digital Ocean	New Jersey	07014
https://user-uuid-2188652440aa480b9b789bc8ba5a2882:5db3017a5b01@184.164.147.6:22222	AS20454 SECURED SERVERS LLC	Tempe	United States	Secured Servers LLC		Secured Servers LLC	Arizona	85281
https://user-uuid-2188652440aa480b9b789bc8ba5a2882:5db3017a5b01@66.85.140.5:22222	AS20454 SECURED SERVERS LLC	Phoenix	United States	Secured Servers LLC		Dolorem Ipsum, s.r.o	Arizona	85001
https://user-uuid-2188652440aa480b9b789bc8ba5a2882:5db3017a5b01@184.164.146.16:22222	AS20454 SECURED SERVERS LLC	Tempe	United States	Secured Servers LLC		Secured Servers LLC	Arizona	85281
https://user-uuid-2188652440aa480b9b789bc8ba5a2882:5db3017a5b01@184.164.133.43:22222	AS20454 SECURED SERVERS LLC	Tempe	United States	Secured Servers LLC		Secured Servers LLC	Arizona	85281
```
