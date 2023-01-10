package newscontroller

import (
	"fmt"
	"golang/constant/constantError"
	"golang/models/dto"
	newsservice "golang/service/newsService"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	Service newsservice.Service
}

func (ctr *Controller) GetAll(c echo.Context) error {
	response, err := ctr.Service.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail get all news",
			"error":   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get all news",
		"news":    response,
	})
}

func (ctr *Controller) GetByID(c echo.Context) error {
	id := c.Param("id")
	response, err := ctr.Service.GetByID(id)
	if err != nil {
		if _, ok := constantError.ErrorCode[err.Error()]; ok {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "fail get news by id",
				"error":   err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail get by id news",
			"error":   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get news by id",
		"news":    response,
	})
}

func (ctr *Controller) Create(c echo.Context) error {

	var input dto.NewsTransaction
	// upload image
	file, err := c.FormFile("image")
	if err != nil {
		input.Image = "https://icon2.cleanpng.com/20180605/ijl/kisspng-computer-icons-image-file-formats-no-image-5b16ff0d2414b5.0787389815282337411478.jpg"
		input.Title = c.FormValue("title")
		input.Description = c.FormValue("description")

		// Validate request body
		if err = c.Validate(input); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "There is an empty field",
				"error":   err.Error(),
			})
		}

		// Call service to create news
		err = ctr.Service.Create(input)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "fail create news",
				"error":   err.Error(),
			})
		}

		// Return response if success
		return c.JSON(http.StatusOK, echo.Map{
			"message": "success create news",
		})
	}
	src, err := file.Open()
	if err != nil {
		return err
	}

	filename := "images/news-images/" + strconv.FormatInt(time.Now().Unix(), 10) + ".png"

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
	input.Title = c.FormValue("title")
	input.Description = c.FormValue("description")

	// Validate request body
	if err = c.Validate(input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "There is an empty field",
			"error":   err.Error(),
		})
	}

	// Call service to create news
	err = ctr.Service.Create(input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail create news",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success create news",
	})
}

func (ctr *Controller) Update(c echo.Context) error {
	var input dto.NewsTransaction

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

		input.Title = c.FormValue("title")
		input.Description = c.FormValue("description")

		err = ctr.Service.Update(id, input)
		if err != nil {
			if _, ok := constantError.ErrorCode[err.Error()]; ok {
				return c.JSON(http.StatusInternalServerError, echo.Map{
					"message": "fail update news",
					"error":   err.Error(),
				})
			}
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "fail update news",
				"error":   err.Error(),
			})
		}

		// Return response if success
		return c.JSON(http.StatusOK, echo.Map{
			"message": "success update news",
		})
	}
	os.Remove(imageDelete)
	src, err := file.Open()
	if err != nil {
		return err
	}

	filename := "images/news-images/" + strconv.FormatInt(time.Now().Unix(), 10) + ".png"

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
	input.Title = c.FormValue("title")
	input.Description = c.FormValue("description")

	// Call service to update news
	err = ctr.Service.Update(id, input)
	if err != nil {
		if _, ok := constantError.ErrorCode[err.Error()]; ok {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "fail update news",
				"error":   err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail update news",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success update news",
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
	// Call service to delete news
	err := ctr.Service.Delete(id)
	if err != nil {
		if _, ok := constantError.ErrorCode[err.Error()]; ok {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "fail delete news",
				"error":   err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail delete news",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success delete news",
	})
}
