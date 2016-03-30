package wink

import "encoding/json"

type lightBulbLastReading struct {
	Brightness                 float64   `json:"brightness"`
	BrightnessChangedAt        EpochTime `json:"brightness_changed_at"`
	BrightnessUpdatedAt        EpochTime `json:"brightness_updated_at"`
	DesiredBrightnessChangedAt EpochTime `json:"desired_brightness_changed_at"`
	DesiredBrightnessUpdatedAt EpochTime `json:"desired_brightness_updated_at"`

	Powered                 bool      `json:"powered"`
	PoweredChangedAt        EpochTime `json:"powered_changed_at"`
	PoweredUpdatedAt        EpochTime `json:"powered_updated_at"`
	DesiredPoweredChangedAt EpochTime `json:"desired_powered_changed_at"`
	DesiredPoweredUpdatedAt EpochTime `json:"desired_powered_updated_at"`

	Connection          bool      `json:"connection"`
	ConnectionUpdatedAt EpochTime `json:"connection_updated_at"`
	ConnectionChangedAt EpochTime `json:"connection_changed_at"`

	FirmwareVersion          string    `json:"firmware_version"`
	FirmwareVersionUpdatedAt EpochTime `json:"firmware_version_updated_at`

	FirmwareDateCode          string    `json:"firmware_date_code"`
	FirmwareDateCodeUpdatedAt EpochTime `json:"firmware_date_code_updated_at`
}

type lightBulbDesiredState struct {
	Brightness float64 `json:"brightness"`
	Powered    bool    `json:"powered"`
}

type LightBulb struct {
	// specific to light bulb
	LightBulbID  string                `json:"light_bulb_id"`
	GangID       string                `json:"gang_id"`
	Order        int                   `json:"order"`
	LastReading  lightBulbLastReading  `json:"last_reading"`
	DesiredState lightBulbDesiredState `json:"desired_state"`

	// base fields that all non-hub devices have
	GenericDevice
}

func newLightBulb(client *Client, data map[string]interface{}) (*LightBulb, error) {
	var lightBulb LightBulb
	buf, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(buf, &lightBulb)
	if err != nil {
		return nil, err
	}
	lightBulb.client = client
	return &lightBulb, nil
}

func (d *LightBulb) SetBrightness(brightness float64) error {
	state := map[string]interface{}{"brightness": brightness}
	if brightness == 0.0 {
		state["powered"] = false
	}
	return d.client.updateDesiredState("light_bulbs", d.LightBulbID, state)
}

func (d *LightBulb) SetPowered(powered bool) error {
	state := map[string]interface{}{"powered": powered}
	return d.client.updateDesiredState("light_bulbs", d.LightBulbID, state)
}
