package wink

import "encoding/json"

type hubLastReading struct {
	PairingMode          string    `json:"pairing_mode"`
	PairingModeUpdatedAt EpochTime `json:"pairing_mode_updated_at"`
	//PairingModeChangedAt        EpochTime `json:"pairing_mode_changed_at"`
	//DesiredPairingMode          string    `json:"desired_pairing_mode"`
	DesiredPairingModeUpdatedAt EpochTime `json:"desired_pairing_mode_updated_at"`

	PairingPrefix          interface{} `json:"pairing_prefix"`
	PairingPrefixUpdatedAt EpochTime   `json:"pairing_prefix_updated_at"`
	//PairingPrefixChangedAt        EpochTime   `json:"pairing_prefix_changed_at"`
	DesiredPairingPrefix          interface{} `json:"desired_pairing_prefix"`
	DesiredPairingPrefixUpdatedAt EpochTime   `json:"desired_pairing_prefix_updated_at"`

	PairingModeDuration          int       `json:"pairing_mode_duration"`
	PairingModeDurationUpdatedAt EpochTime `json:"pairing_mode_duration_updated_at"`
	//PairingModeDurationChangedAt        EpochTime `json:"pairing_mode_duration_changed_at"`
	//DesiredPairingModeDuration          int       `json:"desired_pairing_mode_duration"`
	DesiredPairingModeDurationUpdatedAt EpochTime `json:"desired_pairing_mode_duration_updated_at"`

	//KiddeRadioCode                 uint8     `json:"kidde_radio_code"`
	KiddeRadioCodeUpdatedAt EpochTime `json:"kidde_radio_code_updated_at"`
	//KiddeRadioCodeChangedAt        EpochTime `json:"kidde_radio_code_changed_at"`
	//DesiredKiddeRadioCode          uint8     `json:"desired_kidde_radio_code"`
	DesiredKiddeRadioCodeUpdatedAt EpochTime `json:"desired_pairing_mode_duration_updated_at"`

	Connection          bool      `json:"connection"`
	ConnectionUpdatedAt EpochTime `json:"connection_updated_at"`

	AgentSessionID          string    `json:"agent_session_id"`
	AgentSessionIDUpdatedAt EpochTime `json:"agent_session_id_updated_at"`
	AgentSessionIDChangedAt EpochTime `json:"agent_session_id_changed_at"`

	UpdatingFirmware          bool      `json:"updating_firmware"`
	UpdatingFirmwareUpdatedAt EpochTime `json:"updating_firmware_updated_at"`

	FirmwareVersion          string    `json:"firmware_version"`
	FirmwareVersionUpdatedAt EpochTime `json:"firmware_version_updated_at"`

	UpdateNeeded          bool      `json:"update_needed"`
	UpdateNeededUpdatedAt EpochTime `json:"update_needed_updated_at"`

	MacAddress          string    `json:"mac_address"`
	MacAddressUpdatedAt EpochTime `json:"mac_address_updated_at"`

	IPAddress          string    `json:"ip_address"`
	IPAddressUpdatedAt EpochTime `json:"ip_address_updated_at"`

	HubVersion          string    `json:"hub_version"`
	HubVersionUpdatedAt EpochTime `json:"hub_version_updated_at"`

	RemotePairable          bool      `json:"remote_pairable"`
	RemotePairableUpdatedAt EpochTime `json:"remote_pairable_updated_at"`

	LocalControlPublicKeyHash          string    `json:"local_control_public_key_hash"`
	LocalControlPublicKeyHashUpdatedAt EpochTime `json:"local_control_public_key_hash_updated_at"`

	LocalControlID          string    `json:"local_control_id"`
	LocalControlIDUpdatedAt EpochTime `json:"local_control_id_updated_at"`
}

type hubDesiredState struct {
	PairingMode         string      `json:"pairing_mode"`
	PairingPrefix       interface{} `json:"pairing_prefix"`
	PairingModeDuration int         `json:"pairing_mode_duration"`
}

type hubCapabilities struct {
	OAuth2Clients       []string `json:"oauth2_clients"`
	HomeSercurityDevice bool     `json:"home_security_device"`
}

type hubConfiguration struct {
	KiddeRadioCode int `json:"kidde_radio_code"`
}

type Hub struct {
	// specific to hub
	HubID         string           `json:"hub_id"`
	LastReading   hubLastReading   `json:"last_reading"`
	DesiredState  hubDesiredState  `json:"desired_state"`
	Capabilities  hubCapabilities  `json:"capabilities"`
	Configuration hubConfiguration `json:"configuration"`
	UpdateNeeded  bool             `json:"updated_needed"`

	// base fields that all devices have
	BaseDevice
}

func newHub(client *Client, data map[string]interface{}) (*Hub, error) {
	var hub Hub
	buf, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(buf, &hub)
	if err != nil {
		return nil, err
	}
	hub.client = client
	return &hub, nil
}
