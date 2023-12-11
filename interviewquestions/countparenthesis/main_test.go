package main_test

import (
	"fmt"
	"github.com/peterehik/goland/interviewquestions/countparenthesis"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCountParenthesis(t *testing.T) {
	testcases := []struct {
		input    string
		expected int
		wantErr  bool
	}{
		{
			input:    "(((())))",
			expected: 0,
		},
		{
			input:    ")()()",
			expected: 1,
		},
		{
			input:    "()(())(",
			expected: 1,
		},
	}
	for _, tc := range testcases {
		tc := tc
		t.Run(fmt.Sprintf("input_%s", tc.input), func(t *testing.T) {
			t.Parallel()
			count, err := main.CountParenthesis(tc.input)
			if tc.wantErr {
				require.Error(t, err)
				return
			}
			require.Equal(t, tc.expected, count)
		})
	}

}
