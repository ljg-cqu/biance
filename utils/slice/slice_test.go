package slice

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSlice(t *testing.T) {
	s := New(5)
	for i := 0; i < 6; i++ {
		s.AddElem(i)
	}

	require.Equal(t, 5, s.Len())
	require.Equal(t, 1, s.Elem(0).(int))
	require.Equal(t, 2, s.Elem(1).(int))
	require.Equal(t, 3, s.Elem(2).(int))
	require.Equal(t, 4, s.Elem(3).(int))
	require.Equal(t, 5, s.Elem(4).(int))
}
