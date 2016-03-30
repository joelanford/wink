package wink

import (
	"strconv"
	"time"
)

type EpochTime time.Time

func (t *EpochTime) UnmarshalJSON(b []byte) error {
	bs := string(b)
	if bs == "null" {
		t = nil
	} else {
		result, err := strconv.ParseFloat(bs, 64)
		if err != nil {
			return err
		}
		*t = EpochTime(time.Unix(0, int64(result*1000)*int64(time.Millisecond)))
	}
	return nil
}

type Field struct {
	Field      string         `json:"field"`
	Type       string         `json:"type"`
	Mutability string         `json:"mutability"`
	Choices    []string       `json:"choices"`
	Range      [2]interface{} `json:"range"`
}

type Capabilities struct {
	Fields []Field `json:"fields"`
}

type Subscription struct {
	Pubnub struct {
		SubscribeKey string `json:"subscribe_key"`
		Channel      string `json:"channel"`
	}
}

type BaseDevice struct {
	CreatedAt               EpochTime              `json:"created_at"`
	Capabilities            Capabilities           `json:"capabilities"`
	DesiredState            map[string]interface{} `json:"desired_state"`
	DeviceManufacturer      string                 `json:"device_manufacturer"`
	HiddenAt                EpochTime              `json:"hidden_at"`
	LastReading             map[string]interface{} `json:"last_reading"`
	LatLng                  [2]float64             `json:"lat_lng"`
	Locale                  string                 `json:"locale"`
	Location                string                 `json:"location"`
	ManufacturerDeviceID    string                 `json:"manufacturer_device_id"`
	ManufacturerDeviceModel string                 `json:"manufacturer_device_model"`
	ModelName               string                 `json:"model_name"`
	Name                    string                 `json:"name"`
	Subscription            Subscription           `json:"subscription"`
	Triggers                []interface{}          `json:"triggers"`
	Units                   map[string]string      `json:"units"`
	UserIDs                 []string               `json:"user_ids"`
	UPCID                   string                 `json:"upc_id"`
	UPCCode                 string                 `json:"upc_code"`
	UUID                    string                 `json:"uuid"`

	client *Client
}

type GenericDevice struct {
	HubID           string `json:"hub_id"`
	RadioType       string `json:"radio_type"`
	LocalID         string `json:"local_id"`
	LinkedServiceID string `json:"linked_service_id"`
	BaseDevice
}
