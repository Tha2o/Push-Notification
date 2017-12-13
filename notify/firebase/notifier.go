package firebase

import (
	"strconv"
	"fmt"
	"github.com/NaySoftware/go-fcm"
)



type Notifier struct{
	Serverkey string
}




func (n *Notifier) 	SendNotification(To string,Title string,Body string) {
	c := fcm.NewFcmClient(n.Serverkey)



	data1 := strconv.Quote(Title)
	data2 := strconv.Quote(Body)
	data3 := strconv.Quote("00:00")

	data := map[string]string{
		"title" : data1,
		"message" : data2,
		"timestamp": data3,
	}



	c.SetTimeToLive(4)
	c.SetPriority("High")


	c.NewFcmMsgTo(To,data)

	status, err := c.Send()
	if err == nil {
		status.PrintResults()
	} else {
		fmt.Println(err)
	}

}
