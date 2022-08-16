package event

import (
	"github.com/gin-gonic/gin"
)

func manageEvents(c *gin.Context) {
	c.Param("org")
}

func manageEvent(c *gin.Context) {

}
