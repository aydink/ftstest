package main

import (
	"fmt"

	"github.com/aydink/fts"
)

func main() {

	/*
		sampleText := "Babası  sıfır otomobil hediye etti! Ekspertize götürünce hayatının şokunu yaşadı. ĞÜŞİÖÇ ÂÎÛ âîû"

		a := fts.NewSimpleAnalyzer(fts.NewSimpleTokenizer())
		tokens := a.Analyze(sampleText)
		fmt.Println(tokens)

		turkishFilter := fts.NewTurkishLowercaseFilter()

		a.AddTokenFilter(turkishFilter)
		tokens = a.Analyze(sampleText)
		fmt.Println(tokens)

		turkishAccentFilter := fts.NewTurkishAccentFilter()
		a.AddTokenFilter(turkishAccentFilter)
		tokens = a.Analyze(sampleText)
		fmt.Println(tokens)
	*/
	//TestAnlyzer()
	TestFields()

}

func TestAnlyzer() {
	sampleText := "TOKEN MOKEN Babası (sıfır-bir) otomobil hediye etti! Ekspertize götürünce hayatının şokunu yaşadı. ĞÜŞİÖÇ ÂÎÛ âîû"

	analyzer := fts.NewSimpleAnalyzer(fts.NewSimpleTokenizer())

	turkishFilter := fts.NewTurkishLowercaseFilter()
	analyzer.AddTokenFilter(turkishFilter)

	turkishAccentFilter := fts.NewTurkishAccentFilter()
	analyzer.AddTokenFilter(turkishAccentFilter)

	tokens := analyzer.Analyze(sampleText)
	fmt.Println(tokens)

}

func TestFields() {

	analyzer := fts.NewSimpleAnalyzer(fts.NewSimpleTokenizer())

	turkishFilter := fts.NewTurkishLowercaseFilter()
	turkishAccentFilter := fts.NewTurkishAccentFilter()

	analyzer.AddTokenFilter(turkishFilter)
	analyzer.AddTokenFilter(turkishAccentFilter)

	index := fts.New("index", analyzer)

	textTitle := "Merhaba dünya, bu bir title dünya"
	textContent := "TOKEN MOKEN Babası (sıfır-bir) otomobil hediye etti! Ekspertize götürünce hayatının şokunu yaşadı. ĞÜŞİÖÇ ÂÎÛ âîû"

	title := fts.NewTextField("title", textTitle, fts.INDEXED, fts.STORED, fts.ANALYZED)
	content := fts.NewTextField("content", textContent, fts.INDEXED, fts.STORED, fts.ANALYZED)

	doc := fts.NewDocument()
	doc.Add(title)
	doc.Add(content)
	doc.AddCatergory("cat1")
	doc.AddCatergory("cat2")

	index.Add(doc)

	textTitle = "İkinci  hamlesi kalkınma, dünya dünya yalan dünya"
	textContent = "sana bir gül hediye etsem nasıl olur, nasıl desem, nasıl anlatsam otomobil mi alsam acaba"

	title = fts.NewTextField("title", textTitle, fts.INDEXED, fts.STORED, fts.ANALYZED)
	content = fts.NewTextField("content", textContent, fts.INDEXED, fts.STORED, fts.ANALYZED)

	doc = fts.NewDocument()
	doc.Add(title)
	doc.Add(content)
	doc.AddCatergory("cat1")
	doc.AddCatergory("cat2")

	index.Add(doc)

	textTitle = "Üçüncü kalkınma hamlesi"
	textContent = "Otomobil ile bir tur atalım seninle otomobil ile"

	title = fts.NewTextField("title", textTitle, fts.INDEXED, fts.STORED, fts.ANALYZED)
	content = fts.NewTextField("content", textContent, fts.INDEXED, fts.STORED, fts.ANALYZED)

	doc = fts.NewDocument()
	doc.Add(title)
	doc.Add(content)
	doc.AddCatergory("cat1")
	doc.AddCatergory("cat2")
	index.Add(doc)

	index.Commit()

	fmt.Printf("%+v\n\n", index)
	posting1, _ := index.ReadPostings("title:kalkınma")
	posting2, _ := index.ReadPostings("title:hamlesi")
	posting := fts.Intersection(posting1.Postings, posting2.Postings)
	fmt.Printf("%+v\n", posting)

}
