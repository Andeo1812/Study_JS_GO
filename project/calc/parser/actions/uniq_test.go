package actions

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetNumber(t *testing.T) {
	var input string = "-10basd"
	var output float64 = -10
	number, length, _ := GetNumber(input)
	require.Equal(t, length, 3, "they should be equal", input)
	require.Equal(t, output, number, "they should be equal", input)

	input = "10)basd"
	output = 10
	number, length, _ = GetNumber(input)
	require.Equal(t, length, 2, "they should be equal", input)
	require.Equal(t, output, number, "they should be equal", input)
}

func TestGetEndExpression(t *testing.T) {
	var input string = "()())"
	var output int = 4
	pos := GetEndExpression(input)
	require.Equal(t, pos, output, "they should be equal", input)

	input = "((())))"
	output = 6
	pos = GetEndExpression(input)
	require.Equal(t, pos, output, "they should be equal", input)

	input = "012()56)"
	output = 7
	pos = GetEndExpression(input)
	require.Equal(t, pos, output, "they should be equal", input)
}
