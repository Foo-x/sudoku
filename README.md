# sudoku

以下リンクの数独ソルバーをGoで実装しました。
http://www.aoky.net/articles/peter_norvig/sudoku.htm


## Usage

```bash
go run cmd/solver/main.go assets/input.txt
```

or

```bash
go build cmd/solver/main.go
./main assets/input.txt
```

input.txtは81文字の1行からなります。  
使用できる文字は0から9およびピリオドです。  
0とピリオドは空白マスを表します。
