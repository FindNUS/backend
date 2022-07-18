package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"gopkg.in/gomail.v2"
)

// Load environment variable for mailer
func GetMailSecret() string {
	mailSecret := os.Getenv("EMAIL_PASSWORD")
	if mailSecret == "" {
		// Read from secrets file
		log.Println("Mail secret not found. Attempting to find secret locally.")
		f, err := os.Open("../../secrets/email.txt")
		if err != nil {
			log.Fatal(err)
		}
		scanner := bufio.NewScanner(f)
		defer f.Close()
		for scanner.Scan() {
			mailSecret = scanner.Text()
		}
	}
	return mailSecret
}

// Send an email to the associated Lostee for possible item matches
func MailSendMessage(esItems []ElasticItem, lostItem Item, toEmail string) bool {
	mail := gomail.NewMessage()
	mail.SetAddressHeader("From", "findnus@outlook.com", "FindNUS")
	mail.SetHeader("To", toEmail)
	mail.SetHeader("Subject", "FindNUS: We found matches for your lost item!")

	messageIntro := fmt.Sprintf(
		`Hi there! <br><br>
		You got FindNUS to keep a lookout for your lost item:
		<b>%s</b><br><br>`, lostItem.Name,
	)

	messageBody := "Our smart minions travelled to the ends of the Earth and found some potential matches!<br>"
	messageBody += "<ol>"
	for _, item := range esItems {
		if item.Item_details == "" {
			messageBody += fmt.Sprintf(`
			<li>
			<b>%s</b><br>
			Item ID: %s<br><br>
			</li>
			`, item.Name, item.Id)
		} else {
			messageBody += fmt.Sprintf(`
			<li>
			<b> %s</b><br>
			Additional Details: %s<br>
			Item ID: %s<br><br>
			</li>
			`, item.Name, item.Item_details, item.Id)
		}
	}
	messageBody += "</ol>"
	messageBody += "To find out more, go to FindNUS.netlify.app and search for the <b>Item ID</b> to see more details about your possible matches!<br><br>"
	messageBody += "To stop FindNUS from searching for your item, login to the site and click the unsubscribe button.<br><br>"
	messageBody += "<b>Yours sincerely,<br>The FindNUS Team</br>"

	message := messageIntro + messageBody
	mail.SetBody("text/html", message)
	dialer := gomail.NewDialer("smtp.office365.com", 587, "findnus@outlook.com", GetMailSecret())
	err := dialer.DialAndSend(mail)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	return true
}
