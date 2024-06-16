package v

import "math"

//
// Full disclosure: This file is a direct copy of raylib's matrix code
//

type Mat struct {
	M0, M4, M8, M12  float32
	M1, M5, M9, M13  float32
	M2, M6, M10, M14 float32
	M3, M7, M11, M15 float32
}

func M(m0, m4, m8, m12, m1, m5, m9, m13, m2, m6, m10, m14, m3, m7, m11, m15 float32) Mat {
	return Mat{m0, m4, m8, m12, m1, m5, m9, m13, m2, m6, m10, m14, m3, m7, m11, m15}
}

// MatrixIdentity - Returns identity matrix
func MatrixIdentity() Mat {
	return M(
		1.0, 0.0, 0.0, 0.0,
		0.0, 1.0, 0.0, 0.0,
		0.0, 0.0, 1.0, 0.0,
		0.0, 0.0, 0.0, 1.0)
}

func MatrixDeterminant(mat Mat) float32 {
	var result float32

	a00 := mat.M0
	a01 := mat.M1
	a02 := mat.M2
	a03 := mat.M3
	a10 := mat.M4
	a11 := mat.M5
	a12 := mat.M6
	a13 := mat.M7
	a20 := mat.M8
	a21 := mat.M9
	a22 := mat.M10
	a23 := mat.M11
	a30 := mat.M12
	a31 := mat.M13
	a32 := mat.M14
	a33 := mat.M15

	result = a30*a21*a12*a03 - a20*a31*a12*a03 - a30*a11*a22*a03 + a10*a31*a22*a03 +
		a20*a11*a32*a03 - a10*a21*a32*a03 - a30*a21*a02*a13 + a20*a31*a02*a13 +
		a30*a01*a22*a13 - a00*a31*a22*a13 - a20*a01*a32*a13 + a00*a21*a32*a13 +
		a30*a11*a02*a23 - a10*a31*a02*a23 - a30*a01*a12*a23 + a00*a31*a12*a23 +
		a10*a01*a32*a23 - a00*a11*a32*a23 - a20*a11*a02*a33 + a10*a21*a02*a33 +
		a20*a01*a12*a33 - a00*a21*a12*a33 - a10*a01*a22*a33 + a00*a11*a22*a33

	return result
}

// MatrixNormalize - Normalize provided matrix
func MatrixNormalize(mat Mat) Mat {
	var result Mat

	det := MatrixDeterminant(mat)

	result.M0 /= det
	result.M1 /= det
	result.M2 /= det
	result.M3 /= det
	result.M4 /= det
	result.M5 /= det
	result.M6 /= det
	result.M7 /= det
	result.M8 /= det
	result.M9 /= det
	result.M10 /= det
	result.M11 /= det
	result.M12 /= det
	result.M13 /= det
	result.M14 /= det
	result.M15 /= det

	return result
}

// MatrixAdd - Add two matrices
func MatrixAdd(left, right Mat) Mat {
	result := MatrixIdentity()

	result.M0 = left.M0 + right.M0
	result.M1 = left.M1 + right.M1
	result.M2 = left.M2 + right.M2
	result.M3 = left.M3 + right.M3
	result.M4 = left.M4 + right.M4
	result.M5 = left.M5 + right.M5
	result.M6 = left.M6 + right.M6
	result.M7 = left.M7 + right.M7
	result.M8 = left.M8 + right.M8
	result.M9 = left.M9 + right.M9
	result.M10 = left.M10 + right.M10
	result.M11 = left.M11 + right.M11
	result.M12 = left.M12 + right.M12
	result.M13 = left.M13 + right.M13
	result.M14 = left.M14 + right.M14
	result.M15 = left.M15 + right.M15

	return result
}

// MatrixSubtract - Subtract two matrices (left - right)
func MatrixSubtract(left, right Mat) Mat {
	result := MatrixIdentity()

	result.M0 = left.M0 - right.M0
	result.M1 = left.M1 - right.M1
	result.M2 = left.M2 - right.M2
	result.M3 = left.M3 - right.M3
	result.M4 = left.M4 - right.M4
	result.M5 = left.M5 - right.M5
	result.M6 = left.M6 - right.M6
	result.M7 = left.M7 - right.M7
	result.M8 = left.M8 - right.M8
	result.M9 = left.M9 - right.M9
	result.M10 = left.M10 - right.M10
	result.M11 = left.M11 - right.M11
	result.M12 = left.M12 - right.M12
	result.M13 = left.M13 - right.M13
	result.M14 = left.M14 - right.M14
	result.M15 = left.M15 - right.M15

	return result
}

// MatrixMultiply - Returns two matrix multiplication
func MatrixMultiply(left, right Mat) Mat {
	var result Mat

	result.M0 = left.M0*right.M0 + left.M1*right.M4 + left.M2*right.M8 + left.M3*right.M12
	result.M1 = left.M0*right.M1 + left.M1*right.M5 + left.M2*right.M9 + left.M3*right.M13
	result.M2 = left.M0*right.M2 + left.M1*right.M6 + left.M2*right.M10 + left.M3*right.M14
	result.M3 = left.M0*right.M3 + left.M1*right.M7 + left.M2*right.M11 + left.M3*right.M15
	result.M4 = left.M4*right.M0 + left.M5*right.M4 + left.M6*right.M8 + left.M7*right.M12
	result.M5 = left.M4*right.M1 + left.M5*right.M5 + left.M6*right.M9 + left.M7*right.M13
	result.M6 = left.M4*right.M2 + left.M5*right.M6 + left.M6*right.M10 + left.M7*right.M14
	result.M7 = left.M4*right.M3 + left.M5*right.M7 + left.M6*right.M11 + left.M7*right.M15
	result.M8 = left.M8*right.M0 + left.M9*right.M4 + left.M10*right.M8 + left.M11*right.M12
	result.M9 = left.M8*right.M1 + left.M9*right.M5 + left.M10*right.M9 + left.M11*right.M13
	result.M10 = left.M8*right.M2 + left.M9*right.M6 + left.M10*right.M10 + left.M11*right.M14
	result.M11 = left.M8*right.M3 + left.M9*right.M7 + left.M10*right.M11 + left.M11*right.M15
	result.M12 = left.M12*right.M0 + left.M13*right.M4 + left.M14*right.M8 + left.M15*right.M12
	result.M13 = left.M12*right.M1 + left.M13*right.M5 + left.M14*right.M9 + left.M15*right.M13
	result.M14 = left.M12*right.M2 + left.M13*right.M6 + left.M14*right.M10 + left.M15*right.M14
	result.M15 = left.M12*right.M3 + left.M13*right.M7 + left.M14*right.M11 + left.M15*right.M15

	return result
}

// MatrixTranslate - Returns translation matrix
func MatrixTranslate(x, y, z float32) Mat {
	return M(
		1.0, 0.0, 0.0, x,
		0.0, 1.0, 0.0, y,
		0.0, 0.0, 1.0, z,
		0, 0, 0, 1.0)
}

// MatrixRotateX - Returns x-rotation matrix (angle in radians)
func MatrixRotateX(angle float32) Mat {
	result := MatrixIdentity()

	cosres := float32(math.Cos(float64(angle)))
	sinres := float32(math.Sin(float64(angle)))

	result.M5 = cosres
	result.M6 = -sinres
	result.M9 = sinres
	result.M10 = cosres

	return result
}

// MatrixRotateY - Returns y-rotation matrix (angle in radians)
func MatrixRotateY(angle float32) Mat {
	result := MatrixIdentity()

	cosres := float32(math.Cos(float64(angle)))
	sinres := float32(math.Sin(float64(angle)))

	result.M0 = cosres
	result.M2 = sinres
	result.M8 = -sinres
	result.M10 = cosres

	return result
}

// MatrixRotateZ - Returns z-rotation matrix (angle in radians)
func MatrixRotateZ(angle float32) Mat {
	result := MatrixIdentity()

	cosres := float32(math.Cos(float64(angle)))
	sinres := float32(math.Sin(float64(angle)))

	result.M0 = cosres
	result.M1 = -sinres
	result.M4 = sinres
	result.M5 = cosres

	return result
}

// MatrixScale - Returns scaling matrix
func MatrixScale(x, y, z float32) Mat {
	result := M(
		x, 0.0, 0.0, 0.0,
		0.0, y, 0.0, 0.0,
		0.0, 0.0, z, 0.0,
		0.0, 0.0, 0.0, 1.0)

	return result
}
