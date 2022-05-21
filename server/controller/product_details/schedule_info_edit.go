package product_details

import (
	"github.com/gofiber/fiber/v2"
)

var SchEdit *ScheduleInfo

func ScheduleInfoEditPage(c *fiber.Ctx) error {
	SchEdit = &SchInfo
	return c.Render("scheduleinfoeditpage", SchEdit)
}
