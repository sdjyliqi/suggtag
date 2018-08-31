package suggtag

import "testing"

//

func TestChkTagLen(t *testing.T) {
	sugg := TagFilter{Tag: "09919291929"}
	if bResult := sugg.ChkTagLen(); bResult == false {
		t.Errorf("ChkTagLen failed !", bResult)
	}
	sugg = TagFilter{Tag: "012345678901234567890123456789"}
	if bResult := sugg.ChkTagLen(); bResult == true {
		t.Errorf("ChkTagLen failed !", bResult)
	}
}

func TestChkTagIsEN(t *testing.T) {
	sugg := TagFilter{Tag: "abcdefg"}
	if bResult := sugg.ChkTagIsEN(); bResult == false {
		t.Errorf("ChkTagIsEN failed !", bResult)
	}
	sugg = TagFilter{Tag: "!#$%^&*!@#$%^"}
	if bResult := sugg.ChkTagIsEN(); bResult == true {
		t.Errorf("ChkTagIsEN failed !", bResult)
	}
}

func TestChkTagIsENCN(t *testing.T) {
	sugg := TagFilter{Tag: "abcdefg"}
	if bResult := sugg.ChkTagIsEN(); bResult == false {
		t.Errorf("ChkTagIsEN failed !", bResult)
	}
	sugg = TagFilter{Tag: "!#$%^&*!@#$%^"}
	if bResult := sugg.ChkTagIsEN(); bResult == true {
		t.Errorf("ChkTagIsEN failed !", bResult)
	}
}

func TestChkTagIsZN(t *testing.T) {
	sugg := TagFilter{Tag: "中国"}
	if bResult := sugg.ChkTagIsCN(); bResult == false {
		t.Errorf("ChkTagIsCN failed !", bResult)
	}
	sugg = TagFilter{Tag: "中国er"}
	if bResult := sugg.ChkTagIsCN(); bResult == true {
		t.Errorf("ChkTagIsCN failed !", bResult)
	}
}

func TestChkTagIsDigital(t *testing.T) {
	sugg := TagFilter{Tag: "01234567"}
	if bResult := sugg.ChkTagIsDigital(); bResult == false {
		t.Errorf("ChkTagIsDigital failed !", bResult)
	}
	sugg = TagFilter{Tag: "中国er"}
	if bResult := sugg.ChkTagIsDigital(); bResult == true {
		t.Errorf("ChkTagIsDigital failed !", bResult)
	}
}
