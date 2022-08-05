package main

import (
	"fmt"
	"os/exec" // OSコマンドの実行
)
// calculate  bcコマンドを呼び出して計算する
func calculate(exp string) ([]byte, error) {
	command := fmt.Sprintf("echo \"%s\" | bc -l", exp)
	// コマンドの文字列を作る。Sprintfの%sの部分がexpの値で置き換わる
	result, err := exec.Command("sh", "-c", command).Output()
	return result, err // 2つの値を戻す
	// result は結果を表すバイト列（byteのスライス）
	// err はエラーがあるかどうかを示す。errが nilならば正常終了
	// これの処理は呼び出し側（calculateAndPrintValue）が行う
}
