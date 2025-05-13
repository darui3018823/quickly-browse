package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"syscall"
	"unsafe"
)

var (
	user32          = syscall.NewLazyDLL("user32.dll")
	procMessageBoxW = user32.NewProc("MessageBoxW")
)

func showHelp() {
	if runtime.GOOS == "windows" {
		showWindowsHelpDialog()
	} else {
		// macOSやLinux用は後で実装
	}
}

func main() {
	flag.Parse()
	args := flag.Args()

	// --helpまたは-hを含んでいたらダイアログでヘルプを表示
	if len(args) == 0 || containsHelpFlag(os.Args[1:]) {
		showHelp()
		return
	}

	// 検索エンジンオプション
	useGoogle := flag.Bool("g", false, "Search with Google (default)")
	useYouTube := flag.Bool("y", false, "Search with YouTube")
	useTwitter := flag.Bool("t", false, "Search with Twitter")
	useDuck := flag.Bool("d", false, "Search with DuckDuckGo")

	// ヘルプテキストの上書き
	flag.Usage = func() {
		fmt.Println("Usage:")
		fmt.Println("  q-brow [options] \"search terms\"")
		fmt.Println()
		fmt.Println("Options:")
		fmt.Println("  -g          Search with Google (default)")
		fmt.Println("  -y          Search with YouTube")
		fmt.Println("  -t          Search with Twitter")
		fmt.Println("  -d          Search with DuckDuckGo")
		fmt.Println("  --help      Show this help message")
	}

	// ヘルプ強制表示
	if len(args) == 0 {
		flag.Usage()
		return
	}

	query := strings.Join(args, " ")
	encoded := url.QueryEscape(query)

	// 検索エンジンの選択
	searchURL := "https://www.google.com/search?q=" + encoded // デフォルト: Google
	switch {
	case *useGoogle:
		searchURL = "https://www.google.com/search?q=" + encoded
	case *useYouTube:
		searchURL = "https://www.youtube.com/results?search_query=" + encoded
	case *useTwitter:
		searchURL = "https://twitter.com/search?q=" + encoded
	case *useDuck:
		searchURL = "https://duckduckgo.com/?q=" + encoded
	}

	// プラットフォーム別の開き方
	openBrowser(searchURL)
}

func openBrowser(url string) {
	switch runtime.GOOS {
	case "windows":
		exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Run()
	case "darwin":
		exec.Command("open", url).Run()
	default:
		exec.Command("xdg-open", url).Run()
	}
}

func containsHelpFlag(args []string) bool {
	for _, arg := range args {
		if arg == "--help" || arg == "-h" {
			return true
		}
	}
	return false
}

func showWindowsHelpDialog() {
	title := "q-brow Help"
	content := `Usage:
  q-brow [options] "search terms"

Options:
  -g        Google
  -y        YouTube
  -t        Twitter
  -d        DuckDuckGo
  --help    Show this message
`

	procMessageBoxW.Call(
		0,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(content))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))),
		0x00000040, // MB_ICONINFORMATION
	)
}
