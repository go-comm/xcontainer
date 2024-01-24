package sortarray

import (
	"sort"

	"github.com/go-comm/xcontainer/xslice"
)

func New(data sort.Interface) *Index {
	index := &Index{data: data}
	if in, ok := data.(Interface); ok {
		index.opt = in
	}
	n := data.Len()
	if n > 0 {
		index.index = make([]int, n)
		for i := 0; i < n; i++ {
			index.index[i] = i
		}
	}
	sort.Sort(index)
	return index
}

type Interface interface {
	sort.Interface

	Push(x interface{})
	Pop() interface{}
}

type Index struct {
	data  sort.Interface
	opt   Interface
	index []int
}

func (n *Index) Swap(i, j int) {
	n.index[i], n.index[j] = n.index[j], n.index[i]
}

func (n *Index) Len() int {
	return len(n.index)
}

func (n *Index) Less(i, j int) bool {
	return n.data.Less(n.index[i], n.index[j])
}

func (n *Index) At(i int) int {
	return n.index[i]
}

func (n *Index) Pop() {
	n.index = n.index[:n.Len()-1]
	n.opt.Pop()
}

func (n *Index) Push(x interface{}) {
	v := n.data.Len()
	n.opt.Push(x)
	n.index = append(n.index, v)
}

func Push(index *Index, x interface{}) {
	index.Push(x)
	xslice.FixSortInterface(index, index.Len()-1)
}

func Remove(index *Index, i int) {
	length := index.Len()
	if i >= length {
		return
	}
	if i == length-1 {
		index.Pop()
		return
	}
	index.Swap(i, length-1)
	index.Pop()
	xslice.FixSortInterface(index, i)
}

func Foreach(index *Index, f func(int) (exit bool)) {
	xslice.Foreach(index.Len(), func(i int) (exit bool) { return f(index.At(i)) })
}

func (n *Index) Foreach(f func(int) (exit bool)) {
	Foreach(n, f)
}

func Find(index *Index, f func(int) bool) int {
	return xslice.Find(index.Len(), func(i int) bool { return f(index.At(i)) })
}

func (n *Index) Find(f func(int) bool) int {
	return Find(n, f)
}

func FindN(index *Index, i int, j int, f func(int) bool) int {
	return xslice.FindN(index.Len(), i, j, func(k int) bool { return f(index.At(k)) })
}

func (n *Index) FindN(i int, j int, f func(int) bool) int {
	return FindN(n, i, j, f)
}

func BinarySearch(index *Index, f func(int) bool) int {
	return xslice.BinarySearch(index.Len(), func(i int) bool { return f(index.At(i)) })
}

func (n *Index) BinarySearch(f func(int) bool) int {
	return BinarySearch(n, f)
}

func BinarySearchN(index *Index, i int, j int, f func(int) bool) int {
	return xslice.BinarySearchN(index.Len(), i, j, func(k int) bool { return f(index.At(k)) })
}

func (n *Index) BinarySearchN(i int, j int, f func(int) bool) int {
	return BinarySearchN(n, i, j, f)
}

func BinaryFind(index *Index, cmp func(int) int) (int, bool) {
	return xslice.BinaryFind(index.Len(), func(i int) int { return cmp(index.At(i)) })
}

func (n *Index) BinaryFind(cmp func(int) int) (int, bool) {
	return BinaryFind(n, cmp)
}

func BinaryFindN(index *Index, i int, j int, cmp func(int) int) (int, bool) {
	return xslice.BinaryFindN(index.Len(), i, j, func(k int) int { return cmp(index.At(k)) })
}

func (n *Index) BinaryFindN(i int, j int, cmp func(int) int) (int, bool) {
	return BinaryFindN(n, i, j, cmp)
}
