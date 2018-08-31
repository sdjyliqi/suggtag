package suggtag

import (
	"bufio"
	"database/sql"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type SuggTag struct {
	FilePath string
}

func (p *SuggTag) GetFilePath() (string, error) {
	return p.FilePath, nil
}

func (p *SuggTag) GetBlcakTagMapFromDB() (map[string]bool, error) {
	dicMap := make(map[string]bool)
	db, err := sql.Open("mysql", "tugele:PwTugeLe2017@tcp(test.tugele.rds.sogou:3306)/tugele?charset=utf8")
	if err != nil {
		return nil, err
	}
	defer db.Close()
	rows, err := db.Query("select name from pic_filter_keywords where type=1 and isDeleted=0")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		dicMap[name] = true
	}
	return dicMap, nil
}

func (p *SuggTag) ChkTagValid(strTagLine string, blackDic map[string]bool) (bool, string, int) {
	tagList := strings.Split(strTagLine, "	")
	strTag := tagList[0]
	if strTag == "" {
		return false, "", 1
	}
	if len(strTag) >= 16 {
	   
		return false, "", 2
	}
	if len(tagList) <= 1 {
		return false, "", 3
	}
	wCnt, err := strconv.Atoi(tagList[1])
	if err != nil {
		return false, "", 4
	}
	if wCnt <= 50 {
		return false, "", 5
	}
	var hzRegexp = regexp.MustCompile("^[\u4e00-\u9fa5]+$")
	if bIsZW := hzRegexp.MatchString(strTag); bIsZW == false {
		return false, "", 6
	}
	//判断string是否是叠词，如果是叠词，暂时按照双字词处理
	rs := []rune(strTag)
	hzRegexp = regexp.MustCompile("^[" + string(rs[0]) + "]+$")
	bIsDZ := hzRegexp.MatchString(strTag)
	if bIsDZ == false {
		//此时判断该次是否在黑名单中
		_, bExist := blackDic[strTag]
		if bExist == false {
			return true, strTag, wCnt
		}
		return false, "", 7
	} else {
		_, bExist := blackDic[strTag]
		if bExist == false {
			return true, strings.Repeat(string(rs[0]), 2), wCnt
		}
		return false, "", 8
	}
	return true, strTag, wCnt
}

func (p *SuggTag) GetSuggTag() (map[string]int, error) {
	dicMap := make(map[string]int)
	blackDic, err := p.GetBlcakTagMapFromDB()
	if err != nil {
		fmt.Println("Error: Get the black search-words failed from mysql.")
		return nil, err
	}
	fi, err := os.Open(p.FilePath)
	if err != nil {
		fmt.Printf("Error: Read the search-words failed.", err)
		return nil, err
	}
	defer fi.Close()
	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
	//	fmt.Println(string(a))
		bValid, tag, wCnt := p.ChkTagValid(string(a), blackDic)
		if bValid == true {
			dicMap[tag] = wCnt
		}
	}
	return dicMap, nil
}

func (p *SuggTag) PrintSuggTag() {
	tagDic, err := p.GetSuggTag()
	if err != nil {
		fmt.Println("Error: some errors ")
	}
	for k, v := range tagDic {
		fmt.Println(k, "	", v)
	}
}
