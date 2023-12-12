package contentssecurity

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
	"math/rand"


	conn "github.com/uecconsecexp/secexp2022/se_go/connector"
)

// ========== kadai1用 ==========

// main.goで使用したい機能はパッケージとして切り出しておくと便利です。
// lib.goに記述する必要はなく、別なファイルにしてもかまいません。
// ただしファイルの行頭はpackage contentssecurityとする必要があります。

// main.goで使用したい関数はパブリックとするため、大文字で始めます。
func Hello() string {
	return "Hello, world!"
}

// CSVを二次元配列に保存する関数
func ReadCSV(filename string) ([][]float64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var lines [][]float64
	isFirstRow := true
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}
		if isFirstRow {
			isFirstRow = false
			continue
		}
		var line []float64
		for i, item := range record {
			if i == 0 {
				continue
			}
			f, err := strconv.ParseFloat(item, 64)
			if err != nil {
				return nil, err
			}
			line = append(line, f)
		}
		lines = append(lines, line)
	}
	return lines, nil
}


// 行列の積計算

func Calc_matrix(a [][]float64, b [][]float64)([][]float64, error) {
	rowsA := len(a)
	colsA := len(a[0])
	colsB := len(b[0])

	result := make([][]float64, rowsA)
	for i := range result {
		result[i] = make([]float64, colsB)
	}

	for i := 0; i < rowsA; i++ {
		for j := 0; j < colsB; j++ {
			for k := 0; k < colsA; k++ {
				result[i][j] += a[i][k] * b[k][j]
			}
		}
	}

	return result, nil
}

//2引数目が合格最低点
func Hantei(a [][]float64, b [][]float64)([][]float64, error){
	rows := 4
	cols := 4



	result := make([][]float64, rows)
	for i := range result {
		result[i] = make([]float64, cols)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			
				if a[i][j] > b[0][j]{
					result[i][j] = 1
				} else {
					result[i][j] = 0
				}
			
		}
	}

	return result, nil
}

// 1,0を合格不合格に変換する関数
func Gouhi_henkan(a [][]float64)([][]string, error){
	rows := 5
	cols := 5



	result := make([][]string, rows)
	for i := range result {
		result[i] = make([]string, cols)
	}

	result[0][0] = "　　 "
	result[1][0] = "生徒1"
	result[2][0] = "生徒2"
	result[3][0] = "生徒3"
	result[4][0] = "生徒4"

	result[0][1] = " 高校A"
	result[0][2] = " 高校B"
	result[0][3] = " 高校C"
	result[0][4] = " 高校D"

    for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if a[i][j] == 1 {
				result[i+1][j+1] = "　合格"
			} else {
				result[i+1][j+1] = "不合格"
			}

		}
	}

	return result, nil
}







// 便利なモジュールやその他必要な情報は README.md にまとめています。

// ========== kadai2用 ==========

// 実際に通信を行う機能もパッケージ側に書いてしまいましょう。
// YobikouSide、ChugakuSideの中身を書き換えてください。

// 6*6乱数行列を生成する関数 [][]float64の二十行列


func Generaterandommatrix() [][]float64 {
	rand.Seed(time.Now().UnixNano()) // 乱数のシードを設定
	matrix := make([][]float64, 6)   // 6x6の行列を初期化

	for i := range matrix {
		matrix[i] = make([]float64, 6)
		for j := range matrix[i] {
			matrix[i][j] = rand.Float64() * 100 // 0から100までの乱数を生成
		}
	}

	return matrix
}

// 行列を左右半分にする関数
func Splitmatrix(matrix [][]float64) ([][]float64, [][]float64) {
    var Mleft, Mright [][]float64

    for i := 0; i < 6; i++ {
        for j := 0; j < 6; j++ {
            if j < 3 {
                Mleft[i][j] = matrix[i][j]
            } else {
                Mright[i][j-3] = matrix[i][j]
            }
        }
    }
    return Mleft, Mright
}

func YobikouSide() {
	yobikou, err := conn.NewYobikouServer()
	// 例外処理: err != nilは必ず確認し、適切に処理してください。panicでかまいません。
	if err != nil {
		panic(err)
	}
	defer yobikou.Close() // TCP通信を終了するため、最後に必ず呼び出す必要があります。消さないでください。

	// バイト列を受け取るには Receive() を使用します。デバッグ用であり最終的なプログラムでは使用する必要はありません。
	m, err := yobikou.Receive()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Received: %s\n", m)

	// 配列を送るには ReceiveTable()を使用します。こちらを使用して最終的なプログラムを作成してください。
	matrix, err := yobikou.ReceiveTable()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Received: %v\n", matrix)

	matrix1 := [][]float64{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	err = yobikou.SendTable(matrix1)
	if err != nil {
		panic(err)
	}

	// 受信する処理のみ書きましたが、ChugakuSideで行っているSend/SendTableも可能です。
}

// 中学校　プログラム
func ChugakuSide(addr string) {
	chugaku, err := conn.NewChugakuClient(addr)
	if err != nil {
		panic(err)
	}
	defer chugaku.Close() // 消さないでください。

	// バイト列を送るには Send() を使用します。デバッグ用であり最終的なプログラムでは使用する必要はありません。
	err = chugaku.Send([]byte("ping"))
	if err != nil {
		panic(err)
	}

	// 配列を送るには SendTable()を使用します。こちらを使用して最終的なプログラムを作成してください。
	matrix := Generaterandommatrix()
	err = chugaku.SendTable(matrix)
	if err != nil {
		panic(err)
	}

	matrix1, err := chugaku.ReceiveTable()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Received: %v\n", matrix1)


	// 送信する処理のみ書きましたが、YobikouSideで行っているReceive/ReceiveTableも可能です。
}
