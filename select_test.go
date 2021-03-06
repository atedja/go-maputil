package maputil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Cheap(key interface{}, value interface{}) bool {
	return value.(int) < 100
}

func TestSelect(t *testing.T) {
	var prices = make(map[interface{}]interface{})
	prices["toothpaste"] = 100
	prices["cookies"] = 80
	prices["watermelons"] = 200
	prices["vodka"] = 400
	result := Select(prices, Cheap)
	assert.Equal(t, 80, result["cookies"].(int))
	assert.Nil(t, result["watermelons"])
	assert.Nil(t, result["toothpaste"])
	assert.Nil(t, result["vodka"])
}
