// Command line tool to open random question
//   as a Google search
// Written by: Cameron Napoli

package main

import (
    "io"
    "bytes"
    "bufio"
    "fmt"
    "log"
    "runtime"
    "os/exec"
    "os"
    "math/rand"
    "time"
    "net/url"
)


func check(e error) {
    if e != nil {
        panic(e)
    }
}


func numLinesInFile(fname string) (int, error) {
    f, err := os.Open(fname)
    check(err)
    r := bufio.NewReader(f)

    buf := make([]byte, 32*1024)
    count := 0
    lineSep := []byte{'\n'}

    for {
        c, err := r.Read(buf)
        count += bytes.Count(buf[:c], lineSep)

        switch {
        case err == io.EOF:
            return count, nil
        case err != nil:
            return count, err
        }
    }
}


func getRandLineInFile(fname string) string {
    num_lines, err := numLinesInFile(fname)
    check(err)

    lines := make([]string, num_lines)

    f, err := os.Open(fname)
    check(err)
    r := bufio.NewReader(f)

    for i := 0; i < num_lines + 1; i++ {
        line, _, err := r.ReadLine()
        if err == io.EOF {
            break
        }
        lines[i] = string(line)
    }

    return lines[rand.Intn(len(lines))]
}

func getSnglrNoun() string {
    // Get random singular noun from file
    return getRandLineInFile("./resources/nouns_singular.txt")
}

func getRandPlurNoun() string {
    // Get random plural noun from file
    return getRandLineInFile("./resources/nouns_plural.txt")
}

func getRandAdj() string {
    // Get random adjective from file
    return getRandLineInFile("./resources/adjectives.txt")
}

func formQueryString() string {
    /*
    Where (is|are) (noun|noun_plural) (from|right now|located)
    Who   (is|are) (noun|noun_plural) (similar to)
    Why   (is|are) (noun|noun_plural) {adjective}
    What  (is|are) (noun|noun_plural)
    */

    w2, noun := "", ""
    is_plural := []bool {true, false}[rand.Intn(2)]
    if is_plural {
        w2 = "are"
        noun = getRandPlurNoun()
    } else {
        w2 = "is"
        noun = getSnglrNoun()
    }

    q_begins := []string {"Where", "Who", "Why", "What"}
    q_type := q_begins[rand.Intn(len(q_begins))]

    query_string := ""
    // Hardcode question building for now
    switch q_type {
    case "Where":
        optional := []string {" from", " right now", " located", ""}[rand.Intn(4)]
        query_string += "Where " + w2 + " " + noun + optional + "?"
    case "Who":
        optional := []string {" similar to", ""}[rand.Intn(2)]
        query_string += "Who " + w2 + " " + noun + optional + "?"
    case "Why":
        adjective := getRandAdj()
        query_string += "Why " + w2 + " " + noun + " " + adjective + "?"
    case "What":
        query_string += "What " + w2 + " " + noun + "?"
    default:
        fmt.Println("UNKNOWN case")
    }

    return query_string
}

// TODO: run this as goroutine
func runQuery(query string) {
    // Validate URL query
    qEsc := url.QueryEscape(query)

    l, e := url.Parse("https://google.com#q=" + qEsc)

    if e != nil {
        fmt.Println("Error with query%s\n", e)
    }

    // Open query in default OS browser
    cmd := new(exec.Cmd)

	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("xdg-open", l.String())
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", l.String())
	case "darwin": // MAC OS
		cmd = exec.Command("open", l.String())
	default:
		fmt.Println("Issue opening browswer")
	}

    err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
    seed := int64(time.Now().UnixNano())
    rand.Seed(seed)

    question := formQueryString()
    fmt.Println("question: ", question)

    runQuery(question)
}
