package bibinfo

import (
	"bufio"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
)

type Article struct {
	citekey string
	author  string
	pages   string
	title   string
	year    string
	journal string
	volume  string
	number  string
}

type Unpublished struct {
	citekey string
	author  string
	note    string
	title   string
	year    string
	month   string
}

type Book struct {
	citekey   string
	author    string
	title     string
	publisher string
	year      string
	volume    string
	series    string
	address   string
	edition   string
	month     string
	note      string
}

type PhDThesis struct {
	citekey string
	author  string
	title   string
	school  string
	year    string
	typ     string
	address string
	month   string
	note    string
}

type InBook struct {
	citekey   string
	author    string
	title     string
	editor    string
	chapter   string
	pages     string
	publisher string
	year      string
	volume    string
	series    string
	typ       string
	address   string
	edition   string
	month     string
	note      string
}

type Booklet struct {
	citekey      string
	title        string
	author       string
	howpublished string
	address      string
	month        string
	year         string
	note         string
}

type InCollection struct {
	citekey   string
	author    string
	title     string
	booktitle string
	editor    string
	chapter   string
	pages     string
	publisher string
	year      string
	volume    string
	series    string
	typ       string
	address   string
	edition   string
	month     string
	note      string
}

type Bibentry struct {
	citekey string
}

func Add_ArticleS() {
	bibliography, err := os.OpenFile(viper.GetString("bibliography"), os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer bibliography.Close()

	smf := new(Article)

	smf.citekey = GenInfo("citekey")
	smf.author = GenInfo("author")
	smf.title = GenInfo("title")
	smf.year = GenInfo("year")
	smf.journal = GenInfo("journal")
	smf.volume = GenInfo("volume")
	smf.number = GenInfo("number")
	smf.pages = GenInfo("pages")

	bibliography.WriteString("\n@article{")
	bibliography.WriteString(smf.citekey)
	bibliography.WriteString(",\n")
	bibliography.WriteString(WriteData("citekey", smf.citekey))
	if smf.author != "" {
		bibliography.WriteString(WriteData("author", smf.author))
	}
	if smf.title != "" {
		bibliography.WriteString(WriteData("title", smf.title))
	}
	if smf.year != "" {
		bibliography.WriteString(WriteData("year", smf.year))
	}
	if smf.journal != "" {
		bibliography.WriteString(WriteData("journal", smf.journal))
	}
	if smf.volume != "" {
		bibliography.WriteString(WriteData("volume", smf.volume))
	}
	if smf.number != "" {
		bibliography.WriteString(WriteData("number", smf.number))
	}
	if smf.pages != "" {
		bibliography.WriteString(WriteData("pages", smf.pages))
	}
	bibliography.WriteString("}\n")
}

func WriteData(field string, a string) string {
	printed_article := field + " = {" + a + "},\n"
	return printed_article
}

func GenInfo(x string) string {
	fmt.Println("Enter", x+strings.TrimSpace(":"))
	reader2 := bufio.NewReader(os.Stdin)
	info, _ := reader2.ReadString('\n')
	info = strings.TrimSuffix(info, "\n")
	return info
}

func OpenTag() string {
	opentag := "{"
	return opentag
}

func CloseTag() string {
	closetag := "},\n"
	return closetag
}

func AddBook() {
	fmt.Println("####### Enter new BOOK entry #######")
	//Open bibliography//
	bibliography, err := os.OpenFile(viper.GetString("bibliography"), os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer bibliography.Close()
	smf := new(Book)
	//Gathers all the information//
	smf.citekey = GenInfo("citekey")
	smf.author = GenInfo("author")
	smf.title = GenInfo("title")
	smf.year = GenInfo("year")
	smf.publisher = GenInfo("publisher")
	smf.volume = GenInfo("volume")
	smf.series = GenInfo("series")
	smf.address = GenInfo("address")
	smf.edition = GenInfo("edition")
	smf.month = GenInfo("month")
	smf.note = GenInfo("note")
	// writes the information to file
	bibliography.WriteString("\n@book{")
	bibliography.WriteString(smf.citekey)
	bibliography.WriteString(",\n")
	if smf.author != "" {
		bibliography.WriteString(WriteData("author", smf.author))
	}
	if smf.title != "" {
		bibliography.WriteString(WriteData("title", smf.title))
	}
	if smf.year != "" {
		bibliography.WriteString(WriteData("year", smf.year))
	}
	if smf.volume != "" {
		bibliography.WriteString(WriteData("volume", smf.volume))
	}
	if smf.publisher != "" {
		bibliography.WriteString(WriteData("publisher", smf.publisher))
	}
	if smf.series != "" {
		bibliography.WriteString(WriteData("series", smf.series))
	}
	if smf.address != "" {
		bibliography.WriteString(WriteData("address", smf.address))
	}
	if smf.edition != "" {
		bibliography.WriteString(WriteData("edition", smf.edition))
	}
	if smf.month != "" {
		bibliography.WriteString(WriteData("month", smf.month))
	}
	if smf.note != "" {
		bibliography.WriteString(WriteData("note", smf.note))
	}
	bibliography.WriteString("}\n")
}

func AddPhdThesis() {
	fmt.Println("####### Enter new PhD THESIS entry #######")
	//Open bibliography//
	bibliography, err := os.OpenFile(viper.GetString("bibliography"), os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer bibliography.Close()
	smf := new(PhDThesis)
	//Gathers all the information//
	smf.citekey = GenInfo("citekey")
	smf.author = GenInfo("author")
	smf.title = GenInfo("title")
	smf.school = GenInfo("school")
	smf.year = GenInfo("year")
	smf.address = GenInfo("address")
	smf.month = GenInfo("month")
	smf.note = GenInfo("note")
	smf.typ = GenInfo("type")
	// writes the information to file
	bibliography.WriteString("\n@phdthesis{")
	bibliography.WriteString(smf.citekey)
	bibliography.WriteString(",\n")
	if smf.author != "" {
		bibliography.WriteString("author = ")
		bibliography.WriteString("{")
		bibliography.WriteString(smf.author)
		bibliography.WriteString(CloseTag())
	}
	if smf.title != "" {
		bibliography.WriteString("title = ")
		bibliography.WriteString("{")
		bibliography.WriteString(smf.title)
		bibliography.WriteString(CloseTag())
	}
	if smf.year != "" {
		bibliography.WriteString("year = ")
		bibliography.WriteString("{")
		bibliography.WriteString(smf.year)
		bibliography.WriteString(CloseTag())
	}
	if smf.school != "" {
		bibliography.WriteString("school = ")
		bibliography.WriteString("{")
		bibliography.WriteString(smf.school)
		bibliography.WriteString(CloseTag())
	}
	if smf.address != "" {
		bibliography.WriteString("address = ")
		bibliography.WriteString("{")
		bibliography.WriteString(smf.address)
		bibliography.WriteString(CloseTag())
	}
	if smf.month != "" {
		bibliography.WriteString("month = ")
		bibliography.WriteString("{")
		bibliography.WriteString(smf.month)
		bibliography.WriteString(CloseTag())
	}
	if smf.note != "" {
		bibliography.WriteString("note = ")
		bibliography.WriteString("{")
		bibliography.WriteString(smf.note)
		bibliography.WriteString(CloseTag())
	}
	if smf.typ != "" {
		bibliography.WriteString("type = ")
		bibliography.WriteString("{")
		bibliography.WriteString(smf.typ)
		bibliography.WriteString(CloseTag())
	}
	bibliography.WriteString("}\n")
}

func AddInBook() {
	fmt.Println("####### Enter new InBOOK entry #######")
	//Open bibliography//
	bibliography, err := os.OpenFile(viper.GetString("bibliography"), os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer bibliography.Close()
	smf := new(InBook)
	//Gathers all the information//
	smf.citekey = GenInfo("citekey")
	smf.author = GenInfo("author")
	smf.title = GenInfo("title")
	smf.year = GenInfo("year")
	smf.publisher = GenInfo("publisher")
	smf.volume = GenInfo("volume")
	smf.series = GenInfo("series")
	smf.address = GenInfo("address")
	smf.edition = GenInfo("edition")
	smf.month = GenInfo("month")
	smf.note = GenInfo("note")
	smf.typ = GenInfo("type")
	// writes the information to file
	bibliography.WriteString("\n@inbook{")
	bibliography.WriteString(smf.citekey)
	bibliography.WriteString(",\n")
	if smf.author != "" {
		bibliography.WriteString("author = ")
		bibliography.WriteString("{")
		bibliography.WriteString(smf.author)
		bibliography.WriteString(CloseTag())
	}
	if smf.title != "" {
		bibliography.WriteString("title = ")
		bibliography.WriteString("{")
		bibliography.WriteString(smf.title)
		bibliography.WriteString(CloseTag())
	}
	if smf.year != "" {
		bibliography.WriteString("year = ")
		bibliography.WriteString("{")
		bibliography.WriteString(smf.year)
		bibliography.WriteString(CloseTag())
	}
	if smf.pages != "" {
		bibliography.WriteString("pages = ")
		bibliography.WriteString("{")
		bibliography.WriteString(smf.pages)
		bibliography.WriteString(CloseTag())
	}
	if smf.chapter != "" {
		bibliography.WriteString("chapter = ")
		bibliography.WriteString("{")
		bibliography.WriteString(smf.chapter)
		bibliography.WriteString(CloseTag())
	}
	if smf.editor != "" {
		bibliography.WriteString("editor = ")
		bibliography.WriteString("{")
		bibliography.WriteString(smf.editor)
		bibliography.WriteString(CloseTag())
	}
	if smf.publisher != "" {
		bibliography.WriteString("publisher = ")
		bibliography.WriteString("{")
		bibliography.WriteString(smf.publisher)
		bibliography.WriteString(CloseTag())
	}
	if smf.volume != "" {
		bibliography.WriteString("volume = ")
		bibliography.WriteString("{")
		bibliography.WriteString(smf.volume)
		bibliography.WriteString(CloseTag())
	}
	if smf.series != "" {
		bibliography.WriteString("series = ")
		bibliography.WriteString("{")
		bibliography.WriteString(smf.series)
		bibliography.WriteString(CloseTag())
	}
	if smf.address != "" {
		bibliography.WriteString("address = ")
		bibliography.WriteString("{")
		bibliography.WriteString(smf.address)
		bibliography.WriteString(CloseTag())
	}
	if smf.edition != "" {
		bibliography.WriteString("edition = ")
		bibliography.WriteString("{")
		bibliography.WriteString(smf.edition)
		bibliography.WriteString(CloseTag())
	}
	if smf.month != "" {
		bibliography.WriteString("month = ")
		bibliography.WriteString("{")
		bibliography.WriteString(smf.month)
		bibliography.WriteString(CloseTag())
	}
	if smf.note != "" {
		bibliography.WriteString("note = ")
		bibliography.WriteString("{")
		bibliography.WriteString(smf.note)
		bibliography.WriteString(CloseTag())
	}
	if smf.typ != "" {
		bibliography.WriteString("type = ")
		bibliography.WriteString("{")
		bibliography.WriteString(smf.typ)
		bibliography.WriteString(CloseTag())
	}
	bibliography.WriteString("}\n")
}

func AddBooklet() {
	fmt.Println("####### Enter new BOOKLET entry #######")
	//Open bibliography//
	bibliography, err := os.OpenFile(viper.GetString("bibliography"), os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer bibliography.Close()
	smf := new(Booklet)
	//Gathers all the information//
	smf.citekey = GenInfo("citekey")
	smf.author = GenInfo("author")
	smf.title = GenInfo("title")
	smf.year = GenInfo("year")
	smf.howpublished = GenInfo("how published")
	smf.address = GenInfo("address")
	smf.month = GenInfo("month")
	smf.note = GenInfo("note")
	// writes the information to file
	bibliography.WriteString("\n@booklet{")
	bibliography.WriteString(smf.citekey)
	bibliography.WriteString(",\n")
	if smf.author != "" {
		bibliography.WriteString("author = ")
		bibliography.WriteString("{")
		bibliography.WriteString(smf.author)
		bibliography.WriteString(CloseTag())
	}
	if smf.title != "" {
		bibliography.WriteString("title = ")
		bibliography.WriteString("{")
		bibliography.WriteString(smf.title)
		bibliography.WriteString(CloseTag())
	}
	if smf.year != "" {
		bibliography.WriteString("year = ")
		bibliography.WriteString("{")
		bibliography.WriteString(smf.year)
		bibliography.WriteString(CloseTag())
	}
	if smf.howpublished != "" {
		bibliography.WriteString("howpublished = ")
		bibliography.WriteString("{")
		bibliography.WriteString(smf.howpublished)
		bibliography.WriteString(CloseTag())
	}
	if smf.address != "" {
		bibliography.WriteString("address = ")
		bibliography.WriteString("{")
		bibliography.WriteString(smf.address)
		bibliography.WriteString(CloseTag())
	}
	if smf.month != "" {
		bibliography.WriteString("month = ")
		bibliography.WriteString("{")
		bibliography.WriteString(smf.month)
		bibliography.WriteString(CloseTag())
	}
	if smf.note != "" {
		bibliography.WriteString("note = ")
		bibliography.WriteString("{")
		bibliography.WriteString(smf.note)
		bibliography.WriteString(CloseTag())
	}
	bibliography.WriteString("}\n")
}

func AddUnpublished() {
	fmt.Println("####### Enter new UNPUBLISHED entry #######")
	//Open bibliography//
	bibliography, err := os.OpenFile(viper.GetString("bibliography"), os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer bibliography.Close()
	smf := new(Unpublished)
	//Gathers all the information//
	smf.citekey = GenInfo("citekey")
	smf.author = GenInfo("author")
	smf.title = GenInfo("title")
	smf.year = GenInfo("year")
	smf.month = GenInfo("month")
	smf.note = GenInfo("note")
	// writes the information to file
	bibliography.WriteString("\n@unpublished{")
	bibliography.WriteString(smf.citekey)
	bibliography.WriteString(",\n")
	if smf.author != "" {
		bibliography.WriteString("author = ")
		bibliography.WriteString("{")
		bibliography.WriteString(smf.author)
		bibliography.WriteString(CloseTag())
	}
	if smf.title != "" {
		bibliography.WriteString("title = ")
		bibliography.WriteString("{")
		bibliography.WriteString(smf.title)
		bibliography.WriteString(CloseTag())
	}
	if smf.year != "" {
		bibliography.WriteString("year = ")
		bibliography.WriteString("{")
		bibliography.WriteString(smf.year)
		bibliography.WriteString(CloseTag())
	}
	if smf.month != "" {
		bibliography.WriteString("month = ")
		bibliography.WriteString("{")
		bibliography.WriteString(smf.month)
		bibliography.WriteString(CloseTag())
	}
	if smf.note != "" {
		bibliography.WriteString("note = ")
		bibliography.WriteString("{")
		bibliography.WriteString(smf.note)
		bibliography.WriteString(CloseTag())
	}
	bibliography.WriteString("}\n")
}

func AddInCollection() {
	fmt.Println("####### Enter new InBOOK entry #######")
	//Open bibliography//
	bibliography, err := os.OpenFile(viper.GetString("bibliography"), os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer bibliography.Close()
	smf := new(InCollection)
	//Gathers all the information//
	smf.citekey = GenInfo("citekey")
	smf.author = GenInfo("author")
	smf.title = GenInfo("title")
	smf.year = GenInfo("year")
	smf.publisher = GenInfo("publisher")
	smf.booktitle = GenInfo("book title")
	smf.editor = GenInfo("editor")
	smf.volume = GenInfo("volume")
	smf.series = GenInfo("series")
	smf.address = GenInfo("address")
	smf.edition = GenInfo("edition")
	smf.month = GenInfo("month")
	smf.note = GenInfo("note")
	smf.typ = GenInfo("type")
	smf.chapter = GenInfo("chapter")
	smf.pages = GenInfo("pages")
	// writes the information to file
	bibliography.WriteString("\n@incollection{")
	bibliography.WriteString(smf.citekey)
	bibliography.WriteString(",\n")
	if smf.author != "" {
		bibliography.WriteString("author = ")
		bibliography.WriteString("{")
		bibliography.WriteString(smf.author)
		bibliography.WriteString(CloseTag())
	}
	if smf.title != "" {
		bibliography.WriteString("title = ")
		bibliography.WriteString("{")
		bibliography.WriteString(smf.title)
		bibliography.WriteString(CloseTag())
	}
	if smf.year != "" {
		bibliography.WriteString("year = ")
		bibliography.WriteString("{")
		bibliography.WriteString(smf.year)
		bibliography.WriteString(CloseTag())
	}
	if smf.booktitle != "" {
		bibliography.WriteString("booktitle = ")
		bibliography.WriteString("{")
		bibliography.WriteString(smf.booktitle)
		bibliography.WriteString(CloseTag())
	}
	if smf.editor != "" {
		bibliography.WriteString("editor = ")
		bibliography.WriteString("{")
		bibliography.WriteString(smf.editor)
		bibliography.WriteString(CloseTag())
	}
	if smf.pages != "" {
		bibliography.WriteString("pages = ")
		bibliography.WriteString("{")
		bibliography.WriteString(smf.pages)
		bibliography.WriteString(CloseTag())
	}
	if smf.chapter != "" {
		bibliography.WriteString("chapter = ")
		bibliography.WriteString("{")
		bibliography.WriteString(smf.chapter)
		bibliography.WriteString(CloseTag())
	}
	if smf.editor != "" {
		bibliography.WriteString("editor = ")
		bibliography.WriteString("{")
		bibliography.WriteString(smf.editor)
		bibliography.WriteString(CloseTag())
	}
	if smf.publisher != "" {
		bibliography.WriteString("publisher = ")
		bibliography.WriteString("{")
		bibliography.WriteString(smf.publisher)
		bibliography.WriteString(CloseTag())
	}
	if smf.volume != "" {
		bibliography.WriteString("volume = ")
		bibliography.WriteString("{")
		bibliography.WriteString(smf.volume)
		bibliography.WriteString(CloseTag())
	}
	if smf.series != "" {
		bibliography.WriteString("series = ")
		bibliography.WriteString("{")
		bibliography.WriteString(smf.series)
		bibliography.WriteString(CloseTag())
	}
	if smf.address != "" {
		bibliography.WriteString("address = ")
		bibliography.WriteString("{")
		bibliography.WriteString(smf.address)
		bibliography.WriteString(CloseTag())
	}
	if smf.edition != "" {
		bibliography.WriteString("edition = ")
		bibliography.WriteString("{")
		bibliography.WriteString(smf.edition)
		bibliography.WriteString(CloseTag())
	}
	if smf.month != "" {
		bibliography.WriteString("month = ")
		bibliography.WriteString("{")
		bibliography.WriteString(smf.month)
		bibliography.WriteString(CloseTag())
	}
	if smf.note != "" {
		bibliography.WriteString("note = ")
		bibliography.WriteString("{")
		bibliography.WriteString(smf.note)
		bibliography.WriteString(CloseTag())
	}
	if smf.typ != "" {
		bibliography.WriteString("type = ")
		bibliography.WriteString("{")
		bibliography.WriteString(smf.typ)
		bibliography.WriteString(CloseTag())
	}
	bibliography.WriteString("}\n")
}
