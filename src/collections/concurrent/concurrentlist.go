package concurrent

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
)

type changeType int

const (
	add    changeType = iota
	remove changeType = iota
)

type ConcurrentList struct {
	data          []interface{}
	mux           sync.Mutex
	chItemChanged chan changeType
}

func NewList[T any]() *ConcurrentList {
	return &ConcurrentList{
		data:          make([]interface{}, 0),
		chItemChanged: make(chan changeType, 1),
	}
}

func (c *ConcurrentList) notifyItemChanged(t changeType) {
	select {
	case c.chItemChanged <- t:
	default:
	}
}

func (c *ConcurrentList) Add(v interface{}) {
	c.mux.Lock()
	c.data = append(c.data, v)
	c.mux.Unlock()

	c.notifyItemChanged(add)
}

func (c *ConcurrentList) AddRange(v []interface{}) {
	c.mux.Lock()
	c.data = append(c.data, v...)
	c.mux.Unlock()

	c.notifyItemChanged(add)
}

func (c *ConcurrentList) Clear() {
	c.mux.Lock()
	c.data = make([]interface{}, 0)
	c.mux.Unlock()

	c.notifyItemChanged(remove)
}

func (c *ConcurrentList) Remove(v interface{}) bool {
	c.mux.Lock()
	result := c.removeWithoutLock(v)
	c.mux.Unlock()

	c.notifyItemChanged(remove)

	return result
}
func (c *ConcurrentList) removeWithoutLock(v interface{}) bool {
	newslice := make([]interface{}, 0, len(c.data))
	isexist := false

	for _, d := range c.data {
		if reflect.DeepEqual(d, v) && !isexist {
			isexist = true
		} else {
			newslice = append(newslice, d)
		}
	}

	if isexist {
		c.data = newslice
		return true
	}

	return false
}

func (c *ConcurrentList) RemoveRange(index, count int) {
	c.mux.Lock()
	c.removeRangeWithoutLock(index, count)
	c.mux.Unlock()

	c.notifyItemChanged(remove)
}
func (c *ConcurrentList) removeRangeWithoutLock(index, count int) {
	newslice := make([]interface{}, 0, len(c.data)-1)
	newslice = append(newslice, c.data[0:index]...)
	newslice = append(newslice, c.data[index+count:len(c.data)]...)

	c.data = newslice
}

// just get ,not remove
func (c *ConcurrentList) Get(index int) (interface{}, error) {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.getWithoutLock(index)
}

func (c *ConcurrentList) getWithoutLock(index int) (interface{}, error) {
	var defaultT interface{}
	if len(c.data) <= index {
		return defaultT, errors.New(fmt.Sprintf("index: %d out of bound,max len: %d", index, len(c.data)))
	}
	return c.data[index], nil
}

// get all and not remove
func (c *ConcurrentList) GetAll() []interface{} {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.data
}

// get one item and remove
func (c *ConcurrentList) Take(index int) (interface{}, error) {
	c.mux.Lock()
	defer c.mux.Unlock()

	d, err := c.getWithoutLock(index)
	if err != nil {
		return d, err
	}
	if !c.removeWithoutLock(d) {
		return d, errors.New("remove fail")
	}
	return d, nil
}
func (c *ConcurrentList) TakeAll() []interface{} {
	c.mux.Lock()
	defer c.mux.Unlock()

	r := c.data
	c.data = make([]interface{}, 0)
	return r
}

// when no item add,it will block
func (c *ConcurrentList) TakeAllBlock() []interface{} {
	for {
		if t := <-c.chItemChanged; t == add {
			break
		}
	}
	return c.TakeAll()
}
