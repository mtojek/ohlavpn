package vpn

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/google/uuid"
)

const (
	baseUri = "https://client.hola.org/client_cgi"

	browser = "chrome"
	extVer  = "1.125.157"
	rmtVer  = "1.2.676"
)

type Client struct {
	uuid string
	key  string
}

type Tunnels struct {
	Login    string
	Password string

	Servers []TunnelSettings
}

type TunnelSettings struct {
	Host  string
	Port  string
	Proto string
}

func (ts *TunnelSettings) String() string {
	return fmt.Sprintf("%s\t%s:%s", ts.Proto, ts.Host, ts.Port)
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

func NewClient() *Client {
	return &Client{
		uuid: strings.Replace(uuid.New().String(), "-", "", -1),
	}
}

func (c *Client) Initialize() error {
	response, err := http.PostForm(baseUri+"/background_init", url.Values{
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

func (c *Client) FindTunnels(countryCode string, limit int) (*Tunnels, error) {
	u := baseUri + "/zgettunnels?" + "uuid=" + c.uuid + "&session_key=" + c.key +
		"&country=" + countryCode + "&rmt_ver=" + rmtVer + "&ext_ver=" + extVer + "&browser=" + browser +
		"&product=cws" + "&lccgi=1" + fmt.Sprintf("&limit=%d", limit)
	response, err := http.Get(u)
	defer response.Body.Close()
	var r zGetTunnelsResponse

	err = json.NewDecoder(response.Body).Decode(&r)
	if err != nil {
		return nil, err
	}

	tunnels := &Tunnels{
		Login:    c.uuid,
		Password: r.AgentKey,
		Servers:  []TunnelSettings{},
	}

	proxyEndpoints, ok := r.Ztun[countryCode]
	if ok {
		for _, endpoint := range proxyEndpoints {
			endpointSplit := strings.SplitN(endpoint, " ", 2)
			hostPort := endpointSplit[1]

			hostPortSplit := strings.SplitN(hostPort, ":", 2)
			hostname := hostPortSplit[0]
			port := hostPortSplit[1]

			ipAddress, foundIpAddress := r.IPList[hostname]
			if !foundIpAddress {
				return nil, fmt.Errorf("IP address not found (hostname: %s)", hostname)
			}

			protocol, foundProtocol := r.Protocol[hostname]
			if !foundProtocol {
				return nil, fmt.Errorf("protocol not found (hostname: %s)", hostname)
			}

			tunnels.Servers = append(tunnels.Servers, TunnelSettings{
				Host:  ipAddress,
				Port:  port,
				Proto: protocol,
			})
		}
	}

	return tunnels, nil
}
