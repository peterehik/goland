package ftfc

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"strings"
	"sync"
	"testing"
	"time"
)

func randName() string {
	randomString := uuid.NewString()
	splits := strings.Split(randomString, "-")
	return splits[0] + " " + splits[1]
}

func generateMessages(numMsgs int) (msgs []*SignupMessage) {
	start := time.Now()
	for i := 0; i < numMsgs; i++ {
		msgs = append(msgs, &SignupMessage{
			MessageBody: "In",
			Name:        randName(),
			PhoneNumber: uuid.NewString(),
			Timestamp:   start.Add(1 * time.Millisecond),
		})
	}
	return msgs
}

//TODO: Test a bunch of bad messages
//func TestMessageNotificationHandlerWBadMEssages {}

func TestMessageNotificationHandler(t *testing.T) {
	messagesToReceive := generateMessages(100)

	tests := []struct {
		name            string
		numMessages     int
		dogfoodMessages int // i.e. pass in some messages after generating some.
		expectedOutput  string
	}{
		{
			name:        "test with 38 players",
			numMessages: 38,
		},
		//TODO: Add some with bad messages
	}

	msgNotificationHandler := NewMessageNotificationHandler(DefaultMinPlayersPerGame, DefaultMaxPlayersPerGame)

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			wg := sync.WaitGroup{}

			//receive messages
			for _, signupMsg := range messagesToReceive[:tc.numMessages] {
				signupMsg := signupMsg
				wg.Add(1)
				go func() {
					defer wg.Done()
					err := msgNotificationHandler.ReceiveMessage(signupMsg)
					//time.Sleep(1 * time.Millisecond)
					require.NoError(t, err)
				}()
			}
			wg.Wait()
			messageToSend, err := msgNotificationHandler.GenerateCompiledMessage()
			require.NoError(t, err)
			fmt.Println(messageToSend)

		})
	}

}
