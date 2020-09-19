# benchmark-go-json
benchmark-go-json

---

## subject

- I want to check unmarshal performance for each json library.

## json libraries (2020-09)

- standard json.Decoder
- standard json.Unmarshal
- :star2: 8.2k https://github.com/json-iterator/go
- :star2: 2.7k https://github.com/pquerna/ffjson/ffjson
- :star2: 7k https://github.com/tidwall/gjson
- :star2: 1.5k https://github.com/ugorji/go
- :star2: 1k https://github.com/valyala/fastjson
- :star2: 6.1k https://github.com/aws/aws-sdk-go
- .. and my `DecodeJson()` function :)

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
go test bench=.
```

## benchmark

```
➜  benchmark-go-json git:(master) ✗ go test -bench=. 
goos: darwin
goarch: amd64
pkg: github.com/heshed/benchmark-go-json
BenchmarkDecodeJson-16                   2746500               435 ns/op            1008 B/op          5 allocs/op
BenchmarkAWSJsonSDK-16                   4326685               278 ns/op             880 B/op          3 allocs/op
BenchmarkStdJsonUnmarshal-16              681595              1755 ns/op             304 B/op          9 allocs/op
BenchmarkStdJsonDecode-16                4951611               238 ns/op             864 B/op          2 allocs/op
BenchmarkFastJsonJustParsing-16           816799              1369 ns/op            1752 B/op         13 allocs/op
BenchmarkGjsonJustParsing-16            27170618                43.9 ns/op            96 B/op          1 allocs/op
BenchmarkJsonIteratorUnmarshal-16        3308511               354 ns/op              16 B/op          2 allocs/op
BenchmarkJsonIteratorDecode-16           5238150               206 ns/op             664 B/op          3 allocs/op
BenchmarkJsonFFJsonUnmarshal-16           662472              1679 ns/op             304 B/op          9 allocs/op
BenchmarkJsonFFJsonDecodeReader-16       4692634               250 ns/op             864 B/op          2 allocs/op
BenchmarkJsonFFJsonDecodeFast-16        12783390                88.4 ns/op            80 B/op          2 allocs/op
BenchmarkJsonFFJsonDecode-16              703741              1780 ns/op             304 B/op          9 allocs/op
BenchmarkUgorjiJsonCodec-16              1000000              1187 ns/op            1088 B/op          4 allocs/op
PASS
ok      github.com/heshed/benchmark-go-json     17.576s
```

## Considerations

Due to its usability, it was excluded from the test.
- https://github.com/mailru/easyjson 

fastjson, gjson didn't provide unmarshaling functionality, so it only did parsing.
