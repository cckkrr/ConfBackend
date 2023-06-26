package coord

import (
	S "ConfBackend/services"
	"errors"
	"github.com/gonum/matrix/mat64"
	"github.com/sirupsen/logrus"
	"math"
)

// doCalcCoordMath 实际计算坐标，p是长度大于等于4的二维数组，为已知节点坐标
// d是和p等长度的二维数组，为已知节点对应的距离。单位均为米/m。
// 返回：x，y，z，error
// 注意：至少要4个已知坐标、距离，才能计算。
// 例如：
//
//	     p := [][]float64{{0.01, 0.01, 0}, {0.05, 0.03, 0}, {0.1, 0.01, 0}, {0.05, 0.01, 0.01}}
//				d := [][]float64{{0.04}, {0.02}, {0.05}, {0.01}}
func doCalcCoordMath(p [][]float64, d [][]float64) (float64, float64, float64, error) {

	/*		p := [][]float64{{0.01, 0.01, 0}, {0.05, 0.03, 0}, {0.1, 0.01, 0}, {0.05, 0.01, 0.01}}
			d := [][]float64{{0.04}, {0.02}, {0.05}, {0.01}}*/

	/*	p := [][]float64{{3.671541, 100.0591, 1.160181}, {84.71916, 158.5611, 4.47283}, {86.89013, -18.87274, 4.721066}, {0.1268933, 31.06314, 0.9591231}, {5.993424, 131.8034, 0.9396151}}
		d := [][]float64{{65.6}, {130.58}, {81.9}, {31.24}, {94.91}}*/
	if len(p) < 4 || len(p) != len(d) {
		S.S.Logger.WithFields(logrus.Fields{
			"len_p": len(p),
			"len_d": len(d),
		}).Infof("doCalc计算坐标入参有误")
		return 0, 0, 0, errors.New("入参有误")
	}

	A := mat64.NewDense(len(p), 4, nil)
	B := mat64.NewDense(len(p), 1, nil)

	for i := 0; i < len(p); i++ {
		A.Set(i, 3, 1)
		A.Set(i, 0, -2*p[i][0])
		A.Set(i, 1, -2*p[i][1])
		A.Set(i, 2, -2*p[i][2])
		B.Set(i, 0, d[i][0]*d[i][0]-p[i][0]*p[i][0]-p[i][1]*p[i][1]-p[i][2]*p[i][2])

	}

	// gen normalized weight
	dT := transposeSingleColumn(d)
	weights := genNormalizedWeight(dT)
	W := genWeightedMatrix(weights)

	x := mat64.NewDense(4, 1, nil)
	// X = (A^TWA)^-1 A^T W B

	// calculate A^TWA
	A_T := mat64.NewDense(4, len(p), nil)
	A_T.Copy(A.T())
	A_T_W := mat64.NewDense(4, len(p), nil)
	A_T_W.Mul(A_T, W)
	A_T_W_A := mat64.NewDense(4, 4, nil)
	A_T_W_A.Mul(A_T_W, A)
	// calculate A^T W B
	A_T_W_B := mat64.NewDense(4, 1, nil)
	A_T_W_B.Mul(A_T_W, B)
	// calculate (A^TWA)^-1
	A_T_W_A_Inv := mat64.NewDense(4, 4, nil)
	A_T_W_A_Inv.Inverse(A_T_W_A)
	// calculate (A^TWA)^-1 A^T W B
	x.Mul(A_T_W_A_Inv, A_T_W_B)

	/*		if err != nil {
			panic(err)

		}*/
	// 输出结果
	//fmt.Println("节点坐标：")
	////fmt.Printf("x = %v, y = %v\n", x.At(0, 0), x.At(1, 0))
	//fmt.Printf("x = %.4f, y = %.4f, z = %.4f\n", roundToFour(x.At(0, 0)), roundToFour(x.At(1, 0)), roundToFour(x.At(2, 0)))

	return roundToFour(x.At(0, 0)), roundToFour(x.At(1, 0)), roundToFour(x.At(2, 0)), nil

}

func roundToFour(x float64) float64 {
	return math.Round(x*10000) / 10000
}

func transposeSingleColumn(in [][]float64) (res []float64) {
	for i := 0; i < len(in); i++ {
		res = append(res, in[i][0])
	}
	return
}

// genNormalizedWeight 产生一个归一化的距离权重，测量越近就可能越准确，对于近的距离赋予大权重
// 公式：wi = exp(-di^2/2*sigma^2)
func genNormalizedWeight(in []float64) (res []float64) {
	// use Gaussian distance weight function wi = exp(-di^2/2*sigma^2)
	// sigma = 1
	sigma := 0.05
	for i := 0; i < len(in); i++ {
		res = append(res, math.Exp(-math.Pow(in[i], 2)/2*math.Pow(sigma, 2)))
	}

	// normalize the res 归一化权重
	sum := 0.0
	for i := 0; i < len(res); i++ {
		sum += res[i]
	}
	for i := 0; i < len(res); i++ {
		res[i] = res[i] / sum
	}

	return

}

// genWeightedMatrix 产生一个加权后的对角矩阵
func genWeightedMatrix(weights []float64) (res *mat64.Dense) {
	res = mat64.NewDense(len(weights), len(weights), nil)
	for i := 0; i < len(weights); i++ {
		res.Set(i, i, weights[i])
	}
	return
}
