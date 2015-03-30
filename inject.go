package maputil

type InjectFunc func(memo interface{}, k interface{}, v interface{}) interface{}

/*
Combines all entries of map by applying an operation, specified by
`InjectFunc`.  For each entry, the `InjectFunc` is passed a memo value and
the key-value pair. The result becomes the memo value for the next iteration.

Example:

	func GallonUsage(memo interface{}, k interface{}, v interface{}) interface{} {
		m := memo.(map[interface{}]interface{})
		m[k] = float64(v.(int)) / 40.0
		return m
	}

	var distances = make(map[interface{}]interface{})
	distances["Las Vegas"] = 269
	distances["San Francisco"] = 382
	distances["San Diego"] = 120
	distances["Sacramento"] = 384
	var gallons = make(map[interface{}]interface{})
	maputil.Inject(distances, gallons, GallonUsage)
	// gallons = { "Las Vegas": 6.725, "San Francisco": 9.55, "San Diego": 3.0, "Sacramento": 9.6 }
*/
func Inject(hash map[interface{}]interface{}, initial interface{}, injectFunc InjectFunc) interface{} {
	if hash == nil || injectFunc == nil {
		return hash
	}

	memo := initial
	for k, v := range hash {
		memo = injectFunc(memo, k, v)
	}
	return memo
}
