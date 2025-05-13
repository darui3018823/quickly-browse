package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func main() {
	// フラグ定義
	useGoogle := flag.Bool("g", false, "Search with Google (default)")
	// 他のオプション（今後）: -y, -t など
	flag.Parse()

	// 検索クエリ取得
	queryArgs := flag.Args()
	if len(queryArgs) == 0 {
		fmt.Println("Usage: q-brow [-g] \"search terms\"")
		return
	}
	query := strings.Join(queryArgs, " ")
	encoded := url.QueryEscape(query)

	// デフォルトはGoogle
	searchURL := "https://www.google.com/search?q=" + encoded
	if *useGoogle {
		// 明示的に -g を指定した場合もGoogle
		searchURL = "https://www.google.com/search?q=" + encoded
	}

	// 実行ファイル名から呼び出し名を取得（optional: for debug/logging）
	binName := strings.ToLower(os.Args[0])
	if strings.Contains(binName, "quickly-browse") || strings.Contains(binName, "q-brow") {
		// OK, future: you could swap behavior by binary name
	}

	// OS別にブラウザで開く
	if runtime.GOOS == "windows" {
		exec.Command("cmd", "/c", "start", "", searchURL).Run()
	} else {
		cmd := "xdg-open"
		if runtime.GOOS == "darwin" {
			cmd = "open"
		}
		exec.Command(cmd, searchURL).Run()
	}

	fmt.Printf("Searching: %s\n", query)
}
