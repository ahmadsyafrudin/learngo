package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	fmt.Println("SetUp Unittest")
	m.Run()
	fmt.Println("TearDown Unittest")
}

func TestTableExample(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "rudi",
			request:  "Rudi",
			expected: "Hello Rudi",
		},
		{
			name:     "ina",
			request:  "Ina",
			expected: "Hello Ina",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)

		})
	}
}
func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Rudi")

	assert.Equal(t, "Hello Rudi", result)

	fmt.Println("Test Hello World Finished")
}
