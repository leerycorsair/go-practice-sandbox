package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type Config struct {
	Key     int
	Numeric bool
	Reverse bool
	Unique  bool
	Month   bool
	Blanks  bool
	Check   bool
	// HumanNumeric bool
}

func parseConfig() *Config {
	kFlag := flag.Int("k", 1, "key")
	nFlag := flag.Bool("n", false, "numeric sort")
	rFlag := flag.Bool("r", false, "reverse")
	uFlag := flag.Bool("u", false, "unique")
	MFlag := flag.Bool("M", false, "month sort")
	bFlag := flag.Bool("b", false, "ignore leading blanks")
	cFlag := flag.Bool("c", false, "check")
	// hFlag := flag.Bool("h", false, "human numeric sort")
	flag.Parse()

	return &Config{
		Key:     *kFlag,
		Numeric: *nFlag,
		Reverse: *rFlag,
		Unique:  *uFlag,
		Month:   *MFlag,
		Blanks:  *bFlag,
		Check:   *cFlag,
		// HumanNumeric: *hFlag,
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
		if line != "" {
			lines = append(lines, line)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func writeLines(filename string, lines []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		line += "\n"
		if _, err := writer.WriteString(line); err != nil {
			return err
		}
	}
	if err := writer.Flush(); err != nil {
		return err
	}
	return nil
}

func removeLeadingBlanks(lines []string) []string {
	for i, line := range lines {
		lines[i] = strings.TrimLeft(line, " \t")
	}
	return lines
}

func removeDuplicates(lines []string) []string {
	seen := make(map[string]struct{})
	result := []string{}
	for _, line := range lines {
		if _, ok := seen[line]; !ok {
			seen[line] = struct{}{}
			result = append(result, line)
		}
	}
	return result
}

func isSorted(lines []string, cmp func(str1, str2 string) bool, checker func(str string) bool) (bool, error) {
	if len(lines) <= 1 {
		return true, nil
	}
	for i := range len(lines) - 1 {
		if checker != nil && (!checker(lines[i]) || !checker(lines[i+1])) {
			return false, fmt.Errorf("cmp error")
		}
		if cmp(lines[i], lines[i+1]) {
			return false, nil
		}
	}
	return true, nil
}

var months = map[string]int{
	"JANUARY":   1,
	"FEBRUARY":  2,
	"MARCH":     3,
	"APRIL":     4,
	"MAY":       5,
	"JUNE":      6,
	"JULY":      7,
	"AUGUST":    8,
	"SEPTEMBER": 9,
	"OCTOBER":   10,
	"NOVEMBER":  11,
	"DECEMBER":  12,
}

func monthCmp(str1, str2 string) bool {
	key1, ok1 := months[strings.ToUpper(str1)]
	key2, ok2 := months[strings.ToUpper(str2)]
	if !ok1 {
		return false
	}
	if !ok2 {
		return true
	}
	return key1 < key2
}

func mySort(lines []string, conf *Config) ([]string, error) {
	if conf.Blanks {
		lines = removeLeadingBlanks(lines)
	}
	if conf.Unique {
		lines = removeDuplicates(lines)
	}

	key := conf.Key - 1
	var cmp func(i, j int) bool
	if conf.Numeric {
		cmp = func(i, j int) bool {
			num1, err1 := strconv.ParseFloat(getField(lines[i], key), 64)
			num2, err2 := strconv.ParseFloat(getField(lines[j], key), 64)
			if err1 != nil {
				return conf.Reverse
			}
			if err2 != nil {
				return !conf.Reverse
			}
			if conf.Reverse {
				return num1 > num2
			}
			return num1 < num2
		}
	} else if conf.Month {
		cmp = func(i, j int) bool {
			month1 := getField(lines[i], key)
			month2 := getField(lines[j], key)
			result := monthCmp(month1, month2)
			if conf.Reverse {
				return !result
			}
			return result
		}
	} else {
		cmp = func(i, j int) bool {
			field1 := getField(lines[i], key)
			field2 := getField(lines[j], key)
			if conf.Reverse {
				return field1 > field2
			}
			return field1 < field2
		}
	}
	// else if conf.HumanNumeric {
	// 	cmp = nil
	// }

	if conf.Check {
		check, err := isSorted(lines, func(str1, str2 string) bool { return !cmp(0, 1) }, nil)
		if err != nil {
			return nil, err
		}
		if check {
			return nil, fmt.Errorf("data is already sorted")
		}
	}

	sort.SliceStable(lines, cmp)
	return lines, nil
}

func getField(line string, key int) string {
	fields := strings.Fields(line)
	if key >= 0 && key < len(fields) {
		return fields[key]
	}
	return ""
}

func main() {
	conf := parseConfig()
	if flag.NArg() == 0 {
		log.Fatalf("no filename")
	}
	inputFile := flag.Args()[0]
	outputFile := flag.Args()[1]

	lines, err := readLines(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	lines, err = mySort(lines, conf)
	if err != nil {
		log.Fatal(err)
	}

	writeLines(outputFile, lines)
}
