package datamodels

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type FinalData struct {
	/*
		struct data for finaldata after unmarshalling
	*/
	Id               int    `json:"id"`
	Reference        string `json:"reference"`
	Amount           uint   `json:"amount"`
	Paid_at          string `json:"paid_at"`
	Created_at       string `json:"created_at"`
	Email            string `json:"email"`
	Mobile           string `json:"mobile"`
	MeterNumber      uint   `json:"meter_number"`
	State            string `json:"state"`
	Status           string `json:"status"`
	Gateway_response string `json:"gateway_response"`
	Fees             int    `json:"paygateway_fees"`
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	Channel          string `json:"channel"`
	Country_code     string `json:"country_code"`
	Bank             string `json:"bank"`
}

type TransactionPaymentData struct {
	/*
		struct data for transactiondata payload from API url
	*/
	Event string
	Data  struct {
		Id               int
		Reference        string
		Amount           uint
		Status           string
		Gateway_response string
		Paid_at          string
		Created_at       string
		Metadata         struct {
			Meternumber string
			State       string
			Phone       string
			First_name  string
			Last_name   string
		}
		Fees     int
		Customer struct {
			Email     string
			Last_name string
			Phone     string
		}
		Authorization struct {
			Channel      string
			Bank         string
			Country_code string
		}
	}
}

func (finalData *FinalData) UnmarshalJSON(data []byte) error {
	/*
		function to parse data and unmarshal to ideal json pair values
	*/
	var TransactionData TransactionPaymentData
	err := json.Unmarshal(data, &TransactionData)
	if err != nil {
		return err
	}
	if TransactionData.Event == "charge.success" {

		td := TransactionData.Data

		finalData.Id = td.Id
		finalData.Reference = td.Reference
		finalData.Gateway_response = td.Gateway_response
		finalData.Status = td.Status
		/*
			Amount comes with extra 00.
			convert amount to ideal values without the extra zeros
			convert uint amount type to string data type
			remove last two characters
			convert from characters to uint64 and back to uint
		*/
		var AmountAsString string = strconv.FormatUint(uint64(td.Amount), 10)
		if last := len(AmountAsString) - 2; last >= 0 && AmountAsString[last] == '0' {
			AmountAsString = AmountAsString[:last]
		}
		AmountToUint64, err := strconv.ParseUint(AmountAsString, 10, 64)
		if err != nil {
			return fmt.Errorf("%s", err.Error())
		}
		AmountToUint := uint(AmountToUint64)
		finalData.Amount = AmountToUint
		finalData.Paid_at = td.Paid_at
		finalData.Created_at = td.Created_at
		//convert meter number from string datatype to uint
		MeterNumberToUint64, _ := strconv.ParseUint(td.Metadata.Meternumber, 10, 64)
		finalData.MeterNumber = uint(MeterNumberToUint64)
		//convert account number from string datatype to uint
		finalData.State = td.Metadata.State
		finalData.Email = td.Customer.Email
		finalData.Mobile = td.Customer.Phone
		finalData.FirstName = td.Metadata.First_name
		finalData.LastName = td.Customer.Last_name
		finalData.Fees = td.Fees
		finalData.Channel = td.Authorization.Channel
		finalData.Bank = td.Authorization.Bank
		finalData.Country_code = td.Authorization.Country_code
	}
	return nil
}
