package main

import (
	"fmt"

	. "github.com/go-gadget/gadget"
	"github.com/go-gadget/gadget/vtree"
)

func main() {
	fmt.Println("Go Go Gadget!")
	AppComponent := GenerateComponent("<div><h1>App</h1></div>", nil, nil)
	HomeComponent := GenerateComponent("<div><h2>Home</h2></div>", nil, nil)
	UserComponent := GenerateComponent("<div><h2>User</h2><router-view></router-view></div>", nil, nil)
	UserProfile := GenerateComponent("<div><h2>Profile</h2></div>", nil, nil)
	UserPosts := GenerateComponent("<div><h2>Posts</h2></div>", nil, nil)

	g := NewGadget(vtree.Builder())
	g.Router(Router{
		Route{
			Path:      "/",
			Name:      "Home",
			Component: HomeComponent,
		},
		Route{
			Path:      "/user/:id",
			Name:      "User",
			Component: UserComponent,
			Children: []Route{
				Route{
					Path:      "profile",
					Name:      "UserProfile",
					Component: UserProfile,
				},
				Route{
					Path:      "posts",
					Name:      "UserPosts",
					Component: UserPosts,
				},
			},
		},
	})

	go g.MainLoop()
	g.Mount(NewComponent(AppComponent))
	select {}
}
