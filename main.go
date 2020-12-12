package main

import (
	"net/http"

	"./controller"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	"./docs"
)

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @termsOfService http://swagger.io/terms/
func main() {

	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
		
	r := gin.Default()
	c := controller.NewController()

	candidates := r.Group("/"); {
		candidates.GET("/readCandidate", c.ReadCandidate)
		candidates.POST("/createCandidate", c.CreateCandidate)
		candidates.DELETE("/deleteCandidate", c.DeleteCandidate)
		candidates.PUT("/denyCandidate", c.DenyCandidate)
		candidates.PUT("/acceptCandidate", c.AcceptCandidate)
		candidates.GET("/findAssigneeIDByName", c.FindAssigneeIDByName)
		candidates.GET("/findAssigneesCandidates", c.FindAssigneesCandidates)
		candidates.PUT("/arrangeMeeting", c.ArrangeMeeting)
		candidates.PUT("/completeMeeting", c.CompleteMeeting)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/swagger/index.html")
	})
	r.Run(":8080")
}