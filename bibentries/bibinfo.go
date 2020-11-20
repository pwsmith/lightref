package bibinfo

import (
	"bufio"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
)

func GetInfo2() string {

	reader2 := bufio.NewReader(os.Stdin)
	fmt.Print("Enter extra: ")
	info, _ := reader2.ReadString('\n')
	info = strings.TrimSuffix(info, "\n")
	return info
}

func Hi() {
	fmt.Println("Hello world from embedded!")
	fmt.Println("This I just added to bibinfo")
}

type Cliche struct {
	firstWord  string
	secondWord string
}

type Unpublished struct {
	citekey string
	author  string
	note    string
	title   string
	year    string
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

//type Article struct {
//	citekey string
//	author  string
//	pages   string
//	title   string
//	year    string
//	journal string
//	volume  string
//	number  string
//}

func GenInfo(x string) string {
	fmt.Println("Enter", x, ":")
	reader2 := bufio.NewReader(os.Stdin)
	info, _ := reader2.ReadString('\n')
	info = strings.TrimSuffix(info, "\n")
	return info
}

//func Add_pages() string {
//	fmt.Println("Enter pages: ")
//	value := GetInfo()
//	return value
//}

func Add_unpublished() {
	fmt.Println("####### Enter new UNPUBLISHED entry #######")
	smf := new(Unpublished)
	smf.citekey = GenInfo("citekey")
	smf.author = GenInfo("Author")
	smf.note = GenInfo("note")
	smf.title = GenInfo("title")
	smf.year = GenInfo("year")

	file, err := os.OpenFile("biblio.bib", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	if _, err := file.WriteString("\n\n@unpublished{" + smf.citekey + ",\n" + "author = {" + smf.author + "},\n" + "note = {" + smf.note + "},\n" + "title = {" + smf.title + "},\n" + "year = {" + smf.year + "},\n}"); err != nil {
		log.Fatal(err)
	}
}

//func AddArticle() string {
//	fmt.Println("####### Enter new ARTICLE entry #######")
//	smf := new(Article)
//	smf.citekey = GenInfo("citekey")
//	smf.author = GenInfo("author")
//	smf.title = GenInfo("title")
//	smf.year = GenInfo("year")
//	smf.journal = GenInfo("journal")
//	smf.volume = GenInfo("volume")
//	smf.number = GenInfo("number")
//	smf.pages = GenInfo("pages")
//
//	//file, err := os.OpenFile("biblio.bib", os.O_APPEND|os.O_WRONLY, 0644)
//	//if err != nil {
//	//	log.Println(err)
//	//}
//	//defer file.Close()
//	if smf.title != "" {
//		fmt.Println("\n", smf.title, "printed as the if clause")
//	} else {
//		fmt.Println("Article title is null")
//	}
//	article_Entry := "\n\n@article{" + smf.citekey + ",\n" + "author = {" + smf.author + "},\n" + "title = {" + smf.title + "},\n" + "journal = {" + smf.journal + "},\n" + "pages = {" + smf.pages + "},\n" + "volume = {" + smf.volume + "},\n" + "number = {" + smf.number + "},\n" + "year = {" + smf.year + "},\n}"
//	return article_Entry
//}

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

func OpenTag() string {
	opentag := "{"
	return opentag
}

func CloseTag() string {
	closetag := "},\n"
	return closetag
}

func AddArticle() {
	fmt.Println("####### Enter new ARTICLE entry #######")
	//Open bibliography//
	bibliography, err := os.OpenFile(viper.GetString("bibliography"), os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer bibliography.Close()
	smf := new(Article)
	//Gathers all the information//
	smf.citekey = GenInfo("citekey")
	smf.author = GenInfo("author")
	smf.title = GenInfo("title")
	smf.year = GenInfo("year")
	smf.journal = GenInfo("journal")
	smf.volume = GenInfo("volume")
	smf.number = GenInfo("number")
	smf.pages = GenInfo("pages")
	// writes the information to file
	bibliography.WriteString("\n@article{")
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
	if smf.journal != "" {
		bibliography.WriteString("journal = ")
		bibliography.WriteString("{")
		bibliography.WriteString(smf.journal)
		bibliography.WriteString(CloseTag())
	}
	if smf.volume != "" {
		bibliography.WriteString("volume = ")
		bibliography.WriteString("{")
		bibliography.WriteString(smf.volume)
		bibliography.WriteString(CloseTag())
	}
	if smf.number != "" {
		bibliography.WriteString("number = ")
		bibliography.WriteString("{")
		bibliography.WriteString(smf.number)
		bibliography.WriteString(CloseTag())
	}
	if smf.pages != "" {
		bibliography.WriteString("pages = ")
		bibliography.WriteString("{")
		bibliography.WriteString(smf.pages)
		bibliography.WriteString(CloseTag())
	}
	bibliography.WriteString("}\n")
}

//type Book struct {
//	author    string
//	title     string
//	publisher string
//	year      string
//	volume    string
//	series    string
//	address   string
//	edition   string
//	month     string
//	note      string
//}

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
	bibliography.WriteString("}\n")
}
