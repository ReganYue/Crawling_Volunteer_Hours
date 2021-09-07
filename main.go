package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

var (
	re_acid   = `\d{10}`
	re_acName = `\"acName\":\"([\s\S]+?)\"`
	re_time   = `\"time\":\"([\s\S]+?)\"`
	re_ygs    = `:(\d{0,2}\.\d)`
)

func main() {
	for stuNo := 2019117000; stuNo <= 2019117015; stuNo++ {
		stuNo := strconv.Itoa(stuNo)
		url_details := "http://119.23.13.122/user/byactid?acid=" + stuNo
		details := GetInfos(url_details)
		fmt.Println(details)

		//encode_name:= url.QueryEscape(name)
		//url_total := "http://119.23.13.122/user/bynameid?id="+ stuNo+ "&name=" + encode_name
		//total := GetInfos(url_total)
		//fmt.Println(total)

		acid_re := regexp.MustCompile(re_acid)
		acid_rets := acid_re.FindAllStringSubmatch(details, -1)

		acName_re := regexp.MustCompile(re_acName)
		acName_rets := acName_re.FindAllStringSubmatch(details, -1)

		time_re := regexp.MustCompile(re_time)
		time_rets := time_re.FindAllStringSubmatch(details, -1)

		ygs_re := regexp.MustCompile(re_ygs)
		ygs_rets := ygs_re.FindAllStringSubmatch(details, -1)

		for _, ret := range acid_rets {
			fmt.Print("\t", ret, "\t\t\t\t\t")
		}
		fmt.Println()
		for _, ret := range acName_rets {
			fmt.Print("\t", ret[1])
		}
		fmt.Println()
		for _, ret := range time_rets {
			fmt.Print("\t", ret[1]+"\t")
		}
		fmt.Println()
		for _, ret := range ygs_rets {
			fmt.Print("\t", ret[1], "\t\t\t\t\t\t\t\t")
		}

		fmt.Println("学号为："+stuNo)
		time.Sleep(5*time.Second)
	}
}

func GetInfos(url string) string {
	resp, _ := http.Get(url)
	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)
	data := string(bytes)
	return data
}
