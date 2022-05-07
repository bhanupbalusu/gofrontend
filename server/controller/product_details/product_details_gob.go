package product_details

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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
	GobItemID         uuid.UUID
	GobInProductName  string
	GobInDescription  string
	GobInBulkQuantity GobInBulkQuantity
	GobInPrice        GobInPrice
}

var gobinpd GobInProductDetails
var GobOutPDList []GobInProductDetails

func AddGobProductsGetPage(c *fiber.Ctx) error {
	fmt.Println("Inside AddGobProductsGetPage")
	if len(GobOutPDList) != 0 {
		fmt.Println("length of GobOutPDList inside AddGobProductsGetPage --- ", len(GobOutPDList))
		return c.Render("productdetailspage", GobOutPDList)
	}
	return c.Render("productdetailspage", nil)
}

func AddGobProductsPostPage(c *fiber.Ctx) error {
	value := c.Context().QueryArgs()
	bval := value.QueryString()

	for i := range GobOutPDList {
		fmt.Println("prod.GobItemID Before if - ", GobOutPDList[i].GobItemID.String())
		fmt.Println("bval Before if - ", string(bval))

		if GobOutPDList[i].GobItemID.String() == string(bval) {
			fmt.Println("prod.GobItemID After if - ", GobOutPDList[i].GobItemID.String())
			fmt.Println("bval After if - ", string(bval))
			fmt.Println("productdesc ---> ", c.FormValue("productdesc"))

			pname := c.FormValue("productname")
			pdesc := c.FormValue("productdesc")
			bqvol := c.FormValue("volume")
			bqunits := c.FormValue("unit")
			amt := c.FormValue("pricevalue")
			curr := c.FormValue("currency")
			perunit := c.FormValue("qtyperunit")
			ppunit := c.FormValue("baseunit")

			gobinpd := gobinpd.NewGobInProductDetails(GobOutPDList[i].GobItemID, pname, pdesc, bqvol, bqunits, amt, curr, perunit, ppunit)

			goboutpd := gobinpd.SerializeDeserialize()

			GobOutPDList[i].GobInProductName = goboutpd.GobInProductName
			GobOutPDList[i].GobInDescription = goboutpd.GobInDescription
			GobOutPDList[i].GobInBulkQuantity.GobInVolume = goboutpd.GobInBulkQuantity.GobInVolume
			GobOutPDList[i].GobInBulkQuantity.GobInBQUnits = goboutpd.GobInBulkQuantity.GobInBQUnits
			GobOutPDList[i].GobInPrice.GobInAmount = goboutpd.GobInPrice.GobInAmount
			GobOutPDList[i].GobInPrice.GobInCurrency = goboutpd.GobInPrice.GobInCurrency
			GobOutPDList[i].GobInPrice.GobInPerUnit = goboutpd.GobInPrice.GobInPerUnit
			GobOutPDList[i].GobInPrice.GobInPUnits = goboutpd.GobInPrice.GobInPUnits

			fmt.Println("PDList inside if - ", GobOutPDList)
			return c.Render("productdetailspage", GobOutPDList)
		}
	}

	id := uuid.New()
	pname := c.FormValue("productname")
	pdesc := c.FormValue("productdesc")
	bqvol := c.FormValue("volume")
	bqunits := c.FormValue("unit")
	amt := c.FormValue("pricevalue")
	curr := c.FormValue("currency")
	perunit := c.FormValue("qtyperunit")
	ppunit := c.FormValue("baseunit")

	gobinpd := gobinpd.NewGobInProductDetails(id, pname, pdesc, bqvol, bqunits, amt, curr, perunit, ppunit)

	goboutpd := gobinpd.SerializeDeserialize()

	GobOutPDList = append(GobOutPDList, goboutpd)

	fmt.Println(GobOutPDList)

	return c.Render("productdetailspage", GobOutPDList)
}

func (gobinpd *GobInProductDetails) SerializeDeserialize() GobInProductDetails {
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	dec := gob.NewDecoder(&network)

	err := enc.Encode(gobinpd)
	if err != nil {
		log.Fatal("encode error:", err)
	}

	goboutpd := GobInProductDetails{}
	err = dec.Decode(&goboutpd)
	if err != nil {
		log.Fatal("decode error:", err)
	}

	return goboutpd
}

func (gobinpd *GobInProductDetails) NewGobInProductDetails(id uuid.UUID, pname, pdesc, bqvol, bqunits, amt, curr, perunit, ppunit string) *GobInProductDetails {
	return &GobInProductDetails{
		id,
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
