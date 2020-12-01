package router

import (

	"net/http"

	"apidemo/handler/sd"
	"apidemo/handler/user"
	"apidemo/handler/vuedata"
	"apidemo/router/middleware"

	"github.com/gin-gonic/gin"

	_ "apidemo/docs"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middlewares.
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)

	g.Use(mw...)
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	// 文档界面访问URL
	// http://XXXXXXXX/swagger/index.html
	swag := g.Group("/swagger")
	{
		swag.GET("/*any",ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	u := g.Group("/api/v1/user")
	{
		u.POST("/login",user.VeaLogin)
		u.Use(middleware.JWTAuth())
		u.GET("/info",user.VeaLoginInfo)
		u.POST("/logout",user.VeaLogOut)
	}



	vue := g.Group("/api/v1/article")

	{
		vue.Use(middleware.JWTAuth())
		vue.GET("/list",vuedata.QueryVueDataAll)
		vue.POST("/update",vuedata.UpdateVueData)
		vue.POST("/create",vuedata.CreateVueData)
		vue.POST("/delete",vuedata.DeleteVueData)

	}


	svcd := g.Group("/api/v1/student")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/class/list",sd.ClassList)
		svcd.GET("/teacher/list",sd.TeacherList)
		svcd.GET("/list", sd.StudentList)
		//svcd.GET("/query/", sd.QueryStudentSingle)
		svcd.POST("/create", sd.CreateStudent)
		svcd.POST("/update", sd.UpdateStudent)
		svcd.POST("/delete",sd.DeleteStudent)

	}
	//svcd.Use(middleware.JWTAuth())

	return g
}
