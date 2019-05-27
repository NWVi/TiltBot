package namegenerator

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

// Webscraper gets a first name from fakenamegenerator.com
func Webscraper(max int) (string, error) {
	link := "https://www.fakenamegenerator.com/gen-random-rucyr-us.php"
	c := colly.NewCollector()

	var newName string

	c.OnHTML("#details .address h3", func(e *colly.HTMLElement) {
		newName = strings.Fields(e.Text)[0]
		if len(newName) >= max {
			fmt.Printf("%s if pretty long... trying to get a new name...\n", newName)
			c.Visit(e.Request.AbsoluteURL(link))
		}
	})

	c.Visit(link)

	if newName == "" {
		return newName, errors.New("Could not get a name")
	}

	return newName, nil
}
