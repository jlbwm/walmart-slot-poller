package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func Send_Gridmail(plainTextContent string) {
	from := mail.NewEmail("Walmart Slot Robort", "ljx249@gmail.com")
	subject := "There are slots available"
	to := mail.NewEmail("Special Customer", os.Getenv("TOEMAIL1"))
	htmlContent := "<strong>" + plainTextContent + "</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("PASSWORD"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
	to2 := mail.NewEmail("Custom2", os.Getenv("TOEMAIL2"))
	message2 := mail.NewSingleEmail(from, subject, to2, plainTextContent, htmlContent)
	response2, err := client.Send(message2)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response2.StatusCode)
		fmt.Println(response2.Body)
		fmt.Println(response2.Headers)
	}

}
