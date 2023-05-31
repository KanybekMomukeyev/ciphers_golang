package matrix_vector

import (
	"fmt"
	"strings"
)

// Матрица представляет собой квадратную матрицу
type Matrix struct {
	order int
	data  [][]int
}

// String реализует интерфейс Stringer
func (m Matrix) String() string {
	var b strings.Builder
	for _, row := range m.data {
		for i, item := range row {
			if i == 0 {
				b.WriteString("|")
			}
			b.WriteString("\t" + fmt.Sprintf("%d", item) + "\t|")
		}
		b.WriteString("\n")
	}
	return b.String()
}

// Порядок возврата квадратной матрицы
func (m *Matrix) Order() int {
	return m.order
}

// NewMatrix возвращает новую квадратную матрицу заданного порядка, загруженную заданными данными.
// Размер входных данных должен быть точно равен квадрату порядка (order^2).
func NewMatrix(order int, data []int) (*Matrix, error) {
	if len(data) != order*order {
		return nil, fmt.Errorf("failed to build square matrix, got invalid data size %d, wantMatrix %d", len(data), order*order)
	}
	m := &Matrix{order: order}
	m.data = make([][]int, order)
	for i := 0; i < order; i++ {
		row := make([]int, order)
		for j := 0; j < order; j++ {
			row[j] = data[(i*order)+j]
		}
		m.data[i] = row
	}
	return m, nil
}

// Определитель возвращает определитель матрицы. Этот алгоритм очень медленный O(n!), так как он
// основывается на наивном подходе. Реализуйте разложение LU для повышения производительности O (n ^ 3).
func (m *Matrix) Determinant() (int, error) {
	if m.order < 1 {
		return 0, fmt.Errorf("determinant is undefined for order < 1")
	}
	if m.order == 1 {
		return m.data[0][0], nil
	}
	sign := 1
	var det int
	for i := 0; i < m.order; i++ {
		subM, _ := Minor(m, 0, i)
		minor, _ := subM.Determinant()
		det += sign * m.data[0][i] * minor
		sign *= -1
	}
	return det, nil
}

// IsInvertibleMod возвращает, является ли матрица обратимой по модулю n. Матрица A с элементами в Zn есть
// обратим по модулю n тогда и только тогда, когда вычет det(A) по модулю n имеет обратный по модулю m. Также,
// A обратим по модулю n, если m и вычет det(A) по модулю n не имеют общих простых множителей.
func (m *Matrix) IsInvertibleMod(n int) bool {
	if n < 2 {
		return false
	}
	for _, col := range m.data {
		for _, x := range col {
			if x < 0 || x >= n {
				return false
			}
		}
	}
	det, err := m.Determinant()
	if err != nil {
		return false
	}
	res := Residue(det, n)
	if _, err := ModularInverse(res, n); err != nil {
		return false
	}

	return true
}

// InverseMod возвращает перевернутую квадратную матрицу по модулю n.
func (m *Matrix) InverseMod(n int) (*Matrix, error) {
	if n < 2 {
		return nil, fmt.Errorf("got modulo < 2")
	}
	if !m.IsInvertibleMod(n) {
		return nil, fmt.Errorf("matrix is not invertible mod %d", n)
	}
	det, _ := m.Determinant() // Neglect error since its checked by IsInvertibleMod
	res := Residue(det, n)
	inverse, _ := ModularInverse(res, n) // Neglect error since its checked by IsInvertibleMod
	adj, err := m.Adjoint()
	if err != nil {
		return nil, fmt.Errorf("failed to compute Adj(\n%s\n); %v", m, err)
	}
	for i := range adj.data {
		for j := range adj.data {
			adj.data[i][j] = Residue(adj.data[i][j]*inverse, n)
		}
	}
	return adj, nil
}

// Adjoint возвращает сопряженную матрицу
func (m *Matrix) Adjoint() (*Matrix, error) {
	cof, err := m.Cofactor()
	if err != nil {
		return nil, fmt.Errorf("failed to compute cofactor matrix for \n%s; %v", m, err)
	}
	return cof.Transpose(), nil
}

// Кофактор возвращает матрицу кофактора
func (m *Matrix) Cofactor() (*Matrix, error) {
	cof := &Matrix{order: m.order, data: make([][]int, m.order)}
	for i := 0; i < m.order; i++ {
		row := make([]int, m.order)
		for j := 0; j < m.order; j++ {
			minor, _ := Minor(m, i, j) // Error is neglected since row & col are always in bound
			detM, err := minor.Determinant()
			if err != nil {
				return nil, fmt.Errorf("failed to compute det(m) for minor at row:%d col:%d\n%s;%v", i, j, m, err)
			}
			if (i+j)%2 == 0 {
				row[j] = detM
			} else {
				row[j] = detM * -1
			}
		}
		cof.data[i] = row
	}
	return cof, nil
}

// Transpose возвращает транспонированную матрицу
func (m *Matrix) Transpose() *Matrix {
	t := &Matrix{order: m.order, data: make([][]int, m.order)}
	for i := 0; i < m.order; i++ {
		t.data[i] = make([]int, m.order)
		for j := 0; j < m.order; j++ {
			t.data[i][j] = m.data[j][i]
		}
	}
	return t
}

// Minor возвращает матрицу сомножителей заданной матрицы в точках p, q (строка, столбец).
func Minor(m *Matrix, p, q int) (*Matrix, error) {
	if 0 > p || p >= m.order || 0 > q || q >= m.order {
		return nil, fmt.Errorf("got row and/or col out of bound")
	}
	if m.order <= 1 {
		return &Matrix{}, nil
	}
	r := &Matrix{order: m.order - 1}
	r.data = make([][]int, 0, r.order)
	for row := 0; row < m.order; row++ {
		tmp := make([]int, 0, r.order)
		for col := 0; col < m.order; col++ {
			if row != p && col != q {
				tmp = append(tmp, m.data[row][col])
			}
		}
		if row != p {
			r.data = append(r.data, tmp)
		}
	}
	return r, nil
}

// VectorProductMod возвращает произведение матрицы на вектор
func (m *Matrix) VectorProductMod(mod int, vector ...int) ([]int, error) {
	if len(vector) != m.order {
		return nil, fmt.Errorf("got invalid vector size %d, want %d", len(vector), m.order)
	}
	if mod < 2 {
		return nil, fmt.Errorf("got modulo < 2")
	}
	vp := make([]int, m.order)
	for i := 0; i < m.order; i++ {
		for j := 0; j < m.order; j++ {
			vp[i] += m.data[i][j] * vector[j]
		}
		vp[i] = vp[i] % mod
	}
	return vp, nil
}
