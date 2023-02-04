package event

import (
	"fmt"
	"github.com/fragforce/event-manager/pkg/filestore"
	"strings"
)

func LoadEvents() (err error) {
	err = filestore.ReadConfig("events/users.yml", &users)
	if err != nil {
		return err
	}

	err = filestore.ReadConfig("events/teams.yml", &teams)
	if err != nil {
		return err
	}

	for tid, team := range teams {
		var events []Event
		err = filestore.ReadConfig(fmt.Sprintf("events/team/%s/events.yml", team.ID), &events)
		if err != nil {
			return err
		}

		for eid, event := range events {
			for sid, _ := range event.Signups {
				for qid, _ := range event.Signups[sid].Questions {
					event.Signups[sid].Questions[qid].LabelLower =
						strings.ReplaceAll(strings.ToLower(event.Signups[sid].Questions[qid].Label), " ", "_")
				}
			}

			var shifts []Shift
			err = filestore.ReadConfig(fmt.Sprintf("events/team/%s/%s/shifts.yml", team.ID, event.ID), &shifts)
			if err != nil {
				return err
			}
			events[eid].Shifts = shifts
		}

		teams[tid].Events = events

	}

	return
}
