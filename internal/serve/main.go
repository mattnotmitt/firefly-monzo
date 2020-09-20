package serve

import (
	"github.com/gin-gonic/gin"
	"github.com/mattnotmitt/firefly-monzo/internal/api"
)

func Serve() {
	r := gin.Default()

	api.SetRoutes(r)
	r.Run(":8082") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
