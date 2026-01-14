package queue

type SimpleQueueType int

const (
	SimpleQueueDurable SimpleQueueType = iota
	SimpleQueueTransient
)

const (
	ExchangeDirect = "shaker_queue_direct"
	ExchangeTopic  = "shaker_queue_topic"
)