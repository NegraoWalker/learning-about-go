package main

import "fmt"

type Product struct {
	id    string
	title string
	price float64
}

func main() {

	// Time to practice what you learned!

	// 1) Create a new array (!) that contains three hobbies you have
	// 		Output (print) that array in the command line.
	// 2) Also output more data about that array:
	//		- The first element (standalone)
	//		- The second and third element combined as a new list

	// 3) Create a slice based on the first element that contains
	//		the first and second elements.
	//		Create that slice in two different ways (i.e. create two slices in the end)

	// 4) Re-slice the slice from (3) and change it to contain the second
	//		and last element of the original array.
	// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
	// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
	// 7) Bonus: Create a "Product" struct with title, id, price and create a
	//		dynamic list of products (at least 2 products).
	//		Then add a third product to the existing list of products.

	//1:
	hobbies1 := [3]string{"Play tennis", "Go out with my friends", "Reading"}
	fmt.Println(hobbies1)
	//2:
	fmt.Println(hobbies1[0])
	fmt.Println(hobbies1[1:3])
	//3:
	hobbies2 := hobbies1[:2]
	fmt.Println(hobbies2)
	//4:
	fmt.Println(cap(hobbies2))
	hobbies2 = hobbies1[1:3]
	fmt.Println(hobbies2)
	//5:
	courseGoals := []string{"Learn Go!", "Learn all the basics"}
	fmt.Println(courseGoals)
	//6:
	courseGoals[1] = "Learn all the detaills!"
	courseGoals = append(courseGoals, "Learn all the basics!")
	fmt.Println(courseGoals)
	//7:
	products := []Product{
		Product{"first-product", "A first product", 12.99},
		Product{"second-product", "A second product", 129.99},
	}

	fmt.Println(products)
	newProduct := Product{"third-product", "A third product", 15.99}
	products = append(products, newProduct)
	fmt.Println(products)
}
