package dateutil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMonthEndDay(t *testing.T) {
	end := GetMonthEndDay()
	expected := "2022-03-31 23:59:59"
	assert.Equal(t, expected, end)
	t.Log("expected:", expected, "actual:", end)
}
