package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Joke struct {
	Setup     string `json:"setup"`
	Punchline string `json:"punchline"`
}

type Response struct {
	Body []Joke `json:"body"`
}

type APIKey struct {
	Key string `json:"api_key"`
}

func main() {
	app := fiber.New()

	app.Get("/", catchword)

	log.Fatal(app.Listen(":8000"))
}

func catchword(c *fiber.Ctx) error {
	URL := "https://dad-jokes.p.rapidapi.com/random/joke"

	req, _ := http.NewRequest("GET", URL, nil)

	filePath := "key.json"
	jsonBytes, _ := ioutil.ReadFile(filePath)

	var apiKeyData APIKey
	json.Unmarshal(jsonBytes, &apiKeyData)

	ApiKEY := apiKeyData.Key
	req.Header.Add("X-RapidAPI-Key", ApiKEY)
	req.Header.Add("X-RapidAPI-Host", "dad-jokes.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return err
	}
	stringBody := string(body)

	var response Response

	err = json.Unmarshal([]byte(stringBody), &response)
	if err != nil {
		log.Println(err)
		return err
	}

	return c.JSON(response.Body)
}
