# [text] - Encoding/Decoding Text Package by [Golang]

[![Build Status](https://travis-ci.org/spiegel-im-spiegel/text.svg?branch=master)](https://travis-ci.org/spiegel-im-spiegel/text)
[![GitHub license](http://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/spiegel-im-spiegel/text/master/LICENSE)
[![GitHub release](http://img.shields.io/github/release/spiegel-im-spiegel/mt.svg)](https://github.com/spiegel-im-spiegel/text/releases/latest)

## Declare [text] module

See [go.mod](https://github.com/spiegel-im-spiegel/gpgpdump/blob/master/go.mod) file. 

## Usage of package

### Import Package

```go
import "github.com/spiegel-im-spiegel/openbd-api"
```

### detect

```go
encoding := detect.EncodingJa(bytes.NewBufferString("こんにちは，世界"))
fmt.Println(encoding)
// Output:
// UTF-8
```

### decode

```go
jisText := []byte{0x1b, 0x24, 0x42, 0x24, 0x33, 0x24, 0x73, 0x24, 0x4b, 0x24, 0x41, 0x24, 0x4f, 0x40, 0x24, 0x33, 0x26, 0x1b, 0x28, 0x42}
res, err := decode.ToUTF8ja(bytes.NewReader(jisText))
if err != nil {
    fmt.Println(err)
    return
}
buf := new(bytes.Buffer)
io.Copy(buf, res)
fmt.Println(buf)
// Output:
// こんにちは世界
```

### encode

```go
utf8Text := "こんにちは，世界\n"
res, err := encode.FromUTF8To(detect.ISO2022JP, bytes.NewBufferString(utf8Text))
if err != nil {
    fmt.Println(err)
    return
}
buf := new(bytes.Buffer)
io.Copy(buf, res)
fmt.Println(buf.Bytes())
// Output:
// [27 36 66 36 51 36 115 36 75 36 65 36 79 33 36 64 36 51 38 27 40 66 10]
```

### convert

```go
jisText := []byte{0x1b, 0x24, 0x42, 0x24, 0x33, 0x24, 0x73, 0x24, 0x4b, 0x24, 0x41, 0x24, 0x4f, 0x40, 0x24, 0x33, 0x26, 0x1b, 0x28, 0x42}
res, err := convert.FromTo(detect.ISO2022JP, detect.UTF8, bytes.NewReader(jisText))
if err != nil {
    fmt.Println(err)
    return
}
buf := new(bytes.Buffer)
io.Copy(buf, res)
fmt.Println(buf)
// Output:
// こんにちは世界
```

### newline

```go
res := newline.Reader(strings.NewReader("こんにちは\nこんにちは\rこんにちは\r\nこんにちは"), newline.LF)
buf := new(bytes.Buffer)
io.Copy(buf, res)
fmt.Println(buf)
// Output:
// こんにちは
// こんにちは
// こんにちは
// こんにちは
```

```go
res := newline.String("こんにちは\nこんにちは\rこんにちは\r\nこんにちは", newline.LF)
fmt.Println(res)
// Output:
// こんにちは
// こんにちは
// こんにちは
// こんにちは
```

### normalize

```go
res := normalize.Reader(strings.NewReader("ﾍﾟﾝｷﾞﾝ"), normalize.NFKC)
buf := new(bytes.Buffer)
io.Copy(buf, res)
fmt.Println(buf)
// Output:
// ペンギン
```

```go
res := normalize.String("ﾍﾟﾝｷﾞﾝ", normalize.NFKC)
fmt.Println(res)
// Output:
// ペンギン
```

### width

```go
res := width.Reader(strings.NewReader("１２３４５６７８９０ｱｲｳｴｵｶｷｸｹｺＡＢＣＤＥＦＧＨＩＪＫ"), width.Fold)
buf := new(bytes.Buffer)
io.Copy(buf, res)
fmt.Println(buf)
// Output:
// 1234567890アイウエオカキクケコABCDEFGHIJK
```

```go
res := width.String("１２３４５６７８９０ｱｲｳｴｵｶｷｸｹｺＡＢＣＤＥＦＧＨＩＪＫ", width.Fold)
fmt.Printlnres)
// Output:
// 1234567890アイウエオカキクケコABCDEFGHIJK
```

## Command Line Interface

### Binaries

See [latest release](https://github.com/spiegel-im-spiegel/text/releases/latest).

### Usage

```
$ gonkf -h
Network Kanji Filter by Golang

Usage:
  gonkf [flags]
  gonkf [command]

Available Commands:
  conv        Convert character encoding of text
  guess       Guess character encoding of text
  help        Help about any command
  norm        Unicode normalization
  nwline      Convert newline of text
  version     Print the version number of gonkf
  width       Convert character width of text

Flags:
  -h, --help   help for gonkf

Use "gonkf [command] --help" for more information about a command.
```

### guess sub-command

```
$ gonkf guess -h
Guess character encoding of text

Usage:
  gonkf guess [flags] [text file]

Flags:
  -h, --help   help for guess

$ echo こんにちは。世界の国から | gonkf guess
UTF-8
```

### conv sub-command

```
$ gonkf conv -h
Convert character encoding of text

Usage:
  gonkf conv [flags] [text file]

Flags:
  -d, --dst-encoding string   encoding of dest [euc|jis|sjis|utf8] (default "utf8")
  -h, --help                  help for conv
  -o, --output string         output file path
  -s, --src-encoding string   encoding of src [euc|jis|sjis|utf8]

$ gonkf conv -d utf8 testdata/SHIFT_JIS.txt
こんにちは。世界の国から。
```

### norm sub-command

```
$ gonkf norm -h
Unicode normalization (UTF-8 text only)

Usage:
  gonkf norm [flags] [text file]

Flags:
  -f, --form string     normalization form [nfc|nfd|nfkc|nfkd] (default "nfc")
  -h, --help            help for norm
  -o, --output string   output file path

$ echo ﾍﾟﾝｷﾞﾝ | gonkf norm -f NFKC
ペンギン
```

### width sub-command

```
$ gonkf width -h
Convert character width of text (UTF-8 text only)

Usage:
  gonkf width [flags] [text file]

Flags:
  -f, --form string     form of width [fold|narrow|widen] (default "fold")
  -h, --help            help for width
  -o, --output string   output file path

$ echo １２３４５６７８９０ｱｲｳｴｵｶｷｸｹｺＡＢＣＤＥＦＧＨＩＪＫ | gonkf width -f fold
1234567890アイウエオカキクケコABCDEFGHIJK
```

[text]: https://github.com/spiegel-im-spiegel/text "spiegel-im-spiegel/text: Encoding/Decoding Text Package by Golang"
[Golang]: https://golang.org/ "The Go Programming Language"
