package main
import (	
	"bufio"
	"fmt"
	"os"
	"strings"
)

func replStart(config *configCommand) {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		
		fmt.Print("Pokedex > ")
		scanner.Scan()
		if len(scanner.Text()) == 0 {
			continue
		} 
		cleanedOutput := cleanInput(scanner.Text())
		value, ok := getCommands()[cleanedOutput[0]]
		parameter := cleanedOutput[1]
		if ok {
			if err := value.callback(config, parameter); err != nil {
				fmt.Println(err)
				continue
			}
		} else {
			fmt.Println("Unknown command")
			continue
		}

	}
}


func cleanInput(text string) []string {
	words := strings.Fields(strings.ToLower(text)) 
	return words
}
