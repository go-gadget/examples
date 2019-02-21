package main

import (
	"fmt"

	. "github.com/go-gadget/gadget"
	"github.com/go-gadget/gadget/vtree"
)

/*
 * ChildComponent - to be nested in Parent Component
 */
type ChildComponent struct {
	BaseComponent
	Text string
}

var b int

func (g *ChildComponent) Init() {
	g.Text = fmt.Sprintf("Hello-%d", b)
	b++
}

func (g *ChildComponent) Template() string {
	return `
	<button g-click="add_dot" g-value="Text">Add</button>
	`
}

func (g *ChildComponent) Handlers() map[string]Handler {
	return map[string]Handler{
		"add_dot": func(Updates chan Action) {
			g.Text = g.Text + "."
		},
	}
}

func ChildComponentFactory() Component {
	s := &ChildComponent{}
	s.SetupStorage(s)
	return s
}

type ParentComponent struct {
	BaseComponent
	Show bool
}

func (g *ParentComponent) Components() map[string]Builder {
	return map[string]Builder{"child-component": ChildComponentFactory}
}

func (g *ParentComponent) Template() string {
	return `<div><h1>Parent!</h1>
	<button g-click="toggle">Toggle</button>
	<child-component g-if="Show"></child-component><br>
	<child-component></child-component>
	</div>
	`
}

func (g *ParentComponent) Handlers() map[string]Handler {
	return map[string]Handler{
		"toggle": func(Updates chan Action) {
			g.Show = !g.Show
		},
	}
}

func ParentComponentFactory() Component {
	s := &ParentComponent{}
	s.SetupStorage(s)
	return s
}
func main() {
	fmt.Println("Go Go Gadget!")

	g := NewGadget(vtree.Builder())

	go g.MainLoop()
	component := NewComponent(ParentComponentFactory)

	g.Mount(component)
	select {}
}
