package main

import (
	"fmt"

	. "github.com/go-gadget/gadget"
	"github.com/go-gadget/gadget/j"
	"github.com/go-gadget/gadget/vtree"
)

/*
 * This is the most full-fledged example of what go-gadget can do
 * (though it doesn't demonstrate nested components)
 */
type ChildComponent struct {
	BaseComponent
	Text string
}

var b int

func (g *ChildComponent) Init() {
	g.Text = fmt.Sprintf("Child component %d - click me!", b)
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
			g.Text = "*" + g.Text + "*"
		},
	}
}

func ChildComponentFactory() Component {
	s := &ChildComponent{}
	s.SetupStorage(s)
	return s
}

type SampleComponent struct {
	BaseComponent
	Todos     []string
	NewTODO   string
	SomeValue string
	Bar       int
	Show      bool
	Color     string
	i         int
	c1        bool
	c2        bool
}

func (g *SampleComponent) Components() map[string]Builder {
	return map[string]Builder{"child-component": ChildComponentFactory}
}

func (g *SampleComponent) Init() {
	g.Todos = []string{"First entry!", "Second Entry!"}
	g.SomeValue = "Some value"
	g.Color = "red"
	g.NewTODO = "x"
}

func (g *SampleComponent) Template() string {
	return `<div id="rootdiv">
	<b g-class="Color">
	  <g-tag g-value="SomeValue">1</g-tag>
	  <g-tag g-value="Bar">2</g-tag>
	</b>
	<br>
	<input type="text" g-bind="NewTODO">
	<button g-click="add_todo">Add</button>
	<i class="red" g-if="Show">
	 Hello!
	</i>
	<ul>
	  <li g-for="Todo in Todos" g-value="Todo">
	    A todo
	  </li>
	</ul>
	<child-component g-if="c1"></child-component>
	<child-component g-if="c2"></child-component>
	</div>`
}

func (g *SampleComponent) Handlers() map[string]Handler {
	return map[string]Handler{
		"add_todo": func(Updates chan Action) {
			g.Doit()
		},
	}
}

func (g *SampleComponent) Doit() {
	i := g.i
	g.i++

	// can't call SetValue() (it's defined on Component),
	// which means we can't track *which* value changed.
	g.Todos = append(g.Todos, g.NewTODO)
	j.J("add_todo called", g.Todos)
	// g.NewTODO = fmt.Sprintf("And another entry %d", i)
	g.SomeValue = "Completely different"
	g.Bar = i + 1000
	g.Show = i%2 == 1

	switch i % 3 {
	case 0:
		g.Color = "red"
		g.c1 = true
		g.c2 = false
	case 1:
		g.Color = "green"
		g.c1 = false
		g.c2 = true
	case 2:
		g.Color = ""
		g.c1 = false
		g.c2 = false
	}
}

func SampleComponentFactory() Component {
	s := &SampleComponent{}
	// s.BaseComponent = BaseComponent{Storage: s}
	s.SetupStorage(s)
	return s
}

func main() {
	fmt.Println("Go Go Gadget!")

	// Create the framework
	g := NewGadget(vtree.Builder())

	// Start the mainloop
	go g.MainLoop()

	// Create a component
	component := NewComponent(SampleComponentFactory)

	// Mount it on 'nil', making it the main component
	g.Mount(component)
	select {}
}
