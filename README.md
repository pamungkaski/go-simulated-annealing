# Simulated Annealing using Golang

## Description
Simulated Annealing is an Alogirthm that search for maximum/minimum value of a function.

This repository is my course work project.

## Cost function
![](https://image.ibb.co/cpKPZz/Screen_Shot_2018_09_21_at_21_56_18.png)

```go
func (a *Annealing) CostFunction(firstValue, secondValue float64) float64 {
	var result float64
	fVSquare := firstValue * firstValue
	sVSquare := secondValue * secondValue

	fVSin := math.Sin(firstValue)
	sVCos := math.Cos(secondValue)

	subtrahend := math.Sqrt(fVSquare+sVSquare) / math.Pi
	insideExponential := math.Abs(1 - subtrahend)
	exponent := math.Exp(insideExponential)

	result = -1 * math.Abs(fVSin*sVCos*exponent)

	return result
}
```

## Main Function
Please open the detailed file [main.go](https://github.com/pamungkaski/go-simulated-annealing/blob/master/app/main.go)

The comment is already tell how the program work.

## Current Parameter
Temperature is being set become 1000 for more accuration purpose. The changes of temperature will effect the number of iterations

* Temperature: 1000

Alpha is being set become 0.8 to 0.95 to maximize the accuration also. the changes of alpha will effect the number of iterations.

* Alpha: 0.8 - 0.95

Minimum iterations is the number of iterations must be done before doing Temperature cooling.

It being set to be 1000 to increase the accuration of the model

* Minimum Iterations: 1000

## Current optimum solution
* x1:= 8.019960 
* x2:= 9.668527 
* E:= -19.196317
* Accuration:= 99.936577%

## Stack
* Go <= 1.10.3
* Already set **$GOPATH** and **$GOROOT**

## Dependency
* [dep](https://github.com/golang/dep)
* [gonum](https://github.com/gonum/gonum)
* [plot](https://github.com/gonum/plot)

## Installation
```
 $ go get github.com/pamungkaski/go-simulated-annealing
 $ cd $GOPATH/src/github.com/pamungkaski/go-simulated-annealing
 $ dep ensure
```

## Running the program
``` 
 $ cd $GOPATH/src/github.com/pamungkaski/go-simulated-annealing
 $ go run app/main.go
```

## Visualization
It is located inside visualization folder on the repository

![](https://svgshare.com/i/8PY.svg)
 
>NB: The **bigger** the circle, the **smaller** the value