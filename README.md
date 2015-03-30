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
    result := maputil.Reject(myMap, Pricey)  // { "toothpaste" : 100, "cookie": 80 }

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
    result := maputil.Select(myMap, Pricey)  // { "watermelons" : 200, "vodka": 400 }

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
