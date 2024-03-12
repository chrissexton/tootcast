package mastodon

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type MastodonCredentials struct {
	serverDomain string
	accessToken  string
}

func Toot(tootText string, credentials MastodonCredentials) error {
	log.Println("posting to mastodon: " + tootText)
	client := http.Client{}
	//endpoint := "https://" + credentials.serverDomain + "/api/v1/statuses"
	u, _ := url.Parse(credentials.serverDomain)
	u = u.JoinPath("/api/v1/statuses")
	endpoint := u.String()

	form := url.Values{}
	form.Add("status", tootText)

	req, err := http.NewRequest("POST", endpoint, strings.NewReader(form.Encode()))
	//generate an error intentionally
	//req, err := http.NewRequest("POST", endpoint, nil)
	if err != nil {
		log.Println("error creating request")
		return err
	}

	req.Header.Add("Authorization", "Bearer "+credentials.accessToken)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		log.Println("error posting to mastodon")
		log.Println(err)
		return err
	}
	if resp.Status != "200 OK" {
		log.Println("error posting to mastodon: " + resp.Status)
		log.Println(resp)
		body, _ := io.ReadAll(resp.Body)
		log.Println(body)
		return err
	}

	log.Println("posted to mastodon")
	return nil
}
