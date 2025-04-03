package joke

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Joke struct {
	Title string
	Joke  string
}

func (j *Joke) String() string {
	return fmt.Sprintf("%s\n%s", j.Title, j.Joke)
}

const (
	url       = "https://www.reddit.com/r/Jokes/top.json?t=day&limit=1"
	userAgent = "motd-joke/0.0.1"
)

type ListingResponse struct {
	Data Data `json:"data"`
}

type Data struct {
	Posts []Post `json:"children"`
}

type Post struct {
	Data PostData `json:"data"`
}

type PostData struct {
	Joke  string `json:"selftext"`
	Title string `json:"title"`
}

func makeRequest() *http.Request {
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Add("User-Agent", userAgent)
	return request
}

func GetJoke() (Joke, error) {
	resp, err := http.DefaultClient.Do(makeRequest())
	if err != nil {
		return Joke{}, err
	}
	defer resp.Body.Close()

	var listing ListingResponse

	if err := json.NewDecoder(resp.Body).Decode(&listing); err != nil {
		return Joke{}, err
	}

	joke := Joke{Joke: listing.Data.Posts[0].Data.Joke, Title: listing.Data.Posts[0].Data.Title}
	return joke, nil
}
