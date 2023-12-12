package helpers

import (
	"fmt"
	"math"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func GenerateRandomString(length int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	s := make([]rune, length)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func GenerateCode(prefix string, number int) string {
	UpperCasePrefix := strings.ToUpper(prefix)
	var suffix string

	if number < 10 {
		suffix = "000" + strconv.Itoa(number)
	} else if number < 100 {
		suffix = "00" + strconv.Itoa(number)
	} else if number < 1000 {
		suffix = "0" + strconv.Itoa(number)
	} else {
		suffix = strconv.Itoa(number)
	}
	return fmt.Sprintf("%s-%s", UpperCasePrefix, suffix)

}

func GetCurrentDate() string {
	currentDate := time.Now().Format("2006-01-02")
	return currentDate
}

func FormatDate(date string) string {
	ok, _ := regexp.Match("^[0-9]{4}-[0-9]{2}-[0-9]{2}$", []byte(date))
	if !ok {
		// log.Printf("date: %s", date)
		parsedDate, _ := time.Parse(time.RFC3339, date)
		// log.Printf("parsedDate: %s", parsedDate)
		return parsedDate.Local().Format("2006-01-02")
	}
	return date

}

func RoundDecimalPlaces(value float64, precision int) float64 {
	shift := math.Pow(10, float64(precision))
	return math.Round(value*shift) / shift
}
