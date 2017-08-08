package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
	//"encoding/xml"
	"encoding/xml"
)

//!+main

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(decodeHandler))
	mux.Handle("/J2X", http.HandlerFunc(J2Xhandler))
	mux.Handle("/X2J", http.HandlerFunc(X2Jhandler))

	log.Fatal(http.ListenAndServe(":8000", mux))
}

type struct_for_JSON map[string]string

type Data_type struct {
	XMLName xml.Name `xml:"main" json:"-"`
	A       string   `xml:"A" json:"A"`
	B       string   `xml:"B" json:"B"`
	CList   []C_type `xml:"C" json:"C"`
}

type C_type struct {
	XMLName xml.Name `xml:"C" json:"C"`
	C1      string   `xml:"C1" json:"C1"`
	C2      string   `xml:"C2" json:"C2"`
}

//!+handler

func J2Xhandler(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	body := &Data_type{}
	if err := decoder.Decode(&body); err != nil {
		defer req.Body.Close()
		panic(err)
	}
	defer req.Body.Close()

	XMLBody, err := xml.Marshal(body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Printf("%s\n", XMLBody)

	//Connect to localdb
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	Jses := session.DB("test").C("XMLcollection")

	err = Jses.Insert(bson.M{"this is XML document": &body})
	if err != nil {
		log.Fatal(err)
	}

	rw.Header().Set("Content-Type", "application/xml") // normal header
	rw.WriteHeader(http.StatusOK)
	rw.Write(XMLBody)
	fmt.Fprintf(rw, "XML response")
}

func X2Jhandler(rw http.ResponseWriter, req *http.Request) {
	decoder := xml.NewDecoder(req.Body)
	body := &Data_type{}
	if err := decoder.Decode(&body); err != nil {
		defer req.Body.Close()
		panic(err)
	}

	///////----ready to marshal
	JSONBody, err := json.Marshal(body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Printf("%s\n", JSONBody)

	//Connect to localdb
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	Jses := session.DB("test").C("JSONcollection")

	err = Jses.Insert(bson.M{"this is JSON document": &body})
	if err != nil {
		log.Fatal(err)
	}

	rw.Header().Set("Content-Type", "application/json") // normal header
	rw.WriteHeader(http.StatusOK)
	rw.Write(JSONBody)
	fmt.Fprintf(rw, "JSON response")
}

func decodeHandler(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var result struct_for_JSON
	if err := decoder.Decode(&result); err != nil {
		defer req.Body.Close()
		panic(err)
	}
	defer req.Body.Close()

	//Connect to localdb
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	Jses := session.DB("test").C("JSONcollection")
	//Xses := session.DB("test").C("XMLcollection")

	err = Jses.Insert(&result)
	if err != nil {
		log.Fatal(err)
	}

	res := []struct_for_JSON{}
	err = Jses.Find(bson.M{}).All(&res)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Phone:", res)

	rw.Header().Set("Content-Type", "text/plain; charset=utf-8") // normal header
	rw.WriteHeader(http.StatusOK)
	fmt.Fprintf(rw, "Hi there, I love you!")
}
