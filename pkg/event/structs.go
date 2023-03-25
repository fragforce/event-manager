package event

import (
	"time"
)

type User struct {
	ID              string `json:"user_id,omitempty"`
	Hash            string `json:"hash,omitempty"`
	DiscordID       string `json:"discord_id,omitempty"`
	SalesforceOrgID string `json:"salesforce_id,omitempty"`
	DonationURL     string `json:"donation_url,omitempty"`
}

type Team struct {
	ID     string  `json:"team_id,omitempty"`
	Hash   string  `json:"team_hash,omitempty"`
	Name   string  `json:"team_name,omitempty"`
	Owner  string  `json:"team_owner,omitempty"`
	URL    string  `json:"team_url,omitempty"`
	Events []Event `json:"events,omitempty"`
}

type Question struct {
	Label       string `json:"label,omitempty"`
	LabelLower  string
	Placeholder string `json:"placeholder,omitempty"`
	Type        string `json:"type,omitempty"`
	Required    bool   `json:"required,omitempty"`
}

type Signup struct {
	Type          string     `json:"type,omitempty"`
	UserID        string     `json:"user_id,omitempty"`
	Questions     []Question `json:"questions,omitempty"`
	HasGame       bool       `json:"has_game,omitempty"`
	Game          string     `json:"game,omitempty"`
	PlayersNeeded bool       `json:"players_needed,omitempty"`
	PlayerCount   int        `json:"player_count,omitempty"`
}

type Shift struct {
	Type       string     `json:"type,omitempty"`
	ID         string     `json:"shift_id,omitempty"`
	Hash       string     `json:"hash,omitempty"`
	UserID     string     `json:"user_id,omitempty"`
	Title      string     `json:"title,omitempty"`
	HasGame    bool       `json:"hasGame,omitempty"`
	Game       string     `json:"game,omitempty"`
	Start      time.Time  `json:"start,omitempty"`
	Length     int        `json:"length,omitempty"`
	MaxPlayers int        `json:"max_players,omitempty"`
	Players    []User     `json:"players,omitempty"`
	Signups    []Signup   `json:"signups,omitempty"`
	Questions  []Question `json:"questions,omitempty"`
}

type Event struct {
	ID            string    `json:"event_id,omitempty"`
	Hash          string    `json:"event_hash,omitempty"`
	Finalized     bool      `json:"finalized,omitempty"`
	Name          string    `json:"event_name,omitempty"`
	Start         time.Time `json:"event_start,omitempty"`
	Length        int       `json:"event_length,omitempty"`
	Participants  []User    `json:"participants,omitempty"`
	Shifts        []Shift   `json:"shifts,omitempty"`
	SignupEnabled bool      `json:"signup_enabled,omitempty"`
	Signups       []Signup
}
