<p align="center">
    <img title="Flutterwave" height="200" src="https://flutterwave.com/images/logo/full.svg" width="50%"/>
</p>

# Rave Go Library

## Introduction
This is a Go wrapper around the [API](https://flutterwavedevelopers.readme.io/v2.0/reference) for [Rave by Flutterwave](https://rave.flutterwave.com).
#### Payment types implemented:
* Card Payments
* Bank Account Payments
* Preauth
* Refund
* Subaccount
* Transfer
* Payment Plan
* Subscription
* USSD Payments
## Installation
To install, run

```go get github.com/anjolabassey/Rave-go/rave```

Note: This is currently under active development
## Import Package
The base class for this package is 'Rave'. To use this class, add:

```go
 import (
 	"github.com/anjolabassey/Rave-go/rave"
 )
 ```

## Initialization

#### To instantiate in sandbox:
To use Rave, instantiate Rave with your public key. We recommend that you store your secret key in an environment variable named, ```RAVE_SECKEY```. However, you can also pass it in here alongside your public key. Instantiating Rave is therefore as simple as:


```go
var r = rave.Rave{
	false,
	"FLWPUBK-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-X",
	"FLWSECK-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-X",
}
```

**Note: If you store your secret key as an environment variable, just pass an empty string "" for the secret field as shown below**

```go
var r = rave.Rave{
	false,
	"FLWPUBK-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-X",
	"",
}
```

#### To instantiate in production:
To initialize in production, simply set the ```production``` flag to ```true```.

```go
var r = rave.Rave{
	true,
	"FLWPUBK-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-X",
	"FLWSECK-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-X",
}
```

# Rave Methods
This is the documentation for all of the components of Rave-go

## ```rave.Card{}```
This is used to facilitate card transactions via rave. ```rave.Card{}``` is of type ```struct``` and requires  ```rave.Rave``` as its only property.

Hence, in order to use it, you need to pass in an instance of ```rave.Rave``` . A sample is shown below

```go
    var card = rave.Card{
    	r,
    }
```
**Methods Included:**

* ```.ChargeCard```

* ```.ValidateCard```

* ```.VerifyCard```

* ```.TokenizedCharge```

### ```.ChargeCard(data CardChargeData) (error error, response map[string]interface{})```
This is called to charge a card. The payload should be of type ```rave.CardChargeData```. See below for  ```rave.CardChargeData``` definition

```go
type CardChargeData struct {
	Cardno               string         `json:"cardno"`
	Cvv                  string         `json:"cvv"`
	Expirymonth          string         `json:"expirymonth"`
	Expiryyear           string         `json:"expiryyear"`
	Pin                  string         `json:"pin"`
	Amount               float64        `json:"amount"`
	Currency             string         `json:"currency"`
	Country              string         `json:"country"`
	CustomerPhone        string         `json:"customer_phone"`
	Firstname            string         `json:"firstname"`
	Lastname             string         `json:"lastname"`
	Email                string         `json:"email"`
	Ip                   string         `json:"IP"`
	Txref		         string	        `json:"txRef"`
	RedirectUrl          string         `json:"redirect_url"`
	Subaccounts          types.Slice    `json:"subaccounts"`
	DeviceFingerprint    string         `json:"device_fingerprint"`
	Meta                 types.Slice    `json:"meta"`
	SuggestedAuth        string         `json:"suggested_auth"`
	BillingZip           string         `json:"billingzip"`
	BillingCity          string         `json:"billingcity"`
	BillingAddress       string         `json:"billingaddress"`
	BillingState         string         `json:"billingstate"`
	BillingCountry       string         `json:"billingcountry"`
	Chargetype		     string	        `json:"charge_type"`

}
```
A sample charge call is:

```go
    payload := rave.CardChargeData{

        Amount:100,
		Txref:"MC-11001993",
		Email:"test@test.com",
		CustomerPhone:"08123456789",
		Currency:"NGN",
		Cardno:"5399838383838381",
		Cvv:"470",
		Expirymonth:"10",
		Expiryyear:"22",
		Pin: "3310",
    }

    err, response := card.ChargeCard(payload)
    if err != nil {
        panic(err)
    }
    fmt.Println(response)
```
#### Sample Response

```sh
map[status:success message:V-COMP data:map[chargeResponseMessage:Please enter the OTP sent to your mobile number 080****** and email te**@rave**.com modalauditid:004855926f1352dcdfbb8f02b8a3376c paymentType:card paymentPlan:<nil> paymentPage:<nil> deletedAt:<nil> orderRef:URF_1543836853430_975035 device_fingerprint:N/A cycle:one-time narration:CARD Transaction  acctvalrespmsg:<nil> AccountId:7813 redirectUrl:N/A settlement_token:<nil> charged_amount:300 IP:::ffff:10.29.81.254 acctvalrespcode:<nil> authurl:N/A is_live:0 vbvrespmessage:Approved. Successful charge_type:normal customercandosubsequentnoauth:false amount:300 appfee:4.2 merchantfee:0 chargeResponseCode:02 currency:NGN id:344583 status:success-pending-validation paymentId:1446 vbvrespcode:00 fraud_status:ok createdAt:2018-12-03T11:34:13.000Z txRef:MC-11001993 flwRef:FLW-MOCK-adaf803863b1ad2de90c506a3c6095db merchantbearsfee:1 raveRef:RV3154383685229161954E6069 authModelUsed:PIN customerId:66735 updatedAt:2018-12-03T11:34:14.000Z customer:map[id:66735 phone:<nil> customertoken:<nil> createdAt:2018-12-03T11:34:12.000Z updatedAt:2018-12-03T11:34:12.000Z fullName:Anonymous customer email:kwaku@gmail.com deletedAt:<nil> AccountId:7813]]]
```

### ```.ValidateCard(data CardValidateData) (error error, response map[string]interface{})```
This is called to validate a card charge. The payload should be of type ```rave.CardValidateData```. See below for  ```rave.CardValidateData``` definition.

After a successful charge, most times you will be asked to verify with OTP. To check if this is required, check the `chargeResponseMessage` key in the response of the charge call.

In the case that an authUrl is returned from your charge call, you may skip the validation step and simply pass your authUrl to the end-user.

authUrl = response["data"].(map[string]interface{})["authUrl"].(string)
```go
type CardValidateData struct {
	Reference	   string	      `json:"transaction_reference"`
	Otp		       string	      `json:"otp"`
	PublicKey      string         `json:"PBFPubKey"``
}
```
The Reference is the `flwRef` gotten from the response of the ChargeCard function. See an example below
ref := response["data"].(map[string]interface{})["flwRef"].(string)

A sample validate call is:

```go
    payload := rave.CardValidateData{
        Otp:"12345",
		Reference: ref,
    }

    err, response := card.ValidateCard(payload)
    if err != nil {
        panic(err)
    }
    fmt.Println(response)
```
#### Sample Response

```sh
map[status:success message:Charge Complete data:map[tx:map[merchantfee:0 authModelUsed:PIN modalauditid:7a32839d67033b0f52208ae7e1c6c451 deletedAt:<nil> id:344594 txRef:MC-11001993 redirectUrl:N/A device_fingerprint:N/A createdAt:2018-12-03T11:40:11.000Z authurl:N/A acctvalrespmsg:<nil> paymentPage:<nil> fraud_status:ok merchantbearsfee:1 paymentId:1446 customer:map[phone:<nil> email:kwaku@gmail.com createdAt:2018-12-03T11:40:10.000Z updatedAt:2018-12-03T11:40:10.000Z deletedAt:<nil> id:66738 fullName:Anonymous customer customertoken:<nil> AccountId:7813] status:successful is_live:0 updatedAt:2018-12-03T11:40:13.000Z AccountId:7813 orderRef:URF_1543837211195_1508135 settlement_token:<nil> charged_amount:2200 chargeResponseMessage:Please enter the OTP sent to your mobile number 080****** and email te**@rave**.com charge_type:normal customerId:66738 flwRef:FLW-MOCK-f5b78794403ce62f0572f1ea8250636a appfee:30.8 narration:CARD Transaction  vbvrespcode:00 paymentType:card chargeToken:map[user_token:b4f74 embed_token:flw-t0-4aaf3ce75844d30eb5323d83d25b865f-m03k] cycle:one-time chargeResponseCode:00 currency:NGN IP:::ffff:10.29.86.227 amount:2200 vbvrespmessage:successful acctvalrespcode:<nil> paymentPlan:<nil> raveRef:RV31543837210006AFAE01EA80] airtime_flag:<nil> data:map[responsemessage:successful responsecode:00 responsetoken:mocktoken]]]
```

### ```.VerifyCard(data CardVerifyData) (error error, response map[string]interface{})```
This is called to validate a card charge. The payload should be of type ```rave.CardVerifyData```. See below for  ```rave.CardVerifyData``` definition
```go
type CardVerifyData struct {
	Reference	   string	      `json:"txref"`
	Amount	       float64	      `json:"amount"`
	Currency       string         `json:"currency"`
	SecretKey      string         `json:"SECKEY"`
}
```
The Reference is the `txRef` which is gotten from the response of the ChargeCard function. See example below
txref := response["data"].(map[string]interface{})["txRef"].(string)

A sample verify call is:

```go
    payload := rave.CardVerifyData{
		Reference: txref,
        Amount: 100,
        Currency: "NGN",

    }

    err, response := card.VerifyCard(payload)
    if err != nil {
        panic(err)
    }
    fmt.Println(response)
```
#### Sample Response

```sh
map[data:map[paymentplan:<nil> raveref:RV31543838107864EEE0EF8E70 card:map[brand:GUARANTY TRUST BANK DEBITSTANDARD card_tokens:[map[embedtoken:flw-t1nf-53b4caf0e406a0924bf8b11bf5113414-m03k shortcode:d44cb expiry:9999999999999]] type:MASTERCARD life_time_token:flw-t1nf-53b4caf0e406a0924bf8b11bf5113414-m03k expirymonth:10 expiryyear:22 cardBIN:539983 last4digits:8381] meta:[] createddayname:MONDAY createdhour:11 custemailprovider:GMAIL acctcode:<nil> acctmessage:<nil> chargetype:normal createdday:1 amountsettledforthistransaction:493 devicefingerprint:N/A chargecode:00 vbvcode:00 custemail:kwaku@gmail.com acctcontactperson:Anjolaoluwa Bassey currency:NGN createdminute:55 custname:Anonymous customer paymentid:1446 acctalias:<nil> createddayispublicholiday:0 paymentpage:<nil> createdmonthname:DECEMBER created:2018-12-03T11:55:09.000Z acctcountry:NG chargedamount:500 merchantbearsfee:1 authurl:N/A amount:500 createdmonth:11 createdyear:2018 custphone:<nil> txid:344611 createdweek:49 custcreated:2018-12-03T11:55:08.000Z createdquarter:4 custnetworkprovider:N/A chargemessage:Please enter the OTP sent
to your mobile number 080****** and email te**@rave**.com ip:::ffff:10.11.233.111 fraudstatus:ok paymenttype:card accountid:7813 merchantfee:0 authmodel:PIN narration:CARD Transaction  acctbearsfeeattransactiontime:1 cycle:one-time status:successful acctbusinessname:Anjola's enterprise createdpmam:am customerid:66748 orderref:URF_1543838109028_2039935 txref:MC-11001993 appfee:7 vbvmessage:successful createdyearisleap:false acctvpcmerchant:N/A acctisliveapproved:0 flwref:FLW-MOCK-d4a1e94548eb0e929dd80dc265c9bcc3 acctparent:1] status:success message:Tx Fetched]
```
### ```.TokenizedCharge(data SaveCardChargeData) (error error, response map[string]interface{})```
This is called to charge a saved card using a token(which can be gotten in the [verify payment response](https://developer.flutterwave.com/v2.0/reference#save-a-card)). The payload should be of type ```rave.TokenizedChargeData```. See below for  ```rave.TokenizedChargeData``` definition
```go
type TokenizedChargeData struct {
    SecretKey      string         `json:"SECKEY"`
	Currency       string         `json:"currency"`
	Token          string         `json:"token"`
	Country        string         `json:"country"`
	Amount	       float64	      `json:"amount"`
	Email          string         `json:"email"`
	Firstname      string         `json:"firstname"`
	Lastname       string         `json:"lastname"`
	Ip             string         `json:"IP"`
	Txref		   string	      `json:"txRef"`
}

```
A sample initiate call is:

```go
    payload := rave.TokenizedChargeData{
        Token: "flw-t1nf-2f00ba4c24b27cbb39e7907c6b72d413-m03k",
		Currency:"NGN",
	 	Country:"NG",
	 	Amount:100,
	 	Email:"test@test.com",
	 	Txref:"MC-0123456789",
    }

    err, response := card.TokenizedCharge(payload)
    if err != nil {
        panic(err)
    }
    fmt.Println(response)
```
#### Sample Response

```sh
map[status:success message:Charge success data:map[cycle:one-time appfee:1.4 raveRef:<nil> modalauditid:cb739edfc461cd31656c85e050bec0d1 vbvrespcode:00 acctvalrespmsg:<nil> createdAt:2018-12-03T10:15:40.000Z txRef:MC-0123456789 charged_amount:100 authModelUsed:noauth acctvalrespcode:<nil> fraud_status:ok updatedAt:2018-12-03T10:15:41.000Z customerId:66699 id:344437 chargeResponseMessage:Approved vbvrespmessage:Approved charge_type:normal deletedAt:<nil> chargeToken:map[user_token:ad805 embed_token:flw-t0-54f282018b0a0709d99bc7ccf7c39cca-m03k] flwRef:FLW-M03K-a8ac6680a157f0288236e4823f6fc64c redirectUrl:http://127.0.0 chargeResponseCode:00 narration:Anjola's enterprise authurl:N/A paymentType:card AccountId:7813 settlement_token:<nil> amount:100 paymentId:1446 device_fingerprint:N/A IP:::127.0.0.1 status:successful is_live:0 orderRef:URF_BEE66669988CB9DDE6AE merchantbearsfee:1 paymentPlan:<nil> merchantfee:0 currency:NGN paymentPage:<nil> customer:map[deletedAt:<nil> id:66699 email:kwaku@gmail.com createdAt:2018-12-03T10:05:14.000Z updatedAt:2018-12-03T10:05:14.000Z AccountId:7813 phone:<nil> fullName:Anonymous customer customertoken:<nil>]]]
```
## ```rave.Account{}```
This is used to facilitate bank account transactions via rave. ```rave.Account{}``` is of type ```struct``` and requires  ```rave.Rave``` as its only property.

Hence, in order to use it, you need to pass in an instance of ```rave.Rave``` . A sample is shown below

```go
    var account = rave.Account{
    	r,
    }
```
**Methods Included:**

* ```.ChargeAccount```

* ```.ValidateAccount```

* ```.VerifyAccount```

### ```.ChargeAccount(data AccountChargeData) (error error, response map[string]interface{})```
This is called to charge a bank account. The payload should be of type ```rave.AccountChargeData```. See below for  ```rave.AccountChargeData``` definition

```go
type AccountChargeData struct {
	Cardno                string         `json:"cardno"`
	Cvv                   string         `json:"cvv"`
	Accountbank           string         `json:"accountbank"`
	Accountnumber         string         `json:"accountnumber"`
	Paymenttype           string         `json:"payment_type"`
	Amount                float64        `json:"amount"`
	Currency              string         `json:"currency"`
	Country               string         `json:"country"`
	Bvn                   string         `json:"bvn"`
	Passcode              string         `json:"passcode"`
	CustomerPhone         string         `json:"customer_phone"`
	Firstname             string         `json:"firstname"`
	Lastname              string         `json:"lastname"`
	Email                 string         `json:"email"`
	IP                    string         `json:"IP"`
	Txref		          string	     `json:"txRef"`
	SuggestedAuth         string         `json:"suggested_auth"`
	RedirectUrl           string         `json:"redirect_url"`
	Subaccounts           types.Slice    `json:"subaccounts"`
	DeviceFingerprint     string         `json:"device_fingerprint"`
	Meta                  types.Slice    `json:"meta"`

}
```
A sample charge call is:

```go
    payload := rave.CardChargeData{
        Accountbank: "044", 
		Accountnumber: "0690000031", 
		Amount: 100, 
		Country: "NG", 
		Currency: "NGN",
		Email: "test@test.com", 
		CustomerPhone: "08123456789", 
		Firstname: "Seun", 
		Lastname: "Alade", 
		Paymenttype: "account", 
		IP: "103.238.105.185", 
		Txref: "MXX-ASC-4578",
    }

    err, response := account.ChargeAccount(payload)
    if err != nil {
        panic(err)
    }
    fmt.Println(response)
```
#### Sample Response

```sh
map[data:map[txRef:MXX-ASC-4578 merchantfee:0 vbvrespcode:N/A updatedAt:2018-12-03T12:05:33.000Z deletedAt:<nil> AccountId:7813 validateInstructions:map[valparams:[OTP] instruction:Please validate with the OTP sent to your mobile or email] orderRef:URF_1543838731024_92535 settlement_token:<nil> charged_amount:100 paymentId:2 customerId:66752 device_fingerprint:N/A amount:100 raveRef:RV315438387304222280DDB025 IP:::ffff:10.29.81.94 narration:Anjola's enterprise authurl:NO-URL acctvalrespmsg:<nil> merchantbearsfee:1 modalauditid:18bcf29fda26782c3906b6246c15c074 acctvalrespcode:<nil> paymentPlan:<nil> paymentType:account paymentPage:<nil> createdAt:2018-12-03T12:05:31.000Z validateInstruction:Please dial *901*4*1# to get your OTP. Enter the OTP gotten in the field below flwRef:ACHG-1543838731792 chargeResponseCode:02 authModelUsed:AUTH currency:NGN is_live:0 redirectUrl:N/A chargeResponseMessage:Pending OTP validation status:success-pending-validation vbvrespmessage:N/A charge_type:normal customer:map[id:66752 fullName:Anjola Bassey createdAt:2018-12-03T12:05:30.000Z AccountId:7813 phone:<nil> customertoken:<nil> email:ajb@yahoo.com updatedAt:2018-12-03T12:05:30.000Z deletedAt:<nil>] id:344628 cycle:one-time appfee:1.4 fraud_status:ok] status:success message:V-COMP]
```

### ```.ValidateAccount(data AccountValidateData) (error error, response map[string]interface{})```
This is called to validate an account charge. The payload should be of type ```rave.AccountValidateData```. See below for  ```rave.AccountValidateData``` definition.

After a successful charge, most times you will be asked to verify with OTP. Check the `validateInstructions` key in the response of the charge call, This object contains the instructions you are meant to show to the customer so they know the next step to take.

In the case that an authUrl is returned from your charge call, you may skip the validation step and simply pass your authUrl to the end-user.

authUrl = response["data"].(map[string]interface{})["authUrl"].(string)
```go
type AccountValidateData struct {
	Reference	   string	      `json:"transaction_reference"`
	Otp		       string	      `json:"otp"`
	PublicKey      string         `json:"PBFPubKey"``
}
```
The Reference is the `flwRef` gotten from the response of the ChargeCard function. See an example below
ref := response["data"].(map[string]interface{})["flwRef"].(string)

A sample validate call is:

```go
    payload := rave.AccountValidateData{
        Otp:"12345",
		Reference: ref,
    }

    err, response := card.ValidateAccount(payload)
    if err != nil {
        panic(err)
    }
    fmt.Println(response)
```
#### Sample Response

```sh
map[status:success message:Charge Complete data:map[createdAt:2018-12-03T12:23:20.000Z amount:100 currency:NGN vbvrespcode:N/A paymentPage:<nil> fraud_status:ok id:344650 device_fingerprint:N/A settlement_token:<nil> acctvalrespmsg:Approved Or Completed Successfully redirectUrl:N/A charge_type:normal updatedAt:2018-12-03T12:23:29.000Z AccountId:7813 txRef:MXX-ASC-4578 merchantfee:0 chargeResponseCode:00 paymentId:2 is_live:0 status:successful paymentType:account customer:map[updatedAt:2018-12-03T12:23:20.000Z phone:<nil> fullName:Anjola Bassey customertoken:<nil> email:ajb@yahoo.com id:66762 createdAt:2018-12-03T12:23:20.000Z deletedAt:<nil> AccountId:7813] orderRef:URF_1543839800474_8711935 cycle:one-time merchantbearsfee:1 raveRef:RV3154383979993709371F69C2 narration:Anjola's enterprise airtime_flag:<nil> flwRef:ACHG-1543839801193 charged_amount:100 authModelUsed:AUTH deletedAt:<nil> appfee:1.4 IP:::ffff:10.137.215.140 modalauditid:510c3e34c8d93a067697919ee20cb571 customerId:66762 chargeResponseMessage:Pending OTP validation vbvrespmessage:N/A authurl:NO-URL acctvalrespcode:00 paymentPlan:<nil>]]
```

### ```.VerifyAccount(data AccountVerifyData) (error error, response map[string]interface{})```
This is called to validate an account charge. The payload should be of type ```rave.AccountVerifyData```. See below for  ```rave.AccountVerifyData``` definition
```go
type AccountVerifyData struct {
	Reference	   string	      `json:"txref"`
	Amount	       float64	      `json:"amount"`
	Currency       string         `json:"currency"`
	SecretKey      string         `json:"SECKEY"`
}
```
The Reference is the `txRef` which is gotten from the response of the ChargeCard function. See example below
// txref := txref := response["data"].(map[string]interface{})["txRef"].(string)

A sample verify call is:

```go
    payload := rave.AccountVerifyData{
		Reference: txref,
        Amount: 100,
        Currency: "NGN",

    }

    err, response := card.VerifyAccount(payload)
    if err != nil {
        panic(err)
    }
    fmt.Println(response)
```
#### Sample Response

```sh
map[data:map[acctbusinessname:Anjola's enterprise raveref:RV315438403686292720093A92 narration:Anjola's enterprise status:successful createdweek:49 accountid:7813 custname:Anjola Bassey custemail:ajb@yahoo.com amountsettledforthistransaction:98.6 amount:100 merchantfee:0 acctmessage:Approved Or Completed Successfully custphone:<nil> flwref:ACHG-1543840370084 createdminute:32 paymenttype:account chargetype:normal createdpmam:pm custnetworkprovider:N/A currency:NGN chargedamount:100 authmodel:AUTH vbvcode:N/A paymentplan:<nil> acctcontactperson:Anjolaoluwa Bassey acctcountry:NG txref:MXX-ASC-4578 appfee:1.4 createddayname:MONDAY createdyearisleap:false created:2018-12-03T12:32:49.000Z custemailprovider:YAHOO createdday:1 customerid:66771 orderref:URF_1543840369244_4970735 chargemessage:Pending OTP validation paymentid:2 createdyear:2018 createdhour:12 acctcode:00 acctparent:1 acctvpcmerchant:N/A acctbearsfeeattransactiontime:1 meta:[] account:map[id:2 account_number:0690000031 account_bank:044 last_name:NO-LNAME updatedAt:2018-12-03T12:33:00.000Z deletedAt:<nil> account_token:map[token:flw-t002ab96cd885af7b8-k3n-mock] first_name:NO-NAME account_is_blacklisted:0 createdAt:2016-12-31T04:09:24.000Z] authurl:NO-URL createdmonthname:DECEMBER createdquarter:4 custcreated:2018-12-03T12:32:49.000Z acctisliveapproved:0 txid:344663 merchantbearsfee:1 ip:::ffff:10.63.159.3 createddayispublicholiday:0 cycle:one-time vbvmessage:N/A acctalias:<nil> paymentpage:<nil> devicefingerprint:N/A chargecode:00 fraudstatus:ok createdmonth:11] status:success message:Tx Fetched]
```

## ```rave.Refund{}```
This allows you initiate refunds for Successful transactions via rave. ```rave.Refund{}``` is of type ```struct``` and requires  ```rave.Rave``` as its only property.

Hence, in order to use it, you need to pass in an instance of ```rave.Rave``` . A sample is shown below

```go
    var transfer = rave.Refund{
    	r,
    }
```

**Methods Included:**

* ```.RefundTransaction```

### ```.RefundTransaction(data RefundData) (error error, response map[string]interface{})```
This is called to initiate refunds for Successful transaction. The payload should be of type ```rave.RefundData```. See below for  ```rave.RefundData``` definition

```go
type RefundData struct {
	Ref		       string	      `json:"ref"`
	Amount         int            `json:"amount"`
	SecretKey      string         `json:"seckey"`
}
```

A sample refund call is:

```go
    payload := rave.RefundData{
        Ref: "FLW-MOCK-476a260e67df43988a2ffeddf8e02cc2",
		Amount: 100,
    }

    err, response := refund.RefundTransaction(payload)
    if err != nil {
        panic(err)
    }
    fmt.Println(response)
```

#### Sample Response

```

```

## ```rave.Preauth{}```
This is used to facilitate preauthorised card transactions via rave. ```rave.Preauth{}``` is of type ```struct``` and requires  ```rave.Rave``` as its only property. This inherits the Card class so any task you can do on Card, you can do with preauth.

Hence, in order to use it, you need to pass in an instance of ```rave.Rave``` . A sample is shown below

```go
    var card = rave.Preauth{
    	r,
		rave.Card{r,},
    }
```
**Methods Included:**

* ```.ChargePreauth```

* ```.VerifyPreauth```

* ```.CapturePreauth```

* ```.RefundOrVoidPreauth```

### ```.ChargePreauth(data TokenizedChargeData) (error error, response map[string]interface{})```
This is called to preauthorize a card. Once you have the token saved from an initial charge on the card using `card.ChargeCard(payload)` you can use the token which looks like this `flw-t1nf-5b0f12d565cd961f73c51370b1340f1f-m03k` to perform preauth charges. The payload should be of type ```rave.TokenizedChargeData```. 

A sample call is:

```go
    payload := rave.TokenizedChargeData{
        Token: "flw-t1nf-4f9aef213b795e694d41407f8abddb8e-m03k",
		Currency:"NGN",
		Country:"NG",
		Amount:200,
		Email:"test@test.com",
		Txref:"MC-0123456789",
    }

    err, response := card.ChargePreauth(payload)
    if err != nil {
        panic(err)
    }
    fmt.Println(response)
```
#### Sample Response

```sh
map[status:success message:Charge success data:map[amount:100 currency:NGN vbvrespmessage:Approved deletedAt:<nil> customerId:66699 updatedAt:2018-12-03T10:22:05.000Z settlement_token:<nil> cycle:one-time chargeResponseMessage:Approved IP:::127.0.0.1 paymentPage:<nil> paymentId:1446 txRef:MC-0123456789 device_fingerprint:N/A raveRef:<nil> charged_amount:100 appfee:<nil> authModelUsed:noauth is_live:0 createdAt:2018-12-03T10:22:05.000Z AccountId:7813 chargeToken:map[user_token:c1db3 embed_token:flw-t0-1a7ecd6245544abe3f3d1427f73d4ae2-m03k] orderRef:<nil> merchantbearsfee:0 narration:TOKEN CHARGE authurl:N/A vbvrespcode:00 fraud_status:ok id:344452 flwRef:FLW-PREAUTH-M03K-e2d5f6523a2c5d1f9e9624a0bea9c578 chargeResponseCode:00 acctvalrespcode:<nil> paymentPlan:<nil> merchantfee:0 status:pending-capture acctvalrespmsg:<nil> paymentType:card charge_type:preauth customer:map[email:kwaku@gmail.com deletedAt:<nil> AccountId:7813 id:66699 phone:<nil> fullName:Anonymous customer customertoken:<nil> createdAt:2018-12-03T10:05:14.000Z updatedAt:2018-12-03T10:05:14.000Z] redirectUrl:http://127.0.0 modalauditid:e24c03499823307af15f3bb6915e02a2]]
```
## ```.VerifyPreauth(data CardVerifyData) (error error, response map[string]interface{})```
This is called to validate a Preauth card charge. See `rave.VerifyCard` above

#### Sample Response

```sh
map[status:success message:Charge success data:map[amount:100 currency:NGN vbvrespmessage:Approved deletedAt:<nil> customerId:66699 updatedAt:2018-12-03T10:22:05.000Z settlement_token:<nil> cycle:one-time chargeResponseMessage:Approved IP:::127.0.0.1 paymentPage:<nil> paymentId:1446 txRef:MC-0123456789 device_fingerprint:N/A raveRef:<nil> charged_amount:100 appfee:<nil> authModelUsed:noauth is_live:0 createdAt:2018-12-03T10:22:05.000Z AccountId:7813 chargeToken:map[user_token:c1db3 embed_token:flw-t0-1a7ecd6245544abe3f3d1427f73d4ae2-m03k] orderRef:<nil> merchantbearsfee:0 narration:TOKEN CHARGE authurl:N/A vbvrespcode:00 fraud_status:ok id:344452 flwRef:FLW-PREAUTH-M03K-e2d5f6523a2c5d1f9e9624a0bea9c578 chargeResponseCode:00 acctvalrespcode:<nil> paymentPlan:<nil> merchantfee:0 status:pending-capture acctvalrespmsg:<nil> paymentType:card charge_type:preauth customer:map[email:kwaku@gmail.com deletedAt:<nil> AccountId:7813 id:66699 phone:<nil> fullName:Anonymous customer customertoken:<nil> createdAt:2018-12-03T10:05:14.000Z updatedAt:2018-12-03T10:05:14.000Z] redirectUrl:http://127.0.0 modalauditid:e24c03499823307af15f3bb6915e02a2]]
```

### ```.CapturePreauth(data PreauthCaptureData) (error error, response map[string]interface{})```
This is called to preauthorize a card. The payload should be of type ```rave.PreauthCaptureData```. 
The Flwref is the `flwRef` gotten from the response of the CapturePreauth function. See an example below
ref := response["data"].(map[string]interface{})["flwRef"].(string)

A sample call is:

```go
    payload := rave.PreauthCaptureData{
        Amount:100,
        Flwref: ref

    }

    err, response := card.CapturePreauth(payload)
    if err != nil {
        panic(err)
    }
    fmt.Println(response)
```
#### Sample Response

```sh
map[status:success message:Capture complete data:map[device_fingerprint:N/A vbvrespcode:00 chargeResponseCode:00 chargeResponseMessage:Approved paymentPlan:<nil> fraud_status:ok txRef:MC-0123456789 redirectUrl:http://127.0.0 settlement_token:<nil> merchantbearsfee:0 narration:TOKEN CHARGE authurl:N/A id:344479 flwRef:FLW-PREAUTH-M03K-f08dff19ab02f20fd0ced1b00e81b5a2 paymentPage:<nil> appfee:2.8 raveRef:<nil> status:successful is_live:0 deletedAt:<nil> customerId:66699 cycle:one-time amount:200 currency:NGN vbvrespmessage:Approved acctvalrespmsg:CAPTURE REFERENCE acctvalrespcode:FLWPREAUTH-M03K-CP-1543833553346 updatedAt:2018-12-03T10:39:13.000Z AccountId:7813 customer:map[id:66699 phone:<nil> fullName:Anonymous customer email:kwaku@gmail.com createdAt:2018-12-03T10:05:14.000Z deletedAt:<nil> customertoken:<nil> updatedAt:2018-12-03T10:05:14.000Z AccountId:7813] orderRef:<nil> authModelUsed:noauth IP:::127.0.0.1 modalauditid:132b27fe184d7b18d998df74be497bfc charge_type:preauth charged_amount:202.8 merchantfee:0 paymentType:card paymentId:1446 createdAt:2018-12-03T10:28:33.000Z]]
```
### ```.RefundOrVoidPreauth(data PreauthRefundData) (error error, response map[string]interface{})```
This is called to refund or void a card. This is the action to be taken i.e. refund or void. The payload should be of type ```rave.PreauthRefundData```. See below for  ```rave.PreauthRefundData``` definition

```go

type PreauthRefundData struct {
	Flwref	       string	          `json:"ref"`
	Action	       string	          `json:"action"`
	SecretKey      string             `json:"SECKEY"`
}
```
The Flwref is the `flwRef` gotten from the response of the ChargePreauth function. See an example below
ref := response["data"].(map[string]interface{})["flwRef"].(string)

A sample call is:

```go
    payload := rave.PreauthCaptureData{
        Action: "refund" <!-- or "void" -->
        Flwref: ref

    }

    err, response := card.RefundOrVoidPreauth(payload)
    if err != nil {
        panic(err)
    }
    fmt.Println(response)
```
#### Sample Response

```sh
map[status:success message:Capture complete data:map[device_fingerprint:N/A vbvrespcode:00 chargeResponseCode:00 chargeResponseMessage:Approved paymentPlan:<nil> fraud_status:ok txRef:MC-0123456789 redirectUrl:http://127.0.0 settlement_token:<nil> merchantbearsfee:0 narration:TOKEN CHARGE authurl:N/A id:344479 flwRef:FLW-PREAUTH-M03K-f08dff19ab02f20fd0ced1b00e81b5a2 paymentPage:<nil> appfee:2.8 raveRef:<nil> status:successful is_live:0 deletedAt:<nil> customerId:66699 cycle:one-time amount:200 currency:NGN vbvrespmessage:Approved acctvalrespmsg:CAPTURE REFERENCE acctvalrespcode:FLWPREAUTH-M03K-CP-1543833553346 updatedAt:2018-12-03T10:39:13.000Z AccountId:7813 customer:map[id:66699 phone:<nil> fullName:Anonymous customer email:kwaku@gmail.com createdAt:2018-12-03T10:05:14.000Z deletedAt:<nil> customertoken:<nil> updatedAt:2018-12-03T10:05:14.000Z AccountId:7813] orderRef:<nil> authModelUsed:noauth IP:::127.0.0.1 modalauditid:132b27fe184d7b18d998df74be497bfc charge_type:preauth charged_amount:202.8 merchantfee:0 paymentType:card paymentId:1446 createdAt:2018-12-03T10:28:33.000Z]]
```



## ```rave.Transfer{}```
This is used to facilitate transfers via rave. ```rave.Transfer{}``` is of type ```struct``` and requires  ```rave.Rave``` as its only property.

Hence, in order to use it, you need to pass in an instance of ```rave.Rave``` . A sample is shown below

```go
    var transfer = rave.Transfer{
    	r,
    }
```

**Methods Included:**

* ```.InitiateSingleTransfer```

* ```.InitiateBulkTransfer```

* ```.FetchTransfer```

* ```.FetchAllTransfers```

* ```.GetTransferFee```

* ```.GetRaveBalance```

* ```.GetBulkTransferStatus```

### ```.InitiateSingleTransfer(payload SinglePaymentData) (error error, response map[string]interface{})```
This is called to initiate a sole transfer. The payload should be of type ```rave.SinglePaymentData```. See below for  ```rave.SinglePaymentData``` definition

```go
type SinglePaymentData struct {
	SecKey          string      `json:"seckey"`
	AccountBank     string      `json:"account_bank"`
	AccountNumber   string      `json:"account_number"`
	Amount          int         `json:"amount"`
	Narration       string      `json:"narration"`
	Currency        string      `json:"currency"`
	Reference       string      `json:"reference"`
	Meta            types.Slice `json:"meta"`
	BeneficiaryName string      `json:"beneficiary_name"`
}
```

A sample initiate call is:

```go
    payload := rave.SinglePaymentData{
        AccountBank: "044",
        AccountNumber: "0690000044",
        Amount:        500,
        SecKey:        r.GetSecretKey(),
        Narration:     "Test Transfer",
        Currency:      "NGN",
        Reference:     time.Now().String(),
    }

    err, response := transfer.InitiateSingleTransfer(payload)
    if err != nil {
        panic(err)
    }
    fmt.Println(response)
```

#### Sample Response

```sh
map[status:success message:TRANSFER-CREATED data:map[id:3603 account_number:0690000044 bank_code:044 date_created:2018-11-27T12:15:37.000Z amount:500 currency:NGN meta:map[] is_approved:1 bank_name:ACCESS BANK NIGERIA fullname:Mercedes Daniel status:NEW reference:2018-11-27 13:15:36.5762772 +0100 WAT m=+0.013132801 narration:Test Transfer requires_approval:0 fee:45 complete_message:]]
```

### ```.InitiateBulkTransfer(payload BulkPaymentData) (error error, response map[string]interface{})```

This is called to initiate a bulk transfer. The payload should be of type ```rave.BulkPaymentData```. See below for  ```rave.BulkPaymentData``` definition

```go
type BulkPaymentData struct {
    SecKey   string              `json:"seckey"`
    Title    string              `json:"title"`
    BulkData []map[string]string `json:"bulk_data"`
}
```

A sample initiate call is:

```go
    payloads := rave.BulkPaymentData{
        SecKey: "FLWSECK-0b1d6669cf375a6208db541a1d59adbb-X",
        Title:  "May Staff Salary",
        BulkData: []map[string]string{
            {
                "Bank":           "044",
                "Account Number": "0690000032",
                "Amount":         "500",
                "Currency":       "NGN",
                "Narration":      "Bulk transfer 1",
                "reference":      time.Now().String(),
            },
            {
                "Bank":           "044",
                "Account Number": "0690000034",
                "Amount":         "500",
                "Currency":       "NGN",
                "Narration":      "Bulk transfer 2",
                "reference":      time.Now().String(),
            },
        },
    }


    err, response := transfer.InitiateBulkTransfer(payload)
    if err != nil {
        panic(err)
    }
    fmt.Println(response)
```

#### Sample Response

```sh
map[status:success message:BULK-TRANSFER-CREATED data:map[id:683 date_created:2018-11-27T12:57:59.000Z approver:N/A]]
```


### ```.FetchTransfer(reference string) (error error, response map[string]interface{})```

This allows you retrieve a single transfer. The reference should be of type ```string```You may or may not pass in a transfer ```reference```. If you do not pass in a reference, all transfers that have been processed will be returned.

A sample fetch call is:

```go
reference := "kkkkkkkkkkkkk"
err, response := transfer.FetchTransfer(reference)
if err != nil {
 panic(err)
}
fmt.Println(response)
```

#### Sample Response

This call returns a dictionary. A sample response is:

 ```sh
 map[status:success message:QUERIED-TRANSFERS data:map[page_info:map[total:1 current_page:1 total_pages:1] transfers:[map[requires_approval:0 debit_currency:<nil> reference:kkkkkkkkkkkkk complete_message:DISBURSE FAILED: undefined is_approved:1 id:3563 account_number:0690000044 bank_code:044 currency:NGN fee:45 meta:map[] approver:<nil> bank_name:ACCESS BANK NIGERIA narration:New transfer fullname:Mercedes Daniel date_created:2018-11-26T21:38:31.000Z amount:500 status:FAILED]]]]
 ```

 ### ```.FetchTransfer(reference string) (error error, response map[string]interface{})```

 This allows you retrieve a single transfer. The reference should be of type ```string```You may or may not pass in a transfer ```reference```. If you do not pass in a reference, all transfers that have been processed will be returned.

 A sample fetch call is:

 ```go
 reference := "kkkkkkkkkkkkk"
 err, response := transfer.FetchTransfer(reference)
 if err != nil {
  panic(err)
 }
 fmt.Println(response)
 ```

 #### Sample Response

 This call returns a dictionary. A sample response is:

  ```sh
  map[status:success message:QUERIED-TRANSFERS data:map[page_info:map[total:1 current_page:1 total_pages:1] transfers:[map[requires_approval:0 debit_currency:<nil> reference:kkkkkkkkkkkkk complete_message:DISBURSE FAILED: undefined is_approved:1 id:3563 account_number:0690000044 bank_code:044 currency:NGN fee:45 meta:map[] approver:<nil> bank_name:ACCESS BANK NIGERIA narration:New transfer fullname:Mercedes Daniel date_created:2018-11-26T21:38:31.000Z amount:500 status:FAILED]]]]
  ```

### ```.FetchAllTransfers() (error error, response map[string]interface{})```

This allows you retrieve all transfers.

A sample fetchall call is:

```go
err, response := transfer.FetchAllTransfers()
if err != nil {
 panic(err)
}
fmt.Println(response)
```

#### Sample Response

This call returns a dictionary. A sample response is:

 ```sh
 map[status:success message:QUERIED-TRANSFERS data:map[page_info:map[total:1 current_page:1 total_pages:1] transfers:[map[requires_approval:0 debit_currency:<nil> reference:kkkkkkkkkkkkk complete_message:DISBURSE FAILED: undefined is_approved:1 id:3563 account_number:0690000044 bank_code:044 currency:NGN fee:45 meta:map[] approver:<nil> bank_name:ACCESS BANK NIGERIA narration:New transfer fullname:Mercedes Daniel date_created:2018-11-26T21:38:31.000Z amount:500 status:FAILED]]]]
 ```


### ```.GetTransferFee(currency string) (error error, response map[string]interface{})```

This allows you get transfer rates for all Rave supported currencies. You may or may not pass in a ```currency```. If you do not pass in a ```currency```, all Rave supported currencies transfer rates will be returned.

A sample getFee call is:

```go
currencies := "NGN"
error, response := transfer.GetTransferFee(currency)
if err != nil {
 panic(err)
}
fmt.Println(response)

```

#### Sample Response

This call returns a dictionary. A sample response is:

 ```sh
map[message:TRANSFER-FEES data:[map[AccountId:1 id:1 fee_type:value currency:NGN fee:45 createdAt:<nil> updatedAt:<nil> deletedAt:<nil>]] status:success]

 ```

### ```.GetRaveBalance(currency string) (error error, response map[string]interface{})```

This allows you get your balance in a specified currency. You may or may not pass in a ```currency```. If you do not pass in a ```currency```, your balance will be returned in the currency specified in yiur rave account

A sample balance call is:

```go
currencies := "NGN"
error, response := transfer.GetRaveBalance(currency)
if err != nil {
 panic(err)
}
fmt.Println(response)
```

#### Returns

This call returns a dictionary. A sample response is:

 ```sh
map[status:success message:WALLET-BALANCE data:map[LedgerBalance:0 Id:32509 ShortName:NGN WalletNumber:4446000147772 AvailableBalance:0]]
 ```

### ```.GetBulkTransferStatus(batch_id string) (error error, response map[string]interface{})```

This allows you get your status of a queued bulk transfer You may or may not pass in a ```batch_id```. If you do not pass in a ```batch_id```, all queued bulk transfers will be returned

A sample bulk transfer status call is:

```go
batchIDs := [2]string{"634", "635"}

error, response := transfer.GetBulkTransferStatus(batchID)
if err != nil {
 panic(err)
}
fmt.Println(response)
```

#### Returns

This call returns a dictionary. A sample response is:

 ```sh
map[message:QUERIED-TRANSFERS data:map[page_info:map[total_pages:1 total:2 current_page:1] transfers:[map[bank_name:ACCESS BANK NIGERIA account_number:0690000032 fullname:Pastor Bright amount:10 reference:<nil> narration:Bulk transfer 1 approver:<nil> id:3542 bank_code:044 requires_approval:0 is_approved:1 date_created:2018-11-26T14:21:44.000Z currency:NGN debit_currency:<nil> fee:45 status:FAILED meta:<nil> complete_message:DISBURSE FAILED: Invalid transfer amount. Minimum is 100] map[bank_code:044 fee:45 meta:<nil> account_number:0690000034 currency:NGN debit_currency:<nil> amount:10 complete_message:DISBURSE FAILED: Invalid transfer amount. Minimum is 100 fullname:Ade Bond date_created:2018-11-26T14:21:44.000Z reference:<nil> narration:Bulk transfer 2 approver:<nil> is_approved:1 id:3543 status:FAILED requires_approval:0 bank_name:ACCESS BANK NIGERIA]]] status:success]
 ```


### ```.ResolveAccount(account_data AccountResolveData) (error error, response map[string]interface{})```
This allows you verify an account to transfer to. ```account_data``` should be of type ```rave.AccountResolveData```. See below for  ```rave.AccountResolveData``` definition

```go
type AccountResolveData struct {
	PublicKey        string `json:"PBFPubKey"`
	RecipientAccount string `json:"recipientaccount"`
	DestBankCode     string `json:"destbankcode"`
	Currency         string `json:"currency"`
	Country          string `json:"country"`
}
```

A sample initiate call is:

```go
    payload := rave.AccountResolveData{
        RecipientAccount: "0690000034",
        DestBankCode:     "044",
        PublicKey:        r.GetPublicKey(),
    }

    err, response := transfer.ResolveAccount(payload)
    if err != nil {
        panic(err)
    }
    fmt.Println(response)
```

#### Sample Response

```sh
map[message:ACCOUNT RESOLVED data:map[data:map[responsecode:00 accountnumber:0690000034 accountname:Ade Bond responsemessage:Approved Or Completed Successfully phonenumber:<nil> uniquereference:FLWT001034195 internalreference:<nil>] status:success] status:success]
```

## ```rave.Subaccount{}```
This is used to facilitate Subaccount operations via rave. ```rave.Subaccount{}``` is of type ```struct``` and requires  ```rave.Rave``` as its only property.

Hence, in order to use it, you need to pass in an instance of ```rave.Rave``` . A sample is shown below

```go
    var transfer = rave.Subaccount{
    	r,
    }
```

**Methods Included:**

* ```.CreateSubaccount```

* ```.ListSubaccount```

* ```.FetchSubaccount```

* ```.DeleteSubaccount```

### ```.CreateSubaccount(data CreateSubaccountData) (error error, response map[string]interface{})```
This is called to create a subaccount. The payload should be of type ```rave.CreateSubaccountData```. See below for  ```rave.CreateSubaccountData``` definition

```go
type CreateSubaccountData struct {
	AccountBank                   string         `json:"account_bank"`
	AccountNumber                 string         `json:"account_number"`
	BusinessName                  string         `json:"business_name"`
	BusinessEmail                 string         `json:"business_email"`
	BusinessContact               string         `json:"business_contact"`
	BusinessMobile                string         `json:"business_mobile"`
	BusinessContactMobile         string         `json:"business_contact_mobile"`
	Seckey                        string         `json:"seckey"`
	Meta                          types.Slice    `json:"meta"`
	SplitType                     string         `json:"split_type"`
	SplitValue                    string         `json:"split_value"`
}
```

A sample create call is:

```go
    payload := rave.CreateSubaccountData{
        AccountBank: "044",
	 	AccountNumber: "0690000035",
	 	BusinessName: "Test",
	 	BusinessEmail: "test@test.com",
	 	BusinessContact: "Seun Alade",
	 	BusinessContactMobile: "09012345678",
	 	BusinessMobile: "09087930123",
	 	SplitType: "flat",
	 	SplitValue: "100",
    }

    err, response := subaccount.CreateSubaccount(payload)
    if err != nil {
        panic(err)
    }
    fmt.Println(response)
```

#### Sample Response

```

```
### ```.ListSubaccount(data ListSubaccountData) (error error, response map[string]interface{})```
This is called to list all or specific subaccounts. The payload should be of type ```rave.ListSubaccountData```. See below for  ```rave.ListSubaccountData``` definition

```go
type ListSubaccountData struct {
	AccountBank               string         `json:"account_bank"`
	AccountNumber             string         `json:"account_number"`
	BankName                  string         `json:"bank_name"`
	Seckey                    string         `json:"seckey"`
}
```

A sample list call is:

```go
    payload := rave.ListSubaccountData{
        AccountNumber: "0690000035",
    }

    err, response := subaccount.ListSubaccount(payload)
    if err != nil {
        panic(err)
    }
    fmt.Println(response)
```

#### Sample Response

```

```
### ```.FetchSubaccount(id string) (error error, response map[string]interface{})```
This allows you fetch a single subaccount using the subaccount ID. The ID should be of type ```string``` and it is required.

A sample fetch call is:


```go
    id = "example_id"

    err, response := subaccount.FetchSubaccount(id)
    if err != nil {
        panic(err)
    }
    fmt.Println(response)
```

#### Sample Response

```

```
### ```.DeleteSubaccount(id string) (error error, response map[string]interface{})```
This allows you to delete a subaccount by the subaccount ID. The ID should be of type ```string``` and it is required.

A sample delete call is:


```go
    id = "example_id"

    err, response := subaccount.DeleteSubaccount(id)
    if err != nil {
        panic(err)
    }
    fmt.Println(response)
```

#### Sample Response

```

```
