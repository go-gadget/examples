package main

import (
	"fmt"

	. "github.com/go-gadget/gadget"
	"github.com/go-gadget/gadget/vtree"
)

type User struct {
	Name    string
	Posts   []string
	Address string
}

/*
 * Steps:
 * - print :id
 * - hardcoded list of users (ids)
 * - <router-link component thingy?>
 */
func main() {
	fmt.Println("Go Go Gadget!")
	AppComponent := GenerateComponent(`<div><h1>App</h1>
    <ul>
      <li> <router-link To="User" Id="123"> John Doe </router-link></li>
      <li> <router-link To="UserProfile" Id="234"> Jane Doe Profile </router-link></li>
      <li> <router-link To="UserPosts" Id="345"> Shane Doe Posts</router-link></li>
    </ul>
    <router-view></router-view></div>`, nil, nil)
	HomeComponent := GenerateComponent("<div><h2>Home</h2></div>", nil, nil)
	// For now handle RouteParams as props. In a sense this is more consistent and allows you to call the component
	// without a route
	UserComponent := GenerateComponent(`<div><h2>User</h2>Id: <span g-value="id">?</span><router-view></router-view></div>`, nil, []string{"id"})
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

	g.Mount(g.NewComponent(AppComponent))
	go g.MainLoop()
	select {}
}
