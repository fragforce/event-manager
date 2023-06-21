package webserver

import (
	"net/http"

	"github.com/fragforce/event-manager/pkg/event"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"github.com/spf13/viper"
)

func Start(config viper.Viper) {
	loadDiscordOauth(config)
	event.LoadEvents(config)

	s := gin.New()
	s.LoadHTMLGlob("templates/html/*/*.html")
	s.Static("/images", "templates/images")
	s.Static("/css", "templates/styles")
	s.Static("/js", "templates/scripts")

	store := cookie.NewStore([]byte(viper.GetString("secrets.webserver.session")))
	s.Use(sessions.Sessions("eventSessions", store))

	// general pages
	s.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Title": "Home",
		})
	})

	// teams
	s.GET("/teams", event.GetTeams)
	s.GET("/:teamid", event.GetTeam)

	// login
	s.GET("/login", func(c *gin.Context) {
		discordRedirect(c)
	})
	//s.GET("/login", discordRedirect)
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
