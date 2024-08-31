package demo

import (
	"fmt"
	// "time"

	"github.com/myBCA-app/nutrislice/network"
)

func RunDemo() {
	// today := time.Now()
	urlFormat := "https://bergen.api.nutrislice.com/menu/api/weeks/school/bergen-academy/menu-type/lunch/%d/%d/%d"
	url := fmt.Sprintf(urlFormat, 2024, 9, 3)
	println(url)

	data, err := network.GetMenuWeekData(url)
	if err != nil {
		panic(err)
	}

	fmt.Println(data.Days[2].MenuItems[1].Food.Name)
}