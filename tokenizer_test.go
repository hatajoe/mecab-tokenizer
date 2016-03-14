package tokenizer_test

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/hatajoe/mecab-tokenizer"
)

var myTokenizer *tokenizer.Tokenizer

func TestMain(m *testing.M) {
	var err error
	myTokenizer, err = tokenizer.NewTokenizer()
	if err != nil {
		os.Exit(1)
	}
	defer myTokenizer.Destroy()
	os.Exit(m.Run())
}

func TestTokenizeAsNoun1(t *testing.T) {
	if err := myTokenizer.Tokenize("寿司食べたい。", tokenizer.OnParsed(func(n *tokenizer.Node) error {
		nouns := n.DistinctNoun()
		if len(nouns) != 1 {
			return errors.New("err: length of nouns is expecting 1")
		}
		if nouns[0] != "寿司" {
			return fmt.Errorf("err: nonus is %s expecting `寿司`", nouns[0])
		}
		return nil
	})); err != nil {
		t.Fatal(err)
	}
}

func TestTokenizeAsNoun2(t *testing.T) {
	if err := myTokenizer.Tokenize("寿司と焼き肉食べたい。", tokenizer.OnParsed(func(n *tokenizer.Node) error {
		nouns := n.DistinctNoun()
		if len(nouns) != 2 {
			return errors.New("err: length of nouns is expecting 2")
		}
		if nouns[0] != "寿司" {
			return fmt.Errorf(fmt.Sprintf("err: nonus is %s expecting `寿司`", nouns[0]))
		}
		if nouns[1] != "焼き肉" {
			return fmt.Errorf(fmt.Sprintf("err: nonus is %s expecting `焼き肉`", nouns[1]))
		}
		return nil
	})); err != nil {
		t.Fatal(err)
	}
}

func TestTokenizeAsNoun3(t *testing.T) {
	if err := myTokenizer.Tokenize("すもももももももものうち。", tokenizer.OnParsed(func(n *tokenizer.Node) error {
		nouns := n.DistinctNoun().Sort()
		if len(nouns) != 3 {
			return errors.New("err: length of nouns is expecting 3")
		}
		if nouns[0] != "うち" {
			return fmt.Errorf(fmt.Sprintf("err: nonus is %s expecting `うち`", nouns[0]))
		}
		if nouns[1] != "すもも" {
			return fmt.Errorf(fmt.Sprintf("err: nonus is %s expecting `すもも`", nouns[1]))
		}
		if nouns[2] != "もも" {
			return fmt.Errorf(fmt.Sprintf("err: nonus is %s expecting `もも`", nouns[2]))
		}
		return nil
	})); err != nil {
		t.Fatal(err)
	}
}

func BenchmarkTokenizeAsNoun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := tokenize(); err != nil {
			b.Fatal(b)
		}
	}
}

func tokenize() error {
	return myTokenizer.Tokenize("すもももももももものうち。", tokenizer.OnParsed(func(n *tokenizer.Node) error {
		n.DistinctNoun().Sort()
		return nil
	}))
}
