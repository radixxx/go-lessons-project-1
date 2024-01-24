package main

type Customer struct {
	id   int
	name string
}

func (c *Customer) GetId() int {
	return c.id
}

func (c *Customer) GetName() string {
	return c.name
}
