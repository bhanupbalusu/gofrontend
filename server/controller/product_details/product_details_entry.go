package product_details

import (
	"github.com/gofiber/fiber/v2"
)

func ProductDetailsEntryPage(c *fiber.Ctx) error {
	return c.Render("productdetailsentrypage", nil)
}
