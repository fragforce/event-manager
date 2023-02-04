package event

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func GetTeams(c *gin.Context) {
	c.HTML(http.StatusOK, "teams.html", gin.H{
		"Title": "Teams",
		"Teams": teams,
	})
}

func GetTeam(c *gin.Context) {
	var thisTeam Team
	var currentEvents, futureEvents, pastEvents []Event
	// get team info
	for _, team := range teams {
		if team.Hash == c.Param("teamid") {
			thisTeam = team
			break
		}
	}

	//fmt.Println(thisTeam.Events)

	// load specific event
	for _, event := range thisTeam.Events {
		switch {
		case event.Start.After(time.Now()):
			futureEvents = append(futureEvents, event)
		case event.Start.Before(time.Now()) && event.Start.Add(time.Duration(event.Length)*time.Hour).After(time.Now()):
			currentEvents = append(currentEvents, event)
		case event.Start.Add(time.Duration(event.Length) * time.Hour).Before(time.Now()):
			pastEvents = append(pastEvents, event)
		default:
			continue
		}
	}

	c.HTML(http.StatusOK, "team.html", gin.H{
		"Title":         fmt.Sprintf("%s Events", thisTeam.Name),
		"teamid":        c.Param("teamid"),
		"teamName":      thisTeam.Name,
		"currentEvents": currentEvents,
		"futureEvents":  futureEvents,
		"pastEvents":    pastEvents,
	})
}

func GetTeamEvents(c *gin.Context) {

}
