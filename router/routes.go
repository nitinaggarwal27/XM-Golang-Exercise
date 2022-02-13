package router

import (
	"nitinaggarwal27/XM-Golang-Exercise/jwtToken"
	"nitinaggarwal27/XM-Golang-Exercise/service"

	"github.com/gin-gonic/gin"
)

//Cors : handle client origin rules
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

//Routes : create endpoints
func Routes() *gin.Engine {
	//Default : debug
	gin.SetMode("debug")

	r := gin.Default()

	//Applying CORS rule
	r.Use(Cors())

	v1 := r.Group("v1")
	{
		v1.POST("/login", service.Login)
		v1.GET("/company", service.GetCompanies)
		v1.GET("/company/:id", service.GetCompany)
		//setting up middleware for protected apis
		authMiddleware := jwtToken.MwInitializer()
		//Protected resources
		v1.Use(authMiddleware.MiddlewareFunc())
		{
			// add middles to check location of the client
			v1.Use(checkLocation)
			{
				v1.POST("/company", service.CreateCompany)
				v1.PUT("/company/:id", service.UpdateCompany)
				v1.DELETE("/company/:id", service.DeleteCompany)
			}
		}

	}

	return r
}
