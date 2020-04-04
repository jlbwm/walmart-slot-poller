package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
	"github.com/parnurzeal/gorequest"
)


type slot struct {
	StartDateTime string
	EndDateTime   string
	Status        string
}

type slotDay struct {
	Slots []slot
}

type response struct {
	SlotDays []slotDay `json:"slotDays"`
}

func sendmail(message string) {
	fmt.Println("sending email " + message)
	Send_Gridmail(message)
}

func checkAvailability(tt time.Time) {
	today := time.Now()
	sevenDayLater := today.AddDate(0, 0, 7)
	layout := "2006-01-02"
	t := sevenDayLater.Format(layout)
	request := gorequest.New()
	_, body, errs := request.Post("https://www.walmart.ca/api/cart-page/accesspoints/slotavailability?grocery=true").
		Set("Content-Type", "application/json").
		Send(`{"startDate": "` + t + `","endDate": "` + t + `","accessPointId": "04d2977e-f1be-4d97-86da-a0c3d03c1e8c","serviceInfo": {"fulfillmentType": "INSTORE_PICKUP"}}`).
		End()
	if errs != nil {
		fmt.Println("error")
	}
	res := response{}
	json.Unmarshal([]byte(body), &res)
	slotDay := res.SlotDays[0]
	for _, slot := range slotDay.Slots {
		if slot.Status == "AVAILABLE" {
			sendmail(fmt.Sprintf("%s -- %s Slot is %s\n", slot.StartDateTime, slot.EndDateTime, slot.Status))
			return
		}
	}
}

func doEvery(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}

func main() {
	period, _ := strconv.ParseInt(os.Getenv("PERIOD"), 32, 32)
	doEvery(time.Duration(period)*time.Second, checkAvailability)
}
