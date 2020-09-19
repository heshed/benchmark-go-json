/*
json unmarshal, decoder benchmark test

    standard json.Decoder
    standard json.Unmarshal
	https://github.com/json-iterator/go
	https://github.com/pquerna/ffjson/ffjson
	https://github.com/tidwall/gjson
	https://github.com/ugorji/go/codec
	https://github.com/valyala/fastjson
	https://github.com/aws/aws-sdk-go
    DecodeJson (my custom function :))
*/
package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"reflect"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/private/protocol/json/jsonutil"
	jsoniter "github.com/json-iterator/go"
	"github.com/pquerna/ffjson/ffjson"
	"github.com/tidwall/gjson"
	"github.com/ugorji/go/codec"
	"github.com/valyala/fastjson"
)

type oauthProvider struct {
	Provider   string `json:"provider"`
	Token      string `json:"token"`
	ExpireTime string `json:"expire_time"`
}

type linkdedAccount struct {
	Code     int             `json:"code"`
	Message  string          `json:"message"`
	Response []oauthProvider `json:"response"`
}

var (
	inputByte = bytes.NewBufferString(`{"code":200, "message":"", "response":[{"provider":"google","token":"token","expire_time":""}]}`)
	input     = strings.NewReader(inputByte.String())
)

func DecodeJson(r io.Reader, v interface{}) (interface{}, error) {
	vp := reflect.New(reflect.TypeOf(v)).Interface()
	decoder := json.NewDecoder(r)
	decoder.UseNumber()
	err := decoder.Decode(vp)
	return reflect.Indirect(reflect.ValueOf(vp)).Interface(), err
}

func BenchmarkDecodeJson(b *testing.B) {
	b.ReportAllocs()
	var out linkdedAccount
	for i := 0; i < b.N; i++ {
		DecodeJson(input, out)
	}
}

// https://github.com/aws/aws-sdk-go
func BenchmarkAWSJsonSDK(b *testing.B) {
	b.ReportAllocs()
	var out linkdedAccount
	for i := 0; i < b.N; i++ {
		jsonutil.UnmarshalJSON(&out, input)
	}
}

// standard json unmarshal
func BenchmarkStdJsonUnmarshal(b *testing.B) {
	b.ReportAllocs()
	var out linkdedAccount
	for i := 0; i < b.N; i++ {
		json.Unmarshal(inputByte.Bytes(), &out)
	}
}

// standard json decoder
func BenchmarkStdJsonDecode(b *testing.B) {
	b.ReportAllocs()
	var out linkdedAccount
	for i := 0; i < b.N; i++ {
		json.NewDecoder(input).Decode(&out)
	}
}

// https://github.com/valyala/fastjson
func BenchmarkFastJsonJustParsing(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		fastjson.Parse(inputByte.String())
	}
}

// https://github.com/tidwall/gjson/
func BenchmarkGjsonJustParsing(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		gjson.ParseBytes(inputByte.Bytes())
	}
}

// https://github.com/json-iterator/go
func BenchmarkJsonIteratorUnmarshal(b *testing.B) {
	b.ReportAllocs()
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var out linkdedAccount
	for i := 0; i < b.N; i++ {
		json.Unmarshal(inputByte.Bytes(), &out)
	}
}

// https://github.com/json-iterator/go
func BenchmarkJsonIteratorDecode(b *testing.B) {
	b.ReportAllocs()
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var out linkdedAccount
	for i := 0; i < b.N; i++ {
		json.NewDecoder(input).Decode(&out)
	}
}

// https://github.com/pquerna/ffjson
func BenchmarkJsonFFJsonUnmarshal(b *testing.B) {
	b.ReportAllocs()
	var out linkdedAccount
	for i := 0; i < b.N; i++ {
		ffjson.Unmarshal(inputByte.Bytes(), &out)
	}
}

// https://github.com/pquerna/ffjson
func BenchmarkJsonFFJsonDecodeReader(b *testing.B) {
	b.ReportAllocs()
	var out linkdedAccount
	for i := 0; i < b.N; i++ {
		ffjson.NewDecoder().DecodeReader(input, &out)
	}
}

// https://github.com/pquerna/ffjson
func BenchmarkJsonFFJsonDecodeFast(b *testing.B) {
	b.ReportAllocs()
	var out linkdedAccount
	for i := 0; i < b.N; i++ {
		ffjson.NewDecoder().DecodeFast(inputByte.Bytes(), &out)
	}
}

// https://github.com/pquerna/ffjson
func BenchmarkJsonFFJsonDecode(b *testing.B) {
	b.ReportAllocs()
	var out linkdedAccount
	for i := 0; i < b.N; i++ {
		ffjson.NewDecoder().Decode(inputByte.Bytes(), &out)
	}
}

// https://github.com/ugorji/go
func BenchmarkUgorjiJsonCodec(b *testing.B) {
	b.ReportAllocs()
	var out linkdedAccount
	for i := 0; i < b.N; i++ {
		codec.NewDecoder(input, &codec.JsonHandle{}).Decode(&out)
	}
}
