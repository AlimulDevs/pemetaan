package routes

import (
	middlewares "golang/app/middlewares"
	middlewareCostumer "golang/app/middlewares/costumer"
	customerController "golang/controllers/customerController"
	"golang/repository/customerRepository"
	"golang/service/costumerService"

	"golang/helper"

	"golang/util"

	"github.com/go-playground/validator/v10"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func New(db *gorm.DB) *echo.Echo {

	customerRepository := customerRepository.NewCustomerRepository(db)

	costumerService := costumerService.NewcostumerService(customerRepository)

	customerController := customerController.CustomerController{
		CostumerService: costumerService,
	}

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

	// -->

	return app
}
