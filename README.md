Mobio Notify SDK
=====================================================

Examples
========
```golang
package main

import "github.com/luongndcoder/notify-sdk-go/notifysdk"

func main() {
	sdk := notifysdk.GetMobioNotifySDK().Config("sale")

	sdk.SendMessageNotify("merchant_id_123", "sale_deal_assign_me", []string{"account_id_1", "account_id_2"}, map[string]interface{}{
		"deal_count": 5,
	})

}
```

Getting Started
===============
Supports Go 1.17+ and librdkafka 2.4.0+.

Using Go Modules
----------------

Import the `notifysdk` package from GitHub in your code:

```golang
import "github.com/luongndcoder/notify-sdk-go/notifysdk"
```

Install the client
------------------


Manual install:
```bash
go get -u github.com/luongndcoder/notify-sdk-go/notifysdk
```

Golang import:
```golang
import "github.com/luongndcoder/notify-sdk-go/notifysdk"
```
