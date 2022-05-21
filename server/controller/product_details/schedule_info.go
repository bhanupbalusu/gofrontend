package product_details

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Schedular struct {
	StartDate string
	EndDate   string
}

type ScheduleInfo struct {
	PDList    []GobInProductDetails
	Schedular Schedular
}

var SchInfo ScheduleInfo
var Sch Schedular

func ScheduleInfoPage(c *fiber.Ctx) error {
	if len(GobOutPDList) == 0 {
		return c.Render("productdetailspage", GobOutPDList)
	}
	return c.Render("scheduleinfopage", nil)
}

func AddScheduleGetPage(c *fiber.Ctx) error {
	fmt.Println("Inside AddScheduleGetPage ScheduleInfo --> ", SchInfo)
	return c.Render("productdetailsschedulepage", SchInfo)
}

// func DeleteScheduleGetPage(c *fiber.Ctx) error {
// 	SchInfo.Schedular.StartDate = ""
// 	SchInfo.Schedular.EndDate = ""
// 	return c.Render("productdetailspage", SchInfo.PDList)
// }

func AddSchedulePostPage(c *fiber.Ctx) error {
	fmt.Println("Inside AddSchedulePostPage GobOutPDList --> ", GobOutPDList)
	SchInfo.PDList = GobOutPDList
	fmt.Println("Inside AddSchedulePostPage SchInfo.PDList --> ", SchInfo.PDList)
	startdate := c.FormValue("startdate")
	enddate := c.FormValue("enddate")

	sch := Sch.NewSchedular(string(startdate), string(enddate))
	sch.SchSerializeDeserialize()
	SchInfo.Schedular.StartDate = sch.StartDate
	SchInfo.Schedular.EndDate = sch.EndDate
	if SchInfo.Schedular.StartDate == "" || SchInfo.Schedular.EndDate == "" {
		return c.Render("scheduleinfopage", SchInfo.Schedular)
	}
	layout := "2006-01-02T15:04:05.000Z"
	t1, err := time.Parse(layout, SchInfo.Schedular.StartDate+":05.000Z")
	if err != nil {
		fmt.Println(err)
	}
	t2, err := time.Parse(layout, SchInfo.Schedular.EndDate+":05.000Z")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("t1 ---> ", t1)
	fmt.Println("t2 ---> ", t2)
	if t1.After(t2) {
		return c.Render("scheduleinfopage", nil)
	}

	fmt.Println("Inside AddSchedulePostPage SchInfo --> ", SchInfo)
	return c.Render("productdetailsschedulepage", SchInfo)
}

func (sch *Schedular) NewSchedular(sdate, edate string) *Schedular {
	return &Schedular{
		sdate,
		edate,
	}
}

func (sch *Schedular) SchSerializeDeserialize() Schedular {
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	dec := gob.NewDecoder(&network)

	err := enc.Encode(sch)
	if err != nil {
		log.Fatal("encode error:", err)
	}

	schout := Schedular{}
	err = dec.Decode(&schout)
	if err != nil {
		log.Fatal("decode error:", err)
	}
	return schout
}
