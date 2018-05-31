package command

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/codegangsta/cli"
	"github.com/valbeat/whatday"
)

// CmdList prints article
func CmdList(c *cli.Context) {
	t := time.Now()
	articles, err := whatday.NewArticles(t)
	if err != nil {
		fmt.Println(err)
	}

	rand.Seed(time.Now().UnixNano())
	for _, article := range articles {
		fmt.Printf("## %s\n> %s\n", article.Title, article.Text)
	}
}
