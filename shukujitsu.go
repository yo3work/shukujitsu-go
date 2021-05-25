// Package shukujitsu は内閣府が提供している祝日一覧 CSV ファイルを取得・解析します。
package shukujitsu

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

// import 省略...
// Entry は祝日1日分の情報を保持する構造体です。
type Entry struct {
	YMD   string
	Name  string
	Year  int
	Month int
	Day   int
}

const csvURL = "https://www8.cao.go.jp/chosei/shukujitsu/syukujitsu.csv"

// AllEntries は内閣府ウェブサイトから祝日 CSV を取得して Entry スライスに変換します。
func AllEntries() ([]Entry, error) {
	resp, err := http.Get(csvURL)
	if err != nil {
		return nil, fmt.Errorf("接続に失敗しました: %w", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("データの取得に失敗しました: %w", err)
	}
	records, err := csv.NewReader(transform.NewReader(bytes.NewReader(body), japanese.ShiftJIS.NewDecoder())).ReadAll()
	if err != nil {
		return nil, fmt.Errorf("データの解析に失敗しました: %w", err)
	}
	var entries []Entry
	for i, row := range records {
		if i == 0 {
			continue // ヘッダー行をスキップ
		}
		if len(row) != 2 {
			return nil, fmt.Errorf("想定外のデータに遭遇しました: 行 %d = %v", i+1, row)
		}

		arr := []string{}
		arr = strings.Split(row[0], "/")
		int_year, _ := strconv.Atoi(arr[0])
		int_month, _ := strconv.Atoi(arr[1])
		int_day, _ := strconv.Atoi(arr[2])

		entries = append(entries, Entry{YMD: row[0], Name: row[1], Year: int_year, Month: int_month, Day: int_day})
	}
	return entries, nil
}
