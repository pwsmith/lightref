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

func BibliographyCheck() {
	selected_bibliography := viper.GetString("bibliography")
	if _, err := os.Stat(selected_bibliography); os.IsNotExist(err) {
		fmt.Printf("Selected bibliography does not exist, creating %s\n", selected_bibliography)
		os.Create(selected_bibliography)
	} else {
		fmt.Printf("Found %s\n", selected_bibliography)
	}
}

func Add_ArticleS() {
	//	selected_bibliography := viper.GetString("bibliography")
	//	if _, err := os.Stat(selected_bibliography); os.IsNotExist(err) {
	//		fmt.Printf("Selected bibliography does not exist, creating %s\n", selected_bibliography)
	//		os.Create(selected_bibliography)
	//	} else {
	//		fmt.Printf("Found %s\n", selected_bibliography)
	//	}
	BibliographyCheck()
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
	if smf.Author != "" {
		bibliography.WriteString(WriteData("Author", smf.Author))
	}
	if smf.Title != "" {
		bibliography.WriteString(WriteData("Title", smf.Title))
	}
	if smf.Year != "" {
		bibliography.WriteString(WriteData("Year", smf.Year))
	}
	if smf.Journal != "" {
		bibliography.WriteString(WriteData("Journal", smf.Journal))
	}
	if smf.Volume != "" {
		bibliography.WriteString(WriteData("Volume", smf.Volume))
	}
	if smf.Number != "" {
		bibliography.WriteString(WriteData("Number", smf.Number))
	}
	if smf.Pages != "" {
		bibliography.WriteString(WriteData("Pages", smf.Pages))
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
	BibliographyCheck()
	fmt.Println("####### Enter new BOOK entry #######")
	//Open bibliography//
	bibliography, err := os.OpenFile(viper.GetString("bibliography"), os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer bibliography.Close()
	smf := new(Book)
	//Gathers all the information//
	smf.Citekey = GenInfo("citekey")
	smf.Author = GenInfo("author")
	smf.Editor = GenInfo("editor")
	smf.Title = GenInfo("title")
	smf.Year = GenInfo("year")
	smf.Publisher = GenInfo("publisher")
	smf.Volume = GenInfo("volume")
	smf.Number = GenInfo("number")
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
	if smf.Editor != "" {
		bibliography.WriteString(WriteData("Editor", smf.Editor))
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
	if smf.Number != "" {
		bibliography.WriteString(WriteData("Number", smf.Number))
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
	BibliographyCheck()
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
	BibliographyCheck()
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
	BibliographyCheck()
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
	BibliographyCheck()
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
		bibliography.WriteString(WriteData("Author", smf.Author))
	}
	if smf.Title != "" {
		bibliography.WriteString(WriteData("Title", smf.Title))
	}
	if smf.Year != "" {
		bibliography.WriteString(WriteData("Year", smf.Year))
	}
	if smf.Month != "" {
		bibliography.WriteString(WriteData("Month", smf.Month))
	}
	if smf.Note != "" {
		bibliography.WriteString(WriteData("Note", smf.Note))
	}
	bibliography.WriteString("}\n")
}

func AddInCollection() {
	BibliographyCheck()
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
	BibliographyCheck()
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
	BibliographyCheck()
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
	BibliographyCheck()
	fmt.Println("####### Enter new MA THESIS entry #######")
	//Open bibliography//
	bibliography, err := os.OpenFile(viper.GetString("bibliography"), os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer bibliography.Close()
	smf := new(Mathesis)
	smf.Citekey = GenInfo("Citekey")
	smf.Author = GenInfo("Author")
	smf.Title = GenInfo("Title")
	smf.School = GenInfo("School")
	smf.Year = GenInfo("Year")
	smf.Typ = GenInfo("Type")
	smf.Address = GenInfo("Address")
	smf.Month = GenInfo("Month")
	smf.Note = GenInfo("Note")
	// writes the information to file
	bibliography.WriteString("\n@mastersthesis{")
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

func Add_Misc() {
	BibliographyCheck()
	fmt.Println("####### Enter new MISC entry #######")
	//Open bibliography//
	bibliography, err := os.OpenFile(viper.GetString("bibliography"), os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer bibliography.Close()
	smf := new(Misc)
	smf.Citekey = GenInfo("Citekey")
	smf.Author = GenInfo("Author")
	smf.Title = GenInfo("Title")
	smf.Year = GenInfo("Year")
	smf.Howpublished = GenInfo("Howpublished")
	smf.Month = GenInfo("Month")
	smf.Note = GenInfo("Note")
	// writes the information to file
	bibliography.WriteString("\n@misc{")
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
	if smf.Month != "" {
		bibliography.WriteString(WriteData("Month", smf.Month))
	}
	if smf.Note != "" {
		bibliography.WriteString(WriteData("Note", smf.Note))
	}
	bibliography.WriteString("}\n")
}

func AddProceedings() {
	BibliographyCheck()
	fmt.Println("####### Enter new PROCEEDINGS entry #######")
	//Open bibliography//
	bibliography, err := os.OpenFile(viper.GetString("bibliography"), os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer bibliography.Close()
	smf := new(Proceedings)
	// writes the information to file

	smf.Citekey = GenInfo("Citekey")
	smf.Title = GenInfo("Title")
	smf.Year = GenInfo("Year")
	smf.Editor = GenInfo("Editor")
	smf.Volume = GenInfo("Volume")
	smf.Number = GenInfo("Number")
	smf.Address = GenInfo("Address")
	smf.Month = GenInfo("Month")
	smf.Organization = GenInfo("Organization")
	smf.Publisher = GenInfo("Publisher")
	smf.Note = GenInfo("Note")

	bibliography.WriteString("\n@misc{")
	bibliography.WriteString(smf.Citekey)
	bibliography.WriteString(",\n")
	if smf.Title != "" {
		bibliography.WriteString(WriteData("Title", smf.Title))
	}
	if smf.Year != "" {
		bibliography.WriteString(WriteData("Year", smf.Year))
	}
	if smf.Editor != "" {
		bibliography.WriteString(WriteData("Editor", smf.Editor))
	}
	if smf.Volume != "" {
		bibliography.WriteString(WriteData("Volume", smf.Volume))
	}
	if smf.Number != "" {
		bibliography.WriteString(WriteData("Number", smf.Number))
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
	if smf.Publisher != "" {
		bibliography.WriteString(WriteData("Publisher", smf.Publisher))
	}
	if smf.Note != "" {
		bibliography.WriteString(WriteData("Note", smf.Note))
	}
	bibliography.WriteString("}\n")
}

func AddTechReport() {
	BibliographyCheck()
	fmt.Println("####### Enter new TECH REPORT entry #######")
	//Open bibliography//
	bibliography, err := os.OpenFile(viper.GetString("bibliography"), os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer bibliography.Close()
	smf := new(Techreport)
	smf.Citekey = GenInfo("Citekey")
	smf.Author = GenInfo("Author")
	smf.Title = GenInfo("Title")
	smf.Institution = GenInfo("Institution")
	smf.Year = GenInfo("Year")
	smf.Typ = GenInfo("Type")
	smf.Number = GenInfo("Number")
	smf.Address = GenInfo("Address")
	smf.Month = GenInfo("Month")
	smf.Note = GenInfo("Note")

	// writes the information to file
	bibliography.WriteString("\n@techreport{")
	bibliography.WriteString(smf.Citekey)
	bibliography.WriteString(",\n")
	if smf.Author != "" {
		bibliography.WriteString(WriteData("Author", smf.Author))
	}
	if smf.Title != "" {
		bibliography.WriteString(WriteData("Title", smf.Title))
	}
	if smf.Institution != "" {
		bibliography.WriteString(WriteData("Institution", smf.Institution))
	}
	if smf.Year != "" {
		bibliography.WriteString(WriteData("Year", smf.Year))
	}
	if smf.Typ != "" {
		bibliography.WriteString(WriteData("Type", smf.Typ))
	}
	if smf.Number != "" {
		bibliography.WriteString(WriteData("Number", smf.Number))
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

func RemoveType(x string) string {
	var Removed string
	if strings.Contains(x, "@article") {
		Removed = strings.Replace(x, "@article", "Citekey = ", -1)
	} else if strings.Contains(x, "@book") {
		Removed = strings.Replace(x, "@book", "Citekey = ", -1)
	} else if strings.Contains(x, "@conference") {
		Removed = strings.Replace(x, "@conference", "Citekey = ", -1)
	} else if strings.Contains(x, "@inbook") {
		Removed = strings.Replace(x, "@inbook", "Citekey = ", -1)
	} else if strings.Contains(x, "@incollection") {
		Removed = strings.Replace(x, "@incollection", "Citekey = ", -1)
	} else if strings.Contains(x, "@inproceedings") {
		Removed = strings.Replace(x, "@inproceedings", "Citekey = ", -1)
	} else if strings.Contains(x, "@phdthesis") {
		Removed = strings.Replace(x, "@phdthesis", "Citekey = ", -1)
	} else if strings.Contains(x, "@misc") {
		Removed = strings.Replace(x, "@misc", "Citekey = ", -1)
	} else if strings.Contains(x, "@unpublished") {
		Removed = strings.Replace(x, "@unpublished", "Citekey = ", -1)
	} else if strings.Contains(x, "@proceedings") {
		Removed = strings.Replace(x, "@proceedings", "Citekey = ", -1)
	} else if strings.Contains(x, "@manual") {
		Removed = strings.Replace(x, "@manual", "Citekey = ", -1)
	} else if strings.Contains(x, "@techreport") {
		Removed = strings.Replace(x, "@techreport", "Citekey = ", -1)
	} else if strings.Contains(x, "@mastersthesis") {
		Removed = strings.Replace(x, "@mastersthesis", "Citekey = ", -1)
	} else {
		Removed = x
	}
	return Removed
}
