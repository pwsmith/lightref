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

type InProceedings struct {
	citekey      string
	author       string
	title        string
	booktitle    string
	year         string
	editor       string
	volume       string
	series       string
	pages        string
	address      string
	month        string
	organization string
	publisher    string
	note         string
	typ          string
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
		bibliography.WriteString(WriteData("author", smf.author))
	}
	if smf.title != "" {
		bibliography.WriteString(WriteData("title", smf.title))
	}
	if smf.school != "" {
		bibliography.WriteString(WriteData("school", smf.school))
	}
	if smf.year != "" {
		bibliography.WriteString(WriteData("year", smf.year))
	}
	if smf.address != "" {
		bibliography.WriteString(WriteData("address", smf.address))
	}
	if smf.month != "" {
		bibliography.WriteString(WriteData("month", smf.month))
	}
	if smf.note != "" {
		bibliography.WriteString(WriteData("note", smf.note))
	}
	if smf.typ != "" {
		bibliography.WriteString(WriteData("type", smf.typ))
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
		bibliography.WriteString(WriteData("author", smf.author))
	}
	if smf.title != "" {
		bibliography.WriteString(WriteData("title", smf.title))
	}
	if smf.year != "" {
		bibliography.WriteString(WriteData("year", smf.year))
	}
	if smf.publisher != "" {
		bibliography.WriteString(WriteData("publisher", smf.publisher))
	}
	if smf.volume != "" {
		bibliography.WriteString(WriteData("volume", smf.volume))
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
	if smf.typ != "" {
		bibliography.WriteString(WriteData("type", smf.typ))
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
		bibliography.WriteString(WriteData("author", smf.author))
	}
	if smf.title != "" {
		bibliography.WriteString(WriteData("title", smf.title))
	}
	if smf.year != "" {
		bibliography.WriteString(WriteData("year", smf.year))
	}
	if smf.howpublished != "" {
		bibliography.WriteString(WriteData("howpublished", smf.howpublished))
	}
	if smf.address != "" {
		bibliography.WriteString(WriteData("address", smf.address))
	}
	if smf.month != "" {
		bibliography.WriteString(WriteData("month", smf.month))
	}
	if smf.note != "" {
		bibliography.WriteString(WriteData("note", smf.note))
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
		bibliography.WriteString(WriteData("author", smf.author))
	}
	if smf.title != "" {
		bibliography.WriteString(WriteData("title", smf.title))
	}
	if smf.year != "" {
		bibliography.WriteString(WriteData("year", smf.year))
	}
	if smf.month != "" {
		bibliography.WriteString(WriteData("month", smf.month))
	}
	if smf.note != "" {
		bibliography.WriteString(WriteData("note", smf.note))
	}
	bibliography.WriteString("}\n")
}

func AddInCollection() {
	fmt.Println("####### Enter new INCOLLECTION entry #######")
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
		bibliography.WriteString(WriteData("author", smf.author))
	}
	if smf.title != "" {
		bibliography.WriteString(WriteData("title", smf.title))
	}
	if smf.year != "" {
		bibliography.WriteString(WriteData("year", smf.year))
	}
	if smf.booktitle != "" {
		bibliography.WriteString(WriteData("booktitle", smf.booktitle))
	}
	if smf.publisher != "" {
		bibliography.WriteString(WriteData("publisher", smf.publisher))
	}
	if smf.editor != "" {
		bibliography.WriteString(WriteData("editor", smf.editor))
	}
	if smf.volume != "" {
		bibliography.WriteString(WriteData("volume", smf.volume))
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
	if smf.typ != "" {
		bibliography.WriteString(WriteData("type", smf.typ))
	}
	if smf.chapter != "" {
		bibliography.WriteString(WriteData("chapter", smf.chapter))
	}
	if smf.pages != "" {
		bibliography.WriteString(WriteData("pages", smf.pages))
	}
	bibliography.WriteString("}\n")
}

func AddInProceedings() {
	fmt.Println("####### Enter new INPROCEEDINGS entry #######")
	//Open bibliography//
	bibliography, err := os.OpenFile(viper.GetString("bibliography"), os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer bibliography.Close()
	smf := new(InProceedings)
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
	smf.month = GenInfo("month")
	smf.organization = GenInfo("Organization")
	smf.note = GenInfo("note")
	smf.typ = GenInfo("type")
	smf.pages = GenInfo("pages")
	// writes the information to file
	bibliography.WriteString("\n@inproceedings{")
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
	if smf.publisher != "" {
		bibliography.WriteString(WriteData("publisher", smf.publisher))
	}
	if smf.booktitle != "" {
		bibliography.WriteString(WriteData("booktitle", smf.booktitle))
	}
	if smf.editor != "" {
		bibliography.WriteString(WriteData("editor", smf.editor))
	}
	if smf.volume != "" {
		bibliography.WriteString(WriteData("volume", smf.volume))
	}
	if smf.series != "" {
		bibliography.WriteString(WriteData("series", smf.series))
	}
	if smf.address != "" {
		bibliography.WriteString(WriteData("address", smf.address))
	}
	if smf.month != "" {
		bibliography.WriteString(WriteData("month", smf.month))
	}
	if smf.organization != "" {
		bibliography.WriteString(WriteData("organization", smf.organization))
	}
	if smf.note != "" {
		bibliography.WriteString(WriteData("note", smf.note))
	}
	if smf.pages != "" {
		bibliography.WriteString(WriteData("pages", smf.pages))
	}
	if smf.typ != "" {
		bibliography.WriteString(WriteData("type", smf.typ))
	}
}
