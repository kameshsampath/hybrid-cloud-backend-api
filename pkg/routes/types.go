package routes

import (
	"fmt"
	"strconv"
)

//Endpoints is a marker for defining REST API routes
type Endpoints struct {
}

//NewEndpoints provides handle to the REST URI defined in this API
func NewEndpoints() *Endpoints {
	return &Endpoints{}
}

//Message handles the message that needs to be processed
type Message struct {
	//RequestId the unique request id
	RequestId string  `json:"requestId"`
	Request   Request `json:"request"`
}

//Request the worker request to process
type Request struct {
	//Text is any text to process
	Text string `json:"text"`
	//Uppercase change the Text to uppercase
	Uppercase bool `json:"upperCase,omitempty"`
	//Reverse   reverse Text
	Reverse bool `json:"reverse,omitempty"`
	//SleepMillis add some sleep to processing
	SleepMillis int `json:"sleepMillis,omitempty"`
}

//Response is the processed Request
type Response struct {
	//RequestId the request id
	RequestId string `json:"requestId"`
	//WorkerId the worker id
	WorkerId string `json:"workerId"`
	//CloudId the cloud which processed the request
	CloudId string `json:"cloudId"`
	//Text the processed text with all applied transformations
	Text string `json:"text"`
}

func (r *Request) String() string {
	return fmt.Sprintf("Request{text=%s, uppercase=%s, reverse=%s}",
		r.Text, strconv.FormatBool(r.Uppercase), strconv.FormatBool(r.Reverse))
}

func (r *Response) String() string {
	return fmt.Sprintf("Response{requestId=%s,workerid=%s,cloudId=%s,Text=%s}",
		r.RequestId, r.WorkerId, r.CloudId, r.Text)
}

func (m *Message) String() string {
	return fmt.Sprintf("Message{request=%s,requestId=%s}",
		m.Request.String(), m.RequestId)
}
