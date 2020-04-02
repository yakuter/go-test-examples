package parallel

import (
	"fmt"
	"testing"
)

var table = []struct {
	x    int
	y    int
	want int
}{
	{2, 2, 4},
	{5, 3, 8},
	{8, 4, 12},
	{12, 5, 17},
}

func TestSumParalel(t *testing.T) {
	t.Parallel()
	for _, row := range table {
		testName := fmt.Sprintf("Test %d+%d", row.x, row.y)
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			got, _ := Sum(row.x, row.y)
			if got != row.want {
				t.Errorf("Test fail! want: '%d', got: '%d'", row.want, got)
			}
		})
	}
}
