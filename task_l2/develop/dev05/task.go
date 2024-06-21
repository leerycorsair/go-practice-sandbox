package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки
-r - "regexp", использовать регулярные выражения для поиска

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type Config struct {
	After      int
	Before     int
	Context    int
	Count      bool
	IgnoreCase bool
	Invert     bool
	Fixed      bool
	LineNum    bool
	Regexp     bool
}

func parseConfig() *Config {
	AFlag := flag.Int("A", 0, "after content")
	BFlag := flag.Int("B", 0, "before content")
	CFlag := flag.Int("C", 0, "around content")
	cFlag := flag.Bool("c", false, "count")
	iFlag := flag.Bool("i", false, "ignore case")
	vFlag := flag.Bool("v", false, "invert")
	FFlag := flag.Bool("F", false, "fixed")
	nFlag := flag.Bool("n", false, "line num")
	rFlag := flag.Bool("r", false, "regexp")
	flag.Parse()

	return &Config{
		After:      *AFlag,
		Before:     *BFlag,
		Context:    *CFlag,
		Count:      *cFlag,
		IgnoreCase: *iFlag,
		Invert:     *vFlag,
		Fixed:      *FFlag,
		LineNum:    *nFlag,
		Regexp:     *rFlag,
	}
}

func readLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func addBefore(lines []string, add int, i int, matchingLines *map[int]string, addedLines *[]int) {
	if i-add < 0 {
		for y := 0; y < i; y++ {
			added := false
			for _, num := range *addedLines {
				if num == y {
					added = true
					break
				}
			}
			if !added {
				(*matchingLines)[y] = lines[y]
				*addedLines = append(*addedLines, y)
			}
		}
	} else {
		sequence := getSequence(i, i-add, false)
		for _, j := range sequence {
			added := false
			for _, num := range *addedLines {
				if num == j {
					added = true
					break
				}
			}
			if !added {
				(*matchingLines)[j] = lines[j]
				*addedLines = append(*addedLines, j)
			}
		}
	}
}

func addAfter(lines []string, add int, i int, matchingLines *map[int]string, addedLines *[]int) {
	if i+add >= len(lines) {
		for y := i; y < len(lines); y++ {
			added := false
			for _, num := range *addedLines {
				if num == y {
					added = true
					break
				}
			}
			if !added {
				(*matchingLines)[y] = lines[y]
				*addedLines = append(*addedLines, y)
			}
		}
	} else {
		sequence := getSequence(i, i+add, true)
		for _, j := range sequence {
			added := false
			for _, num := range *addedLines {
				if num == j {
					added = true
					break
				}
			}
			if !added {
				(*matchingLines)[j] = lines[j]
				*addedLines = append(*addedLines, j)
			}
		}
	}
}

func getSequence(start int, finish int, up bool) []int {
	sequence := []int{}
	if up {
		for i := start; i <= finish; i++ {
			sequence = append(sequence, i)
		}
	} else {
		for i := start; i >= finish; i-- {
			sequence = append(sequence, i)
		}
	}
	return sequence
}

func findMatches(lines []string, pattern string, conf *Config) (map[int]string, []int) {
	matchingLines := make(map[int]string, 0)
	var re *regexp.Regexp
	var err error

	if conf.Regexp {
		if conf.IgnoreCase {
			re, err = regexp.Compile("(?i)" + pattern)
		} else {
			re, err = regexp.Compile(pattern)
		}
		if err != nil {
			log.Fatal(err)
		}
	}

	addedLines := []int{}
	for i, line := range lines {
		match := false
		if conf.Fixed {
			if conf.IgnoreCase {
				match = strings.EqualFold(line, pattern)
			} else {
				match = line == pattern
			}
		} else if conf.Regexp {
			match = re.MatchString(line)
		} else {
			if conf.IgnoreCase {
				match = strings.Contains(strings.ToLower(line), strings.ToLower(pattern))
			} else {
				match = strings.Contains(line, pattern)
			}
		}

		if (conf.Invert && !match) || (!conf.Invert && match) {
			if conf.Before > 0 || conf.Context > 0 {
				var add int
				if conf.Before > conf.Context {
					add = conf.Before
				} else {
					add = conf.Context
				}
				addBefore(lines, add, i, &matchingLines, &addedLines)
			}
			matchingLines[i] = line
			addedLines = append(addedLines, i)
			if conf.After > 0 || conf.Context > 0 {
				var add int
				if conf.After > conf.Context {
					add = conf.After
				} else {
					add = conf.Context
				}
				addAfter(lines, add, i, &matchingLines, &addedLines)
			}
		}
	}
	addedLines = removeDuplicates(addedLines)

	return matchingLines, addedLines
}

func removeDuplicates(values []int) []int {
	seen := make(map[int]struct{})
	result := []int{}
	for _, value := range values {
		if _, ok := seen[value]; !ok {
			seen[value] = struct{}{}
			result = append(result, value)
		}
	}
	return result
}

func main() {
	conf := parseConfig()
	args := flag.Args()
	if len(args) < 2 {
		log.Fatal("Usage: go run main.go [flags] pattern filename")
	}
	pattern := args[0]
	filename := args[1]

	lines, err := readLines(filename)
	if err != nil {
		log.Fatal(err)
	}

	matchingLines, addedLines := findMatches(lines, pattern, conf)

	if conf.Count {
		fmt.Println(len(matchingLines))
	} else {
		for _, line := range addedLines {
			if conf.LineNum {
				if strings.Contains(matchingLines[line], pattern) {
					fmt.Printf("%d:%s\n", line+1, matchingLines[line])
				} else {
					fmt.Printf("%d-%s\n", line+1, matchingLines[line])
				}
			} else {
				fmt.Println(matchingLines[line])
			}
		}
	}
}
