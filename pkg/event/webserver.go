package event

import "github.com/gin-gonic/gin"

func Start() {
	loadEvents()

	s := gin.New()
	s.LoadHTMLGlob("templates/html/*.html")
	s.Static("/css", "templates/styles")

	//event management
	s.GET("/:org/event/", getSchedules)
	s.GET("/:org/event/:eventid", getSchedule)

	// signup management
	s.GET("/:org/event/:eventid/signup/:position/:userid", getSignup)
	s.POST("/:org/event/:eventid/signup/:position/:userid", postSignup)

	//event management
	s.GET("/:org/event/manage/", manageEvents)
	s.GET("/:org/event/manage/:event", manageEvent)

	s.Run()
}
