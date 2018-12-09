package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"os/exec"
	"runtime"
	"time"
	"math/rand"
	"strconv"
	"reflect"
)

var username string //stores player name
var player_location = [2]int{0,0} //stores player location on the map as 2 coordinates
var player_hp = 50 //Player starts with 10 hp that will change during the game.
var player_max_hp = 50 //Stores the value of maximum possible health for healing purposes.
var player_str = 0 //Strength currently represents player level. Certain actions might increase the base str of the player.
var player_weapon = 0 //This indicates if or which weapon the player has in the hand. This will be used to find the weapon information from the weapon array.

var map_size = 5 //holds the size of the game field, used for map printing function loop.
var game_map [5][5]int //stores playing field. Map currently will consist of randomly filled fields just for test purposes.
var available_direction [4]string //will store all the available directions of movement at any point in the game.

var directions = [4]string{"up", "right", "down", "left"}

//
//
//
//
var story = "Though they were once pioneers of video game storytelling,"+ 
"traditional RPGs have long relied on cliches that were well-worn"+ 
"even in the early 1990s. That/'s not to say these tropes are necessarily a bad thing,"+ 
"but they do make these games more of a predictable comfort food than an exciting unknowable adventure."
//
//
//
//

var clear map[string]func() //create a map for storing clear funcs

//--------------------------------------------------------------------------------
//functions required for the clear screen

func init() {
    clear = make(map[string]func()) //Initialize it
    clear["linux"] = func() { 
        cmd := exec.Command("clear") //Linux example, its tested
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
    clear["darwin"] = func() { 
        cmd := exec.Command("clear") //Darwin example, its tested
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
    clear["windows"] = func() {
        cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested 
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
}

func CallClear() {
    value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
    if ok { //if we defined a clear func for that platform:
        value()  //we execute it
    }
}

//--------------------------------------------------------------------------------

func initMap() {
	for i := 0; i < map_size; i++ {
		for j := 0; j < map_size; j++ {
			s1 := rand.NewSource(time.Now().UnixNano())
    		r1 := rand.New(s1)
			game_map[i][j] = r1.Intn(3)
		}
	}
}

func printMap() {
	for i := 0; i < map_size; i++ {
		for j := 0; j < map_size; j++ {
			if player_location[0] == i && player_location[1] == j {
				print_color_text("x ", "red")
			} else {
				print_color_text(strconv.Itoa(game_map[i][j]), "white") //sends the value of the array location as color and text
				fmt.Print(" ")
			}
			
		}
		fmt.Println()
	}
}

func login() {
	fmt.Print("Enter your username: ")
	reader1 := bufio.NewReader(os.Stdin)
	username, _ = reader1.ReadString('\n')
	username = strings.TrimRight(username, "\r\n")
}

func print_color_text(text string, color string) {
	switch color {
		case "red", "bad", "danger", "0":
			fmt.Print("\x1b[31;1m")
		case "white", "default", "2": 
			fmt.Print("\x1b[0m")
		case "orange", "warning":
			fmt.Print("\x1b[91m")
		case "good", "green", "1":
			fmt.Print("\x1b[32m")
	}
	fmt.Print(text)
	fmt.Print("\x1b[0m")
	
}
func monsters() {
	
}

func decode_map() {
	
}

func clear_screen() {
	cmd := exec.Command("clear") //Linux example, its tested
    cmd.Stdout = os.Stdout
    cmd.Run()
}

func check_available_directions() {
	for i:=0; i<4; i++ {
		available_direction[i] = ""
	}

	if player_location[1] > 0 {
		available_direction[0] = "left"
	}

	if player_location[1] < map_size-1 {
		available_direction[1] = "right"
	}

	if player_location[0] > 0 {
		available_direction[2] = "up"
	}

	if player_location[0] < map_size-1 {
		available_direction[3] = "down"
	}
}

func move(direction string) {
	switch direction {
		case "up":
		case "right":
		case "down":
		case "left":
	}
}

func in_array(val interface{}, array interface{}) (exists bool, index int) {

    exists = false
    index = -1
    
    switch reflect.TypeOf(array).Kind() {
    default:
        s := reflect.ValueOf(array)

        for i := 0; i < s.Len(); i++ {

            if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
                index = i
                exists = true
                return
            }
        }
    }
    return
}


func main() {
	login()
	initMap()
	CallClear()
	fmt.Println("Greetings "+username+"!\n")
	fmt.Println(story+"\n")
	fmt.Println("Objective\n")

	printMap()

	fmt.Println("Name:", username)
	fmt.Println("HP:", player_hp)
	fmt.Println("Max HP:", player_max_hp)
	fmt.Println("STR:", player_str)
	fmt.Println("Weapon:", player_weapon)
	fmt.Println("Location: X:", player_location[0],"Y:", player_location[1])


	sum := 1
	for sum < 20 {


	check_available_directions()

	for i:=0; i<4; i++ {
		if available_direction[i] != "" {
			fmt.Print("You can go: ")
			print_color_text(available_direction[i]+"\n", "green")
		}
		
	}

	fmt.Print("Enter direction: ")
	reader1 := bufio.NewReader(os.Stdin)
	direction, _ := reader1.ReadString('\n')
	direction = strings.TrimRight(direction, "\r\n")

	direction = strings.ToLower(direction)

	switch direction {
		case "right":
			right, _:=in_array(direction, available_direction)
			if right == true {
				player_location[1] = player_location[1] + 1
			} else {
				print_color_text("You can't go this way!\n", "red")
			}
		case "left":
			left, _:=in_array(direction, available_direction)
			if left == true {
				player_location[1] = player_location[1] - 1
			} else {
				print_color_text("You can't go this way!\n", "red")
			}
		case "up":
			up, _:=in_array(direction, available_direction)
			if up == true {
				player_location[0] = player_location[0] - 1
			} else {
				print_color_text("You can't go this way!\n", "red")
			}
		case "down":
			down, _:=in_array(direction, available_direction)
			if down == true {
				player_location[0] = player_location[0] + 1
			} else {
				print_color_text("You can't go this way!\n", "red")
			}
	}


	printMap()

	sum = sum + 1
	}

}