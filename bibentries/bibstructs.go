package bibinfo

type Hello struct {
	cliche string
}

type Article struct {
	Citekey string
	Author  string
	Pages   string
	Title   string
	Year    string
	Journal string
	Volume  string
	Number  string
}

type Unpublished struct {
	Citekey string
	Author  string
	Note    string
	Title   string
	Year    string
	Month   string
}

type Book struct {
	Citekey   string
	Author    string
	Title     string
	Publisher string
	Year      string
	Volume    string
	Series    string
	Address   string
	Edition   string
	Month     string
	Note      string
}

type PhDThesis struct {
	Citekey string
	Author  string
	Title   string
	School  string
	Year    string
	Typ     string
	Address string
	Month   string
	Note    string
}

type InBook struct {
	Citekey   string
	Author    string
	Title     string
	Editor    string
	Chapter   string
	Pages     string
	Publisher string
	Year      string
	Volume    string
	Series    string
	Typ       string
	Address   string
	Edition   string
	Month     string
	Note      string
}

type Booklet struct {
	Citekey      string
	Title        string
	Author       string
	Howpublished string
	Address      string
	Month        string
	Year         string
	Note         string
}

type InCollection struct {
	Citekey   string
	Author    string
	Title     string
	Booktitle string
	Editor    string
	Chapter   string
	Pages     string
	Publisher string
	Year      string
	Volume    string
	Series    string
	Typ       string
	Address   string
	Edition   string
	Month     string
	Note      string
}

type InProceedings struct {
	Citekey      string
	Author       string
	Title        string
	Booktitle    string
	Year         string
	Editor       string
	Volume       string
	Series       string
	Pages        string
	Address      string
	Month        string
	Organization string
	Publisher    string
	Note         string
	Typ          string
}

type Manual struct {
	Citekey      string
	Title        string
	Author       string
	Organization string
	Address      string
	Edition      string
	Month        string
	Year         string
	Note         string
}

type Mathesis struct {
	Citekey string
	Author  string
	Title   string
	School  string
	Year    string
	Typ     string
	Address string
	Month   string
	Note    string
}

type Bibentry struct {
	citekey string
}
