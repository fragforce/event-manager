package webserver

import (
	"fmt"
	"github.com/fragforce/event-manager/pkg/event"
	"github.com/fragforce/event-manager/pkg/filestore"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Start() {
	s := gin.New()
	s.LoadHTMLGlob("templates/html/*/*.html")
	s.Static("/images", "templates/images")
	s.Static("/css", "templates/styles")
	s.Static("/js", "templates/scripts")

	// general pages
	s.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// teams
	s.GET("/teams", event.GetTeams)

	s.GET("/:teamid", event.GetTeam)

	// login
	s.GET("/login", discordRedirect)
	s.GET("/login/discord/callback", discordCallback)

	// event
	s.GET("/:teamid/event/:eventid", event.GetEvent)

	// signup management
	s.GET("/:teamid/event/:eventid/signup/:position", event.GetSignup)
	s.POST("/:teamid/event/:eventid/signup/:position", event.PostSignup)
	s.PATCH("/:teamid/event/:eventid/signup/:position", event.PatchSignup)

	//event management
	s.GET("/:teamid/manage/", event.GetTeamManage)
	s.POST("/:teamid/manage/", event.PostTeamManage)
	s.GET("/:teamid/manage/:eventid", event.GetTeamEvents)
	s.POST("/:teamid/manage/:eventid", event.PostTeamEvent)

	s.Run()
}

func Test() {
	shiftStart, err := time.Parse("Mon Jan 02 3:04 AM 2006 MST", "Fri Nov 18 8:00 AM 2022 EST")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(shiftStart.Format("Mon, Jan 02 2006 3:04:05 PM +MST"))

	shifts := []event.Shift{
		event.Shift{
			ID:     "f403fa5d-fe2c-4424-b0e2-8edab1b51858",
			Type:   "manager",
			Title:  "Manager Shift 1",
			Length: 6,
			Start:  shiftStart.UTC(),
		},
	}
	err = filestore.UpdateConfig("events/shifts.yml", &shifts)
	if err != nil {
		fmt.Println(err)
	}

}
