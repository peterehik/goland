package whatsapp

import "time"

/**

{
  "object": "whatsapp_business_account",
  "entry": [
    {
      "id": "222847807579466",
      "changes": [
        {
          "value": {
            "messaging_product": "whatsapp",
            "metadata": {
              "display_phone_number": "15551019660",
              "phone_number_id": "243929788793960"
            },
            "contacts": [
              {
                "profile": {
                  "name": "Peter Ehikhuemen"
                },
                "wa_id": "18177166256"
              }
            ],
            "messages": [
              {
                "from": "18177166256",
                "id": "wamid.HBgLMTgxNzcxNjYyNTYVAgASGBQzQTIyN0NCQ0QzRkUwMjJBQkE3MAA=",
                "timestamp": "1706502337",
                "text": {
                  "body": "Don’t respond"
                },
                "type": "text"
              }
            ]
          },
          "field": "messages"
        }
      ]
    }
  ]
}
*/

type MessageEvent struct {
	Object string   `json:"object"`
	Entry  []*Entry `json:"entry"`
}

type Entry struct {
	ID      string    `json:"id"`
	Changes []*Change `json:"changes"`
}

type Change struct {
	Value *Value `json:"value"`
	Field string `json:"field"`
}

type Value struct {
	MessagingProduct string     `json:"messaging_product"`
	Metadata         *Metadata  `json:"metadata"`
	Contacts         []*Contact `json:"contacts"`
	Messages         []*Message `json:"messages"`
}

type Metadata struct {
	DisplayPhoneNumber string `json:"display_phone_number"`
	PhoneNumberID      string `json:"phone_number_id"`
}

type Contact struct {
	Profile *Profile `json:"profile"`
	WaID    string   `json:"wa_id"`
}

type Profile struct {
	Name string `json:"name"`
}

type Message struct {
	From      string    `json:"from"`
	ID        string    `json:"id"`
	Timestamp Timestamp `json:"timestamp"` //TODO: replace with Int
	Text      *Text     `json:"text"`
	Type      string    `json:"type"`
}

type Timestamp struct {
	time.Time
}

func (t *Timestamp) UnmarshalJSON(b []byte) error {
	ts, err := time.Parse("20060102", string(b))
	if err != nil {
		return err
	}
	t.Time = ts
	return nil
}

type Text struct {
	Body string `json:"body"`
}
