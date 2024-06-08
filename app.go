package main

import (
	"example.com/go-essential/category"
	"example.com/go-essential/post"
	"fmt"
	"strings"
)

// initial interface
type Getter interface {
	Get() string
}

type Setter interface {
	Set(t string, d string)
}

// find use case from the created interface above
func Save(setter Setter, t string, d string) {
	setter.Set(t, d)
}

func Get(getter Getter) string {
	return getter.Get()
}

// method, and interface
func main() {
	p := post.New("Short Post Title", "Long Post Description")
	po := p.Get()
	fmt.Println(po)

	c := category.New("Short Category Title", "Long Category Description")
	co := c.Get()
	fmt.Println(co)

	Save(p, "A", "B")
	fmt.Println(Get(p))

	Save(c, "C", "D")
	fmt.Println(Get(c))

	fmt.Println("**********")
	fmt.Println(strings.Count("demos", ""))
	fmt.Println(len("demos"))
	fmt.Println(len("demo"))

}
