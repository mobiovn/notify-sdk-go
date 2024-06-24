package notifysdk

import (
	"errors"
	"fmt"
	"time"
)

type MobioNotifySDK struct {
	source string
}

func (sdk *MobioNotifySDK) Config(source string) *MobioNotifySDK {
	sdk.source = source
	return sdk
}

func (sdk *MobioNotifySDK) ValidateDataSendNotify(kwargs map[string]interface{}) error {
	if sdk.source == "" {
		return errors.New(fmt.Sprintf("[NOTIFY_SDK_ERR] %s is required", Param.Source))
	}

	fmt.Printf("[NOTIFY_SDK]: validate_data_send_notify: kwargs: %v\n", kwargs)
	for _, field := range Param.ListFieldRequired {
		if _, ok := kwargs[field]; !ok {
			return errors.New(fmt.Sprintf("[NOTIFY_SDK_ERR] %s is required", field))
		}
	}
	return nil
}

func getDataKwargs(dataKwargs map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for key, val := range dataKwargs {
		result[key] = val
	}
	return result
}

func (sdk *MobioNotifySDK) buildMessageSendNotify(messageType, merchantID, keyConfig string, accountIDs []string, kwargs map[string]interface{}) map[string]interface{} {
	message := map[string]interface{}{
		Param.MessageType: messageType,
		Param.Source:      sdk.source,
		Param.MerchantId:  merchantID,
		Param.KeyConfig:   keyConfig,
		Param.AccountIds:  accountIDs,
		Param.DataKwargs:  getDataKwargs(kwargs),
	}

	return message
}

func (sdk *MobioNotifySDK) SendMessageNotify(merchantID, keyConfig string, accountIDs []string, kwargs map[string]interface{}) {
	startTime := time.Now()
	messageType := MessageType.SendAll
	if err := sdk.ValidateDataSendNotify(map[string]interface{}{
		"merchant_id": merchantID,
		"key_config":  keyConfig,
		"account_ids": accountIDs,
	}); err != nil {
		panic(err)
	}

	dataSendProducer := sdk.buildMessageSendNotify(messageType, merchantID, keyConfig, accountIDs, kwargs)

	pushMessageToTopicReceive(dataSendProducer)
	fmt.Printf("[NOTIFY_SDK]: send_message_notify_sdk: time_process:: %v\n", time.Since(startTime))
}

func (sdk *MobioNotifySDK) SendMessageNotifySocket(merchantID, keyConfig string, accountIDs []string, kwargs map[string]interface{}) {
	startTime := time.Now()

	messageType := MessageType.SendSocket
	if err := sdk.ValidateDataSendNotify(map[string]interface{}{
		"merchant_id": merchantID,
		"key_config":  keyConfig,
		"account_ids": accountIDs,
	}); err != nil {
		panic(err)
	}

	dataSendProducer := sdk.buildMessageSendNotify(messageType, merchantID, keyConfig, accountIDs, kwargs)

	pushMessageToTopicReceive(dataSendProducer)
	fmt.Printf("[NOTIFY_SDK]: send_message_notify_socket: time_process:: %v\n", time.Since(startTime))
}

func (sdk *MobioNotifySDK) SendMessageNotifyEmail(merchantID, keyConfig string, accountIDs []string, kwargs map[string]interface{}) {
	startTime := time.Now()

	messageType := MessageType.SendEmail
	if err := sdk.ValidateDataSendNotify(map[string]interface{}{
		"merchant_id": merchantID,
		"key_config":  keyConfig,
		"account_ids": accountIDs,
	}); err != nil {
		panic(err)
	}

	dataSendProducer := sdk.buildMessageSendNotify(messageType, merchantID, keyConfig, accountIDs, kwargs)
	pushMessageToTopicReceive(dataSendProducer)
	fmt.Printf("[NOTIFY_SDK]: send_message_notify_email: time_process:: %v\n", time.Since(startTime))
}

func (sdk *MobioNotifySDK) SendMessageNotifyPushIDMobileApp(merchantID, keyConfig string, accountIDs []string, kwargs map[string]interface{}) {
	startTime := time.Now()

	messageType := MessageType.SendPushIdMobileApp
	if err := sdk.ValidateDataSendNotify(map[string]interface{}{
		"merchant_id": merchantID,
		"key_config":  keyConfig,
		"account_ids": accountIDs,
	}); err != nil {
		panic(err)
	}

	dataSendProducer := sdk.buildMessageSendNotify(messageType, merchantID, keyConfig, accountIDs, kwargs)
	pushMessageToTopicReceive(dataSendProducer)
	fmt.Printf("[NOTIFY_SDK]: send_message_notify_push_id_mobile_app: time_process:: %v\n", time.Since(startTime))
}
