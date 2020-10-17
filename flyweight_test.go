package flyweight_test

import (
	"flyweight"
	"testing"

	"github.com/stretchr/testify/assert"
)

type RecursiveStruct struct {
	Name  string
	Value []byte

	ChildPtr  *RecursiveStruct
	ChildPtrs []*RecursiveStruct

	Leaf   Leaf
	Leaves []Leaf

	LeafPtr  *Leaf
	LeafPtrs []*Leaf

	NonResetter NonResetter
}

func (r *RecursiveStruct) Reset() { flyweight.Reset(r) }

type Leaf struct {
	Name string
}

func (r *Leaf) Reset() { flyweight.Reset(r) }

type NonResetter struct {
	Name string
}

func TestFlyweightRecursive(t *testing.T) {
	item := &RecursiveStruct{
		Name:  "top level",
		Value: []byte("hello world"),

		ChildPtr: &RecursiveStruct{Name: "first child", Value: []byte("hello child")},
		ChildPtrs: []*RecursiveStruct{
			{Name: "first child"},
		},

		Leaf: Leaf{Name: "first leaf"},
		Leaves: []Leaf{
			{Name: "first leaf arr"},
		},
		LeafPtr: &Leaf{Name: "second leaf"},
		LeafPtrs: []*Leaf{
			{Name: "first leaf ptr arr"},
		},
		NonResetter: NonResetter{Name: "non-resetter"},
	}

	item.Reset()

	assert.Equal(t, "", item.Name)
	assert.Equal(t, []byte{}, item.Value)

	assert.Equal(t, "", item.ChildPtr.Name)
	assert.Equal(t, []byte{}, item.ChildPtr.Value)
	assert.Equal(t, 0, len(item.ChildPtrs))

	item.ChildPtrs = item.ChildPtrs[:cap(item.ChildPtrs)]
	assert.Equal(t, 1, len(item.ChildPtrs))

	underlyingChild := item.ChildPtrs[0]
	assert.Equal(t, "", underlyingChild.Name)

	assert.Equal(t, "", item.Leaf.Name)
	assert.Equal(t, "", item.LeafPtr.Name)

	assert.Equal(t, 0, len(item.Leaves))
	item.Leaves = item.Leaves[:cap(item.Leaves)]
	assert.Equal(t, 1, len(item.Leaves))
	underlyingLeaf := item.Leaves[0]
	assert.Equal(t, "", underlyingLeaf.Name)

	assert.Equal(t, 0, len(item.LeafPtrs))
	item.LeafPtrs = item.LeafPtrs[:cap(item.LeafPtrs)]
	assert.Equal(t, 1, len(item.LeafPtrs))
	underlyingLeafPtr := item.LeafPtrs[0]
	assert.Equal(t, "", underlyingLeafPtr.Name)

	assert.Equal(t, "", item.NonResetter.Name)
}

func TestEmptyStruct(t *testing.T) {
	item := &RecursiveStruct{}
	item.Reset()
}
