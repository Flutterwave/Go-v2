package rave

import "time"

type BillPayments struct {
	Rave
}

type BillCategory struct {
	Id                int       `json:"id"`
	BillerCode        string    `json:"biller_code"`
	Name              string    `json:"name"`
	DefaultCommission float64   `json:"default_commission"`
	DateAdded         time.Time `json:"date_added"`
	Country           string    `json:"country"`
	IsAirtime         bool      `json:"is_airtime"`
	BillerName        string    `json:"biller_name"`
	ItemCode          string    `json:"item_code"`
	ShortName         string    `json:"short_name"`
	Fee               int       `json:"fee"`
	CommissionOnFee   bool      `json:"commission_on_fee"`
	LabelName         string    `json:"label_name"`
	Amount            int       `json:"amount"`
}
type BillCategoriesFilter struct {
	Airtime    int32 `json:"airtime"`
	DataBundle int32 `json:"data_bundle"`
	Power      int32 `json:"power"`
	Internet   int32 `json:"internet"`
	Toll       int32 `json:"toll"`
	Cable      int32 `json:"cable"`
	BillerCode int32 `json:"biller_code"`
}

type ValidationData struct {
	ItemCode string `json:"item_code"`
	Code     string `json:"code"`
	Customer string `json:"customer"`
}

type ValidationResponse struct {
	ResponseCode    string      `json:"response_code"`
	Address         interface{} `json:"address"`
	ResponseMessage string      `json:"response_message"`
	Name            string      `json:"name"`
	BillerCode      string      `json:"biller_code"`
	Customer        string      `json:"customer"`
	ProductCode     string      `json:"product_code"`
	Email           interface{} `json:"email"`
	Fee             int         `json:"fee"`
	Maximum         int         `json:"maximum"`
	Minimum         int         `json:"minimum"`
}

type BillPaymentRequest struct {
	Country    string `json:"country"`
	Customer   string `json:"customer"`
	Amount     string `json:"amount"`
	Recurrence string `json:"recurrence"`
	Type       string `json:"type"`
	Reference  string `json:"reference"`
	BillerName string `json:"biller_name"`
}

type BillPaymentResponse struct {
	PhoneNumber string      `json:"phone_number"`
	Amount      int         `json:"amount"`
	Network     string      `json:"network"`
	FlwRef      string      `json:"flw_ref"`
	TxRef       string      `json:"tx_ref"`
	Reference   interface{} `json:"reference"`
	Fee         int         `json:"fee"`
	Currency    *string     `json:"currency"`
	Extra       interface{} `json:"extra"`
	Token       interface{} `json:"token"`
}

func (b Billpayment) GetBillCategories(filters *BillCategoriesFilter) (categories []BillCategory, err error) {
}

func (b Billpayment) ValidateBillCategory(data *ValidationData) (response ValidationResponse, err error) {
}

func (b Billpayment) Create(req *BillPaymentRequest) (response BillPaymentResponse, err error) {}

func (b Billpayment) Status(ref string) (response BillPaymentResponse, err error) {}
