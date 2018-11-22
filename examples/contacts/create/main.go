package main

import (
	"github.com/BANKEX/go-primetrust"
	"github.com/BANKEX/go-primetrust/models"
	"github.com/satori/go.uuid"
	"log"
	"os"
)

func main() {
	primetrust.Init(true, os.Getenv("PRIMETRUST_LOGIN"), os.Getenv("PRIMETRUST_PASSWORD"))

	if accountId, err := uuid.FromString(os.Getenv("PRIMETRUST_ACCOUNT")); err == nil {
		contact := models.NewNaturalPersonContact(accountId)

		contact.Data.Attributes.DateOfBirth = "1981-10-31"
		contact.Data.Attributes.Email = "name@example.com"
		contact.Data.Attributes.Sex = "male"
		contact.Data.Attributes.Name = "Sample Person"
		contact.Data.Attributes.TaxCountry = "US"
		contact.Data.Attributes.Label = "CEO"
		contact.Data.Attributes.TaxIDNumber = "123123123"
		contact.Data.Attributes.PrimaryAddress = models.Address{
			Type:       models.AddressTypeHome,
			City:       "Beverly Hills",
			Country:    "US",
			PostalCode: "90210",
			Region:     "CA",
			Street1:    "1 Sunset Blvd.",
		}
		contact.Data.Attributes.PrimaryPhoneNumber = models.PhoneNumber{
			Number: "+15555555555",
		}

		if ok, err := primetrust.CreateNewContact(contact); ok == false {
			log.Println("Error creating new contact:", err)
		} else {
			log.Println("Contact created OK")
		}

		log.Println("Done")
	} else {
		log.Println("Error:", err)
	}
}
