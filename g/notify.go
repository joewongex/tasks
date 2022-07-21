package g

import (
	"strings"
)

func BarkNotify(title, body string) (err error) {
	data := map[string]string{"title": title, "body": body}
	url := strings.TrimSuffix(Config.Notify.Bark.Url, "/") + "/" + strings.TrimPrefix(Config.Notify.Bark.DeviceKey, "/")
	res, err := Post(url, data, ContentTypeJSON)
	if err != nil {
		return
	}
	defer res.Body.Close()

	return
}
