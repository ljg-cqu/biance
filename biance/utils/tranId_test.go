package utils

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTranID(t *testing.T) {
	id := TranID()
	require.True(t, len(id) == 20)
	fmt.Println(id)
}
