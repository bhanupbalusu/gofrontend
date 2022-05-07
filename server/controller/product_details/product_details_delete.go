package product_details

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func AddGobProductsDeletePage(c *fiber.Ctx) error {
	fmt.Println("------ Inside AddGobProductsDeletePage -------")

	value := c.Context().QueryArgs()
	bval := value.QueryString()

	fmt.Println(string(bval))
	fmt.Println("length of GobOutPDList --- ", len(GobOutPDList))

	temp := GobOutPDList[:0]
	for i, prod := range GobOutPDList {
		fmt.Println("i ---> ", i)
		if GobOutPDList[i].GobItemID.String() != string(bval) {
			temp = append(temp, prod)
		}
	}
	GobOutPDList = temp
	return AddGobProductsGetPage(c)
}
