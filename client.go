package wink

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/oauth2"
)

func init() {
	oauth2.RegisterBrokenAuthHeaderProvider("https://api.wink.com/")
}

type Client struct {
	config     *oauth2.Config
	httpClient *http.Client
}

func NewClient(clientId, clientSecret string) *Client {
	return &Client{
		config: &oauth2.Config{
			ClientID:     clientId,
			ClientSecret: clientSecret,
			Endpoint: oauth2.Endpoint{
				AuthURL:  "https://api.wink.com/oauth2/auth",
				TokenURL: "https://api.wink.com/oauth2/token",
			},
		},
	}
}

func (c *Client) Authenticate(username, password string) error {
	token, err := c.config.PasswordCredentialsToken(oauth2.NoContext, username, password)
	if err != nil {
		return err
	}

	c.httpClient = c.config.Client(oauth2.NoContext, token)
	return nil
}

func (c *Client) GetDevicesJSON() (string, error) {
	var err error
	resp, err := c.httpClient.Get("https://api.wink.com/users/me/wink_devices")
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)

	var indented bytes.Buffer
	json.Indent(&indented, body, "", "    ")
	return string(indented.Bytes()), nil
}

func (c *Client) GetDevices() ([]interface{}, error) {
	var err error
	resp, err := c.httpClient.Get("https://api.wink.com/users/me/wink_devices")
	if err != nil {
		return nil, err
	}

	var data struct{ Data []map[string]interface{} }
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&data)
	if err != nil {
		return nil, err
	}

	devices := make([]interface{}, len(data.Data))
	for i, d := range data.Data {
		if _, ok := d["light_bulb_id"]; ok {
			devices[i], err = newLightBulb(c, d)
		} else if _, ok := d["thermostat_id"]; ok {
			devices[i], err = newThermostat(c, d)
		} else if _, ok := d["hub_id"]; ok {
			devices[i], err = newHub(c, d)
		} else {
			devices[i] = d
		}
		if err != nil {
			return nil, err
		}

	}
	return devices, nil
}

func (c *Client) GetLightBulbs() ([]*LightBulb, error) {
	devices, err := c.GetDevices()
	if err != nil {
		return nil, err
	}

	var lightBulbs []*LightBulb
	for _, device := range devices {
		if lb, ok := device.(*LightBulb); ok {
			lightBulbs = append(lightBulbs, lb)
		}
	}
	return lightBulbs, nil
}

func (c *Client) GetHubs() ([]*Hub, error) {
	devices, err := c.GetDevices()
	if err != nil {
		return nil, err
	}

	var hubs []*Hub
	for _, device := range devices {
		if hub, ok := device.(*Hub); ok {
			hubs = append(hubs, hub)
		}
	}
	return hubs, nil
}

func (c *Client) GetThermostats() ([]*Thermostat, error) {
	devices, err := c.GetDevices()
	if err != nil {
		return nil, err
	}

	var thermostats []*Thermostat
	for _, device := range devices {
		if t, ok := device.(*Thermostat); ok {
			thermostats = append(thermostats, t)
		}
	}
	return thermostats, nil
}

func (c *Client) updateDesiredState(objectType, objectID string, state map[string]interface{}) error {
	var data struct {
		DesiredState map[string]interface{} `json:"desired_state"`
	}
	data.DesiredState = state

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	fmt.Println(string(jsonData))
	buf := bytes.NewReader(jsonData)
	req, err := http.NewRequest("PUT", "https://api.wink.com/"+objectType+"/"+objectID+"/desired_state", buf)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	} else if resp.StatusCode >= 400 {
		response, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return errors.New(string(response))
	}
	return err
}
