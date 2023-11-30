package day5

type Stack interface {
	Push(val string) int
	Pop() string
	Peek() string
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

func (c *CrateStack) IsEmpty() bool {
	return c.Size() == 0
}

func (c *CrateStack) Size() int {
	return len(c.crates)
}
