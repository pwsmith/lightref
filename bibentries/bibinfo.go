package bibinfo

import (
	"bufio"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
)

//func GetBiblio() {
//	bib, _ := addCmd.Flags().GetString("bibliography")
//	if bib != "" {
//		fmt.Println("Bibliography =", bib)
//		viper.Set("bibliography", bib)
//	}
//	fmt.Println("File chosen by viper is", viper.GetString("bibliography"))
//}
func Add_ArticleS() {
	bibliography, err := os.OpenFile(viper.GetString("bibliography"), os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}

	smf := new(Article)

	smf.Citekey = GenInfo("citekey")
	smf.Author = GenInfo("author")
	smf.Title = GenInfo("title")
	smf.Year = GenInfo("year")
	smf.Journal = GenInfo("journal")
	smf.Volume = GenInfo("volume")
	smf.Number = GenInfo("number")
	smf.Pages = GenInfo("pages")

	bibliography.WriteString("\n@article{")
	bibliography.WriteString(smf.Citekey)
	bibliography.WriteString(",\n")
	bibliography.WriteString(WriteData("citekey", smf.Citekey))
	if smf.Author != "" {
		bibliography.WriteString(WriteData("author", smf.Author))
	}
	if smf.Title != "" {
		bibliography.WriteString(WriteData("title", smf.Title))
	}
	if smf.Year != "" {
		bibliography.WriteString(WriteData("year", smf.Year))
	}
	if smf.Journal != "" {
		bibliography.WriteString(WriteData("journal", smf.Journal))
	}
	if smf.Volume != "" {
		bibliography.WriteString(WriteData("volume", smf.Volume))
	}
	if smf.Number != "" {
		bibliography.WriteString(WriteData("number", smf.Number))
	}
	if smf.Pages != "" {
		bibliography.WriteString(WriteData("pages", smf.Pages))
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
	viper.Set("bibliography", "/home/pwsmith/otherbib.bib")
	bibliography, err := os.OpenFile(viper.GetString("bibliography"), os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer bibliography.Close()
	smf := new(Book)
	//Gathers all the information//
	smf.Citekey = GenInfo("citekey")
	smf.Author = GenInfo("author")
	smf.Title = GenInfo("title")
	smf.Year = GenInfo("year")
	smf.Publisher = GenInfo("publisher")
	smf.Volume = GenInfo("volume")
	smf.Series = GenInfo("series")
	smf.Address = GenInfo("address")
	smf.Edition = GenInfo("edition")
	smf.Month = GenInfo("month")
	smf.Note = GenInfo("note")
	// writes the information to file
	bibliography.WriteString("\n@book{")
	bibliography.WriteString(smf.Citekey)
	bibliography.WriteString(",\n")
	if smf.Author != "" {
		bibliography.WriteString(WriteData("Author", smf.Author))
	}
	if smf.Title != "" {
		bibliography.WriteString(WriteData("Title", smf.Title))
	}
	if smf.Year != "" {
		bibliography.WriteString(WriteData("Year", smf.Year))
	}
	if smf.Volume != "" {
		bibliography.WriteString(WriteData("Volume", smf.Volume))
	}
	if smf.Publisher != "" {
		bibliography.WriteString(WriteData("Publisher", smf.Publisher))
	}
	if smf.Series != "" {
		bibliography.WriteString(WriteData("Series", smf.Series))
	}
	if smf.Address != "" {
		bibliography.WriteString(WriteData("Address", smf.Address))
	}
	if smf.Edition != "" {
		bibliography.WriteString(WriteData("Edition", smf.Edition))
	}
	if smf.Month != "" {
		bibliography.WriteString(WriteData("Month", smf.Month))
	}
	if smf.Note != "" {
		bibliography.WriteString(WriteData("Note", smf.Note))
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
	smf.Citekey = GenInfo("citekey")
	smf.Author = GenInfo("author")
	smf.Title = GenInfo("title")
	smf.School = GenInfo("school")
	smf.Year = GenInfo("year")
	smf.Address = GenInfo("address")
	smf.Month = GenInfo("month")
	smf.Note = GenInfo("note")
	smf.Typ = GenInfo("type")
	// writes the information to file
	bibliography.WriteString("\n@phdthesis{")
	bibliography.WriteString(smf.Citekey)
	bibliography.WriteString(",\n")
	if smf.Author != "" {
		bibliography.WriteString(WriteData("Author", smf.Author))
	}
	if smf.Title != "" {
		bibliography.WriteString(WriteData("Title", smf.Title))
	}
	if smf.School != "" {
		bibliography.WriteString(WriteData("School", smf.School))
	}
	if smf.Year != "" {
		bibliography.WriteString(WriteData("Year", smf.Year))
	}
	if smf.Address != "" {
		bibliography.WriteString(WriteData("Address", smf.Address))
	}
	if smf.Month != "" {
		bibliography.WriteString(WriteData("Month", smf.Month))
	}
	if smf.Note != "" {
		bibliography.WriteString(WriteData("Note", smf.Note))
	}
	if smf.Typ != "" {
		bibliography.WriteString(WriteData("Type", smf.Typ))
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
	smf.Citekey = GenInfo("citekey")
	smf.Author = GenInfo("author")
	smf.Title = GenInfo("title")
	smf.Year = GenInfo("year")
	smf.Publisher = GenInfo("publisher")
	smf.Volume = GenInfo("volume")
	smf.Series = GenInfo("series")
	smf.Address = GenInfo("address")
	smf.Edition = GenInfo("edition")
	smf.Month = GenInfo("month")
	smf.Note = GenInfo("note")
	smf.Typ = GenInfo("type")
	// writes the information to file
	bibliography.WriteString("\n@inbook{")
	bibliography.WriteString(smf.Citekey)
	bibliography.WriteString(",\n")
	if smf.Author != "" {
		bibliography.WriteString(WriteData("Author", smf.Author))
	}
	if smf.Title != "" {
		bibliography.WriteString(WriteData("Title", smf.Title))
	}
	if smf.Year != "" {
		bibliography.WriteString(WriteData("Year", smf.Year))
	}
	if smf.Publisher != "" {
		bibliography.WriteString(WriteData("Publisher", smf.Publisher))
	}
	if smf.Volume != "" {
		bibliography.WriteString(WriteData("Volume", smf.Volume))
	}
	if smf.Series != "" {
		bibliography.WriteString(WriteData("Series", smf.Series))
	}
	if smf.Address != "" {
		bibliography.WriteString(WriteData("Address", smf.Address))
	}
	if smf.Edition != "" {
		bibliography.WriteString(WriteData("Edition", smf.Edition))
	}
	if smf.Month != "" {
		bibliography.WriteString(WriteData("Month", smf.Month))
	}
	if smf.Note != "" {
		bibliography.WriteString(WriteData("Note", smf.Note))
	}
	if smf.Typ != "" {
		bibliography.WriteString(WriteData("Type", smf.Typ))
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
	smf.Citekey = GenInfo("citekey")
	smf.Author = GenInfo("author")
	smf.Title = GenInfo("title")
	smf.Year = GenInfo("year")
	smf.Howpublished = GenInfo("how published")
	smf.Address = GenInfo("address")
	smf.Month = GenInfo("month")
	smf.Note = GenInfo("note")
	// writes the information to file
	bibliography.WriteString("\n@booklet{")
	bibliography.WriteString(smf.Citekey)
	bibliography.WriteString(",\n")
	if smf.Author != "" {
		bibliography.WriteString(WriteData("Author", smf.Author))
	}
	if smf.Title != "" {
		bibliography.WriteString(WriteData("Title", smf.Title))
	}
	if smf.Year != "" {
		bibliography.WriteString(WriteData("Year", smf.Year))
	}
	if smf.Howpublished != "" {
		bibliography.WriteString(WriteData("Howpublished", smf.Howpublished))
	}
	if smf.Address != "" {
		bibliography.WriteString(WriteData("Address", smf.Address))
	}
	if smf.Month != "" {
		bibliography.WriteString(WriteData("Month", smf.Month))
	}
	if smf.Note != "" {
		bibliography.WriteString(WriteData("Note", smf.Note))
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
	smf.Citekey = GenInfo("citekey")
	smf.Author = GenInfo("author")
	smf.Title = GenInfo("title")
	smf.Year = GenInfo("year")
	smf.Month = GenInfo("month")
	smf.Note = GenInfo("note")
	// writes the information to file
	bibliography.WriteString("\n@unpublished{")
	bibliography.WriteString(smf.Citekey)
	bibliography.WriteString(",\n")
	if smf.Author != "" {
		bibliography.WriteString(WriteData("author", smf.Author))
	}
	if smf.Title != "" {
		bibliography.WriteString(WriteData("title", smf.Title))
	}
	if smf.Year != "" {
		bibliography.WriteString(WriteData("year", smf.Year))
	}
	if smf.Month != "" {
		bibliography.WriteString(WriteData("month", smf.Month))
	}
	if smf.Note != "" {
		bibliography.WriteString(WriteData("note", smf.Note))
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
	smf.Citekey = GenInfo("citekey")
	smf.Author = GenInfo("author")
	smf.Title = GenInfo("title")
	smf.Year = GenInfo("year")
	smf.Publisher = GenInfo("publisher")
	smf.Booktitle = GenInfo("book title")
	smf.Editor = GenInfo("editor")
	smf.Volume = GenInfo("volume")
	smf.Series = GenInfo("series")
	smf.Address = GenInfo("address")
	smf.Edition = GenInfo("edition")
	smf.Month = GenInfo("month")
	smf.Note = GenInfo("note")
	smf.Typ = GenInfo("type")
	smf.Chapter = GenInfo("chapter")
	smf.Pages = GenInfo("pages")
	// writes the information to file
	bibliography.WriteString("\n@incollection{")
	bibliography.WriteString(smf.Citekey)
	bibliography.WriteString(",\n")
	if smf.Author != "" {
		bibliography.WriteString(WriteData("Author", smf.Author))
	}
	if smf.Title != "" {
		bibliography.WriteString(WriteData("Title", smf.Title))
	}
	if smf.Year != "" {
		bibliography.WriteString(WriteData("Year", smf.Year))
	}
	if smf.Booktitle != "" {
		bibliography.WriteString(WriteData("Booktitle", smf.Booktitle))
	}
	if smf.Publisher != "" {
		bibliography.WriteString(WriteData("Publisher", smf.Publisher))
	}
	if smf.Editor != "" {
		bibliography.WriteString(WriteData("Editor", smf.Editor))
	}
	if smf.Volume != "" {
		bibliography.WriteString(WriteData("Volume", smf.Volume))
	}
	if smf.Series != "" {
		bibliography.WriteString(WriteData("Series", smf.Series))
	}
	if smf.Address != "" {
		bibliography.WriteString(WriteData("Address", smf.Address))
	}
	if smf.Edition != "" {
		bibliography.WriteString(WriteData("Edition", smf.Edition))
	}
	if smf.Month != "" {
		bibliography.WriteString(WriteData("Month", smf.Month))
	}
	if smf.Note != "" {
		bibliography.WriteString(WriteData("Note", smf.Note))
	}
	if smf.Typ != "" {
		bibliography.WriteString(WriteData("Type", smf.Typ))
	}
	if smf.Chapter != "" {
		bibliography.WriteString(WriteData("Chapter", smf.Chapter))
	}
	if smf.Pages != "" {
		bibliography.WriteString(WriteData("Pages", smf.Pages))
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
	smf.Citekey = GenInfo("citekey")
	smf.Author = GenInfo("author")
	smf.Title = GenInfo("title")
	smf.Year = GenInfo("year")
	smf.Publisher = GenInfo("publisher")
	smf.Booktitle = GenInfo("book title")
	smf.Editor = GenInfo("editor")
	smf.Volume = GenInfo("volume")
	smf.Series = GenInfo("series")
	smf.Address = GenInfo("address")
	smf.Month = GenInfo("month")
	smf.Organization = GenInfo("Organization")
	smf.Note = GenInfo("note")
	smf.Typ = GenInfo("type")
	smf.Pages = GenInfo("pages")
	// writes the information to file
	bibliography.WriteString("\n@inproceedings{")
	bibliography.WriteString(smf.Citekey)
	bibliography.WriteString(",\n")
	if smf.Author != "" {
		bibliography.WriteString(WriteData("Author", smf.Author))
	}
	if smf.Title != "" {
		bibliography.WriteString(WriteData("Title", smf.Title))
	}
	if smf.Year != "" {
		bibliography.WriteString(WriteData("Year", smf.Year))
	}
	if smf.Publisher != "" {
		bibliography.WriteString(WriteData("Publisher", smf.Publisher))
	}
	if smf.Booktitle != "" {
		bibliography.WriteString(WriteData("Booktitle", smf.Booktitle))
	}
	if smf.Editor != "" {
		bibliography.WriteString(WriteData("Editor", smf.Editor))
	}
	if smf.Volume != "" {
		bibliography.WriteString(WriteData("Volume", smf.Volume))
	}
	if smf.Series != "" {
		bibliography.WriteString(WriteData("Series", smf.Series))
	}
	if smf.Address != "" {
		bibliography.WriteString(WriteData("Address", smf.Address))
	}
	if smf.Month != "" {
		bibliography.WriteString(WriteData("Month", smf.Month))
	}
	if smf.Organization != "" {
		bibliography.WriteString(WriteData("Organization", smf.Organization))
	}
	if smf.Note != "" {
		bibliography.WriteString(WriteData("Note", smf.Note))
	}
	if smf.Pages != "" {
		bibliography.WriteString(WriteData("Pages", smf.Pages))
	}
	if smf.Typ != "" {
		bibliography.WriteString(WriteData("Type", smf.Typ))
	}
	bibliography.WriteString("}\n")
}

func AddManual() {
	fmt.Println("####### Enter new MANUAL entry #######")
	//Open bibliography//
	bibliography, err := os.OpenFile(viper.GetString("bibliography"), os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer bibliography.Close()
	smf := new(Manual)
	smf.Citekey = GenInfo("Citekey")
	smf.Title = GenInfo("Title")
	smf.Author = GenInfo("Author")
	smf.Organization = GenInfo("Organization")
	smf.Address = GenInfo("Address")
	smf.Edition = GenInfo("Edition")
	smf.Month = GenInfo("Month")
	smf.Year = GenInfo("Year")
	smf.Note = GenInfo("Note")
	// writes the information to file
	bibliography.WriteString("\n@manual{")
	bibliography.WriteString(smf.Citekey)
	bibliography.WriteString(",\n")
	if smf.Title != "" {
		bibliography.WriteString(WriteData("Title", smf.Title))
	}
	if smf.Author != "" {
		bibliography.WriteString(WriteData("Author", smf.Author))
	}
	if smf.Organization != "" {
		bibliography.WriteString(WriteData("Organization", smf.Organization))
	}
	if smf.Address != "" {
		bibliography.WriteString(WriteData("Address", smf.Address))
	}
	if smf.Edition != "" {
		bibliography.WriteString(WriteData("Edition", smf.Edition))
	}
	if smf.Month != "" {
		bibliography.WriteString(WriteData("Month", smf.Month))
	}
	if smf.Year != "" {
		bibliography.WriteString(WriteData("Year", smf.Year))
	}
	if smf.Note != "" {
		bibliography.WriteString(WriteData("Note", smf.Note))
	}
	bibliography.WriteString("}\n")
}
func AddMAThesis() {
	fmt.Println("####### Enter new MA THESIS entry #######")
	//Open bibliography//
	bibliography, err := os.OpenFile(viper.GetString("bibliography"), os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer bibliography.Close()
	smf := new(Mathesis)
	smf.Author = GenInfo("Author")
	smf.Title = GenInfo("Title")
	smf.School = GenInfo("School")
	smf.Year = GenInfo("Year")
	smf.Typ = GenInfo("Type")
	smf.Address = GenInfo("Address")
	smf.Month = GenInfo("Month")
	smf.Note = GenInfo("Note")
	// writes the information to file
	bibliography.WriteString("\n@manual{")
	bibliography.WriteString(smf.Citekey)
	bibliography.WriteString(",\n")
	if smf.Author != "" {
		bibliography.WriteString(WriteData("Author", smf.Author))
	}
	if smf.Title != "" {
		bibliography.WriteString(WriteData("Title", smf.Title))
	}
	if smf.School != "" {
		bibliography.WriteString(WriteData("School", smf.School))
	}
	if smf.Year != "" {
		bibliography.WriteString(WriteData("Year", smf.Year))
	}
	if smf.Typ != "" {
		bibliography.WriteString(WriteData("Type", smf.Typ))
	}
	if smf.Address != "" {
		bibliography.WriteString(WriteData("Address", smf.Address))
	}
	if smf.Month != "" {
		bibliography.WriteString(WriteData("Month", smf.Month))
	}
	if smf.Note != "" {
		bibliography.WriteString(WriteData("Note", smf.Note))
	}
	bibliography.WriteString("}\n")
}
