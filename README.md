# maputil
--
    import "github.com/atedja/go-maputil"


## Usage

#### func  Inject

```go
func Inject(hash map[interface{}]interface{}, initial interface{}, injectFunc InjectFunc) interface{}
```
Combines all entries of map by applying an operation, specified by `InjectFunc`.
For each entry, the `InjectFunc` is passed a memo value and the key-value pair.
The result becomes the memo value for the next iteration.

Example:

    func Sum(memo interface{}, e interface{}) interface{} {
    	return memo.(int) + e.(int)
    }

    var myArray = []interface{}{1, 2, 3, 4}
    result := Inject(myArray, nil, Sum)     // result is 10
    result := Inject(myArray, 10, Sum)      // result is 20

#### func  Reject

```go
func Reject(hash map[interface{}]interface{}, rejectFunc RejectFunc) map[interface{}]interface{}
```
* Invokes `RejectFunc` for each k,v pair in the map, deleting elements for which
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
    result := Reject(myMap, Pricey)  // { "toothpaste" : 100, "cookie": 80 }

#### func  Select

```go
func Select(hash map[interface{}]interface{}, selectFunc SelectFunc) map[interface{}]interface{}
```
* Invokes `SelectFunc` for each k,v pair in the map, keeping elements for which
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

#### type InjectFunc

```go
type InjectFunc func(memo interface{}, k interface{}, v interface{}) interface{}
```


#### type RejectFunc

```go
type RejectFunc func(interface{}, interface{}) bool
```


#### type SelectFunc

```go
type SelectFunc func(interface{}, interface{}) bool
```
