package network

import (
	"encoding/json"
	"net/http"

	"github.com/myBCA-app/nutrislice/schema"
)

func GetMenuWeekData(url string) (schema.MenuWeek, error) {
	res, err := http.Get(url)
	if err != nil {
		return schema.MenuWeek{}, err
	}

	var body []byte
	_, err = res.Body.Read(body)
	if err != nil {
		return schema.MenuWeek{}, err
	}
	
	var data schema.MenuWeek
	err = json.Unmarshal(body, &data)

	if err != nil {
		return schema.MenuWeek{}, err
	}

	return data, nil
}