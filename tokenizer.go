package tokenizer

import (
	"sort"
	"strings"

	"github.com/bluele/mecab-golang"
)

const (
	featureTypeTokenIndex = iota
	featureTypeNounKey    = "名詞"
)

// Tokenizer is mecab-golang wrapper
type Tokenizer struct {
	m *mecab.MeCab
	t *mecab.Tagger
}

type OnParsed func (*Node) error

// NewTokenizer initialize mecab-golang
func NewTokenizer(opt ...string) (*Tokenizer, error) {
	m, err := mecab.New(opt...)
	if err != nil {
		return nil, err
	}
	t, err := m.NewTagger()
	if err != nil {
		return nil, err
	}
	return &Tokenizer{
		m: m,
		t: t,
	}, nil
}

// Destroy tokenizer
func (m *Tokenizer) Destroy() {
	if m.m != nil {
		m.m.Destroy()
	}
	if m.t != nil {
		m.t.Destroy()
	}
}

// Tokenize create *mecab.Lattice and call OnParsed with *Node args
func (m *Tokenizer) Tokenize(text string, cb OnParsed) error {
	lt, err := m.m.NewLattice(text)
	if err != nil {
		return err
	}
	defer lt.Destroy()
	if cb != nil {
		if err := cb(&Node{m.t.ParseToNode(lt)}); err != nil {
			return err
		}
	}
	return nil
}

// Node is alias of *mecab.Node
type Node struct {
	*mecab.Node
}

// DistinctNoun return distinct noun surfaces
func (m *Node) DistinctNoun() (nouns Tokens) {
	distinctNouns := map[string]bool{}
	for {
		if m.Next() == mecab.StopIteration {
			break
		}
		features := strings.Split(m.Feature(), ",")
		if len(features) <= 0 {
			continue
		}
		if features[featureTypeTokenIndex] != featureTypeNounKey {
			continue
		}
		if _, ok := distinctNouns[m.Surface()]; !ok {
			distinctNouns[m.Surface()] = true
			nouns = append(nouns, m.Surface())
		}
	}
	return nouns
}

// Tokens is node surfaces
type Tokens []string

// Sort as mecab.Node.Surface ASC
func (m Tokens) Sort() Tokens {
	sort.Sort(&m)
	return m
}

// Len is implementation of sort.Sort
func (m Tokens) Len() int {
	return len(m)
}

// Swap is implementation of sort.Sort
func (m Tokens) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

// Less is implementation of sort.Sort
func (m Tokens) Less(i, j int) bool {
	return m[i] < m[j]
}
