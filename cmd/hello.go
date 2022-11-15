// サンプルプログラム

package main // 以下の内容がmainパッケージに含まれることを示す

import "fmt" // fmtパッケージを使うことを宣言

/*
 * mainメソッド
 */
func main() { // Goのプログラムは、mainメソッドの先頭行から実行される決まり
  fmt.Println("Hello Go!") // fmtパッケージの中にある Println()メソッド
	fmt.Println("hello", 10, 'x')
}
