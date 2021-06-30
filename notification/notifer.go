package notification

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"sync"

	"github.com/JinlongWukong/DevLab/config"
	"github.com/JinlongWukong/DevLab/manager"
	"github.com/JinlongWukong/DevLab/utils"
)

var notificationKind = "webex"
var queueSize = 1000

//Message channel buffer
var MessageChan = make(chan Message, queueSize)

// MessageCreateRequest is the Create Message Request Parameters
type Message struct {
	Target string `json:"target,omitempty"` // Target ID.
	Text   string `json:"text,omitempty"`   // Message in plain text format.
}

type Notifier struct {
}

var _ manager.Manager = Notifier{}

//initialize configuration
func init() {
	if config.Notification.Kind != "" {
		notificationKind = config.Notification.Kind
	}
	if config.Notification.QueueSize > 0 {
		queueSize = config.Notification.QueueSize
	}
}

//notficaiont internal api interface
func SendNotification(message Message) {

	select {
	case MessageChan <- message:
	default:
		log.Println("Message buffer is full!!!")
	}

}

//controller loop
func (n Notifier) Control(ctx context.Context, wg *sync.WaitGroup) {

	log.Println("Notification manager started")
	defer func() {
		log.Println("Notification manager exited")
		wg.Done()
	}()

	myToken := os.Getenv("BOT_TOKEN")

	for {
		select {
		case <-ctx.Done():
			return
		case message := <-MessageChan:
			formatedMessage := WebexMessageRequest{ToPersonEmail: message.Target, Text: message.Text}
			payload, _ := json.Marshal(formatedMessage)
			err, _ := utils.HttpSendJsonDataWithAuthBearer("https://webexapis.com/v1/messages", "POST", myToken, payload)
			if err != nil {
				log.Println(err)
			}
		}
	}

}
