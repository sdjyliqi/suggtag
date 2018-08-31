package suggtag

import (
	"regexp"

	_ "github.com/go-sql-driver/mysql"
)

const wMaxTagLen = 15

type SearchTag struct {
	Tag    string
	WCnt   int64
	Level  int64
	RoleId int64
	DirtId int64
}

type TagFilter struct {
	Tag string
}

//判断搜素关键字是否满足长度限制
func (p *TagFilter) ChkTagLen() bool {
	if len(p.Tag) == 0 || len(p.Tag) > wMaxTagLen {
		return false
	}
	return true
}

//判断搜索关键字是否是纯英文
func (p *TagFilter) ChkTagIsEN() bool {
	var hzRegexp = regexp.MustCompile("^[a-zA-Z]+$")
	if bInvalid := hzRegexp.MatchString(p.Tag); bInvalid == false {
		return false
	}
	return true
}

////判断搜索关键字是否是纯数字
func (p *TagFilter) ChkTagIsDigital() bool {
	var hzRegexp = regexp.MustCompile("^[0-9]+$")
	if bInvalid := hzRegexp.MatchString(p.Tag); bInvalid == false {
		return false
	}
	return true
}

////判断搜索关键字是否是纯中文
func (p *TagFilter) ChkTagIsCN() bool {
	var hzRegexp = regexp.MustCompile("^[\u4e00-\u9fa5]+$")
	if bInvalid := hzRegexp.MatchString(p.Tag); bInvalid == false {
		return false
	}
	return true
}

////判断搜索关键字是否是英文+ 数字组成
func (p *TagFilter) ChkTagIsENDigital() bool {
	var hzRegexp = regexp.MustCompile("^[a-zA-Z0-9]+$")
	if bInvalid := hzRegexp.MatchString(p.Tag); bInvalid == false {
		return false
	}
	return true
}

////判断搜索关键字是否是中文+ 英文
func (p *TagFilter) ChkTagIsENCN() bool {
	var hzRegexp = regexp.MustCompile("^[a-zA-z\u4e00-\u9fa5]+$")
	if bInvalid := hzRegexp.MatchString(p.Tag); bInvalid == false {
		return false
	}
	return true
}

////判断搜索关键字是否是中文+ 英文
func (p *TagFilter) ChkTagIsCNDigital() bool {
	var hzRegexp = regexp.MustCompile("^[0-9\u4e00-\u9fa5]+$")
	if bInvalid := hzRegexp.MatchString(p.Tag); bInvalid == false {
		return false
	}
	return true
}

////判断搜索关键字是否是中文+ 英文 + 数字
func (p *TagFilter) ChkTagIsENCNDigital() bool {
	var hzRegexp = regexp.MustCompile("^[a-zA-Z0-9\u4e00-\u9fa5]+$")
	if bInvalid := hzRegexp.MatchString(p.Tag); bInvalid == false {
		return false
	}
	return true
}
