package maps

import "fmt"

func DynamicMaps() {
	fmt.Println("############################")
	fmt.Println("Maps")
	fmt.Println("############################")

	var brands = map[string]string{}
	var brandsInitial = make(map[string]string)

	fmt.Println(brands)

	brands["facebook"] = "facebook.com"
	brands["youtube"] = "youtube.com"
	brands["google"] = "google.com"

	fmt.Println(brands)
	fmt.Println(brands["facebook"])

	fmt.Println(brandsInitial)
	brandsInitial["twitter"] = "twitter.com"
	fmt.Println(brandsInitial)
}
