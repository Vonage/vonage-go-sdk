package nexmo
type Links struct {
    Self Link `json:"self,omitempty"`
    Next Link `json:"next,omitempty"`
    Prev Link `json:"prev,omitempty"`
    First Link `json:"first,omitempty"`
    Last Link `json:"last,omitempty"`
}


type Link struct {
    Href string `json:"href,omitempty"`
}


type BasicInsightRequest struct {
    APIKey string `json:"api_key,omitempty"`
    APISecret string `json:"api_secret,omitempty"`
    Number string `json:"number,omitempty"`
    Country string `json:"country,omitempty"`
}

func (r *BasicInsightRequest) setApiCredentials(apiKey, apiSecret string) {
    r.APIKey = apiKey
    r.APISecret = apiSecret
}

type BasicInsightResponse struct {
    Status int64 `json:"status,omitempty"`
    StatusMessage string `json:"status_message,omitempty"`
    ErrorText string `json:"error_text,omitempty"`
    RequestId string `json:"request_id,omitempty"`
    InternationalFormatNumber string `json:"international_format_number,omitempty"`
    NationalFormatNumber string `json:"national_format_number,omitempty"`
    CountryCode string `json:"country_code,omitempty"`
    CountryCodeIso3 string `json:"country_code_iso3,omitempty"`
    CountryName string `json:"country_name,omitempty"`
    CountryPrefix string `json:"country_prefix,omitempty"`
}


type CallEndpoint struct {
}


type PhoneCallEndpoint struct {
    Type string `json:"type"`
    Number string `json:"number"`
    Dtmfanswer string `json:"dtmfAnswer,omitempty"`
}


type WebSocketCallEndpoint struct {
    Type string `json:"type"`
    URI string `json:"uri"`
    ContentType string `json:"content-type,omitempty"`
    Headers interface{} `json:"headers,omitempty"`
}


type SIPCallEndpoint struct {
    Type string `json:"type"`
    URI string `json:"uri"`
}


type SendSMSRequest struct {
    APIKey string `json:"api_key"`
    APISecret string `json:"api_secret"`
}

func (r *SendSMSRequest) setApiCredentials(apiKey, apiSecret string) {
    r.APIKey = apiKey
    r.APISecret = apiSecret
}

type SendSMSResponse struct {
    Status int64 `json:"status"`
    StatusMessage string `json:"status_message,omitempty"`
    ErrorText string `json:"error_text,omitempty"`
}


type CreateCallRequest struct {
    To []interface{} `json:"to"`
    From interface{} `json:"from"`
    AnswerURL []string `json:"answer_url"`
    AnswerMethod string `json:"answer_method,omitempty"`
    EventURL string `json:"event_url,omitempty"`
    EventMethod string `json:"event_method,omitempty"`
    MachineDetection string `json:"machine_detection,omitempty"`
    LengthTimer int64 `json:"length_timer,omitempty"`
    RingingTimer int64 `json:"ringing_timer,omitempty"`
}


type CreateCallResponse struct {
    UUID string `json:"uuid"`
    ConversationUUID string `json:"conversation_uuid"`
    Direction string `json:"direction"`
    Status string `json:"status"`
}


type SearchCallsRequest struct {
    Status string `url:"status,omitempty"`
    DateStart string `url:"date_start,omitempty"`
    DateEnd string `url:"date_end,omitempty"`
    PageSize int64 `url:"page_size,omitempty"`
    RecordIndex int64 `url:"record_index,omitempty"`
    Order string `url:"order,omitempty"`
    ConversationUUID string `url:"conversation_uuid,omitempty"`
}


type SearchCallsResponse struct {
    Count int64 `json:"count,omitempty"`
    PageSize int64 `json:"page_size,omitempty"`
    RecordIndex int64 `json:"record_index,omitempty"`
    Links Links `json:"_links,omitempty"`
    Embedded EmbeddedCalls `json:"_embedded,omitempty"`
}


type EmbeddedCalls struct {
    Calls []CallInfo `json:"calls,omitempty"`
}


type CallInfo struct {
    UUID string `json:"uuid,omitempty"`
    ConversationUUID string `json:"conversation_uuid,omitempty"`
    To interface{} `json:"to,omitempty"`
    From interface{} `json:"from,omitempty"`
    Status string `json:"status,omitempty"`
    Direction string `json:"direction,omitempty"`
    Rate string `json:"rate,omitempty"`
    Price string `json:"price,omitempty"`
    Duration string `json:"duration,omitempty"`
    Network string `json:"network,omitempty"`
    StartTime string `json:"start_time,omitempty"`
    EndTime string `json:"end_time,omitempty"`
}


type SimpleModifyCallRequest struct {
    Action string `json:"action,omitempty"`
}


type TransferCallRequest struct {
    Action string `json:"action,omitempty"`
    Destination TransferDestination `json:"destination,omitempty"`
}


type ModifyCallResponse struct {
    Message string `json:"message,omitempty"`
}


type TransferDestination struct {
    Type string `json:"type,omitempty"`
    URL []string `json:"url,omitempty"`
}


type CallErrorResponse struct {
    Type string `json:"type,omitempty"`
    ErrorTitle string `json:"error_title,omitempty"`
}

