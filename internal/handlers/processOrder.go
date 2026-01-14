package handlers

import(
	"fmt"
	"github.com/jaharbaugh/ShakerQueue/internal/app"
	"github.com/jaharbaugh/ShakerQueue/internal/models"
	"github.com/jaharbaugh/ShakerQueue/internal/queue"
) 

func HandleProcessOrder(deps app.Dependencies) func(models.OrderEvent) queue.Acktype {
    return func(event models.OrderEvent) queue.Acktype {
        // Placeholder logic
        fmt.Printf("Received order: %+v\n", event)
        return queue.Ack
    }
}