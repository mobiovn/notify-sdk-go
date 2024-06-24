package notifysdk

import (
	"os"
)

type ConsumerTopic struct {
	TopicNotifySdkReceiveMessage string
}

type MobioEnvironment struct {
	KafkaBroker string
}

var (
	Topic = ConsumerTopic{
		TopicNotifySdkReceiveMessage: "nm-sdk-receive-message",
	}

	Env = MobioEnvironment{
		KafkaBroker: os.Getenv("KAFKA_BROKER"),
	}
)
