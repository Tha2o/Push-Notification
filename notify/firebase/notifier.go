package firebase

import (
	"strconv"
	"fmt"
	"github.com/NaySoftware/go-fcm"
)

//const	serverKey = "AAAA1taAHqo:APA91bHOLq8051Ia5lybfcv5OWaBPGguxzRwNaDgODX-CmTuxqzscTEoyDs1BHsivK5WnAkLE5PfyY28HQQopYrjszmwz1AEk9bLmhydhiGs_HPFL-KH14eWxjhpefwYIbKJbj6RDc-P"

//type Notifier struct {
//	To []string
//	Title string
//	Body string
//}

type Notifier struct{
	Serverkey string
}

//var notification fcm.NotificationPayload


func (n *Notifier) 	SendNotification(To string,Title string,Body string) {
	c := fcm.NewFcmClient(n.Serverkey)


	//notification.Body = n.Body;
	//notification.Title = n.Title;


	data1 := strconv.Quote(Title)
	data2 := strconv.Quote(Body)
	data3 := strconv.Quote("00:00")

	data := map[string]string{
		"title" : data1,
		"message" : data2,
		"timestamp": data3,
	}



	//c.SetNotificationPayload(&notification);
	c.SetTimeToLive(4)
	c.SetPriority("High")

	//c.NewFcmRegIdsMsg(n.To, data)

	c.NewFcmMsgTo(To,data)

	status, err := c.Send()
	if err == nil {
		status.PrintResults()
	} else {
		fmt.Println(err)
	}

}
