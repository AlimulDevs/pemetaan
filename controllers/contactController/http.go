package contactcontroller

import (
	"fmt"
	"golang/constant/constantError"
	"golang/models/dto"
	contactservice "golang/service/contactService"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

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
	// upload image
	file, err := c.FormFile("image")
	if err != nil {
		input.Image = "https://upload.wikimedia.org/wikipedia/commons/thumb/2/2c/Default_pfp.svg/1200px-Default_pfp.svg.png"
		input.Name = c.FormValue("name")
		input.PhoneNumber = c.FormValue("phone_number")

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
	src, err := file.Open()
	if err != nil {
		return err
	}

	filename := "images/contact-images/" + strconv.FormatInt(time.Now().Unix(), 10) + ".png"

	dst, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	//  request body to struct
	basePath, _ := os.Getwd()
	image := fmt.Sprintf(`%s/%s`, basePath, filename)
	input.Image = image
	input.Name = c.FormValue("name")
	input.PhoneNumber = c.FormValue("phone_number")

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

	// Get id from url
	id := c.Param("id")

	response, _ := ctr.Service.GetByID(id)

	basePath, _ := os.Getwd()
	imageDelete := strings.Replace(response.Image, basePath, basePath, 1)

	if response.ID == "" {
		return c.JSON(http.StatusNotFound, echo.Map{
			"message": "data not found",
		})
	}

	// upload image
	file, err := c.FormFile("image")
	if err != nil {

		input.Name = c.FormValue("name")
		input.PhoneNumber = c.FormValue("phone_number")

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
	os.Remove(imageDelete)
	src, err := file.Open()
	if err != nil {
		return err
	}

	filename := "images/contact-images/" + strconv.FormatInt(time.Now().Unix(), 10) + ".png"

	dst, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	//  request body to struct

	image := fmt.Sprintf(`%s/%s`, basePath, filename)
	input.Image = image
	input.Name = c.FormValue("name")
	input.PhoneNumber = c.FormValue("phone_number")

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
	response, _ := ctr.Service.GetByID(id)

	basePath, _ := os.Getwd()
	imageDelete := strings.Replace(response.Image, basePath, basePath, 1)

	if response.ID == "" {
		return c.JSON(http.StatusNotFound, echo.Map{
			"message": "data not found",
		})
	}
	os.Remove(imageDelete)
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
