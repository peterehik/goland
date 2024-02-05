package ftfc

import (
	"fmt"
	"slices"
	"strings"
	"sync"
	"time"
)

const (
	DefaultMaxPlayersPerGame = 18
	DefaultMinPlayersPerGame = 14
)

type SignupMessage struct {
	Timestamp   time.Time
	MessageBody string
	Name        string
	PhoneNumber string
}

type MessageNotificationHandler struct {
	mutex                         sync.Mutex
	playersMap                    map[string]*SignupMessage
	msgReceivedSinceLastGenerated bool
	lastCompiledMessage           string
	maxPlayersPerGame             int
	minPlayersPerGame             int
}

func NewMessageNotificationHandler(minPlayersPerGame int, maxPlayersPerGame int) *MessageNotificationHandler {
	return &MessageNotificationHandler{
		minPlayersPerGame: minPlayersPerGame,
		maxPlayersPerGame: maxPlayersPerGame,
		playersMap:        make(map[string]*SignupMessage),
	}
}

func (handler *MessageNotificationHandler) ReceiveMessage(message *SignupMessage) error {
	defer handler.mutex.Unlock()
	handler.mutex.Lock()
	if isInMessage(message.MessageBody) {
		handler.playersMap[message.PhoneNumber] = message
		handler.msgReceivedSinceLastGenerated = true
		return nil
	}
	if isOutMessage(message.MessageBody) {
		delete(handler.playersMap, message.PhoneNumber)
		handler.msgReceivedSinceLastGenerated = true
		return nil
	}
	return fmt.Errorf("bad message: %+v", message)
}

func (handler *MessageNotificationHandler) GenerateCompiledMessage() (string, error) {
	if !handler.msgReceivedSinceLastGenerated {
		return handler.lastCompiledMessage, nil
	}

	players := make([]*SignupMessage, 0)

	handler.mutex.Lock()
	for _, v := range handler.playersMap {
		players = append(players, v)
	}
	handler.mutex.Unlock()

	slices.SortFunc(players, func(a, b *SignupMessage) int {
		return a.Timestamp.Compare(b.Timestamp)
	})

	var lastPlayer int
	if len(players) >= handler.maxPlayersPerGame*2 {
		lastPlayer = handler.maxPlayersPerGame * 2
	} else if len(players) >= handler.minPlayersPerGame*2 {
		lastPlayer = handler.minPlayersPerGame * 2
	} else {
		lastPlayer = handler.maxPlayersPerGame
	}

	result := ""
	for i := 0; i < len(players); i++ {
		num := i + 1
		player := players[i]
		result += fmt.Sprintf("%2d. %s\n", num, player.Name)
		if num == lastPlayer {
			result += "\n============BENCH================\n"
		}
	}
	return result, nil
}

func isOutMessage(message string) bool {
	return strings.ToUpper(strings.Trim(message, " ")) == "OUT"
}

func isInMessage(message string) bool {
	return strings.ToUpper(strings.Trim(message, " ")) == "IN"
}
