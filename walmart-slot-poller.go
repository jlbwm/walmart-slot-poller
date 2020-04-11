package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

const (
	Toronto = "5b755848-3e2e-490e-ba80-56bf3fbc5bb6"
)
const (
	Subscription = "Congratulations on your successful subscription to the Walmart(Toronto) slot reminder service.
	
	We will automatically check the available slots every hour and remind you by email"
)

type slot struct {
	StartDateTime string
	EndDateTime   string
	Status        string
	SlotId        string
}

type slotDay struct {
	Slots         []slot
	SlotDate      string
	AccessPointId string
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
	eightDayLater := today.AddDate(0, 0, 8)
	layout := "2006-01-02"
	startDate := sevenDayLater.Format(layout)
	endDate := eightDayLater.Format(layout)
	fmt.Printf("checking availibility for %s - %s \n", startDate, endDate)
	requestBody, err := json.Marshal(map[string]interface{}{
		"startDate":     startDate,
		"endDate":       endDate,
		"accessPointId": Toronto,
		"serviceInfo": map[string]string{
			"fulfillmentType": "DELIVERY",
		},
		"customerInfo": map[string]interface{}{
			"customerAddress": map[string]interface{}{
				"postalCode":          "M5V 2Y6",
				"stateOrProvinceName": "ON",
				"latitude":            43.6439056396,
				"longitude":           -79.402961731,
			},
		},
	})
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := http.Post("https://www.walmart.ca/api/cart-page/accesspoints/slotavailability?grocery=true", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	res := response{}
	json.Unmarshal([]byte(string(body)), &res)
	for _, slotDay := range res.SlotDays {
		for _, slot := range slotDay.Slots {
			if slot.Status == "AVAILABLE" {
				sendmail(fmt.Sprintf("%s -- %s Slot is %s\n", slot.StartDateTime, slot.EndDateTime, slot.Status))
				fmt.Println(fmt.Sprintf("%s -- %s Slot is %s\n", slot.StartDateTime, slot.EndDateTime, slot.Status))
				return
			}
		}
	}
}

func doEvery(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}

func main() {
	sendmail(Subscription)
	fmt.Printf("Start app and waiting for %s \n", os.Getenv("PERIOD"))
	period, _ := strconv.ParseInt(os.Getenv("PERIOD"), 10, 32)
	doEvery(time.Duration(period)*time.Second, checkAvailability)
}
