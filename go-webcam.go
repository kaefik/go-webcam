// go-webcam
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/parnurzeal/gorequest"
	"golang.org/x/net/html/charset"
)

//получение страницы из урла url
func gethtmlpage(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("HTTP error:", err)
		panic("HTTP error")
	}
	defer resp.Body.Close()
	// вот здесь и начинается самое интересное
	utf8, err := charset.NewReader(resp.Body, resp.Header.Get("Content-Type"))
	if err != nil {
		fmt.Println("Encoding error:", err)
		panic("Encoding error")
	}
	body, err := ioutil.ReadAll(utf8)
	if err != nil {
		fmt.Println("IO error:", err)
		panic("IO error")
	}
	return body
}

//сохранить в файл
func Savestrtofile(namef string, str string) int {
	file, err := os.Create(namef)
	if err != nil {
		// handle the error here
		return -1
	}
	defer file.Close()

	file.WriteString(str)
	return 0
}

func getimagefromcamera(url string, user string, passw string) string {
	request := gorequest.New().SetBasicAuth(user, passw)
	_, body, errs := request.Get(url).End()
	if errs != nil {
		return ""
	}
	return body
}

func main() {
	var (
		us    string
		passw string
	)
	urls := "http://192.168.0.2/image/jpeg.cgi"

	fmt.Println("Введите пользователя для доступа к камере: ")
	fmt.Scanf("%s", &us)
	fmt.Println("Введите пароль для доступа к камере: ")
	fmt.Scanf("%s", &passw)

	res := getimagefromcamera(urls, us, passw)
	Savestrtofile("image.jpg", res)

}
