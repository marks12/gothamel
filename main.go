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
            Text: "Some content 1",
        },
        tag.Div{
            Text: "Some content 2",
        },
        tag.Div{
            Text: "Some content 3",
        },
        tag.P{
            Text: "Some content P0 in P",
        },
        tag.Div{
            Children: &[]tag.Nodes{
                tag.Div{
                    Text: "Some content 3.1",
                },
            },
        },
        tag.Div{
            Text: "Some content 4",
        },
    }}

    doc.SetContent(content).
        AddHeader(tag.Header{}).
        AddFooter(tag.Footer{})

    htmlPage := doc.GetHtmlPage(true)

    os.Stdout.Write([]byte(htmlPage))
    os.Stdout.Write([]byte("\n"))
}
