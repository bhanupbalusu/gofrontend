package main

import (
	"log"

	c "github.com/bhanupbalusu/gofrontend/server/controller/product_details"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

type Data struct {
	Items []Devicevalue_view
}

type Devicevalue_view struct {
	Devicetype string
	Iddevice   string
	Devicename string
	Oidname    string
	Value      string
}

func main() {
	engine := html.New("./static_files/html", ".html")
	app := fiber.New(fiber.Config{
		Views: engine, //set as render engine
	})
	app.Static("/", "./static_files")

	app.Get("/homepage", homePage)
	app.Get("/preorderrequestspage", preorderRequestsPage)
	app.Get("/productdetailspage", c.AddProductsGetPage)
	app.Get("/productdetailsentrypage", c.ProductDetailsEntryPage)
	app.Post("/productdetailspage", c.AddProductsModPage)
	app.Get("/scheduleinfopage", scheduleInfoPage)

	log.Fatal(app.Listen(":5000"))
}

func homePage(c *fiber.Ctx) error {

	return c.Render("homepage", nil)
}

// func addProductsPage(c *fiber.Ctx) error {

// 	return c.Render("productdetailspage", nil)
// }

// func productsDetailsEntryPage(c *fiber.Ctx) error {

// 	return c.Render("productsdetailsentrypage", nil)
// }

func scheduleInfoPage(c *fiber.Ctx) error {

	return c.Render("scheduleinfopage", nil)
}

func preorderRequestsPage(c *fiber.Ctx) error {
	data := Data{}
	for i := 1; i < 10; i++ {
		view := Devicevalue_view{
			Devicetype: "devicetype",
			Iddevice:   "iddevice",
			Devicename: "devicename",
			Oidname:    "oidname",
			Value:      "value",
		}

		data.Items = append(data.Items, view)
	}

	return c.Render("preorderrequestspage", data)
}
