# study-go

## 準備
1. UbuntuMATE16.04 desktopをVirtualBox上に設置
2. OSアップデート  
   ```shell
  $ sudo apt update && sudo apt -y dist-upgrade
  ```
3. [チュートリアル](http://golang.jp/go_tutorial)http://golang.jp/go_tutorialを閲覧
4. gccgoをインストール  
   ```shell
   $ sudo apt install -y gccgo
   ```
5. 定番のhelo-worldを書く
   ```shell
   $ vi hello-world.go
   ```
   以下をプログラムhello-world.goに書く。
   ```go
   package main
   
   import fmt "fmt" // 入出力フォーマットを実装したパッケージ
   
   func main() {
     fmt.Printf("Hello, world\n")
   }
   ```
6. コンパイルする
   ```shell
   $ gccgo hello-world.go
   $ ls
   README.md  a.out  hello-wold.go
      ```
7. 実行してみる
   ```shell
   $ ./a.out
   Hello, world
   ```
8. 上記のように"Hello, world"と表示されたら成功！
   ```
