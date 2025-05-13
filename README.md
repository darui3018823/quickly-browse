# quickly-browse (`q-brow`)

**quickly-browse** は、コマンドラインや `Win + R` から一瞬で検索できる軽量CLIツールです。  <br>
Google、YouTube、Twitter、DuckDuckGo などの検索に対応し、指定したキーワードを即座に既定ブラウザで開きます。<br>
macOS向けにも提供を行いますが、現時点ではWindows向けに最適化を行い、macOSは後になっております。

## 特徴

- `Win + R` から即起動（ターミナル非表示）
- Google / YouTube / Twitter / DuckDuckGo 対応
- クロスプラットフォーム（Windows / macOS Intel / Apple Silicon）
- 単一バイナリ、インストール不要

## 使い方

```sh
q-brow [options] "検索キーワード"
````

例：

```sh
q-brow "Golang CLI"
q-brow -y "Lofi Chill"
q-brow -t "@apple"
q-brow -d "プライバシー重視 ブラウザ"
```

### オプション一覧

| オプション    | 説明               |
| -------- | ---------------- |
| `-g`     | Google検索（省略時も同様） |
| `-y`     | YouTube検索        |
| `-t`     | Twitter検索        |
| `-d`     | DuckDuckGo検索     |
| `--help` | ヘルプ（GUIダイアログ表示）  |

## ダウンロード

| プラットフォーム            | バイナリ                             |
| ------------------- | -------------------------------- |
| Windows (.exe)      | [q-brow-win.zip](releases)       |
| macOS Intel         | [q-brow-mac-intel.zip](releases) |
| macOS Apple Silicon | [q-brow-mac-arm.zip](releases)   |

> 中身のファイルをそのまま `Win + R` から呼び出せる場所に置いてください（`C:\Windows`などにコピーすると便利です）

## ビルド方法（開発者向け）

```sh
# Windows
go build -ldflags="-H=windowsgui" -o q-brow.exe

# macOS Intel
GOOS=darwin GOARCH=amd64 go build -o q-brow-mac-intel

# macOS ARM (M1/M2/M3)
GOOS=darwin GOARCH=arm64 go build -o q-brow-mac-arm
```

## ライセンス

[MIT License](./LICENSE)

## その他
配布しているバイナリが実行できない際は[Issues](https://github.com/darui3018823/quickly-browse/issues)、または[こちら](https://daruks.com/contact/)へお問い合わせをお願いいたします。