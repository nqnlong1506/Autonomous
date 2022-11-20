package controller

import (
	"fmt"
	"net/smtp"
)

const (
	companyEmail = "taobiettaophailamgi@gmail.com"
	appPassword  = "fzfidhmsogkovbum"
)

func SendMailForUpdatingInventory(vendorEmail, vendorName, productName string) error {
	// from := "taobiettaophailamgi@gmail.com"
	// password := "fzfidhmsogkovbum"

	// toEmailAddress := "kanenguyen1506@autonomous.nyc"
	to := []string{vendorEmail}

	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port

	subject := "Subject: Update inventory urgently for more orders!!!\n"
	body := fmt.Sprintf("Hi %s team,\n\nYour product %s is selling fastly and going to be out of stock. Please help us to update the stock as soon as possible if your inventory is still available to sell. If not, when the product reach the lowest stock, we will make the one become Pre-Order for more 1 months and you can track the order later.\nLooking forward to hearing from you soon.\n\nThanks and best regards,\nGalvin.", vendorName, productName)
	message := []byte(subject + body)

	auth := smtp.PlainAuth("", companyEmail, appPassword, host)

	err := smtp.SendMail(address, auth, companyEmail, to, message)
	if err != nil {
		return err
	}
	return nil
}
