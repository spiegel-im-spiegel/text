# [text] - Encoding/Decoding Text Package by [Golang]

[![Build Status](https://travis-ci.org/spiegel-im-spiegel/text.svg?branch=master)](https://travis-ci.org/spiegel-im-spiegel/text)

## Install

```
$ go get github.com/spiegel-im-spiegel/text
```

Installing by [dep].

```
$ dep ensure -add github.com/spiegel-im-spiegel/text
```

## Usage

```go
encoding := detect.EncodingJa([]byte("こんにちは世界"))
fmt.Println(encoding)
// Output:
// UTF-8
```

```go
jisText := []byte{0x1b, 0x24, 0x42, 0x24, 0x33, 0x24, 0x73, 0x24, 0x4b, 0x24, 0x41, 0x24, 0x4f, 0x40, 0x24, 0x33, 0x26, 0x1b, 0x28, 0x42}
res, err := decode.ToUTF8ja(jisText)
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println(string(res))
// Output:
// こんにちは世界
```

## Dependencies

```
dep status -dot | dot -Tpng -o dependency.png
```

[![Dependencies](dependency.png)](dependency.png)

[text]: https://github.com/spiegel-im-spiegel/text "spiegel-im-spiegel/text: Encoding/Decoding Text Package by Golang"
[Golang]: https://golang.org/ "The Go Programming Language"
