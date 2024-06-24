package notifysdk

type ParamMessage struct {
	Source            string
	MerchantId        string
	MessageType       string
	KeyConfig         string
	AccountIds        string
	DataKwargs        string
	ListFieldRequired []string
}

type MessageTypeValue struct {
	SendAll             string
	SendSocket          string
	SendEmail           string
	SendPushIdMobileApp string
}

var Param = ParamMessage{
	Source:            "source",
	MerchantId:        "merchant_id",
	MessageType:       "message_type",
	KeyConfig:         "key_config",
	AccountIds:        "account_ids",
	DataKwargs:        "data_kwargs",
	ListFieldRequired: []string{"merchant_id", "key_config"},
}

var MessageType = MessageTypeValue{
	SendAll:             "send_all",
	SendSocket:          "send_socket",
	SendEmail:           "send_email",
	SendPushIdMobileApp: "send_push_id_mobile_app",
}
