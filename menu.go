package main

import "fmt"

func main() {
	var showmenu = menu()
	showmenu()
}
func menu() func() {
	var message string = ""
	fun := func() {
		for {
			fmt.Print(">>")
			fmt.Scan(&message)
			switch message {
			case "help":
				fmt.Println("help")
			case "version":
				fmt.Println("version")
			case "quit":
				return
			default:
				fmt.Println("error command")
			}
		}
	}
	return fun

}
