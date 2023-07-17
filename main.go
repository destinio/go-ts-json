package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	// https://pkg.go.dev/net/http
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		fmt.Println("error GETing end point")
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error reading body")
		os.Exit(1)
	}

	// Unmarshal
	// https://pkg.go.dev/encoding/json#Unmarshal
	posts := []Post{}

	err = json.Unmarshal(body, &posts)
	if err != nil {
		fmt.Println("error unmarshalling", err)
		os.Exit(1)
	}

	fmt.Println(posts[0].Title)

	// Marshal
	// https://pkg.go.dev/encoding/json#Marshal
	newPost := Post{
		UserID: 1,
		ID:     1,
		Title:  "New Post",
		Body:   "New Post Body",
	}

	newPostJSON, err := json.Marshal(newPost)
	if err != nil {
		fmt.Println("error marshalling", err)
		os.Exit(1)
	}

	// make POST request
	// https://pkg.go.dev/net/http#Post

	bufferedPost := bytes.NewBuffer(newPostJSON)

	resp, err = http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", bufferedPost)
	if err != nil {
		fmt.Println("error POSTing end point")
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error reading body")
		os.Exit(1)
	}

	if resp.StatusCode != 201 {
		fmt.Println("error creating post", resp.StatusCode)
		os.Exit(1)
	}

	fmt.Println(string(body))

	// cutomize request headers
	// https://pkg.go.dev/net/http#Request
	req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bufferedPost)
	if err != nil {
		fmt.Println("error creating request")
		os.Exit(1)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-My-Header", "my value")

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("error POSTing end point")
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error reading body")
		os.Exit(1)
	}

	if resp.StatusCode != 201 {
		fmt.Println("error creating post", resp.StatusCode)
		os.Exit(1)
	}

	fmt.Println(string(body))

}
