package routes

import (
	middlewares "golang/app/middlewares"
	middlewareCostumer "golang/app/middlewares/costumer"
	contactcontroller "golang/controllers/contactController"
	customerController "golang/controllers/customerController"
	newscontroller "golang/controllers/newsController"
	contactrepository "golang/repository/contactRepository"
	"golang/repository/customerRepository"
	newsrepository "golang/repository/newsRepository"
	contactservice "golang/service/contactService"
	"golang/service/costumerService"
	newsservice "golang/service/newsService"

	"golang/helper"

	"golang/util"

	"github.com/go-playground/validator/v10"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func New(db *gorm.DB) *echo.Echo {
	//repository

	customerRepository := customerRepository.NewCustomerRepository(db)
	contactRepository := contactrepository.NewRepo(db)
	newsRepository := newsrepository.NewRepo(db)
	//end repository

	//service

	costumerService := costumerService.NewcostumerService(customerRepository)
	contactService := contactservice.NewService(contactRepository)
	newsService := newsservice.NewService(newsRepository)
	//end service

	//controller

	customerController := customerController.CustomerController{
		CostumerService: costumerService,
	}
	contactController := contactcontroller.Controller{
		Service: contactService,
	}
	newsController := newscontroller.Controller{
		Service: newsService,
	}
	//end controller

	app := echo.New()

	app.Validator = &helper.CustomValidator{
		Validator: validator.New(),
	}
	configLogger := middlewares.ConfigLogger{
		Format: "[${time_rfc3339}] ${status} ${method} ${host} ${path} ${latency_human}" + "\n",
	}
	configCustomer := middleware.JWTConfig{
		Claims:     &middlewareCostumer.JwtCostumerClaims{},
		SigningKey: []byte(util.GetConfig("TOKEN_SECRET")),
	}

	app.Use(configLogger.Init())
	app.Use(middleware.CORS())

	// costumer
	customer := app.Group("/customer")
	customer.POST("/register", customerController.Register)
	customer.POST("/verifikasi", customerController.Verifikasi)
	customer.POST("/login", customerController.Login)

	privateCustomer := app.Group("/customer", middleware.JWTWithConfig(configCustomer))
	privateCustomer.Use(middlewares.CheckTokenMiddlewareCustomer)

	// private costumer access
	privateCustomer.POST("/logout", customerController.Logout)

	// routes contact
	privateCustomer.GET("/contact/get-all", contactController.GetAll)
	privateCustomer.GET("/contact/get-by-id/:id", contactController.GetByID)
	privateCustomer.POST("/contact/create", contactController.Create)
	privateCustomer.PUT("/contact/update/:id", contactController.Update)
	privateCustomer.DELETE("/contact/delete/:id", contactController.Delete)
	//end routes contact

	// routes news
	privateCustomer.GET("/news/get-all", newsController.GetAll)
	privateCustomer.GET("/news/get-by-id/:id", newsController.GetByID)
	privateCustomer.POST("/news/create", newsController.Create)
	privateCustomer.PUT("/news/update/:id", newsController.Update)
	privateCustomer.DELETE("/news/delete/:id", newsController.Delete)
	//end routes news

	// -->

	return app
}
