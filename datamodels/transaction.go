package datamodels

import (
	"encoding/json"
	"fmt"
	"strconv"
)


type RealData struct {
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
}

type TransactionData struct {
	/*
		struct data for transactiondata payload from API url
	*/
	Status  bool
	Message string
	Data    []struct {
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
			Email string
		}
	}
	Meta json.RawMessage
}

func (realData *RealData) UnmarshalJSON(data []byte) error {
	/*
		function to parse data and unmarshal to ideal json pair values
	*/
	var TransactionData TransactionData
	err := json.Unmarshal(data, &TransactionData)
	if err != nil {
		return err
	}
	for tdindex, td := range TransactionData.Data {

		if tdindex == 0 {
			realData.Id = td.Id
			realData.Reference = td.Reference
			realData.Gateway_response = td.Gateway_response
			realData.Status = td.Status
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
			realData.Amount = AmountToUint
			realData.Paid_at = td.Paid_at
			realData.Created_at = td.Created_at
			//convert meter number from string datatype to uint
			MeterNumberToUint64, _ := strconv.ParseUint(td.Metadata.Meternumber, 10, 64)
			realData.MeterNumber = uint(MeterNumberToUint64)
			//convert account number from string datatype to uint
			realData.State = td.Metadata.State
			realData.Email = td.Customer.Email
			realData.Mobile = td.Metadata.Phone
			realData.FirstName = td.Metadata.First_name
			realData.LastName = td.Metadata.Last_name
			realData.Fees = td.Fees
		}
	}

	return nil
}