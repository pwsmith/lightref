# Why Lightref?

Lightref is a small, cross platform CLI for adding entries to a bibtex bibliography.
It is designed as a command line based alternative to the GUI bibliograph management alternatives such as Jabref and Bibdesk.
Lightref is **not** intended as a serious competitor to these: they are wonderful programs with vastly more options than Lightref offers.
However, for a simple task such as adding an entry to a bibtex file, opening a full GUI is often overkill, hence Lightref.

# Current Satus (November 2020)

Lightref is very much under development, and as such should not be considered a finished product.
Currently it offers the ability to:

- [X] add article entries
- [X] add book entries
- [X] add booklet entries
- [ ] add conference entries
- [X] add inbook entries
- [X] add incollection entries
- [X] add inproceedings entries
- [X] add manual entries
- [X] add mastersthesis entries
- [X] add misc entries
- [X] add phdthesis entries
- [X] add proceedings entries
- [X] add techreport entries
- [X] add unpublished entries

The remaining classes will be added in the following weeks.

- [X] Provide functionality to print the bibliography in shortened form to output
- [] Allow for a long listing of the bibliography


# Installation

Lightref is written in Go using [Cobra](https://github.com/spf13/cobra).
Since it is still under development there is currently no binary for download.
This will be made available when ready.

It is possible to build from source.
You must have [Go](https://golang.org), installed following the instructions [here](https://golang.org/dl/).
You should then install Cobra following instructions [here](https://github.com/spf13/cobra#installing).
Then, fork the Lightref repo and run:

```
$ cd /path/to/lightref/repo

$ go install lightref
```

# Instructions

## Main Commands

### Add

```bash
lightref add ENTRY_TYPE
```

To add an entry to the bibliography, use the `add` command followed by the relevant bibtex entry type.
For instance:

```bash
lightref add article ## to add an article
lightref add inbook ## to add a book chapter
```

You will then be prompted to add in the information for the fields that are relevant for that type.
For example, if you run `lightref add article`, you will be prompted for:

- citekey
- author
- title
- year
- journal
- volume
- number
- pages

Lightref does **not** enforce the filling out of any fields, so any of them can be left blank if you wish.
Bibtex does have a concept of obligatory fields that should be filled out for each type, however, and will warn the user at run time if they are not entered.

### Print

```bash
lightref print
```

Provides a short summary of each entry in the file, containing the (highlighted) citekey, the Author and the Title.


## Global Flags

`-b /path/to/biblio` | `--bibliography /path/to/biblio`: allows the user to specify an alternative bibliography to the default in the `lightref.yaml` file.

# Features to be added 

- search bibliography and display results to terminal
- citekey checking to avoid duplication: currently there exists nothing to stop you using duplicate citekeys in your bibliography, which is not ideal
