package event

import (
	"time"

	"github.com/google/uuid"
)

// FIXME: Add FK Constraints: `constraint:OnUpdate:CASCADE,OnDelete:CASCADE;`

type User struct {
	ID          uuid.UUID `json:"user_id,omitempty" gorm:"type:uuid;default:uuid_generate_v4();not null;primaryKey"`
	DiscordID   string    `json:"discord_id,omitempty" gorm:"uniqueIndex;size:18;not null"`
	DonationURL string    `json:"donation_url,omitempty" gorm:"size:8192"`
	Teams       []*Team   `gorm:"many2many:user_team_membership"`
}

type Team struct {
	ID              uuid.UUID `json:"team_id,omitempty" gorm:"type:uuid;default:uuid_generate_v4();not null;primaryKey"`
	Name            string    `json:"team_name,omitempty" gorm:"uniqueIndex;not null"`
	OwnerID         uuid.UUID `json:"team_owner,omitempty" gorm:"not null"`
	Owner           *User     `gorm:"foreignKey:ID;not null"`
	SalesforceOrgID string    `json:"salesforce_id,omitempty" gorm:"size:18"`
	URL             string    `json:"team_url,omitempty"`
	Events          []*Event  `json:"events,omitempty" gorm:"many2many:team_events;"`
	Members         []*User   `gorm:"many2many:user_team_membership;"`
}

type Event struct {
	ID            string    `json:"event_id,omitempty" gorm:"not null;primaryKey"`
	Finalized     bool      `json:"finalized,omitempty"`
	Name          string    `json:"event_name,omitempty" gorm:"uniqueIndex;not null"`
	Start         time.Time `json:"event_start,omitempty"`
	Length        int       `json:"event_length,omitempty"`
	Shifts        []*Shift  `json:"shifts,omitempty" gorm:"foreignKey:ID"`
	SignupEnabled bool      `json:"signup_enabled,omitempty"`
	Teams         []*Team   `gorm:"many2many:team_events"`
}

type Question struct {
	Label       string `json:"label,omitempty"`
	LabelLower  string
	Placeholder string `json:"placeholder,omitempty"`
	Type        string `json:"type,omitempty"`
	Required    bool   `json:"required,omitempty"`
}

type Signup struct {
	Type          string      `json:"type,omitempty"`
	OwnerID       string      `json:"user_id,omitempty"`
	Questions     []*Question `json:"questions,omitempty"`
	HasGame       bool        `json:"has_game,omitempty"`
	Game          string      `json:"game,omitempty"`
	PlayersNeeded bool        `json:"players_needed,omitempty"`
	PlayerCount   int         `json:"player_count,omitempty"`
}

type Shift struct {
	Type       string      `json:"type,omitempty"`
	ID         string      `json:"shift_id,omitempty"`
	OwnerID    string      `json:"user_id,omitempty"`
	Title      string      `json:"title,omitempty"`
	HasGame    bool        `json:"hasGame,omitempty"`
	Game       string      `json:"game,omitempty"`
	Start      time.Time   `json:"start,omitempty"`
	Length     int         `json:"length,omitempty"`
	MaxPlayers int         `json:"max_players,omitempty"`
	Players    []*User     `json:"players,omitempty"`
	Signups    []*Signup   `json:"signups,omitempty"`
	Questions  []*Question `json:"questions,omitempty"`
}
