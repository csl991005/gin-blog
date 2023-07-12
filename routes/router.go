package routes

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	v1 "github.com/ginblog/api/v1"
	"github.com/ginblog/middleware"
	"github.com/ginblog/utils"
	"net/http"
)

func createMyRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles("admin", "web/admin/dist/index.html")
	r.AddFromFiles("front", "web/front/dist/index.html")
	return r
}

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	r.Use(middleware.Cors())

	// 允许不安全代理
	_ = r.SetTrustedProxies(nil)

	r.HTMLRender = createMyRender()

	r.Static("/admin", "./web/admin/dist")
	r.Static("/static", "./web/front/dist/static")
	r.StaticFile("/favicon.ico", "/web/front/dist/favicon.ico")

	r.GET("admin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin", nil)
	})
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "front", nil)
	})

	// r.LoadHTMLGlob("static/admin/index.html")
	// r.Static("admin/assets", "static/admin/assets")
	// r.StaticFile("admin/favicon.ico", "static/admin/favicon.ico")
	//
	// r.GET("admin", func(c *gin.Context) {
	// 	c.HTML(200, "index.html", nil)
	// })

	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		// 用户模块的路由接口
		// auth.GET("user/:id", v1.UserExist)
		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)
		// 分类模块的路由接口
		auth.POST("category/add", v1.AddCategory)
		auth.PUT("category/:id", v1.EditCategory)
		auth.DELETE("category/:id", v1.DeleteCategory)
		// 文章模块的路由接口
		auth.POST("article/add", v1.AddArticle)
		auth.PUT("article/:id", v1.EditArticle)
		auth.DELETE("article/:id", v1.DeleteArticle)
		// 上传接口
		auth.POST("upload", v1.UpLoad)
		auth.PUT("profile/:id", v1.UpdateProfile)
	}
	router := r.Group("api/v1")
	{
		router.GET("category/:id", v1.GetCateInfo)
		router.POST("user/add", v1.AddUser)
		router.GET("users", v1.GetUsers)
		router.GET("user/:id", v1.GetUserInfo)
		router.GET("category", v1.GetCategory)
		router.GET("article", v1.GetArticle)
		router.GET("article/list/:id", v1.GetCateArt)
		router.GET("article/info/:id", v1.GetArtInfo)
		router.POST("login", v1.Login)
		router.GET("profile/:id", v1.GetProfile)
	}

	r.Run(utils.HttpPort)
}
