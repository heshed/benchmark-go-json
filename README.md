# benchmark-go-json
benchmark-go-json

---

## subject

- I want to check unmarshal performance for each json library.

## json libraries (2020-09)

- [standard json.Decoder](https://golang.org/pkg/encoding/json/#Decoder)
- [standard json.Unmarshal](https://golang.org/pkg/encoding/json/#Unmarshal)
- :star2: 8.2k https://github.com/json-iterator/go
- :star2: 2.7k https://github.com/pquerna/ffjson
- :star2: 7k https://github.com/tidwall/gjson
- :star2: 1.5k https://github.com/ugorji/go
- :star2: 1k https://github.com/valyala/fastjson
- :star2: 6.1k https://github.com/aws/aws-sdk-go
- .. and my [`DecodeJson()`][DecodeJson] function :)

## machine environment

```
Model Name: MacBook Pro
Model Identifier: MacBookPro16,1
Processor Name: 8-Core Intel Core i9
Processor Speed: 2.4 GHz
Number of Processors: 1
Total Number of Cores: 8
L2 Cache (per Core): 256 KB
L3 Cache: 16 MB
Hyper-Threading Technology: Enabled
Memory: 32 GB

Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
```

## usage

```console
# go version go1.15 darwin/amd64
go test bench=.
```

## benchmark

```
➜  benchmark-go-json git:(master) ✗ go test -bench=. 
BenchmarkJsonIteratorDecode-16                   5480703               216 ns/op             664 B/op          3 allocs/op
BenchmarkStdJsonDecode-16                        5088841               234 ns/op             864 B/op          2 allocs/op
BenchmarkJsonFFJsonDecodeReader-16               4593306               248 ns/op             864 B/op          2 allocs/op
BenchmarkAWSJsonSDK-16                           4017230               282 ns/op             880 B/op          3 allocs/op
BenchmarkJsonIteratorUnmarshal-16                3103564               357 ns/op              16 B/op          2 allocs/op
BenchmarkDecodeJson-16                           2826174               445 ns/op            1008 B/op          5 allocs/op
BenchmarkGjsonUnmarshal-16                       1231782              1007 ns/op             687 B/op          4 allocs/op
BenchmarkUgorjiJsonCodec-16                      1000000              1218 ns/op            1088 B/op          4 allocs/op
BenchmarkJsonFFJsonDecode-16                      682381              1731 ns/op             304 B/op          9 allocs/op
BenchmarkStdJsonUnmarshal-16                      710850              1737 ns/op             304 B/op          9 allocs/op
BenchmarkJsonFFJsonUnmarshal-16                   699499              1772 ns/op             304 B/op          9 allocs/op
BenchmarkFastJsonParseAndUnmarshal-16             643364              1888 ns/op            2018 B/op         15 allocs/op
ok      github.com/heshed/benchmark-go-json     17.796s
```

## Considerations

Due to its usability, it was excluded from the test.
- https://github.com/mailru/easyjson 

~fastjson, gjson didn't provide unmarshaling functionality, so it only did parsing.~

[DecodeJson]: https://github.com/heshed/benchmark-go-json/blob/master/benchmark_test.go#L49-L55