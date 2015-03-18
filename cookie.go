package main

import (
	"io"
	"mega/go-util/erro"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"time"
)

func main() {

	jar, err := cookiejar.New(nil)
	erro.Trata(err)
	cookie := http.Cookie{Name: "_jira_ondemand_plugin_session",
		Domain: "jiraplugin.zendesk.com",
		Value:  `Y1A2elBhcE9hYVZVbThQb0M0YzdjT1dJaEVQaERZZ2VmaitYbWJBQXlXWHlrY3U4K1lUV1VYQUhscWJJelYxR0R1cStFbnNvcEE2WWJ1WWRrR240QnFVT2xvODNGd2UyRmpaL1hGaWdoeVJnWkpSaFJtZXpiU3l1cGRsOU9Za1NDNjRLakE5OWxZUjN5Tm9GYXBJWFhONmJibGNpSlRpZUhlTW1JeHFhdG1tdE9xTGV6cCsrUUh0U292OW9BeS9qTWFaeXZYNFIraDEyelM3aDgxekVoZz09LS11RldRckRVR0xQNFNMeHhUWGJrb1N3PT0%3D--9b6ca352ea8b0ae012c297792976ac9e08bf87d3`,
	}
	cookies := []*http.Cookie{&cookie}
	urlcookie, err := url.Parse("jiraplugin.zendesk.com")
	erro.Trata(err)
	jar.SetCookies(urlcookie, cookies)

	client := http.Client{nil, nil, jar, time.Minute}

	req, err := http.NewRequest("GET", "https://jiraplugin.zendesk.com/integrations/jira/account/megasistemas/links", nil)
	erro.Trata(err)
	resp, err := client.Do(req)
	erro.Trata(err)
	io.Copy(os.Stdout, resp.Body)
}
