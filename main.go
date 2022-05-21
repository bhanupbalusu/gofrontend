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
	//engine := pug.New("./views", ".pug")
	engine := html.New("./template", ".html")
	app := fiber.New(fiber.Config{
		Views: engine, //set as render engine
	})

	// app.Static("/", "./template", fiber.Static{
	// 	Compress: true,
	// })

	app.Get("/homepage", homePage)
	app.Get("/preorderrequestspage", preorderRequestsPage)
	//app.Get("/productdetailspage", productDetailsGetPage)
	app.Get("/productdetailspage", c.AddGobProductsGetPage)
	app.Post("/productdetailspage", c.AddGobProductsPostPage)
	app.Post("/productdetailsschedulepage", c.AddSchedulePostPage)
	app.Get("/productdetailsschedulepage", c.AddScheduleGetPage)
	//app.Get("/productdetailsdeleteschedulepage", c.DeleteScheduleGetPage)
	app.Get("/productdetailsdeletepage", c.AddGobProductsDeletePage)
	app.Get("/productdetailsentrypage", c.ProductDetailsEntryPage)
	app.Get("/productdetailseditpage", c.ProductDetailsEditPage)
	app.Get("/scheduleinfopage", c.ScheduleInfoPage)
	app.Get("/scheduleinfoeditpage", c.ScheduleInfoEditPage)

	log.Fatal(app.Listen(":5000"))
}

func homePage(c *fiber.Ctx) error {

	return c.Render("homepage", nil)
}

func productDetailsGetPage(c *fiber.Ctx) error {

	return c.Render("productdetailspage", nil)
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
