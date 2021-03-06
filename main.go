package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	"./docs"
	"./repository"
	"./api"
)

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @termsOfService http://swagger.io/terms/
func main() {

	docs.SwaggerInfo.Title = "BS Candidate Manager API"
	docs.SwaggerInfo.Description = "This is a simple boilerplate of credential management API written in Go"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
		
	r := gin.Default()

	repository.Register()

	candidates := r.Group("/"); {
		candidates.GET("/readCandidate", api.ReadCandidate)
		candidates.POST("/createCandidate", api.CreateCandidate)
		candidates.DELETE("/deleteCandidate", api.DeleteCandidate)
		candidates.PUT("/denyCandidate", api.DenyCandidate)
		candidates.PUT("/acceptCandidate", api.AcceptCandidate)
		candidates.GET("/findAssigneeIDByName", api.FindAssigneeIDByName)
		candidates.GET("/findAssigneesCandidates", api.FindAssigneesCandidates)
		candidates.PUT("/arrangeMeeting", api.ArrangeMeeting)
		candidates.PUT("/completeMeeting", api.CompleteMeeting)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/swagger/index.html")
	})
	r.Run(":8080")
}