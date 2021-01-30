# Why Lightref?

Lightref is a small, cross platform CLI for adding entries to a bibtex bibliography.
It is designed as a command line based alternative to the GUI bibliograph management alternatives such as Jabref and Bibdesk.
Lightref is **not** intended as a serious competitor to these: they are wonderful programs with vastly more options than Lightref offers.
However, for a simple task such as adding an entry to a bibtex file, opening a full GUI is often overkill, hence Lightref.

# Installation

Lightref is written in Go using [Cobra](https://github.com/spf13/cobra) and so is a cross-platform utility.
Precompiled binaries for Linux, Mac and Windows are available on the [releases page](https://github.com/pwsmith/lightref/releases).
Note that the windows binary is **not** tested, as I don't have a windows system at hand to test it on. 
It should work as I don't think there's anything platform specific to cause a problem, but I can't say for sure.
Simply download the relevant binary, make it executable, and place it in your $PATH.

## Linux/Mac

```bash
## to download the binary, or just use the link above
$ wget https://github.com/pwsmith/lightref/releases/download/v0.1.0/lightref-v0.1.0-linux #for linux 
$ wget https://github.com/pwsmith/lightref/releases/download/v0.1.0/lightref-v0.1.0-mac #for mac

## to make the command shorter
$ mv lightref-v0.1.0-linux lightref 

## to make the binary exectuable
$ chmod u+x lightref

## or place it wherever you want
$ sudo mv lightref /usr/local/bin 

## if /usr/local/bin isn't already on your $PATH
$ echo 'export PATH='$PATH:/usr/local/bin' >> ~/.bashrc' 
```

Alternatively, you can build from the source code.
You must have [Go](https://golang.org), installed following the instructions [here](https://golang.org/dl/).
You should then install Cobra following instructions [here](https://github.com/spf13/cobra#installing).
Then, fork the Lightref repo and run:

```
$ cd /path/to/lightref/repo

$ go install lightref
```

# Instructions

## Configuration

A default bibliography can be set in the file `~/.lightref.yaml` on Unix systems.
The only relevant field at this point is the bibliography value, which you should set with the absolute path:

```yaml
bibliography: "/path/to/bibliography/file/on/system"
```

Note that this is **optional**, and the bibliography can be set at the command using the `-b`/`--bibliography` flag as detailed below.

## Main Commands

### Add

```bash
lightref add ENTRY_TYPE
```

To add an entry to the bibliography, use the `add` command followed by the relevant bibtex entry type.
Thus, the following are all add commands:

```bash
lightref add article ## to add an article
lightref add book ## add a book entry
lightref add booklet ## add a booklet entry
lightref add inbook ## to add an inbook entry
lightref add incollection ## add an incollection entry
lightref add inproceedings ## add an inproceedings entry
lightref add manual ## add a manual entry
lightref add mastersthesis ## add a mastersthesis entry
lightref add misc ## add a misc entry
lightref add phdthesis ## add a phdthesis entry
lightref add proceedings ## add a proceedings entry
lightref add techreport ## add a techreport entry
lightref add unpublished ## add an unpublished entry
```

You will then be prompted to add in the information for the fields that are relevant for that type.
Lightref does **not** enforce the filling out of any fields, so any of them can be left blank if you wish.
Bibtex does have a concept of obligatory fields that should be filled out for each type, however, and will warn the user at run time if they are not entered.

### Print

```bash
lightref print
```

Provides a short summary of each entry in the file, containing the (highlighted) citekey, the Author and the Title.
`Print` can be combined with two flags:

- `-l`/`--long` to show the long output of all entries, with symbols like `{}"` removed.
-  `-r`/`--raw` to show exactly what is in the bibliography file, i.e. no formatting.

**Tip**: The `print` command is best used in conjunction with a paginator such as `less`, combined with the search function.

## Global Flags

`-b`/`--bibliography /path/to/biblio`: allows the user to specify an alternative bibliography to the default in the `lightref.yaml` file.

# Changelog

## 0.1.1

- Added 'booktitle' field to `inbook` entry.
- Added in the `--help` flag information for the root command.

# Features to be added in future releases

- search bibliography and display results to terminal
- citekey checking to avoid duplication: currently there exists nothing to stop you using duplicate citekeys in your bibliography, which is not ideal
- the ability to input data into `yaml` and `json` format.
