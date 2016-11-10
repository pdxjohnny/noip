package main

import (
	"encoding/base64"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

const (
	Server = "https://dynupdate.no-ip.com"
	Path   = "/nic/update"
)

func Update() error {
	c := &http.Client{}
	v := url.Values{}

	v.Set("hostname", os.Getenv("hostname"))

	username := os.Getenv("username")
	password := os.Getenv("password")
	basicAuth := base64.StdEncoding.EncodeToString(
		[]byte(username + ":" + password),
	)

	u, err := url.ParseRequestURI(Server)
	if err != nil {
		return err
	}
	u.Path = Path
	u.RawQuery = v.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	req.Header.Add("Authorization", "Basic: "+basicAuth)

	email := os.Getenv("email")
	req.Header.Add("User-Agent", "gonoipclient/v0.0.1 "+email)

	res, err := c.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	_, err = io.Copy(os.Stdout, res.Body)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := Update()
	if err != nil {
		log.Fatal(err)
	}

	c := time.Tick(10 * time.Minute)
	for range c {
		err := Update()
		if err != nil {
			log.Println(err)
		}
	}
}
