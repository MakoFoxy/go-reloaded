package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// функция записи
func WriteF(filename2 string, data string) {
	databyte := []byte(data)
	err := ioutil.WriteFile(filename2, databyte, 0o644)
	if err != nil {
		fmt.Println(err)
	}
}

// функция чтения
func ReadF(filename string) string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	str := string(data)
	return str
}

func main() {
	if len(os.Args) != 3 {
		return
	}
	input := ReadF(os.Args[1])
	fieldsSplitter := Splitter(input)
	// fmt.Println("----fieldsData", fieldsSplitter)
	// printer(fieldsSplitter)

	for i := 0; i < len(fieldsSplitter); i++ {
		txtWords := fieldsSplitter[i]

		if txtWords == "(hex)" {
			if fieldsSplitter[i-1] == string(fieldsSplitter[i-1]) {
				tmpHex, err := strconv.ParseInt(fieldsSplitter[i-1], 16, 64)
				if err != nil {
					fmt.Println("ERROR")
				}

				fieldsSplitter[i-1] = strconv.Itoa(int(tmpHex))
			}
			fieldsSplitter = append(fieldsSplitter[:i], fieldsSplitter[i+1:]...)
			i--
		}
		if txtWords == "(bin)" {
			if fieldsSplitter[i-1] == string(fieldsSplitter[i-1]) {
				tmpBin, err := strconv.ParseInt(fieldsSplitter[i-1], 2, 64)
				if err != nil {
					fmt.Println("ERROR")
				}
				fieldsSplitter[i-1] = strconv.Itoa(int(tmpBin))
			}
			fieldsSplitter = append(fieldsSplitter[:i], fieldsSplitter[i+1:]...)
			i--
		}
		if txtWords == "(up)" {
			fieldsSplitter[i-1] = strings.ToUpper(fieldsSplitter[i-1])
			fieldsSplitter = append(fieldsSplitter[:i], fieldsSplitter[i+1:]...)
			i--
		}
		if txtWords == "(low)" {
			fieldsSplitter[i-1] = strings.ToLower(fieldsSplitter[i-1])
			fieldsSplitter = append(fieldsSplitter[:i], fieldsSplitter[i+1:]...)
			i--
		}
		if txtWords == "(cap)" {
			cap := fieldsSplitter[i-1]
			if cap[len(cap)-2:] == "'t" {
				// defer fmt.Print("s")
				fieldsSplitter[i-1] = strings.ToLower(fieldsSplitter[i-1])
				fieldsSplitter = append(fieldsSplitter[:i], fieldsSplitter[i+1:]...)
				if fieldsSplitter[i-1] == strings.ToLower(fieldsSplitter[i-1]) {
					fieldsSplitter[i-1] = strings.Title(fieldsSplitter[i-1][0:1]) + fieldsSplitter[i-1][1:]
				}
			} else {
				fieldsSplitter[i-1] = strings.Title(strings.ToLower(fieldsSplitter[i-1]))
				fieldsSplitter = append(fieldsSplitter[:i], fieldsSplitter[i+1:]...)
				i--
			}
		}

		if i > 0 && i <= len(fieldsSplitter) && fieldsSplitter[i-1] == "(low," && string(txtWords[len(txtWords)-1]) == ")" {

			nmbrstr := ""
			for _, k := range txtWords {
				if k >= '0' && k <= '9' || k == '-' {
					nmbrstr = nmbrstr + string(k)
				}
			}
			num, err := strconv.Atoi(nmbrstr)
			if err != nil || num <= 0 {
				fmt.Println("ERROR")
			}
			if num >= i {
				for m := 0; m < i-1; m++ {
					fieldsSplitter[m] = strings.ToLower(fieldsSplitter[m]) // функция будет работать по строке до  m < i-1
				}
			} else {
				for m := i - 2; m >= i-1-num; m-- {
					fieldsSplitter[m] = strings.ToLower(fieldsSplitter[m]) //   функция будет работать по заданным значениям
				}
			}
			fieldsSplitter = append(fieldsSplitter[:i-1], fieldsSplitter[i+1:]...)
			i--

		}
		if i > 0 && i <= len(fieldsSplitter) && fieldsSplitter[i-1] == "(up," && string(txtWords[len(txtWords)-1]) == ")" {
			nmbrstr := ""
			for _, k := range txtWords {
				if k >= '0' && k <= '9' || k == '-' {
					nmbrstr = nmbrstr + string(k)
				}
			}
			num, err := strconv.Atoi(nmbrstr)
			if err != nil || num <= 0 {
				fmt.Println("ERROR")
			}
			if num >= i {
				for m := 0; m < i-1; m++ {
					fieldsSplitter[m] = strings.ToUpper(fieldsSplitter[m])
				}
			} else {
				for m := i - 2; m >= i-1-num; m-- {
					fieldsSplitter[m] = strings.ToUpper(fieldsSplitter[m])
				}
			}
			fieldsSplitter = append(fieldsSplitter[:i-1], fieldsSplitter[i+1:]...)
			i--
		}
		if i > 0 && i <= len(fieldsSplitter) && fieldsSplitter[i-1] == "(cap," && string(txtWords[len(txtWords)-1]) == ")" {
			nmbrstr := ""
			for _, k := range txtWords {
				if k >= '0' && k <= '9' || k == '-' {
					nmbrstr = nmbrstr + string(k)
				}
			}
			num, err := strconv.Atoi(nmbrstr)
			if err != nil || num <= 0 {
				fmt.Println("ERROR")
			}
			if num >= i {
				for m := 0; m < i-1; m++ {
					fieldsSplitter[m] = strings.Title(strings.ToLower(fieldsSplitter[m]))
				}
			} else {
				for m := i - 2; m >= i-1-num; m-- {
					fieldsSplitter[m] = strings.Title(strings.ToLower(fieldsSplitter[m]))
				}
			}
			fieldsSplitter = append(fieldsSplitter[:i-1], fieldsSplitter[i+1:]...)
			i--
		}
		for i, artcl := range fieldsSplitter {
			if i < len(fieldsSplitter)-1 && (artcl == "a" || artcl == "A") {
				if fieldsSplitter[i+1][0] == 'a' || fieldsSplitter[i+1][0] == 'e' || fieldsSplitter[i+1][0] == 'y' || fieldsSplitter[i+1][0] == 'u' || fieldsSplitter[i+1][0] == 'i' || fieldsSplitter[i+1][0] == 'o' || fieldsSplitter[i+1][0] == 'h' {
					fieldsSplitter[i] += "n"
				}
			}
		}
		// fmt.Println(i, txtWords)
	}
	// fmt.Println(fieldsSplitter)
	fieldsSplitter = Punc(fieldsSplitter)
	result := strings.Join(fieldsSplitter, " ")
	WriteF(os.Args[2], result)
}

func Splitter(s string) []string {
	var arr []string
	// arr = strings.Fields(string(s))
	arr = strings.Split(s, " ")
	var newarr []string
	for _, k := range arr {
		if len(k) != 0 {
			newarr = append(newarr, k)
		}
	}
	return newarr
}

func printer(fieldsSplitter []string) {
	for _, arr := range fieldsSplitter {
		fmt.Println(arr)
	}
}

// функция пунктуации
func Punc(fieldsSplitter []string) []string {
	punctuationMarks := []string{",", ".", "!", "?", ":", ";"}
	// проверяет пунктуацию в начале строки
	for i := 1; i < len(fieldsSplitter); i++ {
		word := fieldsSplitter[i]
		for _, punctuation := range punctuationMarks {
			if i != 0 && (string(word[0]) == punctuation) && (string(word[len(word)-1]) != punctuation) {
				fieldsSplitter[i-1] = fieldsSplitter[i-1] + punctuation
				// Начинает новое слово с индекса 1, чтобы оставить пунктуацию с позиции 0
				fieldsSplitter[i] = word[1:]
				i--
			}
		}
		// Проверяет знаки препинания внутри фрагмента и при необходимости выравниваем их рядом со словом
		for i := 1; i < len(fieldsSplitter); i++ {
			word := fieldsSplitter[i]
			for _, punctuation := range punctuationMarks {
				if (string(word[0]) == punctuation) && (string(word[len(word)-1]) == punctuation) && (fieldsSplitter[i] != fieldsSplitter[len(fieldsSplitter)-1]) {
					fieldsSplitter[i-1] = fieldsSplitter[i-1] + word
					// удаляет пустой индекс
					fieldsSplitter = append(fieldsSplitter[:i], fieldsSplitter[i+1:]...)
					i--
				}
			}
		}
		for i, word := range fieldsSplitter {
			for _, punctuation := range punctuationMarks {
				if (string(word[0]) == punctuation) && (fieldsSplitter[len(fieldsSplitter)-1] == fieldsSplitter[i]) {
					fieldsSplitter[i-1] = fieldsSplitter[i-1] + word
					fieldsSplitter = fieldsSplitter[:len(fieldsSplitter)-1]
				}
			}
		}
		countone := 0
		for i := 0; i < len(fieldsSplitter); i++ {
			word := fieldsSplitter[i]
			if string(word[0]) == string(rune(39)) {
				if countone%2 != 0 {
					if i != 0 {
						fieldsSplitter[i-1] = fieldsSplitter[i-1] + string(rune(39))
					}
					if len(word) > 1 {
						fieldsSplitter[i] = fieldsSplitter[i][1:]
					} else {
						fieldsSplitter = append(fieldsSplitter[:i], fieldsSplitter[i+1:]...)
					}
				} else if countone%2 == 0 && len(word) == 1 {
					if i != len(fieldsSplitter)-1 {
						fieldsSplitter[i+1] = string(rune(39)) + fieldsSplitter[i+1]
						fieldsSplitter = append(fieldsSplitter[:i], fieldsSplitter[i+1:]...)
						i++
					}
				}
				countone++
			}
			if len(word) > 1 && string(word[len(word)-1]) == string(rune(39)) {
				if countone%2 == 0 {
					fieldsSplitter[i+1] = string(rune(39)) + fieldsSplitter[i+1]
					fieldsSplitter[i] = fieldsSplitter[i][:len(fieldsSplitter[i])-1]
					i++
				}
				countone++
			}
		}
	}
	counttwo := 0

	for i := 0; i < len(fieldsSplitter); i++ {
		word := fieldsSplitter[i]
		if string(word[0]) == string(rune(34)) {
			if counttwo%2 != 0 {
				if i != 0 {
					fieldsSplitter[i-1] = fieldsSplitter[i-1] + string(rune(34))
				}
				if len(word) > 1 {
					fieldsSplitter[i] = fieldsSplitter[i][1:]
				} else {
					fieldsSplitter = append(fieldsSplitter[:i], fieldsSplitter[i+1:]...)
				}
			} else if counttwo%2 == 0 && len(word) == 1 {
				if i != len(fieldsSplitter)-1 {
					fieldsSplitter[i+1] = string(rune(34)) + fieldsSplitter[i+1]
					fieldsSplitter = append(fieldsSplitter[:i], fieldsSplitter[i+1:]...)
					i++
				}
			}
			counttwo++
		}
		if len(word) > 1 && string(word[len(word)-1]) == string(rune(34)) {
			if counttwo%2 == 0 {
				fieldsSplitter[i+1] = string(rune(34)) + fieldsSplitter[i+1]
				fieldsSplitter[i] = fieldsSplitter[i][:len(fieldsSplitter[i])-1]
				i++
			}
			counttwo++
		}
	}
	return fieldsSplitter
}
