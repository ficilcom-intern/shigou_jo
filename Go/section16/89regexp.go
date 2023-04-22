package main

import (
	"fmt"
	"regexp"
)

func main() {
	match, _ := regexp.MatchString("[0-9]", "2")
	fmt.Println(match)

	re1, _ := regexp.Compile("A")
	match = re1.MatchString("ABC")
	fmt.Println(match)

	re2 := regexp.MustCompile("A")
	match = re2.MatchString("ABC")
	fmt.Println(match)

	regexp.MustCompile("\\d")
	regexp.MustCompile(`\d`)

	//flags
	//i 大小文字区別しない
	//m マルチラインモード
	//s .が\nにマッチ
	//U 最小マッチへの変換(x* -> x*?, x+ -> x+?)
	re3 := regexp.MustCompile(`(?i)abc`)
	match = re3.MatchString("ABC")
	fmt.Println(match)

	re4 := regexp.MustCompile(`^ABC$`)
	match = re4.MatchString("ABC")
	fmt.Println(match)

	match = re4.MatchString("  ABC  ")
	fmt.Println(match)

	re5 := regexp.MustCompile("a+b*")
	fmt.Println(re5.MatchString("ab"))
	fmt.Println(re5.MatchString("a"))
	fmt.Println(re5.MatchString("aaaaaabbbb"))
	fmt.Println(re5.MatchString("b"))

	re6 := regexp.MustCompile(`[XYZ]`)
	fmt.Println(re6.MatchString("Y"))

	re7 := regexp.MustCompile(`^[0-9A-Za-z_]{3}$`)
	fmt.Println(re7.MatchString("ABC"))
	fmt.Println(re7.MatchString("abcdefg"))

	re8 := regexp.MustCompile(`[^0-9A-Za-z_]`)
	fmt.Println(re8.MatchString("ABC"))
	fmt.Println(re8.MatchString("あ"))

	re9 := regexp.MustCompile(`(abc|ABC)(xyz|XYZ)`)
	fmt.Println(re9.MatchString("abcxyz"))
	fmt.Println(re9.MatchString("ABCXYZ"))
	fmt.Println(re9.MatchString("ABCxyz"))
	fmt.Println(re9.MatchString("ABCabc"))

	re := regexp.MustCompile(`(abc|ABC)(xyz|XYZ)`)

	fs1 := re.FindString("AAAAABCXYZZZZ")
	fmt.Println(fs1)
	fs2 := re.FindAllString("ABCXYZABCXYZ", -1)
	fmt.Println(fs2)

	re10 := regexp.MustCompile(`(\d+)-(\d+)-(\d+)`)
	s := `
		0123-456-7889
		111-222-333
		556-787-899
		`
	ms := re10.FindAllStringSubmatch(s, -1)
	fmt.Println(ms)

	for _, v := range ms {
		fmt.Println(v)
	}

	fmt.Println(ms[0][0])
	fmt.Println(ms[0][1])
	fmt.Println(ms[0][2])
	fmt.Println(ms[0][3])

	re11 := regexp.MustCompile(`\s+`)
	fmt.Println(re11.ReplaceAllString("AAA BBB CCC", ","))
	re12 := regexp.MustCompile(`、|。`)
	fmt.Println(re12.ReplaceAllString("私は、Golangを使用する、プログラマ。", ""))

	re13 := regexp.MustCompile(`(abc|ABC)(xyz|XYZ)`)
	fmt.Println(re13.Split("ASHDKFJABCXYZKFDJHGSABCXYZ", -1))

	re14 := regexp.MustCompile(`\s+`)
	fmt.Println(re14.Split("aaaaaaaa    bbbbbbb cccc", -1))
}
