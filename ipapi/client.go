package ipapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const baseUri = "http://ip-api.com/json"

type IPApi struct{}

type GeoIPData struct {
	As         string
	City       string
	Country    string
	ISP        string
	Org        string
	RegionName string
	Zip        string
}

func (g *GeoIPData) String() string {
	return fmt.Sprintf("%s\t%s\t%s\t%s\t\t%s\t%s\t%s", g.As, g.City, g.Country, g.ISP, g.Org, g.RegionName,
		g.Zip)
}

func NewIPApi() *IPApi {
	return new(IPApi)
}

func (api *IPApi) GeoIP(ip string) (*GeoIPData, error) {
	response, err := http.Get(baseUri + "/" + ip)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	var r GeoIPData

	err = json.NewDecoder(response.Body).Decode(&r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
