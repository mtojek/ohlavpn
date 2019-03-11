package ipapi

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const baseURI = "http://ip-api.com/json"

// Client is responsible for calling IP-API.
type Client struct {
	httpClient *http.Client
}

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
	return &Client{
		httpClient: http.DefaultClient,
	}
}

// WithProxy method uses given proxy server to tunnel HTTP requests to IP-API.
func (c *Client) WithProxy(anURL *url.URL) *Client {
	pu := http.ProxyURL(anURL)
	c.httpClient = &http.Client{
		Transport: &http.Transport{
			Proxy:           pu,
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Timeout: 5 * time.Second,
	}
	return c
}

// GeoIP method returns geolocation data.
func (c *Client) GeoIP(ip string) (*GeoIPData, error) {
	response, err := c.httpClient.Get(baseURI + "/" + ip)
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
