package event

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	teams []Team
	users []User
)

func GetEvent(c *gin.Context) {
	var thisTeam Team
	var thisEvent Event
	shifts := make(map[string][]Shift)

	for _, team := range teams {
		if team.Hash == c.Param("teamid") {
			thisTeam = team
			break
		}
	}

	for _, event := range thisTeam.Events {
		if event.Hash == c.Param("eventid") {
			thisEvent = event
			break
		}
	}

	for _, shift := range thisEvent.Shifts {
		shifts[shift.Type] = append(shifts[shift.Type], shift)
	}

	c.HTML(http.StatusOK, "event.html", gin.H{
		"Title":  fmt.Sprintf("Event Schedule"),
		"Team":   thisTeam,
		"Event":  thisEvent,
		"Shifts": shifts,
	})
}
