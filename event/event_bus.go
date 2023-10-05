package event

type EventName string

type Event interface {
	GetPayload() map[string]interface{}
}

type EventBus interface {
	Publish(eventName EventName, event Event) error
	Subscribe(eventName EventName, subscriber func(Event) error) error
}
