package main

import "fmt"

func main() {
	fmt.Println("ARRAYS")

	var animeList [2]string
	fmt.Println("Anime length Before appending: ", len(animeList))

	animeList[0] = "jujustu kaisan"
	animeList[1] = "Demon Slayer"
	fmt.Println("Anime List: ", animeList)
	fmt.Println("Anime length After appending: ", len(animeList))

	var songList = [4]string{"Starboy - weeknd", "Tu hai kahan - AUR", "Calling - Metro Boomin", "Cirle - Post Malone"}
	fmt.Println("Song List: ", songList)
}
