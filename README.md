# Why Lightref?

Lightref is a small, cross platform CLI for adding entries to a bibtex bibliography.
It is designed as a command line based alternative to the GUI bibliograph management alternatives such as Jabref and Bibdesk.
Lightref is **not** intended as a serious competitor to these: they are wonderful programs with vastly more options than Lightref offers.
However, for a simple task such as adding an entry to a bibtex file, opening a full GUI is often overkill, hence Lightref.

# Current Satus (November 2020)

Lightref is very much under development, and as such should not be considered a finished product.
Currently it offers the ability to:

- add book entries
- add article entries
- add PhD theses

The remaining classes will be added in the following weeks.

# Features to be added 

- search bibliography and display results to terminal
- citekey checking to avoid duplication: currently there exists nothing to stop you using duplicate citekeys in your bibliography, which is not ideal
- remove entries from bibliography

# Installation

Lightref is written in Go using [Cobra](https://github.com/spf13/cobra).
Since it is still under development there is currently no binary for download.
This will be made available when ready.

It is possible to build from source.
In which case, you should install Cobra following instructions [here](https://github.com/spf13/cobra#installing).
Then, fork the Lightref repo and run:

```
$ cd /path/to/lightref/repo

$ go install lightref
```

# Instructions

Currently, Lightref supports adding entries to the bibliography through the following commands:

- `lightref add article`
- `lightref add book`
- `lightref add phd`
