// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package client

import (
	"github.com/m3db/m3db/interfaces/m3db"
	"github.com/m3db/m3db/pool"
)

type readerSliceOfSlicesIteratorPool interface {
	// Init pool
	Init()

	// Get an iterator
	Get() *readerSliceOfSlicesIterator

	// Put an iterator
	Put(it *readerSliceOfSlicesIterator)
}

type poolOfReaderSliceOfSlicesIterator struct {
	pool m3db.ObjectPool
}

func newReaderSliceOfSlicesIteratorPool(size int) readerSliceOfSlicesIteratorPool {
	p := pool.NewObjectPool(size)
	return &poolOfReaderSliceOfSlicesIterator{p}
}

func (p *poolOfReaderSliceOfSlicesIterator) Init() {
	p.pool.Init(func() interface{} {
		return newReaderSliceOfSlicesIterator(nil, p)
	})
}

func (p *poolOfReaderSliceOfSlicesIterator) Get() *readerSliceOfSlicesIterator {
	return p.pool.Get().(*readerSliceOfSlicesIterator)
}

func (p *poolOfReaderSliceOfSlicesIterator) Put(it *readerSliceOfSlicesIterator) {
	p.pool.Put(it)
}