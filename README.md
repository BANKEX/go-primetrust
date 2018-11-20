# PrimeTrust go library

## Running the Example

```
PRIMETRUST_LOGIN=user@example.com PRIMETRUST_PASSWORD=password go run examples/main.go
```

## Usage in own application

```go
package main

import (
	"github.com/BANKEX/go-primetrust"
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
}
```