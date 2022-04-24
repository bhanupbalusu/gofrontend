package product_details

import (
	"fmt"

	m "github.com/bhanupbalusu/gofrontend/server/model"

	"github.com/gofiber/fiber/v2"
)

type pdl []m.ProductDetails

var isCalled bool

func AddProductsGetPage(c *fiber.Ctx) error {

	// pDetailsList := []*m.ProductDetails{}
	var pd m.ProductDetails
	pdlist := pdl{}
	//c.Accepts("html")

	if len(pdlist) != 0 && isCalled == false {
		isCalled = true
		pd.ProductName = c.FormValue("productname")
		pd.Description = c.FormValue("productdesc")
		pd.BulkQuantity.Volume = c.FormValue("volume")
		pd.BulkQuantity.BQUnits = c.FormValue("unit")
		pd.Price.Amount = c.FormValue("pricevalue")
		pd.Price.Currency = c.FormValue("currency")
		pd.Price.PerUnit = c.FormValue("qtyperunit")
		pd.Price.PUnits = c.FormValue("baseunit")

		pdlist = append(pdlist, pd)
	}

	return c.Render("productdetailspage", pdlist)
}

func AddProductsPostPage(c *fiber.Ctx) error {
	pDetailsList := []m.ProductDetails{}
	var pd m.ProductDetails
	pdlist := pdl{}

	//c.Accepts("html")
	fmt.Println("pdl --> ", pdlist)
	fmt.Println("pDetailsList --> ", pDetailsList)
	//pdlist = pdl
	fmt.Println("pdlist --> ", pdlist)
	pd.ProductName = c.FormValue("productname")
	pd.Description = c.FormValue("productdesc")
	pd.BulkQuantity.Volume = c.FormValue("volume")
	pd.BulkQuantity.BQUnits = c.FormValue("unit")
	pd.Price.Amount = c.FormValue("pricevalue")
	pd.Price.Currency = c.FormValue("currency")
	pd.Price.PerUnit = c.FormValue("qtyperunit")
	pd.Price.PUnits = c.FormValue("baseunit")

	fmt.Println("pdlist --> ", pdlist)
	fmt.Println("pDetailsList --> ", pDetailsList)
	pdlist = append(pdlist, pd)

	//fmt.Println("pdl --> ", pdl)
	fmt.Println("pdlist --> ", pdlist)
	//pdl = append(pdl, pd)
	//fmt.Println("pdl --> ", pdl)
	fmt.Println("pDetailsList --> ", pDetailsList)
	return c.Render("productdetailspage", pdlist)
}

type p struct {
	pd m.ProductDetails
}
type pl struct {
	pdl []p
}
