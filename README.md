# mecab-tokenizer

[![GoDoc](https://godoc.org/github.com/hatajoe/mecab-tokenizer?status.svg)](https://godoc.org/github.com/hatajoe/mecab-tokenizer)
[![Circle CI](https://circleci.com/gh/hatajoe/mecab-tokenizer.svg?style=svg)](https://circleci.com/gh/hatajoe/mecab-tokenizer)
[![Coverage Status](https://coveralls.io/repos/github/hatajoe/mecab-tokenizer/badge.svg?branch=master)](https://coveralls.io/github/hatajoe/mecab-tokenizer?branch=master)

```
This is under development
```

mecab-tokenizer is text tokenizer wrapped bluele/mecab-golang

## Dependencies

- [MeCab: Yet Another Part-of-Speech and Morphological Analyzer](https://mecab.googlecode.com/svn/trunk/mecab/doc/index.html#install)
- [bluele/mecab-golang: A golang wrapper for mecab.](https://github.com/bluele/mecab-golang)

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
    myTokenizer, err := tokenizer.NewTokenizer()
    if err != nil {
        panic(err)
    }
    if err := myTokenizer.Tokenize(
        "すもももももももものうち。", 
        tokenizer.OnParsed(func (n *tokenizer.Node) error {
            fmt.Println(n.DistinctNoun().Sort()) // [うち すもも もも]
        }),
    ); err != nil {
        panic(err)
    }
}
```
