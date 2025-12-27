package main

import problems "project-euler-go/problems" // Importing problems package

var Solutions = map[int]func(int) int{
	1: problems.MultiplesOf3and5, 	//Multiples of 3 and 5
	2: problems.FibonacciSum,        //Even Fibonacci numbers
	3: problems.LargestPrimeFactor,  //Largest prime factor
}
