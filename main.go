package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	// fmt.Println(mypkg.scraping())
	err := godotenv.Load(fmt.Sprintf("./%s.env", os.Getenv("GO_ENV")))
		Token_Me := os.Getenv("LINE_TOKEN_MINE")
		// Token_F := os.Getenv("LINE_TOKEN_F")
		// Token_Y := os.Getenv("LINE_TOKEN_Y")
		// Token_N := os.Getenv("LINE_TOKEN_N")
		// Token_GM := os.Getenv("LINE_TOKEN_GM")

		// accessToken_GM := Token_GM
		accessToken_Me := Token_Me
		msg := "message"
		URL := "https://notify-api.line.me/api/notify"
		u, err := url.ParseRequestURI(URL)
		if err != nil {
				log.Fatal(err)
		}
		c := &http.Client{}
		form := url.Values{}
		form.Add("message", msg)
		body := strings.NewReader(form.Encode())
		req, err := http.NewRequest("POST", u.String(), body)
		if err != nil {
				log.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		// req.Header.Set("Authorization", "Bearer "+ accessToken_GM)
		req.Header.Set("Authorization", "Bearer "+ accessToken_Me)

		_, err = c.Do(req)
		if err != nil {
				log.Fatal(err)
		}
}