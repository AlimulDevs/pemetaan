package contactcontroller

import (
	"golang/constant/constantError"
	"golang/models/dto"
	contactservice "golang/service/contactService"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	Service contactservice.Service
}

func (ctr *Controller) GetAll(c echo.Context) error {
	response, err := ctr.Service.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail get all contact",
			"error":   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get all contact",
		"contact": response,
	})
}

func (ctr *Controller) GetByID(c echo.Context) error {
	id := c.Param("id")
	response, err := ctr.Service.GetByID(id)
	if err != nil {
		if _, ok := constantError.ErrorCode[err.Error()]; ok {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "fail get contact by id",
				"error":   err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail get by id contact",
			"error":   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get contact by id",
		"contact": response,
	})
}

func (ctr *Controller) Create(c echo.Context) error {
	var input dto.ContactTransaction

	// Binding request body to struct
	err := c.Bind(&input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail bind data",
			"error":   err.Error(),
		})
	}
	// Validate request body
	if err = c.Validate(input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "There is an empty field",
			"error":   err.Error(),
		})
	}

	// Call service to create contact
	err = ctr.Service.Create(input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail create contact",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success create contact",
	})
}

func (ctr *Controller) Update(c echo.Context) error {
	var input dto.ContactTransaction
	// Binding request body to struct
	err := c.Bind(&input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail bind data",
			"error":   err.Error(),
		})
	}

	// Get id from url
	id := c.Param("id")

	// Call service to update contact
	err = ctr.Service.Update(id, input)
	if err != nil {
		if _, ok := constantError.ErrorCode[err.Error()]; ok {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "fail update contact",
				"error":   err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail update contact",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success update contact",
	})
}

func (ctr *Controller) Delete(c echo.Context) error {
	// Get id from url
	id := c.Param("id")

	// Call service to delete contact
	err := ctr.Service.Delete(id)
	if err != nil {
		if _, ok := constantError.ErrorCode[err.Error()]; ok {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "fail delete contact",
				"error":   err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail delete contact",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success delete contact",
	})
}
