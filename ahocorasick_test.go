package ahocorasick

import (
	"fmt"
	"math/rand"
	"regexp"
	"testing"
	"time"
)

var RegMatcher = regexp.MustCompile(BuildSensitiveStr("./sensitive_words.csv"))
var words = BuildSensitiveArray("./sensitive_words.csv")
var AcMatcher = NewMatcher(words)

var hasSensStr string
var noSensStr string

func init() {
	rand.Seed(time.Now().Unix())
	hasSensStr = fmt.Sprintf("AA%sBB%s%sCC", words[rand.Intn(len(words))], words[rand.Intn(len(words))], words[rand.Intn(len(words))])
	noSensStr = "你真是个伟人哈哈呵呵火啊物abcd"
}

func TestACMatcher_Match(t *testing.T) {
	ret1 := AcMatcher.Match(hasSensStr)
	if len(ret1) == 0 {
		t.Fatal(hasSensStr)
	}
	ret2 := AcMatcher.Match(noSensStr)
	if len(ret2) != 0 {
		t.Fatal(noSensStr)
	}
}

func TestACMatcher_Replace(t *testing.T) {
	fmt.Println("origin:" + hasSensStr)
	rep := AcMatcher.Replace(hasSensStr, "*")
	fmt.Println("AC replace:" + rep)
	fmt.Println("Regexp replace:", RegMatcher.ReplaceAllString(hasSensStr, "*"))
}

func TestACMatcher_Has(t *testing.T) {
	ret1 := AcMatcher.Has(hasSensStr)
	if ret1 == false {
		t.Fatal(hasSensStr)
	}
	ret2 := AcMatcher.Has(noSensStr)
	if ret2 == true {
		t.Fatal(noSensStr)
	}
}

func BenchmarkRegMatcher_Reg_Has(b *testing.B) {
	for idx := 1; idx < 50; idx++ {
		RegMatcher.MatchString(hasSensStr)
	}
}

func BenchmarkRegMatcher_Reg_No(b *testing.B) {
	for idx := 1; idx < 50; idx++ {
		RegMatcher.MatchString(noSensStr)
	}
}

func BenchmarkACMatcher_Ac_Has(b *testing.B) {
	for idx := 1; idx < 50000; idx++ {
		AcMatcher.Has(hasSensStr)
	}
}

func BenchmarkACMatcher_Ac_No(b *testing.B) {
	for idx := 1; idx < 50000; idx++ {
		AcMatcher.Has(noSensStr)
	}
}

func BenchmarkACMatcher_Ac_MatchHas(b *testing.B) {
	for idx := 1; idx < 50000; idx++ {
		AcMatcher.Match(hasSensStr)
	}
}

func BenchmarkACMatcher_Ac_MatchNo(b *testing.B) {
	for idx := 1; idx < 50000; idx++ {
		AcMatcher.Match(noSensStr)
	}
}

func BenchmarkRegMatcher_Replace(b *testing.B) {
	for idx := 1; idx < 50; idx++ {
		RegMatcher.ReplaceAllString(hasSensStr, "*")
	}
}

func BenchmarkACMatcher_Replace(b *testing.B) {
	for idx := 1; idx < 50000; idx++ {
		AcMatcher.Replace(hasSensStr, "*")
	}
}
