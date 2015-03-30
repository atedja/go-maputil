package maputil

type InjectFunc func(memo interface{}, k interface{}, v interface{}) interface{}

/*
Combines all entries of map by applying an operation, specified by
`InjectFunc`.  For each entry, the `InjectFunc` is passed a memo value and
the key-value pair. The result becomes the memo value for the next iteration.

Example:

	func Sum(memo interface{}, e interface{}) interface{} {
		return memo.(int) + e.(int)
	}

	var myArray = []interface{}{1, 2, 3, 4}
	result := Inject(myArray, nil, Sum)     // result is 10
	result := Inject(myArray, 10, Sum)      // result is 20
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
