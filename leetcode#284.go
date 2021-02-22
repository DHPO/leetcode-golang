package main

type Iterator struct {
}

func (this *Iterator) hasNext() bool {
	// Returns true if the iteration has more elements.
}

func (this *Iterator) next() int {
	// Returns the next element in the iteration.
}

type PeekingIterator struct {
	iter         *Iterator
	reg          int
	regAvailable bool
}

func Constructor(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{iter, 0, false}
}

func (this *PeekingIterator) hasNext() bool {
	return this.regAvailable || this.iter.hasNext()
}

func (this *PeekingIterator) next() int {
	if this.regAvailable {
		defer func() { this.regAvailable = false }()
		return this.reg
	} else {
		return this.iter.next()
	}
}

func (this *PeekingIterator) peek() int {
	if this.regAvailable {
		return this.reg
	} else {
		this.reg = this.iter.next()
		this.regAvailable = true
		return this.reg
	}
}
