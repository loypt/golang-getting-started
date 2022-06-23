package main

import (
    "fmt"
    "math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
    return fmt.Sprintf("not avaiable for negative number and complex number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
    if(x < 0){
        return 0, ErrNegativeSqrt(x)
    }
    z := 1.0
    for {
        if math.Abs(z-(z-(z*z-x)/(z*2))) < 0.00000000001 { 
            return z, nil
        }else{
            temp := (z*z-x)/(z*2)
            fmt.Printf("temp: %v\n",temp)
            z = z - temp
        }
    }
}

func main() {
    fmt.Println(Sqrt(2))
    fmt.Println(Sqrt(-2))
}
