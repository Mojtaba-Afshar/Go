package routes

import (
	"github.com/Mojtaba-Afshar/ekart/controllers"
	"github.com/gin-gonic/gin"
)

func AppRoutes(r *gin.Engine) {
	//User Routes
	user := r.Group("/api/user")
	user.POST("/register", controllers.HandleRegister)
	user.POST("/login", controllers.HandleLogin)
	user.GET("/logout", controllers.HandleLogout)
	user.GET("/profile/:id", controllers.HandleProfile)
	//product Routes
	product := r.Group("/api/product")
	product.POST("/add", controllers.HandleRegister)
	product.POST("/edit", controllers.HandleLogin)
	product.GET("/search", controllers.HandleLogout)
	product.GET("/product/:id", controllers.HandleProfile)

	//Cart routes
	//Shipping routes
	//reporting routes
}
