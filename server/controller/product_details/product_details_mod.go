package product_details

import (
	"fmt"

	m "github.com/bhanupbalusu/gofrontend/server/model"

	"github.com/gofiber/fiber/v2"
)

type pq struct {
	pqVol   string
	pqUnits string
}

type pprice struct {
	pAmt     string
	pCurr    string
	pPerUnit string
	pPUnits  string
}

type pDetails struct {
	pname  string
	pdesc  string
	pq     pq
	pprice pprice
}

var Prods []m.ProdDetails
var Data []m.ProdDetails

func AddProductsModPage(c *fiber.Ctx) error {
	c.Accepts("html")

	fmt.Printf("Begin Prods --> %p\n", &Prods)
	fmt.Printf("Begin Data --> %p\n", &Data)

	pName := c.FormValue("productname")
	pDesc := c.FormValue("productdesc")
	pQVolume := c.FormValue("volume")
	pQUnits := c.FormValue("unit")
	pPAmount := c.FormValue("pricevalue")
	pPCurrency := c.FormValue("currency")
	pPPerUnit := c.FormValue("qtyperunit")
	pPPUnits := c.FormValue("baseunit")

	fmt.Println("Middle --> ", Prods)
	prod := m.ProdDetails{
		PName: pName,
		PDesc: pDesc,
		PQ: m.PQ{
			Vol:     pQVolume,
			PQUnits: pQUnits,
		},
		PPrice: m.PPrice{
			PAmount:   pPAmount,
			PCurrency: pPCurrency,
			PPerUnit:  pPPerUnit,
			PPUnits:   pPPUnits,
		},
	}

	Prods = append(Prods, prod)

	fmt.Println("End Data Before Append --> ", Data)
	// fmt.Println("End Prods[len(Prods)-1] --> ", Prods[len(Prods)-1])

	//Data = append(Data, Prods[len(Prods)-1])

	fmt.Println("End Data After Append --> ", Data)

	fmt.Printf("End Prods --> %p\n", &Prods[len(Prods)-1])
	fmt.Printf("End Data --> %p\n", &Data[len(Prods)-1])
	//fmt.Println("End PData --> ", pdata)
	return c.Render("productdetailspage", Data)
}
