package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnection(t *testing.T) {
	testCases := []struct {
		Name      string
		EnvPath   string
		WantError bool
	}{
		{
			Name:      "success",
			EnvPath:   "../../.env",
			WantError: false,
		}, {
			Name:      "failed",
			EnvPath:   "../../../.env",
			WantError: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			dbConnection, err := NewConnection(testCase.EnvPath)
			if testCase.WantError {
				assert.NotNil(t, err)
				assert.Nil(t, dbConnection)
			} else {
				assert.NotNil(t, dbConnection)
				assert.Nil(t, err)
			}
		})
	}
}
