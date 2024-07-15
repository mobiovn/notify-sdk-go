package main

import "gitlab.com/nguyenluongdev/notify-sdk-go/notifysdk"

func main() {
	sdk := notifysdk.GetMobioNotifySDK().Config("sale")

	sdk.SendMessageNotify("merchant_id_123", "sale_deal_assign_me", []string{"account_id_1", "account_id_2"}, map[string]interface{}{
		"deal_count": 5,
	})

}
