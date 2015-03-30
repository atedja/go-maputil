package maputil

type RejectFunc func(k interface{}, v interface{}) bool

/**
Invokes `RejectFunc` for each k,v pair in the map, deleting elements for which
the function returns true. Opposite of Select().

Example:

	func Pricey(k interface{}, v interface{}) bool {
		return v.(int) > 100
	}
	var prices = make(map[interface{}]interface{})
	prices["toothpaste"] = 100
	prices["cookies"] = 80
	prices["watermelons"] = 200
	prices["vodka"] = 400
	result := maputil.Reject(myMap, Pricey)  // { "toothpaste" : 100, "cookie": 80 }
*/
func Reject(hash map[interface{}]interface{}, rejectFunc RejectFunc) map[interface{}]interface{} {
	if hash == nil || rejectFunc == nil {
		return hash
	}

	result := make(map[interface{}]interface{})
	for k, v := range hash {
		if !rejectFunc(k, v) {
			result[k] = v
		}
	}
	return result
}
