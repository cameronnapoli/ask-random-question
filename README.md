# Ask Random Question
Command line script which opens a google query with a randomly generated question. Built with Golang.

### Notes

The question is built by using the nouns/adjectives found in `resources/` files. Then the pieces are joined together into Who/What/Where/Why questions.

To modify nouns and adjectives in the question, add entries to the following files:

* ```resources/adjectives.txt```
* ```resources/nouns_plural.txt```
* ```resources/nouns_singular.txt```

### Usage

Run with:

```
go run ask_question.go
```
