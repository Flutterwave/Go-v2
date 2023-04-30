package rave

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

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

type BillCategoryFilter string

const (
	Airtime    BillCategoryFilter = "airtime"
	DataBundle BillCategoryFilter = "data_bundle"
	Power      BillCategoryFilter = "power"
	Internet   BillCategoryFilter = "internet"
	Toll       BillCategoryFilter = "toll"
	Cable      BillCategoryFilter = "cable"
	BillerCode BillCategoryFilter = "biller_code"
)

type ValidationData struct {
	ItemCode *string `json:"item_code"`
	Code     *string `json:"code"`
	Customer *string `json:"customer"`
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

func (b BillPayments) GetBillCategories(filter ...BillCategoryFilter) (categories []BillCategory, err error) {
	var params map[string]string
	url := b.GetBaseURL() + b.GetEndpoint("Billspayments", "categories")
	if len(filter) > 0 {
		if len(filter) > 1 {
			err = errors.New("You cannot select more than one filter")
			return
		}
		params = map[string]string{
			string(filter[0]): strconv.Itoa(1),
		}
	}
	err = b.Get(url, params, &categories)
	return
}

func (b BillPayments) ValidateBillCategory(data *ValidationData) (response ValidationResponse, err error) {
	if data == nil {
		err = errors.New("Please provide category information")
		return
	}
	var params map[string]string
	url := b.GetBaseURL() + b.GetEndpoint("Billspayments", "validate")

	if data.ItemCode != nil {
		url = strings.Replace(url, ":item_code", *data.ItemCode, -1)
	} else {
		err = errors.New("Please provide item_code")
		return
	}
	if data.Code != nil {
		params["code"] = *data.Code
	} else {
		err = errors.New("Please provide a biller code")
		return
	}
	if data.Customer != nil {
		params["customer"] = *data.Customer
	} else {
		err = errors.New("Please provide a customer identifier")
		return
	}
	err = b.Get(url, params, &response)
	return

}

func (b BillPayments) Create(req *BillPaymentRequest) (response BillPaymentResponse, err error) {
	url := b.GetBaseURL() + b.GetEndpoint("Billspayments", "create")
	err = b.Post(url, nil, req, &response)
	return
}

func (b BillPayments) Status(ref string) (response BillPaymentResponse, err error) {
	if ref == "" {
		err = errors.New("Please specify a tx_ref")
		return
	}
	url := b.GetBaseURL() + b.GetEndpoint("Billspayments", "fetch")
	err = b.Get(fmt.Sprintf("%v/%v", url, ref), nil, &response)
	return
}
