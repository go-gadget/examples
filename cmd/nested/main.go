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
	Text string
}

var b int

func (g *ChildComponent) Init() {
	g.Text = fmt.Sprintf("Hello-%d", b)
	b++
}

func (g *ChildComponent) Components() map[string]Builder {
	return nil
}

func (g *ChildComponent) Data() interface{} {
	return g
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
	return s
}

type ParentComponent struct {
	Show bool
}

func (g *ParentComponent) Init() {
}

func (g *ParentComponent) Components() map[string]Builder {
	return map[string]Builder{"child-component": ChildComponentFactory}
}

func (g *ParentComponent) Data() interface{} {
	return g
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
	return s
}
func main() {
	fmt.Println("Go Go Gadget!")

	g := NewGadget(vtree.Builder())

	go g.MainLoop()
	component := g.BuildComponent(ParentComponentFactory)

	g.Mount(component, nil)
	select {}
}
