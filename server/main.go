package main

import "github.com/BoruTamena/server-go/server/server1"

func main() {

	stringxml := `

		<soapenv:Envelope
		xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/"
		xmlns:c2b="http://cps.huawei.com/cpsinterface/c2bpayment">
		<soapenv:Header/>
		<soapenv:Body>
		<c2b:C2BPaymentQueryResult>
		<ResultCode>2</ResultCode>
		<ResultDesc>Failed</ResultDesc>
		<TransID>10111</TransID>
		<BillRefNumber>12233</BillRefNumber>
		<UtilityName>sddd</UtilityName>
		<CustomerName>wee</CustomerName>
		<Amount>30</Amount>
		</c2b:C2BPaymentQueryResult>
		</soapenv:Body>
		</soapenv:Envelope>

	`

	server1.SendDataToServer2(stringxml)

}
