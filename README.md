# study-go

## 準備
1. UbuntuMATE16.04 desktopをVirtualBox上に設置
2. OSアップデート  
   ```shell
  $ sudo apt update && sudo apt -y dist-upgrade
  ```
3. [チュートリアルを参照](http://golang.jp/go_tutorial)(http://golang.jp/go_tutorial)
4. gccgoをインストール  
   ```shell
   $ sudo apt install -y gccgo
   ```
   
## Hello worldをプログラム＆コンパイル＆実行してみる
1. まずは、定番のheloworldを書く  
   ```shell
   $ vi helloworld.go
   ```
   以下をプログラムhello-world.goに書く。  
   ```go
   package main
   
   import fmt "fmt" // 入出力フォーマットを実装したパッケージ
   
   func main() {
     fmt.Printf("Hello, world\n")
   }
   ```
2. コンパイルする  
   ```shell
   $ gccgo hello-world.go
   $ ls
   README.md  a.out  hello-wold.go
      ```
3. 実行してみる  
   ```shell
   $ ./a.out
   Hello, world
   ```
4. 上記のように"Hello, world"と表示されたら成功！

