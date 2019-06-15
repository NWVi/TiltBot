package namegenerator

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"time"
)

var names []string

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// Picks and returns a random name
func PickName() (string, error) {
	if len(names) < 1 {
		createSlice()
	}
	i := rand.Intn(len(names))
	name := names[i]
	// names = append(names[:i], names[i+1:]...)
	// writeFile(names)
	// if name == "" {
	// return name, errors.New("Could not get a name")
	// }

	return name, nil
}

func createSlice() {
	file, err := os.Open("list.txt")
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		names = append(names, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
