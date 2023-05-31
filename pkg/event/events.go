package event

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetEvent(c *gin.Context) {
	var thisTeam Team
	var thisEvent Event

	shifts := make(map[string][]Shift)
	signups := make(map[string][]Signup)

	//for _, team := range teams.Get("") {
	//	if team.Hash == c.Param("teamid") {
	//		thisTeam = team
	//		break
	//	}
	//}

	for _, event := range thisTeam.Events {
		if event.Hash == c.Param("eventid") {
			thisEvent = event
			break
		}
	}

	for _, shift := range thisEvent.Shifts {
		shifts[shift.Type] = append(shifts[shift.Type], shift)
	}

	for _, signup := range thisEvent.Signups {
		signups[signup.Type] = append(signups[signup.Type], signup)
	}

	c.HTML(http.StatusOK, "event.html", gin.H{
		"Title":   fmt.Sprintf("Event Schedule"),
		"Team":    thisTeam,
		"Event":   thisEvent,
		"Shifts":  shifts,
		"Signups": signups,
	})
}
