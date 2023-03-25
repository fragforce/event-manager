package event

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSignup(c *gin.Context) {
	var thisTeam Team
	var thisEvent Event
	var shifts []Shift
	var questions []Question

	// get team info
	for _, team := range teams {
		if team.Hash == c.Param("teamid") {
			thisTeam = team
			break
		}
	}

	// load specific event
	for _, event := range thisTeam.Events {
		if event.Hash == c.Param("eventid") {
			thisEvent = event
			break
		}
	}

	for _, signup := range thisEvent.Signups {
		if signup.Type == c.Param("position") {
			questions = signup.Questions
		}
	}

	// load shifts for this event
	for _, shift := range thisEvent.Shifts {
		if shift.Type == c.Param("position") {
			// if signup shares type then set has game
			for _, signup := range thisEvent.Signups {
				if signup.Type == shift.Type {
					shift.HasGame = signup.HasGame
				}
			}
			shifts = append(shifts, shift)
		}
	}

	c.HTML(http.StatusOK, "signup.html", gin.H{
		"Title":     "Event Signup",
		"eventid":   thisEvent.Hash,
		"team":      thisTeam.Name,
		"teamid":    thisTeam.Hash,
		"event":     thisEvent.Name,
		"position":  c.Param("position"),
		"questions": questions,
		"shifts":    shifts,
	})
}

func PostSignup(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Print(err)
	}
	fmt.Printf(string(body))

	fmt.Println(c.Query("email"))
	fmt.Println(c.Query("Twitch+Username"))
}

func PatchSignup(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Print(err)
	}
	fmt.Printf(string(body))

	fmt.Println(c.Query("email"))
	fmt.Println(c.Query("Twitch+Username"))
}
