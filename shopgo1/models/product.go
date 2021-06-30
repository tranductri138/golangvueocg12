package models

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Price       float64 `json:"price"`
}

func GetProducts(c *fiber.Ctx) error {
	var product []Product
	DB.Find(&product)
	return c.JSON(&product)
}

func GetProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product Product
	DB.Find(&product, id)
	return c.JSON(&product)
}

func SaveProduct(c *fiber.Ctx) error {
	product := new(Product)
	if err := c.BodyParser(product); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	DB.Create(&product)
	return c.JSON(&product)
}

func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product Product
	DB.First(&product, id)
	DB.Delete(&product)
	return c.SendString("Product is deleted!!!")
}

func UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	product := new(Product)
	DB.First(&product, id)
	if err := c.BodyParser(&product); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	DB.Save(&product)
	return c.JSON(&product)
}

// func Paginate(r *http.Request) func(db *gorm.DB) *gorm.DB {
// 	return func(db *gorm.DB) *gorm.DB {
// 		page, _ := strconv.Atoi(r.Query("page"))
// 		if page == 0 {
// 			page = 1
// 		}

// 		pageSize, _ := strconv.Atoi(r.Query("page_size"))
// 		switch {
// 		case pageSize > 100:
// 			pageSize = 100
// 		case pageSize <= 0:
// 			pageSize = 10
// 		}

// 		offset := (page - 1) * pageSize
// 		return db.Offset(offset).Limit(pageSize)
// 	}
// }
