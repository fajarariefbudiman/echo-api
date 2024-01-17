package router

import (
	"echo-api/controller"
	"echo-api/middleware"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/api", controller.HomeController)
	// e.GET("/set-cookie", controller.SetCookie)

	// Inisialisasi middleware
	loginauth := middleware.AuthLogin
	products := e.Group("/products")
	// products.Use(loginauth)
	products.GET("", controller.GetAllProducts)
	products.GET("/:slug", controller.GetProductsBySlug)
	products.POST("", controller.CreateProducts)
	products.PUT("/:slug", controller.UpdateProducts)
	products.DELETE("/:slug", controller.DeleteProducts)

	categories := e.Group("/categories")
	categories.Use(loginauth)
	categories.GET("", controller.GetAllCategories)
	categories.GET("/:slug", controller.GetCategoriesId)
	categories.POST("", controller.CreateCategories)
	categories.PUT("/:slug", controller.UpdateCategories)
	categories.DELETE("/:id", controller.DeleteCategories)

	e.POST("/register", controller.CreateUsers)
	e.POST("/login", controller.CheckLogin)
	e.POST("/reset", controller.ResetPassword)
	e.POST("/forgot-password", controller.ForgotPassword)
	// e.POST("/logout", controller.Logout)

	users := e.Group("/users")
	// users.Use(loginauth)
	users.GET("", controller.GetUsers)
	users.GET("/:id", controller.GetUsersId)
	users.PUT("/:id", controller.UpdateUsers)
	users.DELETE("/:id", controller.DeleteUsers)
	users.GET("/:userName/addresses", controller.GetAddressesUsers)
	users.POST("/:userName/addresses", controller.AddAddressesUsers)

	orders := e.Group("/orders")
	orders.POST("", controller.AddOrders)
	orders.PUT("/:orderId", controller.UpdateOrders)
	orders.GET("/:orderId", controller.GetOrderById)
	orders.POST("/:orderId/order-items", controller.AddOrderItems)

	carts := e.Group("/carts")
	carts.POST("", controller.AddCarts)
	carts.GET("/:userId", controller.GetCartByUserId)
	carts.POST("/:cartId/items", controller.AddCartItems)
	carts.GET("/:cartId/items", controller.GetCartItems)

	return e
}
