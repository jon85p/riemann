package main

import (
	"fmt"
	"math"
	"math/cmplx"

	"gonum.org/v1/gonum/integrate/quad"
)

var s complex128 = complex(2.0, 0.0)

func f() complex128 {
	// Regresa la primera parte de la estimación
	return cmplx.Pow(2, s-1) / (1 - cmplx.Pow(2, 1-s))
}

func g(t float64) complex128 {
	var frac1 complex128 = cmplx.Cos(s * complex(math.Atan(t), 0))
	var frac2_1 complex128 = cmplx.Pow(complex(t*t+1, 0), s/2)
	var frac2_2 float64 = math.Cosh(math.Pi * t / 2)
	return frac1 / (frac2_1 * complex(frac2_2, 0))
}

func gReal(t float64) float64 {
	var res complex128 = g(t)
	return real(res)
}

func gImag(t float64) float64 {
	var res complex128 = g(t)
	return imag(res)
}

func zeta(x complex128) complex128 {
	s = x
	// Ahora vamos a realizar el cálculo
	var real float64
	var imaginaria float64
	var fs = f()
	real = quad.Fixed(gReal, 0, math.Inf(1), 1000, nil, 0)
	imaginaria = quad.Fixed(gImag, 0, math.Inf(1), 1000, nil, 0)
	var out complex128 = fs * complex(real, imaginaria)
	return out
}

func main() {
	fmt.Println("Cálculo de Riemann")
	// Ahora vamos a probar una variable global, jé
	s = complex(-2.0, 0)
	// test := zeta(s)
	// fmt.Println(test)
	// La prueba ahora será estimar al menos algún valor para el cual
	// la función sea un cero no trivial
	var z complex128
	var umbral float64 = 1e-18
	for i := 0; i < 500000; i++ {
		if i%10000 == 0 {
			fmt.Printf("Vamos acá %v\n", i)
		}
		var pR float64 = 0.5
		var pI float64 = float64(i) / 100000000
		s = complex(pR, 14.134725109900+pI)
		z = zeta(s)
		// fmt.Printf("%v \n", z)
		if real(z)*real(z)+imag(z)*imag(z) < umbral {
			fmt.Printf("Hemos encontrado un posible cero en %v, con valor %v \n", s, z)
			return
		}
	}
}