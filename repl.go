package main
import (	
	"bufio"
	"fmt"
	"os"
	"strings"
)

func replStart(config *configCommand) {

	scanner := bufio.NewScanner(os.Stdin)
	parameters := []string{} 
	for {
		
		fmt.Print("Pokedex > ")
		scanner.Scan()
		if len(scanner.Text()) == 0 {
			continue
		} 
		cleanedOutput := cleanInput(scanner.Text())
		value, ok := getCommands()[cleanedOutput[0]]
		if len(cleanedOutput) > 1 {
			parameters = cleanedOutput[1:]
		}

		if ok {
			if err := value.callback(config, parameters...); err != nil {
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
