package restapi

import "bitmex/conf"
import "fmt"

func Announcement() {
	data := new(RestData)
	data.uri = conf.Announcement
	data.method = 1
	//data.data = make(map[string]string){}
	res, err := getData(data)
	fmt.Println(res)
	fmt.Println(err)

}
