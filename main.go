package main

import (
"encoding/xml"
"fmt"
"os"
)

func main() {
    ExampleMarshalIndent()
}

func ExampleMarshalIndent() {

    Header := "<!DOCTYPE html>" + "\n"

    type Node struct {
        Id    string          `xml:"id,attr"`
    }

    type P struct {
        XMLName   xml.Name      `xml:"P"`
        Node
        Text    string
    }

    type Tags interface {
    }

    type Body struct {
        XMLName   xml.Name      `xml:"Body"`
        Tags
    }

    type Head struct {
        XMLName   xml.Name      `xml:"Head"`
        Title string
    }

    type Html struct {
        XMLName     xml.Name     `xml:"Html"`
        Head        Head         `xml:"Head"`
        Body        interface{}         `xml:"Body"`
    }

    content := []Tags{
        P{Text: "SomeText1"},
        P{Text: "SomeText2"},
        P{Text: "SomeText3"},
    }

    b, err := xml.Marshal(content)

    v := Html{
        Head: Head{Title: "Some title"},
        Body: b,
    }

    output, err := xml.MarshalIndent(v, "", "    ")
    if err != nil {
        fmt.Printf("error: %v\n", err)
    }

    fmt.Println(Header)
    fmt.Println(string(output))

    os.Stdout.Write([]byte("\n"))
}

