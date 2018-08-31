package main

import (
	"fmt"
	"searchsugg/suggtag"

	"github.com/texttheater/golang-levenshtein/levenshtein"
)

func main() {
	source := "我的亲爱"
	target := "我的亲爱的"
	distance := levenshtein.DistanceForStrings([]rune(source), []rune(target), levenshtein.DefaultOptions)
	fmt.Printf(`Distance between "%s" and "%s" computed as %d`, source, target, distance)
	fmt.Println("-----------------")
	sugg := suggtag.SuggTag{"./data/doutu.txt"}
	filePath, err := sugg.GetFilePath()
	fmt.Println(filePath, err)
	// sugg.PrintSuggTag()
	suggmod := suggtag.TagFilter{Tag: "Iloveyou"}
	fmt.Println(suggmod.ChkTagIsEN())
	suggmod = suggtag.TagFilter{Tag: "中国"}
	fmt.Println(suggmod.ChkTagIsEN())
	fmt.Println(suggmod.ChkTagIsCN())
	suggmod = suggtag.TagFilter{Tag: "C中国"}
	fmt.Println(suggmod.ChkTagIsEN())
	fmt.Println(suggmod.ChkTagIsCN())
	fmt.Println(suggmod.ChkTagIsENCN())
}
