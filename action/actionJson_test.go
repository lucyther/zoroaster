package action

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewActionJson(t *testing.T) {
	var s = `{"UserId": 1, "TriggerId": 30, "ActionType": "webhook_post", "Attributes": {"URI": "https://webhook.site/202d0fac-4bfa-43f5-8ad0-c791cf051e5f"}}`
	_, err := NewActionJson([]byte(s))
	if err != nil {
		t.Error(err)
	}
}

func TestAction_Webhook(t *testing.T) {
	var s = `{"UserId": 1, "TriggerId": 30, "ActionType": "webhook_post", "Attributes": {"URI": "https://webhook.site/202d0fac-4bfa-43f5-8ad0-c791cf051e5f"}}`
	a := Action{}

	err := json.Unmarshal([]byte(s), &a)
	if err != nil {
		t.Error(err)
	}

	_, ok := a.Attribute.(AttributeWebhookPost)
	assert.True(t, ok)
}

func TestAction_Email(t *testing.T) {
	var s = `
	{  
	   "UserId":1,
	   "TriggerId":30,
	   "ActionType":"email",
	   "Attributes":{  
		  "To":[  
			 "manlio.poltronieri@gmail.com",
			 "marco@atomic.eu.com"
		  ],
		  "Subject":"YO from Zoroaster",
		  "Body":"yo yo yo and a bottle of rum"
	   }
	}`
	a := Action{}

	err := json.Unmarshal([]byte(s), &a)
	if err != nil {
		t.Error(err)
	}

	_, ok := a.Attribute.(AttributeEmail)
	assert.True(t, ok)
}
