package MyDbms

import (
	goCache "github.com/patrickmn/go-cache"
	"os"
)

func main() {
	args := os.Args[1:]

	c2 := goCache.New(goCache.NoExpiration, -1)

	c2.Save(nil)
}
