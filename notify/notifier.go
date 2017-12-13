package notify

type Notifier interface {
	SendNotification(To string,Title string, Body string)
}