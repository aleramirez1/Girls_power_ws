package repository

type WSNotifier interface {
	NotifyUser(usuarioID int, event string, payload interface{})
	NotifyMultiple(usuarioIDs []int, event string, payload interface{})
}