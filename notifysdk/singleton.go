package notifysdk

var instance *MobioNotifySDK

func GetMobioNotifySDK() *MobioNotifySDK {
	if instance == nil {
		instance = &MobioNotifySDK{}
	}
	return instance
}
