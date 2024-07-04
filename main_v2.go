package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Capitalize(s string) string {
	var word string
	var validStr string
	for _, char := range s {
		if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' || char >= '0' && char <= '9' || string(char) == "'" {
			if len(word) == 0 && char >= 'a' && char <= 'z' {
				word += string(char - 32)
			} else if len(word) > 0 && char >= 'A' && char <= 'Z' {
				word += string(char + 32)
			} else {
				word += string(char)
			}
		} else {
			validStr += word + string(char)
			word = ""
		}
	}
	validStr += word
	return validStr
}

func ReplaceAToAn(text string) string {
	re := regexp.MustCompile(`(?i)(a)(\s+[aeiouh])`)

	matchesSlc := re.FindAllStringSubmatch(text, -1)

	newText := text

	for _, matches := range matchesSlc {
		newText = strings.Replace(newText, matches[0], matches[1]+"n"+matches[2], -1)
	}

	return newText
}

func SingleQuotation(text string) string {
	pattern := regexp.MustCompile(`'\s*([^']+?)\s*'`) // [^'] negate symbol
	return pattern.ReplaceAllString(text, "'$1'")
}

func Punctuations(s string) string {
	str := [6]string{".", ",", "!", "?", ":", ";"}
	arr := strings.Split(string(s), " ")
	for i := 0; i < len(arr); i++ {
		for _, ch := range str {
			s = strings.ReplaceAll(s, " "+ch, ch+" ")
			s = strings.ReplaceAll(s, ch+"  ", ch+" ")
		}
	}
	s = strings.Trim(s, " ")
	return s
}

func SpaceEmpty(str string) string {
	run := []rune(str)
	sr := ""
	for i := 0; i < len(run); i++ {
		if run[0] == ' ' {
			continue
		}
		if run[i] == ' ' && run[i-1] == ' ' {
			continue
		}
		sr += string(run[i])
	}
	return sr
}

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("ERROR")
		return
	}
	simple, _ := ioutil.ReadFile(args[0])
	arr := strings.Split(string(simple), " ")
	var tmp []string
	for _, v := range arr { // skip the spaces bfr operations!
		if v != "" {
			tmp = append(tmp, v)
		}
	}
	arr = tmp
	for i := 0; i < len(arr); i++ {

		if i > 0 && arr[i] == "(cap)" {
			arr[i-1] = Capitalize(arr[i-1])
			arr[i] = ""
		} else if arr[0] == "(cap)" {
			arr[i] = ""
		}

		if i > 0 && arr[i] == "(up)" {
			arr[i-1] = strings.ToUpper(arr[i-1])
			arr[i] = ""

		} else if arr[0] == "(up)" {
			arr[i] = ""
		}

		if i > 0 && arr[i] == "(low)" {
			arr[i-1] = strings.ToLower(arr[i-1])
			arr[i] = ""

		} else if arr[0] == "(low)" {
			arr[i] = ""
		}

		if i > 0 && arr[i] == "(bin)" {

			data, _ := strconv.ParseInt(arr[i-1], 2, 0)
			arr[i-1] = strconv.FormatInt(data, 10)
			arr[i] = ""

		} else if arr[0] == "(bin)" {
			arr[i] = ""
		}

		if i > 0 && arr[i] == "(hex)" {

			data, _ := strconv.ParseInt(arr[i-1], 16, 0)
			arr[i-1] = strconv.FormatInt(data, 10)
			arr[i] = ""

		} else if arr[0] == "(hex)" {
			arr[i] = ""
		}

		if i > 0 && strings.Contains(arr[i], "(low,") && i != len(arr)-1 {
			if arr[i+1][len(arr[i+1])-1] != ')' {
				break
			}
			nbr, _ := strconv.Atoi(strings.TrimRight(arr[i+1], ")"))
			for j := nbr; j > 0; j-- {
				if i-j >= 0 {
					arr[i-j] = strings.ToLower(arr[i-j])
				}
			}
			arr[i] = ""
			arr[i+1] = ""
		} else if arr[0] == "(low," && i != len(arr)-1 {
			arr[i] = ""
			arr[i+1] = ""
		}

		if i > 0 && strings.Contains(arr[i], "(cap,") && i != len(arr)-1 {
			if arr[i+1][len(arr[i+1])-1] != ')' {
				break
			}
			nbr, _ := strconv.Atoi(strings.TrimRight(arr[i+1], ")"))
			for j := nbr; j > 0; j-- {
				if i-j >= 0 {
					arr[i-j] = Capitalize(arr[i-j])
				}
			}
			arr[i] = ""
			arr[i+1] = ""
		} else if arr[0] == "(cap," && i != len(arr)-1 {
			arr[i] = ""
			arr[i+1] = ""
		}

		if i > 0 && strings.Contains(arr[i], "(up,") && i != len(arr)-1 {
			if arr[i+1][len(arr[i+1])-1] != ')' {
				break
			}
			nbr, _ := strconv.Atoi(strings.TrimRight(arr[i+1], ")"))

			for j := nbr; j > 0; j-- {
				if i-j >= 0 {
					arr[i-j] = strings.ToUpper(arr[i-j])
				}
			}
			arr[i] = ""
			arr[i+1] = ""
		} else if arr[0] == "(up," && i != len(arr)-1 {
			arr[i] = ""
			arr[i+1] = ""
		}
		str := strings.Join(arr, " ")
		str = Punctuations(str)
		str = SingleQuotation(str)
		str = ReplaceAToAn(str) // call the vowels function
		str = SpaceEmpty(str)
		err := ioutil.WriteFile("result.txt", []byte(str), 0644)
		if err != nil {
			fmt.Println("ERROR")
		}
	}
}
