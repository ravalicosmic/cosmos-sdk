package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGasMeter(t *testing.T) {
	cases := []struct {
		limit Gas
		usage []Gas
	}{
		{10, []Gas{1, 2, 3, 4}},
		{1000, []Gas{40, 30, 20, 10, 900}},
		{100000, []Gas{99999, 1}},
		{100000000, []Gas{50000000, 40000000, 10000000}},
		{65535, []Gas{32768, 32767}},
		{65536, []Gas{32768, 32767, 1}},
	}

	for _, tc := range cases {
		meter := NewGasMeter(tc.limit)
		used := int64(0)

		for _, usage := range tc.usage {
			used += usage
			require.NotPanics(t, func() { meter.ConsumeGas(usage, "") })
			require.Equal(t, used, meter.GasConsumed())
		}

		require.Panics(t, func() { meter.ConsumeGas(1, "") })
		break

	}
}
