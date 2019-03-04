package ipapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const baseURI = "http://ip-api.com/json"

// Client is responsible for calling IP-API.
type Client struct{}

// GeoIPData stores geolocation data.
type GeoIPData struct {
	As         string
	City       string
	Country    string
	ISP        string
	Org        string
	RegionName string
	Zip        string
}

// String method returns string representative of the struct.
func (g *GeoIPData) String() string {
	return fmt.Sprintf("%s\t%s\t%s\t%s\t\t%s\t%s\t%s", g.As, g.City, g.Country, g.ISP, g.Org, g.RegionName,
		g.Zip)
}

// NewClient creates new instance of the IP-API client.
func NewClient() *Client {
	return new(Client)
}

// GeoIP method returns geolocation data.
func (c *Client) GeoIP(ip string) (*GeoIPData, error) {
	response, err := http.Get(baseURI + "/" + ip)
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
