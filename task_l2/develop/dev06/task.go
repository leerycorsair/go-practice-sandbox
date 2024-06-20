package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type Config struct {
	Fields    []int
	Delimeter string
	Separated bool
}

func parseConfig() (*Config, error) {
	fieldsFlag := flag.String("f", "", "fields (columns)")
	delimiter := flag.String("d", "\t", "delimiter")
	separated := flag.Bool("s", false, "delimeter only strings")
	flag.Parse()

	fields := strings.Split(*fieldsFlag, ",")
	if len(fields) == 0 {
		return nil, fmt.Errorf("invalid fields cnt")
	}

	conf := Config{}
	for _, field := range fields {
		f, err := strconv.Atoi(field)
		if err != nil || f < 1 {
			return nil, fmt.Errorf("invalid field")
		}
		conf.Fields = append(conf.Fields, f-1)
	}
	conf.Delimeter = *delimiter
	conf.Separated = *separated
	return &conf, nil
}

func cutLines(lines []string, conf *Config) []string {
	var cutFields []string

	for _, line := range lines {
		if conf.Separated && !strings.Contains(line, conf.Delimeter) {
			continue
		}

		fields := strings.Split(line, conf.Delimeter)
		var currLine []string
		for _, fieldId := range conf.Fields {
			if fieldId < len(fields) {
				currLine = append(currLine, fields[fieldId])
			}
		}
		cutFields = append(cutFields, strings.Join(currLine, conf.Delimeter))
	}
	return cutFields
}

func main() {
	conf, err := parseConfig()
	if err != nil {
		log.Fatal(err)
	}

	var lines []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		lines = append(lines, line)
	}

	separatedLines := cutLines(lines, conf)
	for _, line := range separatedLines {
		fmt.Println(line)
	}
}
