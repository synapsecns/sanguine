# Bytemap

The bytemap package provides an implementation of a map that uses []byte keys, called ByteSliceMap. The map is implemented as a Trie data structure. It is based off of [triemap](https://github.com/google/triemap), but adds support for generics.

## Usage

To use the ByteSliceMap, first import the package:

```go
import "github.com/synapsecns/sanguine/core/bytemap"
```

Then, create a new ByteSliceMap with the desired value type:

```go
m := &bytemap.ByteSliceMap[string]{}
```

### Inserting Values

Values can be inserted into the map using either `[]byte` or `string` keys. To insert a value with a `string` key, use the `PutString` method:

```go
m.PutString("key", "value")
```

To insert a value with a []byte key, use the Put method:
```go
m.Put([]byte{0x01, 0x02}, "value")
```


### Retrieving values

Values can be retrieved from the map using either `[]byte` or `string` keys. To retrieve a value with a string key, use the GetString method:

```go
value, exists := m.GetString("key")
```

To retrieve a value with a `[]byte` key, use the `Get` method:
```go
value, exists := m.Get([]byte{0x01, 0x02})
```

The `exists` return value indicates whether the key exists in the map.

## Performance

The ByteSliceMap seems to perform worse than a regular `map[string]V`, even when casting `[]byte` to string. Therefore, it may not be suitable for performance-critical applications.




