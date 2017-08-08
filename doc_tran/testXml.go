package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

type Data struct {
	XMLName xml.Name `xml:"main" json:"-"`
	A       string   `xml:"A" json:"A"`
	B       string   `xml:"B" json:"B"`
	CList   []C      `xml:"C" json:"C"`
}

type C struct {
	XMLName xml.Name `xml:"C" json:"C"`
	C1      string   `xml:"C1" json:"C1"`
	C2      string   `xml:"C2" json:"C2"`
}

func main() {
	rawXMLData := []byte(`<?xml version="1.0" encoding="UTF-8" ?>
	 					<main>
	 						<A>xxx</A>
	 						<B>yyy</B>
	 						<C>
	 							<C1>some_text1</C1>
	 							<C2>some_text2</C2>
	 						</C>
	 						<person>
	 							<firstname>Maria</firstname>
	 							<lastname>Raboy</lastname>
	 						</person>
	 					</main>`)

	rawJSONData := []byte(`{"A": "xxx", "B": "yyy", "C": [{"C1": "some_data1"}, {"C2": "some_data2"}], "person": "Maria"}`)

	data := &Data{}
	err := xml.Unmarshal(rawXMLData, data)
	if nil != err {
		fmt.Println("Error unmarshalling from XML", err)
		return
	}

	result, err := json.Marshal(data)
	if nil != err {
		fmt.Println("Error marshalling to JSON", err)
		return
	}

	fmt.Printf("%s\n", result)

	data1 := &Data{}
	err = json.Unmarshal(rawJSONData, data1)
	if nil != err {
		fmt.Println("Error unmarshalling from JSON", err)
		return
	}

	resultJ, err := xml.Marshal(data)
	if nil != err {
		fmt.Println("Error marshalling to XML", err)
		return
	}

	fmt.Printf("%s\n", resultJ)
}
