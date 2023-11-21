package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getHH() {
	//url := "https://hh.ru/search/vacancy?text=php&area=0" // замените на нужный URL
	url := "https://hh.ru/search/vacancy?text=php&area=0" // замените на нужный URL

	// Отправляем GET-запрос на URL
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Инициализируем goquery для парсинга HTML-страницы
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Находим все элементы с указанным CSS классом и выводим их содержимое
	// doc.Find(".bloko-header-section-3").Each(func(i int, s *goquery.Selection) {
	// 	fmt.Println(strings.TrimSpace(s.Text()))
	// })

	doc.Find(".bloko-header-section-3").First().Each(func(i int, s *goquery.Selection) {
		fmt.Println(strings.TrimSpace(s.Text()))
	})
}

func convertNumberString(numberString string) string {
	numberMap := map[string]float64{
		"k": 1000,
		"M": 1000000,
		"B": 1000000000,
	}

	numberRegex := regexp.MustCompile(`^(\d+(\.\d+)?)([kMB])?$`)
	matches := numberRegex.FindStringSubmatch(numberString)
	if len(matches) > 0 {
		number, _ := strconv.ParseFloat(matches[1], 64)
		unit := matches[3]

		if multiplier, ok := numberMap[unit]; ok {
			return strconv.FormatFloat(number*multiplier, 'f', -1, 64)
		}

		return strconv.FormatFloat(number, 'f', -1, 64)
	}

	return ""
}

func getLangGitHub(url string) (map[string]string, error) {
	res, err := http.Get(fmt.Sprintf("https://github.com/search?q=language%%3A%s&type=repositories", url))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(body)))
	if err != nil {
		return nil, err
	}

	linksElement := doc.Find(".Box-sc-g0xbh4-0.cgQapc").First()
	linksLength := linksElement.Text()

	return map[string]string{
		"count": linksLength,
		"lang":  url,
	}, nil
}

func parse() {
	url := []string{
		"JavaScript",
		"TypeScript",
		"Php",
		"Python",
		"Go",
	}
	result := []map[string]string{}
	for i := 0; i < len(url); i++ {
		getRes, err := getLangGitHub(url[i])
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		getRes["count"] = convertNumberString(strings.Split(getRes["count"], " ")[0])
		result = append(result, getRes)
	}
	fmt.Println(result)
}

func main() {
	parse()
}
