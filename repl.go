package main
import (	
	"bufio"
	"fmt"
	"os"
	"strings"
)

func replStart(config *configCommand) {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Pokedex > ")

	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			continue
		} 
		cleanedOutput := cleanInput(scanner.Text())
		value, ok := getCommands()[cleanedOutput[0]]
		if ok {
			if err := value.callback(&config); err != nil {
				fmt.Errorf("error issuing command %v", err)
				continue
			}
		} else {
			fmt.Println("Unknown command")
		}

		fmt.Print("Pokedex > ")
	}
}


func cleanInput(text string) []string {
	words := strings.Fields(strings.ToLower(text)) 
	return words
}
