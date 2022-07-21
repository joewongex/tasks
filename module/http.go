package module

import (
	"net/http"
	"net/http/cookiejar"
	"time"

	"golang.org/x/net/publicsuffix"
)

func NewClient() (client *http.Client, err error) {
	var jar *cookiejar.Jar
	cookieOptions := cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	}
	jar, err = cookiejar.New(&cookieOptions)
	if err != nil {		
		return
	}
	client = &http.Client{
		Jar:     jar,
		Timeout: 10 * time.Second,
	}
	return
}
