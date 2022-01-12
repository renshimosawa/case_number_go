package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	get_url_info, err := goquery.NewDocument("https://www.pref.aomori.lg.jp/soshiki/kenko/hoken/covid19-press.html")
	if err != nil {
		fmt.Println("cannot get HTML")
	}
	

	result1 := get_url_info.Find("div.inner > table.bc > tbody > tr.textleft").First()
	result1.Each(func(index int, s *goquery.Selection) {
		tabledata := (s.Text())
		day := strings.Split(tabledata, "\n")[1]
		cases := strings.Split(tabledata, "\n")[3]
		formated := strings.Replace(cases,"新型コロナウイルス感染症患者（","",1)
		number := strings.Replace(formated, "例）を確認", "",1)
		message := day + "公表" + "\n青森県新規感染者数 " + number + "人"
		fmt.Println(message)
	})

}
