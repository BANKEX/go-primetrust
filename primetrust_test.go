package primetrust

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// Init
	Init(true, os.Getenv("PRIMETRUST_LOGIN"), os.Getenv("PRIMETRUST_PASSWORD"))

	// Running
	retCode := m.Run()

	// Teardown
	// ...

	// Exiting
	os.Exit(retCode)
}

func TestPrimetrustVersion(t *testing.T) {
	assert.Equal(t, "1.0.2", Version, "Primetrust version does not match")
}

func TestGetAccounts(t *testing.T) {
	_, err := GetAccounts()
	assert.Equal(t, err, nil, "Error getting accounts")
}
