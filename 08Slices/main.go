package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("SLICES")

	var animeList = []string{"jujustu kaisan", "Demon Slayer"}
	fmt.Println("Anime Slices List: ", animeList)
	fmt.Printf("Anime List type is %T\n", animeList)

	// Append in Slices
	animeList = append(animeList, "Naruto")
	fmt.Println("Anime Slices List After appending: ", animeList)

	// Slicing slices
	animeList = animeList[1:]
	fmt.Println("Anime Slices List After slicing: ", animeList)

	// create slice using make
	songList := make([]string, 3)

	songList[0] = "Starboy - weeknd"
	songList[1] = "Tu hai kahan - AUR"
	songList[2] = "Calling - Metro Boomin"

	// if your can not directly assign out of index for example songList[3]="Circle - Post Malone"
	// songList[3] = "Circle - Post Malone" // this line will give you " panic : runtime error: index out of range [3] with length 3"
	fmt.Println("Song List: ", songList)

	// but you can use append to append to the slice
	songList = append(songList, "Circle - Post Malone")

	fmt.Println("Song List: ", songList)

	fmt.Println("Is Sorted: ", sort.StringsAreSorted(songList))
	sort.Strings(songList)
	fmt.Println("After sorting: ", songList)
	fmt.Println("Is Sorted: ", sort.StringsAreSorted(songList))

	// remove an element from slice
	var courseList = []string{"C", "C++", "Java", "Python", "Go", "C#"}
	fmt.Println("Course List: ", courseList)
	index := 4
	courseList = append(courseList[:index], courseList[index+1:]...)
	fmt.Println("Course List after removing element: ", courseList)
}
