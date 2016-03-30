package wink

import "encoding/json"

type thermostatLastReading struct {
	MaxSetPoint                 float64   `json:"max_set_point"`
	MaxSetPointUpdatedAt        EpochTime `json:"max_set_point_updated_at"`
	MaxSetPointChangedAt        EpochTime `json:"max_set_point_changed_at"`
	DesiredMaxSetPointUpdatedAt EpochTime `json:"desired_max_set_point_updated_at"`
	DesiredMaxSetPointChangedAt EpochTime `json:"desired_max_set_point_changed_at"`

	MinSetPoint                 float64   `json:"min_set_point"`
	MinSetPointUpdatedAt        EpochTime `json:"min_set_pont_updated_at"`
	MinSetPointChangedAt        EpochTime `json:"min_set_point_changed_at"`
	DesiredMinSetPointUpdatedAt EpochTime `json:"desired_min_set_point_updated_at"`
	DesiredMinSetPointChangedAt EpochTime `json:"desired_min_set_point_changed_at"`

	Powered                 bool      `json:"powered"`
	PoweredUpdatedAt        EpochTime `json:"powered_updated_at"`
	PoweredChangedAt        EpochTime `json:"powered_changed_at"`
	DesiredPoweredUpdatedAt EpochTime `json:"desired_powered_updated_at"`
	DesiredPoweredChangedAt EpochTime `json:"desired_powered_changed_at"`

	UsersAway                 bool      `json:"users_away"`
	UsersAwayUpdatedAt        EpochTime `json:"users_away_updated_at"`
	UsersAwayChangedAt        EpochTime `json:"users_away_changed_at"`
	DesiredUsersAwayUpdatedAt EpochTime `json:"desired_users_away_updated_at"`
	DesiredUsersAwayChangedAt EpochTime `json:"desired_users_away_changed_at"`

	FanTimerActive                 bool      `json:"fan_timer_active"`
	FanTimerActiveUpdatedAt        EpochTime `json:"fan_timer_active_updated_at"`
	FanTimerActiveChangedAt        EpochTime `json:"fan_timer_active_changed_at"`
	DesiredFanTimerActiveUpdatedAt EpochTime `json:"desired_fan_timer_active_updated_at"`
	//DesiredFanTimerActiveChangedAt EpochTime `json:"desired_fan_timer_active_changed_at"`

	Units          string    `json:"units"`
	UnitsUpdatedAt EpochTime `json:"units_updated_at"`
	//UnitsChangedAt EpochTime `json:"units_changed_at"`

	Temperature          float64   `json:"temperature"`
	TemperatureUpdatedAt EpochTime `json:"temperature_updated_at"`
	TemperatureChangedAt EpochTime `json:"temperature_changed_at"`

	ExternalTemperature          float64   `json:"external_temperature"`
	ExternalTemperatureUpdatedAt EpochTime `json:"external_temperature_updated_at"`
	//ExternalTemperatureChangedAt EpochTime `json:"external_temperature_changed_at"`

	MinMinSetPoint          float64   `json:"min_min_set_point"`
	MinMinSetPointUpdatedAt EpochTime `json:"min_min_set_point_updated_at"`
	//MinMinSetPointChangedAt EpochTime `json:"min_min_set_point_changed_at"`

	MaxMinSetPoint          float64   `json:"max_min_set_point"`
	MaxMinSetPointUpdatedAt EpochTime `json:"max_min_set_point_updated_at"`
	//MaxMinSetPointChangedAt EpochTime `json:"max_min_set_point_changed_at"`

	MinMaxSetPoint          float64   `json:"min_max_set_point"`
	MinMaxSetPointUpdatedAt EpochTime `json:"min_max_set_point_updated_at"`
	//MinMaxSetPointChangedAt EpochTime `json:"min_max_set_point_changed_at"`

	MaxMaxSetPoint          float64   `json:"max_max_set_point"`
	MaxMaxSetPointUpdatedAt EpochTime `json:"max_max_set_point_updated_at"`
	//MaxMaxSetPointChangedAt EpochTime `json:"max_max_set_point_changed_at"`

	Deadband          float64   `json:"deadband"`
	DeadbandUpdatedAt EpochTime `json:"deadband_updated_at"`
	//DeadbandChangedAt EpochTime `json:"deadband_changed_at"`

	EcoTarget          bool      `json:"eco_target"`
	EcoTargetUpdatedAt EpochTime `json:"eco_target_updated_at"`
	EcoTargetChangedAt EpochTime `json:"eco_target_changed_at"`

	ManufacturerStructureID          string    `json:"manufacturer_structure_id"`
	ManufacturerStructureIDUpdatedAt EpochTime `json:"manufacturer_structure_id_updated_at"`
	//ManufacturerStructureIDChangedAt EpochTime `json:"manufacturer_structure_id_changed_at"`

	HasFan          bool      `json:"has_fan"`
	HasFanUpdatedAt EpochTime `json:"has_fan_updated_at"`
	//HasFanChangedAt EpochTime `json:"has_fan_changed_at"`

	FanDuration          int       `json:"fan_duration"`
	FanDurationUpdatedAt EpochTime `json:"fan_duration_updated_at"`
	//FanDurationChangedAt EpochTime `json:"fan_duration_changed_at"`

	LastError          string    `json:"last_error"`
	LastErrorUpdatedAt EpochTime `json:"last_error_updated_at"`
	LastErrorChangedAt EpochTime `json:"last_error_changed_at"`

	Connection          bool      `json:"connection"`
	ConnectionUpdatedAt EpochTime `json:"connection_updated_at"`
	ConnectionChangedAt EpochTime `json:"connection_changed_at"`

	ShortName                 string    `json:"short_name"`
	ShortNameUpdatedAt        EpochTime `json:"short_name_updated_at"`
	ShortNameChangedAt        EpochTime `json:"short_name_changed_at"`
	DesiredShortNameUpdatedAt EpochTime `json:"desired_short_name_updated_at"`
	DesiredShortNameChangedAt EpochTime `json:"desired_short_name_changed_at"`

	Mode                 string    `json:"mode"`
	ModeUpdatedAt        EpochTime `json:"mode_updated_at"`
	ModeChangedAt        EpochTime `json:"mode_changed_at"`
	DesiredModeUpdatedAt EpochTime `json:"desired_mode_updated_at"`
	DesiredModeChangedAt EpochTime `json:"desired_mode_changed_at"`

	ModesAllowed          []string  `json:"modes_allowed"`
	ModesAllowedUpdatedAt EpochTime `json:"modes_allowed_updated_at"`
	//ModesAllowedChangedAt EpochTime `json:"modes_allowed_changed_at"`
}

type thermostatDesiredState struct {
	MaxSetPoint    float64 `json:"max_set_point"`
	MinSetPoint    float64 `json:"min_set_point"`
	Powered        bool    `json:"powered"`
	UsersAway      bool    `json:"users_away"`
	FanTimerActive bool    `json:"fan_timer_active"`
	Mode           string  `json:"mode"`
	ShortName      string  `json:"short_name"`
}

type thermostatCapabilities struct {
	Fields             []Field  `json:"fields"`
	NotificationRobots []string `json:"notification_robots"`
}

type Thermostat struct {
	ThermostatID         string                 `json:"thermostat_id"`
	SmartScheduleEnabled bool                   `json:"smart_schedule_enabled"`
	LastReading          thermostatLastReading  `json:"last_reading"`
	DesiredState         thermostatDesiredState `json:"desired_state"`
	Capabilities         thermostatCapabilities `json:"capabilities"`
	GenericDevice
}

func newThermostat(client *Client, data map[string]interface{}) (*Thermostat, error) {
	var thermostat Thermostat
	buf, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(buf, &thermostat)
	if err != nil {
		return nil, err
	}
	thermostat.client = client
	return &thermostat, nil
}
