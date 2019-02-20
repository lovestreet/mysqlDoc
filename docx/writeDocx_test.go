package docx

import (
	"testing"

	"baliance.com/gooxml/document"
)

func TestWriteText(t *testing.T) {

	var doc = document.New()
	para := doc.AddParagraph()
	run := para.AddRun()
	run.AddText("Hello World")
	doc.SaveToFile("hello.docx")
}

func TestWriteTable(t *testing.T) {

	var doc = document.New()
	addTable(doc)
	addTable(doc)
	doc.SaveToFile("hello.docx")
}
