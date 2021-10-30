package main

type Node interface {
	Execute()
}

type And struct {
	Value1 Node
	Value2 Node
}

func (node *And) Execute() {

}

func main() {
	var node Node
	node = &And{}
	println(node)
}
