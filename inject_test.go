package maputil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func GallonUsage(memo interface{}, k interface{}, v interface{}) interface{} {
	m := memo.(map[interface{}]interface{})
	m[k] = float64(v.(int)) / 40.0
	return m
}

func TestInject(t *testing.T) {
	var distances = make(map[interface{}]interface{})
	distances["Las Vegas"] = 269
	distances["San Francisco"] = 382
	distances["San Diego"] = 120
	distances["Sacramento"] = 384
	var gallons = make(map[interface{}]interface{})
	Inject(distances, gallons, GallonUsage)
	assert.Equal(t, 6.725, gallons["Las Vegas"])
	assert.Equal(t, 9.55, gallons["San Francisco"])
	assert.Equal(t, 3.0, gallons["San Diego"])
	assert.Equal(t, 9.6, gallons["Sacramento"])
}
