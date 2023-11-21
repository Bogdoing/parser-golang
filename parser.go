// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"regexp"
// 	"strconv"
// 	"strings"
// 	"time"
// )

// var langGithub = []string{"javascript", "python", "java", "ruby", "go"}

// func convertNumberString(numberString string) string {
// 	numberMap := map[string]float64{
// 		"k": 1000,
// 		"M": 1000000,
// 		"B": 1000000000,
// 	}

// 	numberRegex := regexp.MustCompile(`^(\d+(\.\d+)?)([kMB])?$`)
// 	matches := numberRegex.FindStringSubmatch(numberString)
// 	if len(matches) > 0 {
// 		number, _ := strconv.ParseFloat(matches[1], 64)
// 		unit := matches[3]

// 		if multiplier, ok := numberMap[unit]; ok {
// 			return strconv.FormatFloat(number*multiplier, 'f', -1, 64)
// 		}

// 		return strconv.FormatFloat(number, 'f', -1, 64)
// 	}

// 	return ""
// }

// func getLangGitHub(url string) (map[string]string, error) {
// 	resp, err := http.Get(fmt.Sprintf("https://github.com/search?q=language%%3A%s&type=repositories", url))
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return nil, err
// 	}

// 	linksRegex := regexp.MustCompile(`<span class="Box-sc-g0xbh4-0 cgQapc">(.+?)</span>`)
// 	linksMatch := linksRegex.FindStringSubmatch(string(body))
// 	if len(linksMatch) > 0 {
// 		linksLength := linksMatch[1]
// 		return map[string]string{
// 			"count": linksLength,
// 			"lang":  url,
// 		}, nil
// 	}

// 	return nil, nil
// }

// func sleep(ms time.Duration) {
// 	time.Sleep(ms * time.Millisecond)
// }

// func main() {
// 	result := make([]map[string]string, 0)
// 	for _, lang := range langGithub {
// 		getRes, err := getLangGitHub(lang)
// 		for getRes == nil || getRes["count"] == "" {
// 			sleep(1000)
// 			getRes, err = getLangGitHub(lang)
// 		}
// 		if err != nil {
// 			fmt.Println("Error:", err)
// 		}

// 		getRes["count"] = convertNumberString(strings.Split(getRes["count"], " ")[0])
// 		fmt.Println(getRes)
// 		result = append(result, getRes)
// 	}

// 	fmt.Println(result)
// }
