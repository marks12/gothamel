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
    
    
```html 

<!DOCTYPE html><Html>
    <Head>
        <Title>Some title</Title>
    </Head>
    <Body>
        <Header></Header>
        <Article>
            <Div>Some content 1</Div>
            <Div>Some content 2</Div>
            <Div>Some content 3</Div>
            <P>Some content P0 in P</P>
            <Div>
                <Div>Some content 3.1</Div>
            </Div>
            <Div>Some content 4</Div>
        </Article>
        <Footer></Footer>
    </Body>
</Html>


```