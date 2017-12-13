package main

import (

	"github.com/andile/go/firebase/notify/firebase"
	"github.com/andile/go/firebase/notify"

)



func main() {


//	isoNotifier := firebase.Notifier{
//		To: []string{
//			"ey0sPpEuxxw:APA91bGbLnKGd60kzuCi7d7VPDkHGE4CmQ93A1X1AG9RiIidUnaSsv6w6Wcwj8XXlXUekYiBpFUTdDDC7sqgKW1pVHZkaU77Mf47BlPSvIwwRvMO4F7laUFo5utNYPK6eV3ruBjk9tzg",

//		},Title:"Title",Body:"Body",
//	}

//		n := NotifierContract{
//		Notifier: &isoNotifier,
//	}
//		n.Notifier.SendNotification()
//

isoNotifier := firebase.Notifier{
	Serverkey: "AAAA1taAHqo:APA91bHOLq8051Ia5lybfcv5OWaBPGguxzRwNaDgODX-CmTuxqzscTEoyDs1BHsivK5WnAkLE5PfyY28HQQopYrjszmwz1AEk9bLmhydhiGs_HPFL-KH14eWxjhpefwYIbKJbj6RDc-P",
	}

n := NotifierContract{
	Notifier:&isoNotifier,
}

n.Notifier.SendNotification("ey0sPpEuxxw:APA91bGbLnKGd60kzuCi7d7VPDkHGE4CmQ93A1X1AG9RiIidUnaSsv6w6Wcwj8XXlXUekYiBpFUTdDDC7sqgKW1pVHZkaU77Mf47BlPSvIwwRvMO4F7laUFo5utNYPK6eV3ruBjk9tzg",
	"Title","Its Alive")

}
ttype NotifierContract struct {
	Notifier notify.Notifier
}

