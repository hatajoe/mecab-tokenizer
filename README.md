# mecab-tokenizer

[![GoDoc](https://godoc.org/github.com/hatajoe/mecab-tokenizer?status.svg)](https://godoc.org/github.com/hatajoe/mecab-tokenizer)
[![Circle CI](https://circleci.com/gh/hatajoe/mecab-tokenizer.svg?style=svg)](https://circleci.com/gh/hatajoe/mecab-tokenizer)
[![Coverage Status](https://coveralls.io/repos/github/hatajoe/mecab-tokenizer/badge.svg?branch=master)](https://coveralls.io/github/hatajoe/mecab-tokenizer?branch=master)

```
This is under development
```

mecab-tokenizer is text tokenizer wrapped ikawaha/kagome

## Dependencies

- [MeCab: Yet Another Part-of-Speech and Morphological Analyzer](https://mecab.googlecode.com/svn/trunk/mecab/doc/index.html#install)
- [ikawaha/kagome: Self-contained Japanese Morphological Analyzer written in pure golang](https://github.com/ikawaha/kagome)

Please install it

## Install

```
% go get -u github.com/hatajoe/mecab-tokenizer
```

## Examples

```go
package main

import (
    "fmt"

    "github.com/hatajoe/mecab-tokenizer"
)

func main() {
	tokens := tokenizer.NewTokenizer().Tokenize("すもももももももものうち。").DistinctByNoun().Sort()
    fmt.Println(tokens) // [うち(10, 12)KNOWN[8024] すもも(0, 3)KNOWN[36163] もも(4, 6)KNOWN[74989]]
}
```
