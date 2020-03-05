package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	c "github.com/TreyBastian/colourize"
)

func main() {
	usage := "Usage: ./hertz /path/to/file"
	words := []string{" it", " thing", "everything", "something", "anything", "due to", "everybody", "everyone", "situation", "they", "them", "this", " he", "she ", "furthermore", "in addition", "moreover", "goes on to say", "goes into", "also", "perceive", "perception", " one", "obviously", "talk", "article", "brings up", " do", "does", "about", " as ", "throughout", "research", "notes", "clearly", "states", "aspect", "factor", "statistics", "beneficial", "very", "extremely", "based on", "utilize", "there were", "there are", "lack of", "three", "four", "finish", " end", "ended up", "ends up", "finally", "last", "begin", "start"}
	path := os.Args[1]
	if args := len(os.Args); args > 2 {
		fmt.Println(usage)
		os.Exit(0)
	}
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	contents := string(file)

	for word := 0; word <= len(words) - 1; word++ {
		results := strings.Count(contents, words[word])
		if results > 3 {
			format := words[word] + " -- "
			fmt.Print(c.Colourize(format, c.Red, c.Bold))
			fmt.Println(results)
		} else if results >= 1 && results < 3 {
			format := words[word] + " -- "
			fmt.Print(c.Colourize(format, c.Yellow, c.Bold))
			fmt.Println(results)
		} else if results == 0 {
			format := words[word] + " -- "
			fmt.Print(c.Colourize(format, c.Green, c.Bold))
			fmt.Println(results)
		}
	}


}
