package event

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	organizations []Organization
)

func loadEvents() {
	log.SetLevel(log.DebugLevel)
	orgs := viper.New()
	orgs.SetConfigName("organizations")
	orgs.SetConfigType("yaml")
	orgs.AddConfigPath("./events/")
	err := orgs.ReadInConfig() // Find and read the config file
	if err != nil {            // Handle errors reading the config file
		log.Error(err)
	}

	log.Debugf("found orgs %v", orgs.GetStringSlice("organizations"))

	// load config for orgs
	for _, orgName := range orgs.GetStringSlice("organizations") {
		var orgEvents []Event
		var orgSignups []Signup
		log.Debugf("loading org %s", orgName)
		org := viper.New()
		org.SetConfigName("org")
		org.SetConfigType("yaml")
		org.AddConfigPath(fmt.Sprintf("./events/%s/", orgName))
		err := org.ReadInConfig() // Find and read the config file
		if err != nil {           // Handle errors reading the config file
			log.Error(err)
		}

		log.Debugf("found events %v", org.GetStringSlice("events"))

		// load events for orgs
		for _, eventID := range org.GetStringSlice("events") {
			var eventShifts []Shift
			var eventUsers []User

			log.Debugf("loading org event %s", eventID)
			event := viper.New()
			event.SetConfigName("event")
			event.SetConfigType("yaml")
			event.AddConfigPath(fmt.Sprintf("./events/%s/%s/", orgName, eventID))
			err := event.ReadInConfig() // Find and read the config file
			if err != nil {             // Handle errors reading the config file
				log.Panic("fatal error config file:", err)
			}

			eventStartTime, err := time.Parse("Mon Jan _2 15:04 PM 2006", event.GetString("start"))
			if err != nil { // Handle errors reading the config file
				log.Error(err)
			}

			log.Debugf("event users loading")
			users := viper.New()
			users.SetConfigName("users")
			users.SetConfigType("yaml")
			users.AddConfigPath(fmt.Sprintf("./events/%s/%s/", orgName, eventID))
			err = users.ReadInConfig() // Find and read the config file
			if err != nil {            // Handle errors reading the config file
				log.Panic("fatal error config file:", err)
			}

			err = users.UnmarshalKey("users", &eventUsers)
			if err != nil { // Handle errors reading the config file
				log.Panic("fatal error marshalling ", err)
			}
			log.Debugf("event users loaded")

			log.Debugf("event shifts loading")
			shifts := viper.New()
			shifts.SetConfigName("shifts")
			shifts.SetConfigType("yaml")
			shifts.AddConfigPath(fmt.Sprintf("./events/%s/%s/", orgName, eventID))
			err = shifts.ReadInConfig() // Find and read the config file
			if err != nil {             // Handle errors reading the config file
				log.Panic("fatal error config file:", err)
			}

			var newShifts []struct {
				Type       string   `mapstructure:"type,omitempty"`
				ShiftID    string   `mapstructure:"shift_id,omitempty"`
				UserID     string   `mapstructure:"user_id,omitempty"`
				Title      string   `mapstructure:"title,omitempty"`
				HasGame    bool     `mapstructure:"hasGame,omitempty"`
				Game       string   `mapstructure:"game,omitempty"`
				Start      string   `mapstructure:"start,omitempty"`
				Length     int      `mapstructure:"length,omitempty"`
				MaxPlayers int      `mapstructure:"max_players,omitempty"`
				Players    []Player `mapstructure:"players,omitempty"`
			}

			err = shifts.UnmarshalKey("shifts", &newShifts)
			if err != nil { // Handle errors reading the config file
				log.Panic("fatal error marshalling ", err)
			}
			log.Debugf("event shifts loaded")

			for _, shift := range newShifts {
				log.Debugf("%+v", shift)

				shiftStartTime, err := time.Parse("Mon Jan _2 15:04 PM 2006", shift.Start)
				if err != nil { // Handle errors reading the config file
					log.Error(err)
				}
				eventShifts = append(eventShifts, Shift{
					Type:       shift.Type,
					ShiftID:    shift.ShiftID,
					UserID:     shift.UserID,
					Title:      shift.Title,
					Game:       shift.Game,
					Start:      shiftStartTime,
					Length:     shift.Length,
					MaxPlayers: shift.MaxPlayers,
					Players:    shift.Players,
				})
			}

			orgEvents = append(orgEvents, Event{
				Name:   event.GetString("name"),
				ID:     eventID,
				Start:  eventStartTime,
				Length: event.GetInt("length"),
				Users:  eventUsers,
				Shifts: eventShifts,
			})
		}

		err = org.UnmarshalKey("signups", &orgSignups)
		if err != nil { // Handle errors reading the config file
			log.Panic("fatal error marshalling ", err)
		}

		organizations = append(organizations, Organization{
			Name:    orgName,
			Events:  orgEvents,
			Signups: orgSignups,
		})
	}
}
