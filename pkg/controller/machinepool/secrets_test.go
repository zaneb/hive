package machinepool

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorkerUserData(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedOutput string
	}{
		{
			name:           "pre 4.10",
			input:          "4.9.0",
			expectedOutput: "worker-user-data",
		},
		{
			name:           "post 4.10",
			input:          "4.10.0",
			expectedOutput: "worker-user-data-managed",
		},
		{
			name:           "4.10 pre-release",
			input:          "4.10.0-fc.3",
			expectedOutput: "worker-user-data-managed",
		},
		{
			name:           "error",
			input:          "unparseable",
			expectedOutput: "worker-user-data-managed",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actualOutput := workerUserData(test.input)
			assert.Equal(t, test.expectedOutput, actualOutput, "unexpected output")
		})
	}
}
