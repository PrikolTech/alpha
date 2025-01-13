package ptr

import (
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/require"
)

func TestTo(t *testing.T) {
	v := gofakeit.Int()
	p := To(v)
	require.Equal(t, v, *p)
}

func TestValue(t *testing.T) {
	var p *int
	v := Value(p)
	require.Equal(t, 0, v)

	pv := gofakeit.Int()
	p = &pv
	v = Value(p)
	require.Equal(t, pv, v)
}
