package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	userNames := make([]string, 2) //Declaração de um slice
	userNames[0] = "Walker"
	userNames = append(userNames, "William")
	fmt.Println(userNames)

	courseRatings := make(map[string]float64, 3)
	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.9

	fmt.Println(courseRatings)

	courseRatings2 := make(floatMap, 2)
	courseRatings2["go"] = 4.7
	courseRatings2["react"] = 4.8

	courseRatings2.output()

	//loops para slice e maps:
	for key, value := range courseRatings {
		fmt.Println("index:", key)
		fmt.Println("value:", value)
	}

	for index, value := range userNames {
		fmt.Println("index:", index)
		fmt.Println("value:", value)
	}

}
