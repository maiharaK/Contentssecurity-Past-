package contentssecurity

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
	"math/rand"
	"gonum.org/v1/gonum/mat"


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
    // MleftとMrightを適切なサイズで初期化します。
    Mleft := make([][]float64, 6)
    Mright := make([][]float64, 6)
    for i := 0; i < 6; i++ {
        Mleft[i] = make([]float64, 3)
        Mright[i] = make([]float64, 3)
    }

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


// 正則であるかを判定
func IsRegular(matrix [][]float64) bool {
	// 2次元スライスを*mat.Dense型に変換します。
	r, c := len(matrix), len(matrix[0])
	data := make([]float64, r*c)
	for i := range matrix {
		for j := range matrix[i] {
			data[i*c+j] = matrix[i][j]
		}
	}
	m := mat.NewDense(r, c, data)

	// 行列式を計算します。
	determinant := mat.Det(m)
	if determinant != 0 {
		return true
	}
	return false
}

func YobikouSide() {
	yobikou, err := conn.NewYobikouServer()
	// 例外処理: err != nilは必ず確認し、適切に処理してください。panicでかまいません。
	if err != nil {
		panic(err)
	}
	defer yobikou.Close() // TCP通信を終了するため、最後に必ず呼び出す必要があります。消さないでください。


	// 配列を送るには ReceiveTable()を使用します。こちらを使用して最終的なプログラムを作成してください。
	randommatrix, err := yobikou.ReceiveTable()
	if err != nil {
		panic(err)
	}

	fmt.Printf("予備校Received1: %v\n", randommatrix)

	a_primematrix, err := yobikou.ReceiveTable()
	if err != nil {
		panic(err)
	}

	fmt.Printf("予備校Received2: %v\n", a_primematrix)



	// 受信する処理のみ書きましたが、ChugakuSideで行っているSend/SendTableも可能です。
}

// 中学校　プログラム
func ChugakuSide(addr string) {
	chugaku, err := conn.NewChugakuClient(addr)
	if err != nil {
		panic(err)
	}
	defer chugaku.Close() // 消さないでください。

	//1
	// 乱数行列を生成
	randommatrix := Generaterandommatrix()

	for !IsRegular(randommatrix){ // 正則であるかを確認
		randommatrix = Generaterandommatrix()
	}

	// 2
	// 乱数行列を送る
	err = chugaku.SendTable(randommatrix)
	if err != nil {
		panic(err)
	}

	// 6
	// Mleft, Mrightを求める
	Mleft, Mright := Splitmatrix(randommatrix)
	
	fmt.Printf("Mleft: %v\n", Mleft)
	fmt.Printf("Mright: %v\n", Mright)

	// 成績データ読み込み
	seiseki_data, err := ReadCSV("seiseki.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("成績: %v\n", seiseki_data)
	
	// 7
	// A'を求める
	A_prime, err := Calc_matrix(seiseki_data, Mleft)
	fmt.Printf("A': %v\n", A_prime)

	// 8
	// A'を送る
	err = chugaku.SendTable(A_prime)
	if err != nil {
		panic(err)
	}

	// ↑ここまで動作確認済み

	
	
	// // B' を受信
	// b_prime, err := chugaku.ReceiveTable()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("中学Received2: %v\n", b_prime)

	// // 10
	// // A''を求める
	// A_dprime, err := Calc_matrix(seiseki_data, Mright)
	// A_dprime = Calc_matrix(seiseki_data, b_prime)

	// // 12
	// // A''を送る
	// err = chugaku.SendTable(A_dprime)
	// if err != nil {
	// 	panic(err)
	// }

	// // 合否行列を受信
	// gouhi_data, err := chugaku.ReceiveTable()
	// if err != nil {
	// 	panic(err)
	// }
	// gouhi, err := Gouhi_henkan(gouhi_data)

	// // 16
    // // 合否表を出力
	// for i := 0; i < 5; i++{
	// 	for j := 0; j < 5; j++{
	// 		fmt.Print(gouhi[i][j])
	// 		fmt.Print(" ")

	// 	}
	// 	fmt.Print("\n")
	// }	

	




	// 送信する処理のみ書きましたが、YobikouSideで行っているReceive/ReceiveTableも可能です。
}
