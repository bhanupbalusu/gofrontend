package product_details

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type GobInBulkQuantity struct {
	GobInVolume  string
	GobInBQUnits string
}

type GobInPrice struct {
	GobInAmount   string
	GobInCurrency string
	GobInPerUnit  string
	GobInPUnits   string
}

type GobInProductDetails struct {
	GobInProductName  string
	GobInDescription  string
	GobInBulkQuantity GobInBulkQuantity
	GobInPrice        GobInPrice
}

var gobinpd GobInProductDetails
var GobOutPDList []GobInProductDetails

func AddGobProductsGetPage(c *fiber.Ctx) error {
	fmt.Println("Inside AddGobProductsGetPage")
	return c.Render("productdetailspage", nil)
}

func AddGobProductsPostPage(c *fiber.Ctx) error {
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	dec := gob.NewDecoder(&network)

	pname := c.FormValue("productname")
	pdesc := c.FormValue("productdesc")
	bqvol := c.FormValue("volume")
	bqunits := c.FormValue("unit")
	amt := c.FormValue("pricevalue")
	curr := c.FormValue("currency")
	perunit := c.FormValue("qtyperunit")
	ppunit := c.FormValue("baseunit")

	gobinpd := gobinpd.NewGobInProductDetails(pname, pdesc, bqvol, bqunits, amt, curr, perunit, ppunit)

	err := enc.Encode(gobinpd)
	if err != nil {
		log.Fatal("encode error:", err)
	}

	goboutpd := GobInProductDetails{}
	err = dec.Decode(&goboutpd)
	if err != nil {
		log.Fatal("decode error:", err)
	}

	GobOutPDList = append(GobOutPDList, goboutpd)

	fmt.Println(GobOutPDList)

	return c.Render("productdetailspage", GobOutPDList)
}

func (gobinpd *GobInProductDetails) NewGobInProductDetails(pname, pdesc, bqvol, bqunits, amt, curr, perunit, ppunit string) *GobInProductDetails {
	return &GobInProductDetails{
		pname,
		pdesc,
		GobInBulkQuantity{
			bqvol,
			bqunits,
		},
		GobInPrice{
			amt,
			curr,
			perunit,
			ppunit,
		},
	}
}
