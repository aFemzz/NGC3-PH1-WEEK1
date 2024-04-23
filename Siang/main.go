package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	// "math/rand"
	// "net/http"
)

func main() {

	// testing()
	// Looping1()
	// Looping2()
	// Palindrome()
	// xoxo()
	// asterisk()
	// asterisk2()
	// asterisk3()
	// asterisk4()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("List : ")
	fmt.Println("1. Looping 1")
	fmt.Println("2. Looping 2")
	fmt.Println("3. Palindrome")
	fmt.Println("4. xoxo")
	fmt.Println("5. Asterisk 1")
	fmt.Println("6. Asterisk 2")
	fmt.Println("7. Asterisk 3")
	fmt.Println("8. Asterisk 4")
	scanner.Scan()
	fmt.Println()
	i := scanner.Text()

	switch i {
	case "1":
		Looping1()
	case "2":
		Looping2()
	case "3":
		Palindrome()
	case "4":
		xoxo()
	case "5":
		asterisk1()
	case "6":
		asterisk2()
	case "7":
		asterisk3()
	case "8":
		asterisk4()
	}

}

func Looping1() {

	people := []map[string]string{
		{"Name": "Hank", "Age": "50", "Job": "Polisi"},
		{"Name": "Heisenberg", "Age": "52", "Job": "Ilmuwan"},
		{"Name": "Skyler", "Age": "48", "Job": "Akuntan"},
	}

	for _, person := range people {
		fmt.Printf("Hi Perkenalkan, Nama saya %s, umur saya %s, dan saya bekerja sebagai %s\n", person["Name"], person["Age"], person["Job"])
	}
}

func Looping2() {

	slice2 := []float64{1.1, 5.4, 6.7, 9.2, 11.3, 25.2, 33.1}
	slice1 := []float64{1, 5, 7, 8, 10, 24, 33}

	n1 := len(slice1)
	n2 := len(slice2)

	var sum1 float64 = 0
	var sum2 float64 = 0

	sort.Float64s(slice1)

	// for counting the sum
	for i := 0; i < n1; i++ {
		sum1 += (slice1[i])
		sum2 += (slice2[i])
	}

	avg1 := (float64(sum1) / float64(n1))
	avg2 := (float64(sum2) / float64(n2))

	fmt.Println("Sum Slice 1:", sum1, "Sum Slice 2:", sum2)
	fmt.Println("avg Slice 1:", avg1, "avg Slice 2:", avg2)

	var median1 float64
	var median2 float64

	if n1%2 == 0 {
		middleIndex := n1 / 2
		median1 = (slice1[middleIndex-1] + slice1[middleIndex]) / 2
	} else {
		median1 = slice1[n1/2]
	}

	if n2%2 == 0 {
		middleIndex := n2 / 2
		median2 = (slice2[middleIndex-1] + slice2[middleIndex]) / 2
	} else {
		median2 = slice2[n2/2]
	}

	fmt.Println("MEDIAN 1 : ", median1)
	fmt.Println("MEDIAN 2 : ", median2)

}

func Palindrome() {

	var str string
	fmt.Printf("Input Text : ")
	fmt.Scan(&str)

	isPalindrome := true
	lastIdx := len(str) - 1
	for i := 0; i < lastIdx/2 && i < (lastIdx-i); i++ {
		if str[i] != str[lastIdx-i] {
			isPalindrome = false
			break
		}
	}

	// Print the result
	if isPalindrome {
		fmt.Printf("'%s' is palindrome\n", str)
	} else {
		fmt.Printf("'%s' is not palindrome\n", str)
	}

}

func xoxo() {
	str := []int{642, 34, 25, 12, 22, 11, 90}

	for i := 0; i < len(str); i++ { //Loop through the str
		for j := 0; j < len(str)-1; j++ { // Loop through the str but this time we dont want to over-range hence the -1
			if str[j] > str[j+1] { //Check if the current data on the index pointed by "j" is bigger than the index beside it

				//if they are bigger, then do this :

				temp := str[j]    //declare temp variable to store the current data on j index
				str[j] = str[j+1] //move the current data on the current index besides it
				str[j+1] = temp   //swap the current str with the temp

			}

		}

	}
	fmt.Println("This is sorted:", str)

}

func asterisk1() {

	var rows1 int

	// Input the number of rows
	fmt.Print("Enter the number of rows: ")
	fmt.Scanln(&rows1)

	for i := 0; i < rows1; i++ {
		fmt.Println("*")
	}

}

func asterisk2() {

	var rows1 int

	// Input the number of rows
	fmt.Print("Enter the number of rows: ")
	fmt.Scanln(&rows1)

	for i := 0; i < rows1; i++ { //Print the required row that's been inputed
		for j := 0; j < rows1; j++ {
			fmt.Print("*") //Print the asterisk sideways (Dont use println() as its gonna create new line everytime it runs this loop)
		}
		println() //So that the rows are visible (E.g not being displayed sideways which are pointless)
	}

}

func asterisk3() {

	var rows1 int

	// Input the number of rows
	fmt.Print("Enter the number of rows: ")
	fmt.Scanln(&rows1)

	for i := 0; i < rows1; i++ {

		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}

		println()
	}

}

func asterisk4() {

	var rows int

	// Input the number of rows
	fmt.Print("Enter the number of rows: ")
	fmt.Scanln(&rows)

	for i := rows; i >= 1; i-- { //Keep looping until it var i reaches 1
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
