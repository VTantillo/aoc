package day5

import "errors"

type Stack interface {
	Push(val string) int
	Pop() (string, error)
	Peek() (string, error)
	IsEmpty() bool
	Size() int
}

type CrateStack struct {
	crates []string
}

func NewCrateStack() *CrateStack {
	return &CrateStack{
		crates: make([]string, 0),
	}
}

func (c *CrateStack) Push(val string) int {
	c.crates = append(c.crates, val)
	return c.Size()
}

func (c *CrateStack) Pop() (string, error) {
	if c.IsEmpty() {
		return "", errors.New("can't pop, stack is empty")
	}

	pop := c.crates[len(c.crates)-1]
	c.crates = c.crates[:len(c.crates)-1]
	return pop, nil
}

func (c *CrateStack) Peek() (string, error) {
	if c.IsEmpty() {
		return "", errors.New("can't peek, stack is empty")
	}

	return c.crates[len(c.crates)-1], nil
}

func (c *CrateStack) IsEmpty() bool {
	return c.Size() == 0
}

func (c *CrateStack) Size() int {
	return len(c.crates)
}
