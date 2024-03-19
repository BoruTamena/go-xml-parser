package server2

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/BoruTamena/server-go/server/server1"
)

var (
	filename     = "success.xml"
	hashpassword = hashPasswordMD5(server1.Password)
)

func hashPasswordMD5(password string) string {
	hasher := md5.New()
	hasher.Write([]byte(password))
	hashedPassword := hex.EncodeToString(hasher.Sum(nil))
	return hashedPassword
}

func Receive_data(w http.ResponseWriter, r *http.Request) {

	var envlope server1.Envelope
	err := json.NewDecoder(r.Body).Decode(&envlope)

	if err != nil {

		log.Fatal(err)
	}

	data := envlope.Body.C2BPaymentQueryResult

	password := data.Password

	// validating the password

	if hashpassword != hashPasswordMD5(password) {
		filename = "failed.xml"
	}

	// writing to file

	xml_byte, err := xml.Marshal(envlope)

	if err != nil {
		http.Error(w, "Failed to marshal xml", http.StatusInternalServerError)
		return
	}

	err = ioutil.WriteFile(filename, xml_byte, 0644)

	if err != nil {
		http.Error(w, "Failed to save xml ", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(xml_byte) // writing xml data back to the endpoint

}
