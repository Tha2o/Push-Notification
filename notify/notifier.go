package notify

type Notifier interface {
	//SendNotification()
	SendNotification(To string,Title string, Body string)
}