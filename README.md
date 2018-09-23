# Simulated Annealing using Golang

## Description
Simulated Annealing is an Algorithm that search for maximum/minimum value of a function.

This repository is my course work project.

The main problem of the project is to search the minimum value of the problem.

## Visualization

![](https://svgshare.com/i/8PY.svg)
 
>NB: The **bigger** the circle, the **smaller** the value

> It is located inside visualization folder on the repository

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

* Temperature: 100

Temperature is being set become 1000 for more accuration purpose. The changes of temperature will effect the number of iterations

* Alpha: 0.8 - 0.95

Alpha is being set become 0.8 to 0.95 to maximize the accuration also. the changes of alpha will effect the number of iterations.

* Minimum Iterations: 1000

Minimum iterations is the number of iterations must be done before doing Temperature cooling.

It being set to be 1000 to increase the accuration of the model

* Goal: -19.208501

The Goal is actually obtained by running The Simulated Annealing with Unlimited Temperatrue.
The best Solution that I've reached is -19.208501


## Current optimum solution
```
Temperature: 100.00
Alpha: 0.95
Minimum Value: -10.00
Maximum Value: 10.00
Goal: -19.208501
Minimum Temperature: 0.0001
Loop: 10
fV:= -8.025963 sV:= 9.697703 E:= -19.189334, Accuration:= 99.900214%
fV:= -7.990458 sV:= 9.825867 E:= -18.902097, Accuration:= 98.404851%
fV:= -8.010391 sV:= -9.619297 E:= -19.168004, Accuration:= 99.789171%
fV:= -8.014729 sV:= -9.612083 E:= -19.164561, Accuration:= 99.771247%
fV:= 8.025812 sV:= 9.530713 E:= -19.022823, Accuration:= 99.033353%
fV:= -7.889956 sV:= -9.697192 E:= -18.937114, Accuration:= 98.587153%
fV:= -7.987430 sV:= -9.484198 E:= -18.844195, Accuration:= 98.103413%
fV:= 8.086753 sV:= 9.655051 E:= -19.197690, Accuration:= 99.943718%
fV:= -8.119189 sV:= -9.708009 E:= -19.147802, Accuration:= 99.683999%
fV:= -8.034423 sV:= -9.605722 E:= -19.169468, Accuration:= 99.796792%

```
> fV = First Value, sV = Second Value, E =  Energy Cost

## Conclusion
How the problem(CostFunction) solved by Simmulated Annealing is by picking a random function inside the range.
The number of picking itself is based on the Temperature cooling and minimal iterations.
Every solution that are being generated randomly are not directly thrown away. To maximize the accuration,
every solution will be randomly picked or not picked by a certain probabilty that being calculated using AcceptanceProbability Function.


The Data visualization shows that the mainimum value of the problem are located at the very end of every quadrant inside of the cartesian axis.
It basically tells that the minimum value can be reached on the maximum or minumum value on both x1 and x2.

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
## Environment Variables
It is located inside .env
```
TEMPERATURE: Temperature that used as the params.
ALPHA: Cooling rate of the Temperature
MIN_ITERATION: Minimum iteration before start cooling
MIN_VALUE: Minimum Value of the solution point
MAX_VALUE: Maximum Value of the solution point
GOAL: Benchmark of the solution
MIN_TEMPERATURE: The value that end the algorithm
LOOP: The number of how many times the Algorithm will run
```

## Running the program
Modify the params by editing the .env file
``` 
 $ cd $GOPATH/src/github.com/pamungkaski/go-simulated-annealing
 $ go run app/main.go
```

## Disclaimer
> **DO NOT USE THIS PROJECT AS YOUR SUBMISSION FOR AI COURSE WORK ASSIGNMENT**