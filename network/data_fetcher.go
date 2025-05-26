package network

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/xtt28/nutrislice/schema"
)

func GetMenuWeekData(url string) (schema.MenuWeek, error) {
	res, err := http.Get(url)
	if err != nil {
		return schema.MenuWeek{}, err
	}
	if res.StatusCode != http.StatusOK {
		errMsg := fmt.Sprintf("menu API endpoint did not return 200; returned %d", res.StatusCode)
		return schema.MenuWeek{}, errors.New(errMsg)
	}

	dec := json.NewDecoder(res.Body)
	var data schema.MenuWeek

	err = dec.Decode(&data)

	if err != nil {
		return schema.MenuWeek{}, err
	}

	return data, nil
}
