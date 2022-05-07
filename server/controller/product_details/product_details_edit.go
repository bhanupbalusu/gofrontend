package product_details

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

var PDEdit *GobInProductDetails

func ProductDetailsEditPage(c *fiber.Ctx) error {
	value := c.Context().QueryArgs()

	bval := value.QueryString()
	fmt.Println(value)
	fmt.Println(string(bval))

	for _, prod := range GobOutPDList {
		if (prod.GobItemID).String() == string(bval) {
			PDEdit = PDEdit.EditGobInProductDetails(prod.GobItemID, prod.GobInProductName,
				prod.GobInDescription,
				prod.GobInBulkQuantity.GobInVolume,
				prod.GobInBulkQuantity.GobInBQUnits,
				prod.GobInPrice.GobInAmount,
				prod.GobInPrice.GobInCurrency,
				prod.GobInPrice.GobInPerUnit,
				prod.GobInPrice.GobInPUnits)

			break
		}

	}

	return c.Render("productdetailseditpage", PDEdit)
}

func (PDEdit *GobInProductDetails) EditGobInProductDetails(id uuid.UUID, pname, pdesc, bqvol, bqunits, amt, curr, perunit, ppunit string) *GobInProductDetails {
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
