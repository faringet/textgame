package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Player struct {
	inventory []string
	location  string
}

type Location struct {
	name        string
	description string
	paths       map[string]string
	items       []string
}

type Game struct {
	locations map[string]Location
	player    Player
}

func NewGame() *Game {
	kitchen := Location{
		name:        "кухня",
		description: "ты находишься на кухне, надо собрать рюкзак и идти в универ. можно пройти - коридор",
		paths: map[string]string{
			"коридор": "ничего интересного. можно пройти - кухня, комната, улица",
		},
	}

	corridor := Location{
		name:        "коридор",
		description: "ничего интересного. можно пройти - кухня, комната, улица",
		paths: map[string]string{
			"кухня":   "кухня, ничего интересного. можно пройти - коридор",
			"комната": "ты в своей комнате. можно пройти - коридор",
			"улица":   "на улице весна. можно пройти - домой",
		},
	}

	room := Location{
		name:        "комната",
		description: "ты в своей комнате. на столе: рюкзак, конспекты, ключи. можно пройти - коридор",
		paths: map[string]string{
			"коридор": "ничего интересного. можно пройти - кухня, комната, улица",
		},
		items: []string{"рюкзак", "конспекты", "ключи"},
	}

	street := Location{
		name:        "улица",
		description: "на улице весна. можно пройти - домой",
		paths: map[string]string{
			"домой": "ты пришел домой",
		},
	}

	locations := map[string]Location{
		"кухня":   kitchen,
		"коридор": corridor,
		"комната": room,
		"улица":   street,
	}

	return &Game{
		locations: locations,
		player: Player{
			inventory: []string{},
			location:  "кухня",
		},
	}
}

func (g *Game) Look() string {
	loc := g.locations[g.player.location]
	desc := loc.description + "\n"

	if len(loc.items) > 0 {
		desc += "на столе: " + strings.Join(loc.items, ", ") + "."
	} else {
		desc += "пустая комната."
	}

	desc += " можно пройти - " + strings.Join(pathsToList(loc.paths), ", ")

	return desc
}

func (g *Game) GoTo(location string) string {
	loc := g.locations[g.player.location]
	dest, ok := loc.paths[location]
	if !ok {
		return "нет пути в " + location
	}

	g.player.location = location
	return dest
}

func (g *Game) Take(item string) string {
	loc := g.locations[g.player.location]
	for i, it := range loc.items {
		if it == item {
			g.player.inventory = append(g.player.inventory, item)
			loc.items = append(loc.items[:i], loc.items[i+1:]...)
			return "предмет добавлен в инвентарь: " + item
		}
	}
	return "нет такого"
}

func pathsToList(paths map[string]string) []string {
	list := make([]string, 0, len(paths))
	for path := range paths {
		list = append(list, path)
	}
	return list
}

func printWithDelay(str string) {
	for _, char := range str {
		fmt.Print(string(char))
		time.Sleep(50 * time.Millisecond)
	}
	fmt.Println()
}

func printFastWithDelay(str string) {
	for _, char := range str {
		fmt.Print(string(char))
		time.Sleep(10 * time.Millisecond)
	}
	fmt.Println()
}

func main() {
	game := NewGame()
	reader := bufio.NewReader(os.Stdin)

	printWithDelay("Текстовая игра")
	printWithDelay("==============")
	printWithDelay("Мам, мне ко второй")
	printFastWithDelay("⠀⢀⣤⣤⣤⣤⣤⣴⡶⠶⠶⠶⠶⠶⠶⠶⠶⠤⠤⢤⣤⣤⣤⣤⣤⣄⣀⣀⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀\n⠀⠀⠀⠀⠀⠀⠀⠀⢀⣴⠟⠋⠀⠀⠀⠀⢀⣀⠤⠖⠚⢉⣉⣉⣉⣉⣀⠀⠀⠀⠀⠀⠀⠈⠉⠩⠛⠛⠛⠻⠷⣦⣄⡀⠀⠀⠀⠀⠀⠀\n⠀⠀⠀⠀⠀⠀⠀⣠⡿⠋⠀⠀⠀⣀⠤⠒⣉⠤⢒⣊⡉⠠⠤⠤⢤⡄⠈⠉⠉⠀⠂⠀⠀⠐⠂⠀⠉⠉⠉⠉⠂⠀⠙⠻⣶⣄⠀⠀⠀⠀\n⠀⠀⠀⠀⠀⠀⣰⡿⠁⠀⠀⡠⠊⢀⠔⣫⠔⠊⠁⠀⠀⠀⠀⠀⠀⠙⡄⠀⠀⠀⠀⠀⠘⣩⠋⠀⠀⠀⠉⠳⣄⠀⠀⠀⠈⢻⡇⠀⠀⠀\n⠀⠀⠀⠀⠀⣰⡿⠁⠀⠀⠀⠀⠀⠁⠜⠁⣀⣤⣴⣶⣶⣶⣤⣤⣀⠀⠃⠀⠀⠀⠀⠀⠀⠁⠀⠀⠀⠀⠀⠀⠈⠆⠀⠀⠀⠸⣧⡀⠀⠀\n⠀⠀⠀⣠⣾⣿⣥⠤⢄⡀⠀⢠⣤⠔⢠⣾⣿⣿⣿⣿⣿⣯⣄⡈⠙⢿⣦⠀⠀⠀⠀⡀⢀⣤⣶⣿⣿⣿⣿⣿⣦⠀⣀⣀⣀⣀⡙⢿⣦⡀\n⠀⣠⡾⣻⠋⢀⣠⣴⠶⠾⢶⣤⣄⡚⠉⠉⠉⠁⣠⣼⠏⠉⠙⠛⠷⡾⠛⠀⠀⠀⠘⠛⢿⡟⠛⠋⠉⠉⠉⠁⠀⠀⠀⠀⠀⠦⣝⠦⡙⣿\n⢰⡟⠁⡇⢠⣾⠋⠀⠀⣼⣄⠉⠙⠛⠷⠶⠶⠿⠋⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⣇⠀⠀⠀⠠⣦⣄⣴⡾⢛⡛⠻⠷⠘⡄⢸⣿\n⢸⡇⠀⡇⢸⣇⢀⣤⣴⣿⠻⠷⣦⣄⣀⠀⠀⠀⢀⡀⠀⣀⠰⣤⡶⠶⠆⠀⠀⠀⠀⠀⠈⠛⢿⣦⣄⠀⠈⠉⠉⠁⢸⣇⠀⠀⣠⠃⢸⣿\n⠸⣿⡀⢇⠘⣿⡌⠁⠈⣿⣆⠀⠀⠉⢻⣿⣶⣦⣤⣀⡀⠀⠀⢻⣦⠰⡶⠿⠶⠄⠀⠀⠀⣠⣾⠿⠟⠓⠦⡄⠀⢀⣾⣿⡇⢈⠡⠔⣿⡟\n⠀⠙⢿⣌⡑⠲⠄⠀⠀⠙⢿⣿⣶⣦⣼⣿⣄⠀⠈⠉⠛⠻⣿⣶⣯⣤⣀⣀⡀⠀⠘⠿⠾⠟⠁⠀⠀⢀⣀⣤⣾⣿⢿⣿⣇⠀⠀⣼⡟⠀\n⠀⠀⠀⠹⣿⣇⠀⠀⠀⠀⠈⢻⣦⠈⠙⣿⣿⣷⣶⣤⣄⣠⣿⠁⠀⠈⠉⠙⢻⡟⠛⠻⠿⣿⠿⠛⠛⢻⣿⠁⢈⣿⣨⣿⣿⠀⢰⡿⠀⠀\n⠀⠀⠀⠀⠈⢻⣇⠀⠀⠀⠀⠀⠙⢷⣶⡿⠀⠈⠙⠛⠿⣿⣿⣶⣶⣦⣤⣤⣼⣧⣤⣤⣤⣿⣦⣤⣤⣶⣿⣷⣾⣿⣿⣿⡟⠀⢸⡇⠀⠀\n⠀⠀⠀⠀⠀⠈⢿⣦⠀⠀⠀⠀⠀⠀⠙⢷⣦⡀⠀⠀⢀⣿⠁⠉⠙⠛⠻⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇⠀⢸⣷⠀⠀\n⠀⠀⠀⠀⠀⠀⠀⠙⢷⣄⠀⢀⡀⠀⣀⡀⠈⠻⢷⣦⣾⡃⠀⠀⠀⠀⠀⢸⡇⠀⠀⠀⢹⡟⠉⠉⣿⠏⢡⣿⠃⣾⣷⡿⠁⠀⠘⣿⠀⠀\n⠀⠀⠀⠀⠀⠀⠀⠀⠀⠙⢷⣤⣉⠒⠤⣉⠓⠦⣀⡈⠉⠛⠿⠶⢶⣤⣤⣾⣧⣀⣀⣀⣿⣄⣠⣼⣿⣤⣿⠷⠾⠟⠋⠀⠀⠀⠀⣿⠀⠀\n⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠙⠿⣶⣄⡉⠒⠤⢌⣑⠲⠤⣀⡀⠀⠀⠀⠈⠍⠉⠉⠉⠉⠉⠁⠀⠀⠀⠀⠀⣠⠏⠀⢰⠀⠀⣿⡄⠀\n⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠙⠛⠿⢷⣦⣄⡉⠑⠒⠪⠭⢄⣀⣀⠀⠐⠒⠒⠒⠒⠀⠀⠐⠒⠊⠉⠀⢀⡠⠚⠀⠀⢸⡇⠀\n⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠙⠻⢷⣦⣀⠀⠀⠀⠀⠀⠀⠉⠉⠉⠉⠉⠉⠓⠒⠒⠒⠊⠁⠀⠀⠀⢠⣿⠃⠀\n⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠙⠛⠛⠷⠶⣶⣦⣄⣀⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣠⣴⠟⠁⠀⠀\n⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠉⠙⠛⠛⠷⠶⠶⠶⠶⠶⠾⠛⠛⠉⠀⠀⠀⠀")
	printWithDelay(game.Look())

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		parts := strings.Split(input, " ")

		switch parts[0] {
		case "осмотреться":
			printWithDelay(game.Look())
		case "идти":
			if len(parts) < 2 {
				printWithDelay("укажите локацию")
				continue
			}
			printWithDelay(game.GoTo(parts[1]))
		case "взять":
			if len(parts) < 2 {
				printWithDelay("укажите предмет")
				continue
			}
			printWithDelay(game.Take(parts[1]))
		default:
			printWithDelay("неизвестная команда")
		}
	}
}
