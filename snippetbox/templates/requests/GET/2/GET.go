package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Data struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

func main() {
	// создали клиент для отправки запросов
	// дефолтный клиент не всегда подойдёт так как сервер может виснуть, а в дефолтном примере например нет таймаута, то есть он ждёт вечно
	client := &http.Client{} // == http.defltclient

	req, err := http.NewRequest("GET", "https://icanhazdadjoke.com", nil)
	if err != nil {
		log.Fatalln(err)
	}
	// установка заголовка запроса, заголовок - key, значение - value;
	// делается для того чтобы показать какого формата ответа ждёт клиент ( отправил json, ждёт json; xml = xml)
	// если отправить просто GET то вернётся формат HTML вместо json
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(b))

	var data Data
	if err := json.Unmarshal(b, &data); err != nil {
		log.Fatal(err)
	}
	fmt.Println(data.Joke)
	fmt.Println("Response Headers: ", resp.Header)
	fmt.Println("Response Status: ", resp.StatusCode)
	fmt.Println("Response Body: ", string(b))
}
