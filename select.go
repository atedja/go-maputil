package maputil

type SelectFunc func(interface{}, interface{}) bool

/**
Invokes `SelectFunc` for each k,v pair in the map, keeping elements for which
the function returns true. Opposite of Reject().

Example:

	func Pricey(k interface{}, v interface{}) bool {
		return v.(int) > 100
	}
	var prices = make(map[interface{}]interface{})
	prices["toothpaste"] = 100
	prices["cookies"] = 80
	prices["watermelons"] = 200
	prices["vodka"] = 400
	result := Select(myMap, Pricey)  // { "watermelons" : 200, "vodka": 400 }
*/
func Select(hash map[interface{}]interface{}, selectFunc SelectFunc) map[interface{}]interface{} {
	if hash == nil || selectFunc == nil {
		return hash
	}

	result := make(map[interface{}]interface{})
	for k, v := range hash {
		if selectFunc(k, v) {
			result[k] = v
		}
	}
	return result
}
