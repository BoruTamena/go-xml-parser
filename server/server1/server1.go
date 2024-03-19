package server1

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	Password = "password"
)

type Envelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Body    Body     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
}

type Body struct {
	C2BPaymentQueryResult C2BPaymentQueryResult `xml:"http://cps.huawei.com/cpsinterface/c2bpayment C2BPaymentQueryResult"`
}

type C2BPaymentQueryResult struct {
	ResultCode    int    `xml:"ResultCode"`
	ResultDesc    string `xml:"ResultDesc"`
	TransID       int    `xml:"TransID"`
	BillRefNumber int    `xml:"BillRefNumber"`
	UtilityName   string `xml:"UtilityName"`
	CustomerName  string `xml:"CustomerName"`
	Amount        int    `xml:"Amount"`
	Password      string `json:"Password,omitempty"`
}

func SendDataToServer2(xmlstring string) {

	var envelope Envelope

	err := xml.Unmarshal([]byte(xmlstring), &envelope)

	if err != nil {
		log.Fatal(err)
	}

	envelope.Body.C2BPaymentQueryResult.Password = Password

	json_data, jerr := json.Marshal(envelope)

	if jerr != nil {
		log.Fatal(jerr)
	}

	fmt.Println(string(json_data))
	cxt, cancel := context.WithTimeout(context.Background(), 10*time.Second) // creating context with time limit

	defer cancel()

	req, reqerr := http.NewRequestWithContext(cxt, http.MethodPost, "http://localhost:8080/receive", bytes.NewBuffer(json_data))

	if reqerr != nil {
		log.Fatal(reqerr)
	}

	req.Header.Set("Content-Type", "application/json")

	client := http.DefaultClient

	resp, rerr := client.Do(req)

	if rerr != nil {
		fmt.Println("Unable to connect to server 2")
		return
	}

	defer resp.Body.Close()

	fmt.Println(resp.StatusCode)

	if resp.StatusCode != http.StatusOK {
		fmt.Println("")
		return
	}

	fmt.Print("Data sent to server 2")

}
