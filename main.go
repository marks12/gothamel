package main

import (
    "os"
    "github.com/marks12/gothamel/layout"
    "github.com/marks12/gothamel/tag"
)

func main() {

    doc := layout.Doc()

    content := tag.Article{Children: &[]tag.Nodes{
        tag.Div{
            Text: "Some content",
        },
    }}

    doc.SetContent(content).
        AddHeader(tag.Header{}).
        AddFooter(tag.Footer{})

    htmlPage := doc.GetHtmlPage(true)

    os.Stdout.Write([]byte(htmlPage))
    os.Stdout.Write([]byte("\n"))
}
