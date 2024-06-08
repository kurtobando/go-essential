package main

import (
	"example.com/go-essential/category"
	"example.com/go-essential/post"
	"fmt"
	"reflect"
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

// accept & return dynamic types
func GetDynamicType(value interface{}) interface{} {
	switch value.(type) {
	case string:
		return value.(string)
	case int:
		return value.(int)
	default:
		return "I do not know"
	}
}

// accept & return dynamic types using generics
func GetDynamicTypeByGeneric[T interface{}](value T) T {
	fmt.Printf("%v: ", reflect.TypeOf(value))
	return value
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

	fmt.Println("**********")
	fmt.Println(GetDynamicType("true"))
	fmt.Println(GetDynamicType(true))
	fmt.Println(GetDynamicType(1))
	fmt.Println(GetDynamicType(-1))
	fmt.Println(GetDynamicType(1.1))

	fmt.Println("**********")
	s := GetDynamicTypeByGeneric("Hello")
	fmt.Println(s)

	i := GetDynamicTypeByGeneric(1)
	fmt.Println(i)

	b := GetDynamicTypeByGeneric(false)
	fmt.Println(b)
}
