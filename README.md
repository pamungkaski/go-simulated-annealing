# Simulated Annealing using Golang

## Description
Simulated Annealing is an Alogirthm that search for maximum/minimum value of a function.

This repository is my course work project.

## Cost function
![](https://image.ibb.co/cpKPZz/Screen_Shot_2018_09_21_at_21_56_18.png)

## Current Parameter
* Temperature: 1000
* Alpha: 0.95
* Minimum Iterations: 1000

## Current optimum solution
* x1:= 8.019960 
* x2:= 9.668527 E:= -19.196317
* Accuration:= 99.936577%

## Stack
* Go <=1.10.3
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