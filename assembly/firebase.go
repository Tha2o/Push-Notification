package main

import (
	"github.com/Tha2o/Push-Notification/notify/firebase"
)



func main() {

	notifier := firebase.Notifier{
		Serverkey: "AAAA1taAHqo:APA91bHOLq8051Ia5lybfcv5OWaBPGguxzRwNaDgODX-CmTuxqzscTEoyDs1BHsivK5WnAkLE5PfyY28HQQopYrjszmwz1AEk9bLmhydhiGs_HPFL-KH14eWxjhpefwYIbKJbj6RDc-P",
	}

	notifier.SendNotification("ey0sPpEuxxw:APA91bGbLnKGd60kzuCi7d7VPDkHGE4CmQ93A1X1AG9RiIidUnaSsv6w6Wcwj8XXlXUekYiBpFUTdDDC7sqgKW1pVHZkaU77Mf47BlPSvIwwRvMO4F7laUFo5utNYPK6eV3ruBjk9tzg",
		"Title","Body of the title")


}
