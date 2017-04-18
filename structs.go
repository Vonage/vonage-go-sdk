package nexmo

type BasicInsightRequest struct {
	APIKey    string `json:"api_key,omitempty"`
	APISecret string `json:"api_secret,omitempty"`
	Number    string `json:"number,omitempty"`
	Country   string `json:"country,omitempty"`
}

func (r *BasicInsightRequest) setApiCredentials(apiKey, apiSecret string) {
	r.APIKey = apiKey
	r.APISecret = apiSecret
}

type BasicInsightResponse struct {
	Status                    int64  `json:"status,omitempty"`
	StatusMessage             string `json:"status_message,omitempty"`
	ErrorText                 string `json:"error_text,omitempty"`
	RequestId                 string `json:"request_id,omitempty"`
	InternationalFormatNumber string `json:"international_format_number,omitempty"`
	NationalFormatNumber      string `json:"national_format_number,omitempty"`
	CountryCode               string `json:"country_code,omitempty"`
	CountryCodeIso3           string `json:"country_code_iso3,omitempty"`
	CountryName               string `json:"country_name,omitempty"`
	CountryPrefix             string `json:"country_prefix,omitempty"`
}

type CallEndpoint struct {
}

type PhoneCallEndpoint struct {
	Type       string `json:"type"`
	Number     string `json:"number"`
	Dtmfanswer string `json:"dtmfAnswer,omitempty"`
}

type WebSocketCallEndpoint struct {
	Type        string      `json:"type"`
	URI         string      `json:"uri"`
	ContentType string      `json:"content-type,omitempty"`
	Headers     interface{} `json:"headers,omitempty"`
}

type SIPCallEndpoint struct {
	Type string `json:"type"`
	URI  string `json:"uri"`
}

type SendSMSRequest struct {
	APIKey    string `json:"api_key"`
	APISecret string `json:"api_secret"`
}

func (r *SendSMSRequest) setApiCredentials(apiKey, apiSecret string) {
	r.APIKey = apiKey
	r.APISecret = apiSecret
}

type SendSMSResponse struct {
	Status        int64  `json:"status"`
	StatusMessage string `json:"status_message,omitempty"`
	ErrorText     string `json:"error_text,omitempty"`
}

type CreateCallRequest struct {
	To               []interface{} `json:"to"`
	From             interface{}   `json:"from"`
	AnswerURL        []string      `json:"answer_url"`
	AnswerMethod     string        `json:"answer_method,omitempty"`
	EventURL         string        `json:"event_url,omitempty"`
	EventMethod      string        `json:"event_method,omitempty"`
	MachineDetection string        `json:"machine_detection,omitempty"`
	LengthTimer      int64         `json:"length_timer,omitempty"`
	RingingTimer     int64         `json:"ringing_timer,omitempty"`
}

type CreateCallResponse struct {
	UUID             string `json:"uuid"`
	ConversationUUID string `json:"conversation_uuid"`
	Direction        string `json:"direction"`
	Status           string `json:"status"`
}

type ListCallsRequest struct {
}

type ListCallsResponse struct {
}

type GetCallRequest struct {
}

type GetCallResponse struct {
}

type ModifyCallRequest struct {
}

type ModifyCallResponse struct {
}

type CallErrorResponse struct {
	Type       string `json:"type,omitempty"`
	ErrorTitle string `json:"error_title,omitempty"`
}
