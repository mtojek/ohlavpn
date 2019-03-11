package vpn

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/google/uuid"
)

const (
	baseURI = "https://client.hola.org/client_cgi"

	browser = "chrome"
	extVer  = "1.125.157"
	rmtVer  = "1.2.676"
)

// Client is responsible for calling Hola VPN API.
type Client struct {
	uuid string
	key  string
}

// Tunnels stores information about proxy servers
type Tunnels struct {
	Servers []TunnelSettings
}

// TunnelSettings stores tunnel related settings.
type TunnelSettings struct {
	Login    string
	Password string

	Host  string
	Port  string
	Proto string
}

// String method returns string representative of the struct.
func (ts *TunnelSettings) String() string {
	return fmt.Sprintf("%s://%s:%s@%s:%s", ts.Proto, ts.Login, ts.Password, ts.Host, ts.Port)
}

// URL method returns URI representative of the struct.
func (ts *TunnelSettings) URL() *url.URL {
	proxyURL, err := url.Parse(ts.String())
	if err != nil {
		log.Fatalf("URL must be parsed: %v", err)
	}
	return proxyURL
}

type initializeResponse struct {
	Key int
}

type zGetTunnelsResponse struct {
	Ztun     map[string][]string
	IPList   map[string]string `json:"ip_list"`
	Protocol map[string]string
	AgentKey string `json:"agent_key"`
}

// NewClient creates new instance of the VPN client.
func NewClient() *Client {
	return &Client{
		uuid: strings.Replace(uuid.New().String(), "-", "", -1),
	}
}

// Initialize method opens new session with the remote API.
func (c *Client) Initialize() error {
	response, err := http.PostForm(baseURI+"/background_init", url.Values{
		"login": []string{"1"},
		"flags": []string{"0"},
		"ver":   []string{extVer},
		"uuid":  []string{c.uuid},
	})
	if err != nil {
		return err
	}

	defer response.Body.Close()
	var r initializeResponse

	err = json.NewDecoder(response.Body).Decode(&r)
	if err != nil {
		return err
	}

	c.key = fmt.Sprintf("%d", r.Key)
	return nil
}

// FindTunnels method returns available proxy servers.
func (c *Client) FindTunnels(countryCode string, limit int) (*Tunnels, error) {
	u := baseURI + "/zgettunnels?" + "uuid=" + c.uuid + "&session_key=" + c.key +
		"&country=" + countryCode + "&rmt_ver=" + rmtVer + "&ext_ver=" + extVer + "&browser=" + browser +
		"&product=cws" + "&lccgi=1" + fmt.Sprintf("&limit=%d", limit)
	response, err := http.Get(u)
	defer response.Body.Close()
	var r zGetTunnelsResponse

	err = json.NewDecoder(response.Body).Decode(&r)
	if err != nil {
		return nil, err
	}

	login := fmt.Sprintf("user-uuid-%s", c.uuid)
	password := r.AgentKey

	tunnels := &Tunnels{
		Servers: []TunnelSettings{},
	}

	proxyEndpoints, ok := r.Ztun[countryCode]
	if ok {
		for _, endpoint := range proxyEndpoints {
			endpointSplit := strings.SplitN(endpoint, " ", 2)
			hostPort := endpointSplit[1]

			hostPortSplit := strings.SplitN(hostPort, ":", 2)
			hostname := hostPortSplit[0]
			port := hostPortSplit[1]

			ipAddress, foundIPAddress := r.IPList[hostname]
			if !foundIPAddress {
				return nil, fmt.Errorf("IP address not found (hostname: %s)", hostname)
			}

			protocol, foundProtocol := r.Protocol[hostname]
			if !foundProtocol {
				return nil, fmt.Errorf("protocol not found (hostname: %s)", hostname)
			}

			tunnels.Servers = append(tunnels.Servers, TunnelSettings{
				Login:    login,
				Password: password,

				Host:  ipAddress,
				Port:  port,
				Proto: protocol,
			})
		}
	}

	return tunnels, nil
}
