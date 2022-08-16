package event

import (
	"time"
)

type Organization struct {
	Name    string   `mapstructure:"name,omitempty"`
	Events  []Event  `mapstructure:"events,omitempty"`
	Signups []Signup `mapstructure:"signups,omitempty"`
}

type Event struct {
	Name   string    `mapstructure:"name,omitempty"`
	ID     string    `mapstructure:"id,omitempty"`
	Start  time.Time `mapstructure:"start,omitempty"`
	Length int       `mapstructure:"length,omitempty"`
	Users  []User    `mapstructure:"users,omitempty"`
	Shifts []Shift   `mapstructure:"shifts,omitempty"`
}

type User struct {
	UserID          string `mapstructure:"user_id,omitempty"`
	DiscordID       string `mapstructure:"discord_id,omitempty"`
	SalesforceOrgID string `mapstructure:"salesforce_id,omitempty"`
}

type Shift struct {
	Type       string    `mapstructure:"type,omitempty"`
	ShiftID    string    `mapstructure:"shift_id,omitempty"`
	UserID     string    `mapstructure:"user_id,omitempty"`
	Title      string    `mapstructure:"title,omitempty"`
	HasGame    bool      `mapstructure:"hasGame,omitempty"`
	Game       string    `mapstructure:"game,omitempty"`
	Start      time.Time `mapstructure:"start,omitempty"`
	Length     int       `mapstructure:"length,omitempty"`
	MaxPlayers int       `mapstructure:"max_players,omitempty"`
	Players    []Player  `mapstructure:"players,omitempty"`
}

type Player struct {
	UserID string `mapstructure:"user_id,omitempty"`
}

type Signup struct {
	Type          string     `mapstructure:"type,omitempty"`
	UserID        string     `mapstructure:"user_id,omitempty"`
	Questions     []Question `mapstructure:"questions,omitempty"`
	HasGame       bool       `mapstructure:"has_game,omitempty"`
	Game          string     `mapstructure:"game,omitempty"`
	PlayersNeeded bool       `mapstructure:"players_needed,omitempty"`
	PlayerCount   int        `mapstructure:"player_count,omitempty"`
}

type Question struct {
	Label       string `mapstructure:"label,omitempty"`
	Placeholder string `mapstructure:"placeholder,omitempty"`
	Type        string `mapstructure:"type,omitempty"`
	Required    bool   `mapstructure:"required,omitempty"`
}
