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
	to := mail.NewEmail("Customer", os.Getenv("TOEMAIL"))
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
	to2 :=
}
