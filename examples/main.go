package main

import (
	"github.com/BANKEX/go-primetrust"
	// "github.com/BANKEX/go-primetrust/models"
	"log"
	"os"
)

func main() {
	primetrust.Init(true, os.Getenv("PRIMETRUST_LOGIN"), os.Getenv("PRIMETRUST_PASSWORD"))

	if accountTypes, err := primetrust.GetAccountTypes(); err != nil {
		log.Println("Error getting account types:", err.Error())
	} else {
		log.Printf("Account Types: %d", len(accountTypes.Data))
		if accountType, err := primetrust.GetAccountType(accountTypes.Data[0].ID); err != nil {
			log.Printf("Error getting \"%s\" account type:", accountTypes.Data[0].ID, err.Error())
		} else {
			log.Printf("Account Type \"%s\": %+v", accountTypes.Data[0].ID, accountType)
		}
	}

	// if accounts, err := primetrust.GetAccounts(); err != nil {
	// 	log.Println("Error getting accounts:", err.Error())
	// } else {
	// 	log.Printf("Accounts: %d", len(accounts.Data))
	// 	if account, err := primetrust.GetAccount(accounts.Data[0].ID); err != nil {
	// 		log.Println("Error getting account:", err.Error())
	// 	} else {
	// 		log.Printf("Account: %+v", account)
	// 	}
	// }

	// if contacts, err := primetrust.GetContacts(); err != nil {
	// 	log.Println("Error getting contacts:", err.Error())
	// } else {
	// 	log.Printf("Contacts: %d", len(contacts.Data))
	// 	if contact, err := primetrust.GetContact(contacts.Data[0].ID); err != nil {
	// 		log.Println("Error getting contact:", err.Error())
	// 	} else {
	// 		log.Printf("Contact: %+v", contact)
	// 	}
	// }

	// contact := models.NewNaturalPersonContact()
	// if ok, err := primetrust.CreateNewContact(contact); ok == false {
	// 	log.Println("Error creating new contact:", err)
	// } else {
	// 	log.Println("Contact created OK")
	// }

	log.Println("Done")
}
