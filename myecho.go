package main

import (
    "os"
    "flag" // コマンドラインオプションのパーサー
)

var omitNewline = flag.Bool("n", false, "don't print final newline")
var ignorespace = flag.Bool("s", false, "ignore spaces between words")


const (
    Space   = " "
    Newline = "\n"
)

func main() {
    flag.Parse() // パラメータリストを調べてflagに設定
    var s string = ""
    for i := 0; i < flag.NArg(); i++ {
        if i > 0 && !*ignorespace {
            s += Space
        }
        s += flag.Arg(i)
    }
    if !*omitNewline {
        s += Newline
    }
    os.Stdout.WriteString(s)
}

