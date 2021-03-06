package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/Sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	formatter := &logrus.TextFormatter{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var line map[string]interface{}
		if err := json.Unmarshal(scanner.Bytes(), &line); err != nil {
			continue
		}
		e := &logrus.Entry{
			Logger:  logger,
			Data:    line,
			Time:    mustParseTime(line["time"].(string)),
			Level:   mustParseLevel(line["level"].(string)),
			Message: line["msg"].(string),
		}

		if b, err := formatter.Format(e); err != nil {
			continue
		} else {
			fmt.Printf("%s\n\r", b)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func mustParseTime(value string) time.Time {
	t, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return time.Now()
	}
	return t
}

func mustParseLevel(value string) logrus.Level {
	l, err := logrus.ParseLevel(value)
	if err != nil {
		return logrus.ErrorLevel
	}
	return l
}
