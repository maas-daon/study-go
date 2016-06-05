package main

import (
    "os"
    "flag" // コマンドラインオプションのパーサー
)

// 引数"-n"の定義
//   意味：改行なし
//   デフォルト：無指定（改行あり）
//   メッセージ（Usage文字列): "don't print final newline"
var omitNewline = flag.Bool("n", false, "don't print final newline")

// 引数"-s"の定義
//   意味：単語間のスペース無視
//   デフォルト：無指定（スペース有り、ただし複数スペースは１スペースに変換される）
//   メッセージ（Usage文字列): "ignore spaces between words"
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

