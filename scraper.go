package main

import (
    "fmt"
    "strings"
    "github.com/gocolly/colly"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func main() {

    db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/testdb")

    if err != nil {
        panic(err.Error())
    }

    defer db.Close()

    verses := make([]string, 32)
    testament := make([]string, 10)
    bible := "New King James Version"

    var url string = "https://www.biblegateway.com/passage/?search=Genesis%201&version=KJV"

    c := colly.NewCollector()	

    for i := 1; i <= 32; i++ {
        c := colly.NewCollector()
        css_verse := fmt.Sprintf("#en-KJV-%d.text.Gen-1-%d", i, i)

        c.OnHTML(css_verse, func(h *colly.HTMLElement) {
            verse_text := h.Text
            trim := fmt.Sprintf("%d", i)
            verse_text1 := strings.TrimLeft(verse_text, trim)
            verse_text2 := strings.TrimSpace(verse_text1)
            verses[i-1] = verse_text2
        })

        c.Visit(url)
    }

    c.Visit(url)

    var verseNum int
    fmt.Println("What verse would you like to see? ")
    fmt.Scanln(&verseNum)
    fmt.Printf("The following is verse %d from %s " + testament[0] + ":\n", verseNum, bible)
    fmt.Println(verses[verseNum-1])
    fmt.Println(testament[0])
    fmt.Print("\n")
}
