package core

type UpdateNotifier struct {
    Updates chan bool
}

func NewUpdateNotifier() *UpdateNotifier {
    return &UpdateNotifier{
        Updates: make(chan bool, 1),
    }
}

func (notifier *UpdateNotifier) NotifyUpdate() {
    select {
    case notifier.Updates <- true:
    default:
    }
}