package xml

import (
	"encoding/xml"
	"log"
	"testing"
)

/*
This is what we want:
<adept:signIn xmlns:adept="http://ns.adobe.com/adept" method="anonymous">
  <adept:signInData>XXXXX</adept:signInData>
  <adept:publicAuthKey>XXXXX</adept:publicAuthKey>
</adept:signIn>
*/

var adeptXML = `` +
	`<adept:signIn xmlns:adept="http://ns.adobe.com/adept" method="anonymous">` +
	`<adept:signInData>XXXXX</adept:signInData>` +
	`<adept:publicAuthKey>XXXXX77</adept:publicAuthKey>` +
	`</adept:signIn>`

type SignIn struct {
	XMLName xml.Name `xml:"http://ns.adobe.com/adept signIn"`
	//Xmlns         string   `xml:"xmlns:adept,attr"`
	Method        string `xml:"method,attr"`
	SignInData    string `xml:"adept signInData"`
	PublicAuthKey string `xml:"adept publicAuthKey"`
}

func TestAdept(t *testing.T) {
	tmp := SignIn{SignInData: "test2"}
	m, _ := MarshalPrefix(tmp)
	log.Println(string(m))
	err := UnmarshalPrefix([]byte(adeptXML), &tmp)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(tmp)
}
