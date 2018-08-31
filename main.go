package main

import (
	"fmt"
	"searchsugg/suggtag"
)

func main(){
	sugg := suggtag.SuggTag{"./data/doutu.txt"}
	filePath,err := sugg.GetFilePath()
	fmt.Println(filePath,err )
	sugg.PrintSuggTag()
}