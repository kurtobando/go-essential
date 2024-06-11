package arrays

import "fmt"

type TLD struct {
	name string
	ext  string
}

func FixedArray() {
	var colors [4]string

	colors[0] = "blue"
	colors[1] = "red"
	colors[2] = "yellow"
	colors[3] = "green"

	fmt.Println("############################")
	fmt.Println("array and slice")
	fmt.Println("############################")

	fmt.Println(colors)
	fmt.Println(len(colors))
	fmt.Println(colors[0:3])
	fmt.Println(colors[0:4])
	fmt.Println(colors[2:])

	fmt.Println("############################")
	fmt.Println("array from slice copy")
	fmt.Println("############################")

	c := colors[0:]
	fmt.Println(c)
	fmt.Println(len(c))
	c[0] = "white"
	fmt.Println(colors)
	fmt.Println(c)
}

func TypeAnyArray() {
	fmt.Println("############################")
	fmt.Println("array with 'any' type")
	fmt.Println("############################")

	var country = [4]any{"ph", "cn", "au", "eu"}
	fmt.Println(country)
	country[0] = 1234
	fmt.Println(country)
	country[1] = true
	fmt.Println(country)
}

func StructArray() {
	var domains = [3]TLD{
		{
			name: "abc",
			ext:  "com",
		},
		{
			name: "google",
			ext:  "com",
		},
		{
			name: "facebook",
			ext:  "com",
		},
	}

	fmt.Println("############################")
	fmt.Println("array of struct")
	fmt.Println("############################")

	fmt.Println(domains)
	fmt.Println(domains[2])
	domains[0] = TLD{name: "youtube", ext: "com"}
	fmt.Println(domains)
	fmt.Println(domains[1:2])
	fmt.Println(domains[:1])
	fmt.Println(domains[0].name)
	fmt.Println(domains[0].ext)

}

func DynamicArray() {
	fmt.Println("############################")
	fmt.Println("dynamic array")
	fmt.Println("############################")

	var dynamicArrayNames []string
	fmt.Println(dynamicArrayNames)

	// https://go.dev/doc/effective_go#append
	dynamicArrayNames = append(dynamicArrayNames, "john")
	updatedDynamicArrayNames := append(dynamicArrayNames, "mary")

	fmt.Println(dynamicArrayNames)
	fmt.Println(updatedDynamicArrayNames, dynamicArrayNames)
}
