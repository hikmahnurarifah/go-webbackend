package controller

import (
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/hikmahnurarifah/webbackend/database"
	"github.com/hikmahnurarifah/webbackend/models"
	"github.com/hikmahnurarifah/webbackend/util"
	"gorm.io/gorm"
)

func CreatePost(c *fiber.Ctx) error {
	var pasienpost models.Pasien
	if err := c.BodyParser(&pasienpost); err != nil {
		fmt.Println("Unable to parse body")
	}
	if err := database.DB.Create(&pasienpost).Error; err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Invalid payload",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Congratulation! Your data is posted",
	})
}

func AllPost(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit := 5
	offset := (page - 1) * limit
	var total int64
	var getpasien []models.Pasien
	database.DB.Offset(offset).Limit(limit).Find(&getpasien)
	database.DB.Model(&models.Pasien{}).Count(&total)
	return c.JSON(fiber.Map{
		"data": getpasien,
		"meta": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": math.Ceil(float64(int(total) / limit)),
		},
	})

}

func DetailPost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var pasienpost models.Pasien
	database.DB.Where("id=?", id).First(&pasienpost)
	return c.JSON(fiber.Map{
		"data": pasienpost,
	})

}

func UpdatePost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	pasien := models.Pasien{
		ID: uint(id),
	}

	if err := c.BodyParser(&pasien); err != nil {
		fmt.Println("Unable to parse body")
	}
	database.DB.Model(&pasien).Updates(pasien)
	return c.JSON(fiber.Map{
		"message": "post updated successfully",
	})

}

func UniquePost(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	id, _ := util.Parsejwt(cookie)
	var pasien []models.Pasien
	database.DB.Model(&pasien).Where("id=?", id).Find(&pasien)

	return c.JSON(pasien)

}

func DeletePost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	pasien := models.Pasien{
		ID: uint(id),
	}
	deleteQuery := database.DB.Delete(&pasien)
	if errors.Is(deleteQuery.Error, gorm.ErrRecordNotFound) {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Opps!, record Not found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "post deleted Succesfully",
	})

}
