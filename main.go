package main

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
"paystack-go/datamodels"
)

const (
	PAYSTACK_SECRET_KEY = "sk_test_3aae1b6bff40cb5a03fa72adf734204861dfcb2c"
)

func main() {
	

	// Create a HTTP server to listen for Paystack webhook requests
	http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		
		if r.Method != "POST" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		paymentDetails, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusInternalServerError)
			return
		}
		
		respData := []byte(paymentDetails)
		payld := datamodels.FinalData{}                          // Instantiate the struct
		json_err := json.Unmarshal(respData, &payld) // Unmarshal the byte data
		if json_err != nil {
			fmt.Printf("error unmarshalling data:%v", json_err.Error())
		}

		// Convert struct type of realdata to map interface
		var paymentdataMap map[string]interface{}
		payld_data, _ := json.Marshal(payld)
		json.Unmarshal(payld_data, &paymentdataMap)
		fmt.Print(paymentdataMap)
fmt.Fprint(w,paymentdataMap)
		// Return a success response
		w.WriteHeader(http.StatusOK)
	})

	// Create a home URL
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to the Paystack webhook server!")
	})

	// Start the HTTP server
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func verifySignature(signature string, body []byte) bool {
	// Compute the HMAC SHA512 signature using the Paystack secret key
	h := hmac.New(sha512.New, []byte(PAYSTACK_SECRET_KEY))
	h.Write(body)
	expectedSignature := hex.EncodeToString(h.Sum(nil))

	// Compare the expected signature with the actual signature
	return hmac.Equal([]byte(signature), []byte(expectedSignature))
}
