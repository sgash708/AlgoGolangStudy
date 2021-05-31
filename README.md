# studyGo
* アルゴリズム等の勉強
* (ついでに) 文法/記法/慣例などの学習

# version
* go 1.16.4

# Todo

```bash:最初にやること
# "github.com/sgash708/AlgoGoStudy"は自作モジュール名
# 一意な名前にしたいので、githubのパスにする慣例がある
go mod init github.com/sgash708/AlgoGoStudy
```

```bash:ファイル追加時にやること
# コードフォーマット
go fmt
# golintで記法を修正
golint main.go
```

# do not forget

* 関数を追加したら、<code>キャメルケース</code>にして呼び出す
* 関数に対するコメントを必ず残す

```go:element/elementCheck.go
// this is element/elementCheck.go

// RemoveElement valと一致しない値に置き換える
func RemoveElement(nums []int, val int) int {
```

* 命名パッケージは、<code>小文字統一</code>

```go:element/elementCheck.go
// this is element/elementCheck.go

// package elementCheckはNG
// B.C.) don't use MixedCaps in package name;
package elementcheck
```

* 関数呼び出しでは、<code>モジュール名/ディレクトリパス</code>をimportする
* (2回目) 呼び出し時には<code>キャメルケース</code>で呼び出し！！
```go:main.go
// this is main.go

import (
	element "github.com/sgash708/AlgoGoStudy/element"
)

func main() {
	fmt.Println(element.RemoveElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}
```