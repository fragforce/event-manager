package event

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getSignup(c *gin.Context) {
	var thisOrg Organization
	var thisEvent Event
	var thisUser User
	var shifts []Shift
	var questions []Question

	// get org info
	for _, org := range organizations {
		if org.Name == c.Param("org") {
			thisOrg = org
			break
		}
	}

	// load specific event
	for _, event := range thisOrg.Events {
		if event.ID == c.Param("eventid") {
			thisEvent = event
			break
		}
	}

	for _, user := range thisEvent.Users {
		if user.UserID == c.Param("userid") {
			thisUser = user
			break
		}
	}
	
	if thisOrg.Name == "" || thisEvent.Name == "" || thisUser.UserID == "" {
		return
	}

	// load shifts for this event
	for _, shift := range thisEvent.Shifts {
		if shift.Type == c.Param("position") {
			shifts = append(shifts, shift)
			// if signup shares type then set has game
			for _, signup := range thisOrg.Signups {
				if signup.Type == shift.Type {
					shift.HasGame = signup.HasGame
					questions = signup.Questions
				}
			}
		}
	}

	c.HTML(http.StatusOK, "signup.html", gin.H{
		"eventid":   thisEvent.ID,
		"event":     thisEvent.Name,
		"position":  c.Param("position"),
		"userid":    c.Param("userid"),
		"user":      c.Param("user"),
		"questions": questions,
		"shifts":    shifts,
	})

}

func postSignup(c *gin.Context) {
	c.Query("")
}
