package main

import (
	"fmt"
	"time"
)

func main() {
	/*resp, err := http.Get("https://practicum.yandex.ru")
	if err != nil {
		os.Exit(1)
	}

	//var body [50]byte
	body, err2 := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err2 != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s", string(body[:500]))

	resp.Body.Close()*/
	//client := &http.Client{
	//	CheckRedirect: func(req *http.Request, via []*http.Request) error {
	//		fmt.Println(req.URL)
	//		return nil
	//	},
	//}
	//response, err := client.Get("https://ya.ru/showcaptcha?cc=1&...")
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//println(response.StatusCode)

	//println(strconv.FormatFloat(10.000, 'f', -1, 64)
	//CronRequest(time.Duration(pollInterval), time.Duration(reportInterval))

	fmt.Printf("", time.Duration(2*1000*1000*1000))

}
