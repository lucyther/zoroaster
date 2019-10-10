package action

import (
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/aws/aws-sdk-go/service/ses/sesiface"
	"github.com/magiconair/properties/assert"
	"io"
	"net/http"
	"testing"
	"zoroaster/trigger"
)

// HTTP Client mock
type mockHttpClient struct{}

func (m mockHttpClient) Post(url, contentType string, body io.Reader) (*http.Response, error) {
	resp := http.Response{StatusCode: 200}
	return &resp, nil
}

func TestHandleWebHookPost(t *testing.T) {

	tg, _ := trigger.NewTriggerFromFile("../resources/triggers/wac1.json")
	url := AttributeWebhookPost{URI: "https://hal.xyz"}
	cnMatch := trigger.CnMatch{
		tg,
		8888,
		1554828248,
		"0x",
		10,
		"matched values",
		"all values"}

	outcome := handleWebHookPost(url, cnMatch, mockHttpClient{})

	expectedPayload := `{"BlockNo":8888,"BlockTimestamp":1554828248,"ReturnedValue":"matched values","AllValues":"all values"}`
	assert.Equal(t, outcome.Payload, expectedPayload)
	assert.Equal(t, outcome.Outcome, `{"StatusCode":200}`)
}

// SESAPI mock
type mockSESClient struct {
	sesiface.SESAPI
}

func (m *mockSESClient) SendEmail(*ses.SendEmailInput) (*ses.SendEmailOutput, error) {
	msg := "mock email success"
	return &ses.SendEmailOutput{MessageId: &msg}, nil
}

func TestHandleEmail1(t *testing.T) {

	email := AttributeEmail{
		From:    "hello@wolrd.com",
		To:      []string{"manlio.poltronieri@gmail.com", "$AllValues$"},
		Subject: "Hello World Test",
		Body:    "body",
	}

	tg, _ := trigger.NewTriggerFromFile("../resources/triggers/wac1.json")

	payload := trigger.CnMatch{
		Trigger:        tg,
		MatchId:        1,
		BlockNo:        1,
		MatchedValues:  "",
		AllValues:      "[\"marco@atomic.eu.com\"",
		BlockTimestamp: 123,
		BlockHash:      "0x",
	}
	outcome := handleEmail(email, payload, &mockSESClient{})
	expectedPayload := `{"Recipients":["manlio.poltronieri@gmail.com","marco@atomic.eu.com"],"Body":"body"}`

	assert.Equal(t, outcome.Payload, expectedPayload)
}

func TestHandleEmail2(t *testing.T) {

	email := AttributeEmail{
		From:    "hello@wolrd.com",
		To:      []string{"manlio.poltronieri@gmail.com", "$AllValues$"},
		Subject: "Hello World Test",
		Body:    "body",
	}

	tg, _ := trigger.NewTriggerFromFile("../resources/triggers/wac1.json")

	payload := trigger.CnMatch{
		Trigger:        tg,
		MatchId:        1,
		BlockNo:        1,
		MatchedValues:  "",
		AllValues:      "[[\"marco@atomic.eu.com\",\"matteo@atomic.eu.com\",\"not and address\"]]",
		BlockTimestamp: 123,
		BlockHash:      "0x",
	}
	outcome := handleEmail(email, payload, &mockSESClient{})
	expectedPayload := `{"Recipients":["manlio.poltronieri@gmail.com","marco@atomic.eu.com","matteo@atomic.eu.com"],"Body":"body"}`
	assert.Equal(t, outcome.Payload, expectedPayload)
}

func TestHandleEmail3(t *testing.T) {

	email := AttributeEmail{
		From:    "hello@wolrd.com",
		To:      []string{"manlio.poltronieri@gmail.com", "$AllValues$"},
		Subject: "Hello World Test",
		Body:    "body",
	}

	tg, _ := trigger.NewTriggerFromFile("../resources/triggers/wac1.json")

	payload := trigger.CnMatch{
		Trigger:        tg,
		MatchId:        1,
		BlockNo:        1,
		MatchedValues:  "",
		AllValues:      "[4#END# \"manlio.poltronieri@gmail.com\"#END# \"hello@world.com\"]",
		BlockTimestamp: 123,
		BlockHash:      "0x",
	}
	outcome := handleEmail(email, payload, &mockSESClient{})
	expectedPayload := `{"Recipients":["manlio.poltronieri@gmail.com","hello@world.com"],"Body":"body"}`
	assert.Equal(t, outcome.Payload, expectedPayload)
}
