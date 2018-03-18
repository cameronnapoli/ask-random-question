# ask-random-question
Opens google query with random question. Built with Golang.

#### Notes

Query string is built with files in `resources/`

To modify nouns and adjectives in the query string, add entries to the following files:

* ```resources/adjectives.txt```
* ```resources/nouns_plural.txt```
* ```resources/nouns_singular.txt```

##### Usage

Run with:

```
go run ask_question.go 
```