package MyDbms

import (
	"fmt"
	goCache "github.com/patrickmn/go-cache"
	"github.com/yalp/jsonpath"
	"os"
)

func main() {
	cache := goCache.New(goCache.NoExpiration, -1)
	if _, err := os.Stat(os.Getenv("ElhamDatabasePath")); os.IsExist(err) {
		err := cache.LoadFile(os.Getenv("ElhamDatabasePath"))
		Error{
			isClientSide: false,
			Error: fmt.Sprintf("Error:\nCould Not Load Data From $ElhamDatabasePath variable\n%s",
				err.Error()),
		}.throw()
		return
	} else {
		fmt.Println("No ")
	}
	file, err := os.OpenFile(os.Getenv("ElhamDatabasePath"), os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeDir)
	if err != nil {
		Error{
			isClientSide: false,
			Error:        fmt.Sprintf("Error:\nCould Not Open $ElhamDataBasePath Variable\n%s", err.Error()),
		}.throw()
	}

	err = cache.Save(file)
}

type (
	Boolean bool
	String  string
	Integer int64
	Float   float64
	Map     map[string]interface{}
	Array   []interface{}
	Set     struct {
		Array `json:"array"`
	}
	Error struct {
		isClientSide bool
		Error        string
	}
)

func (e Error) throw() {
	fmt.Printf("Error:\nisClientSide:%v\n%s", e.isClientSide, e.Error)
}

func (s Set) makeUnique() {
	for idx, value := range s.Array {
		if firstIndex(s.Array, value) != lastIndex(s.Array, value) {
			removeIndex(s.Array, len(s.Array)-idx)
		}
	}
}

func firstIndex(arr []interface{}, toSearch interface{}) int {
	for i, value := range arr {
		if value == toSearch {
			return i
		}
	}
	return -1
}
func lastIndex(arr []interface{}, toSearch interface{}) int {
	length := len(arr)
	for i := range arr {
		if arr[i-length] == toSearch {
			return i
		}
	}
	return -1
}
func removeIndex(s []interface{}, index int) []interface{} {
	ret := make([]interface{}, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func getByJsonPath(obj Map, path String) (value interface{}, err error) {
	value, err = jsonpath.Read(obj, string(path))
	return
}
