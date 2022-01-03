package cses

import (
	"dsAlgo/array"
	"fmt"
	"math"
)

// Searching Sorting
//Distinct numbers

// You are given a list of n integers, and your task is to calculate the number of distinct values in the list.

// Input

// The first input line has an integer n: the number of values.

// The second line has n integers x1,x2,…,xn.

// Output

// Print one integers: the number of distinct values.

// Constraints

//     1≤n≤2⋅105

// 1≤xi≤109

// Example

// Input:
// 5
// 2 3 2 2 3

// Output:
// 2

func DistictValues(ar []int) int {
	c := 0
	occurence := make(map[int]int)
	for _, val := range ar {
		if occurence[val] == 0 {
			occurence[val]++
			c++
		}
	}
	return c
}

//Appartmrnts

// There are n applicants and m free apartments. Your task is to distribute the apartments so that as many applicants as possible will get an apartment.

// Each applicant has a desired apartment size, and they will accept any apartment whose size is close enough to the desired size.

// Input

// The first input line has three integers n, m, and k: the number of applicants, the number of apartments, and the maximum allowed difference.

// The next line contains n integers a1,a2,…,an: the desired apartment size of each applicant. If the desired size of an applicant is x, he or she will accept any apartment whose size is between x−k and x+k.

// The last line contains m integers b1,b2,…,bm: the size of each apartment.

// Output

// Print one integer: the number of applicants who will get an apartment.

// Constraints

//     1≤n,m≤2⋅105

// 0≤k≤109

// 1≤ai,bi≤109

// Example

// Input:
// 4 3 5
// 60 45 80 60
// 30 60 75

// Output:
// 2

func Apartments(applicants, appartments []int, diff int) int {
	c := 0
	y := appartments
	for _, val := range applicants {
		if pres, index := array.BinarySearch(y, val+diff, 0, len(y)-1); pres {
			y = pop(appartments, index)

		}
		if pres, index := array.BinarySearch(y, val-diff, 0, len(y)-1); pres {
			c++
			y = pop(appartments, index)

		}
		if pres, index := array.BinarySearch(y, val, 0, len(y)-1); pres {
			c++
			y = pop(appartments, index)

		}
	}
	return c
}

func pop(arr []int, i int) []int {

	ar := append(arr[:i], arr[i+1:]...)
	return ar
}

// There are n children who want to go to a Ferris wheel, and your task is to find a gondola for each child.

// Each gondola may have one or two children in it, and in addition, the total weight in a gondola may not exceed x. You know the weight of every child.

// What is the minimum number of gondolas needed for the children?

// Input

// The first input line contains two integers n and x: the number of children and the maximum allowed weight.

// The next line contains n integers p1,p2,…,pn: the weight of each child.

// Output

// Print one integer: the minimum number of gondolas.

// Constraints

//     1≤n≤2⋅105

// 1≤x≤109

// 1≤pi≤x

// Example

// Input:
// 4 10
// 7 2 3 9

// Output:
// 3

func FerrisWheel(studentWeights []int, maxWeight int) int {
	c := 0
	skipIndex := make(map[int]bool)
	i := 0
	j := i + 1
	for i < len(studentWeights)-1 && j < len(studentWeights) {
		want := maxWeight - studentWeights[i]
		if !skipIndex[j] {
			if want >= studentWeights[j] {
				c++
				skipIndex[j] = true
				i++
			}
		}
		j++
		i++
	}
	return len(studentWeights) - c
}

// There are n concert tickets available, each with a certain price. Then, m customers arrive, one after another.

// Each customer announces the maximum price they are willing to pay for a ticket, and after this, they will get a ticket with the nearest possible price such that it does not exceed the maximum price.

// Input

// The first input line contains integers n and m: the number of tickets and the number of customers.

// The next line contains n integers h1,h2,…,hn: the price of each ticket.

// The last line contains m integers t1,t2,…,tm: the maximum price for each customer in the order they arrive.

// Output

// Print, for each customer, the price that they will pay for their ticket. After this, the ticket cannot be purchased again.

// If a customer cannot get any ticket, print −1.

// Constraints

//     1≤n,m≤2⋅105

// 1≤hi,ti≤109

// Example

// Input:
// 5 3
// 5 3 7 8 5
// 4 8 3

// Output:
// 3
// 8
// -1

func ConcertTickets(customers, tickets []int) []int {
	c := []int{}
	skipIndex := make(map[int]bool)
	for i := 0; i < len(customers); i++ {
		min := math.MaxInt
		for j := 0; j < len(tickets); j++ {
			if !skipIndex[j] {
				if tickets[j] <= customers[i] {
					if customers[i]-tickets[j] < min {
						min = customers[i] - tickets[j]
					}
					skipIndex[j] = true
				}
			}
		}
		if customers[i]-min < 0 {
			c = append(c, -1)
		} else {
			c = append(c, customers[i]-min)
		}

	}
	return c
}

// ou are given the arrival and leaving times of n customers in a restaurant.

// What was the maximum number of customers in the restaurant at any time?

// Input

// The first input line has an integer n: the number of customers.

// After this, there are n lines that describe the customers. Each line has two integers a and b: the arrival and leaving times of a customer.

// You may assume that all arrival and leaving times are distinct.

// Output

// Print one integer: the maximum number of customers.

// Constraints

//     1≤n≤2⋅105

// 1≤a<b≤109

// Example

// Input:
// 3
// 5 8
// 2 4
// 3 9

// Output:
// 2

func ResturantCustomers(arrival, departure []int) int {
	c := 0
	skipIndex := make(map[int]bool)
	for i := 0; i < len(arrival); i++ {
		skipIndex[i] = true
		for j := 0; j < len(departure); j++ {
			if !skipIndex[j] {
				if arrival[i] < departure[j] && arrival[j] < arrival[i] {
					c++
				}
			}
			skipIndex[j] = false
		}
	}
	return c
}

// In a movie festival n movies will be shown. You know the starting and ending time of each movie. What is the maximum number of movies you can watch entirely?

// Input

// The first input line has an integer n: the number of movies.

// After this, there are n lines that describe the movies. Each line has two integers a and b: the starting and ending times of a movie.

// Output

// Print one integer: the maximum number of movies.

// Constraints

//     1≤n≤2⋅105

// 1≤a<b≤109

// Example

// Input:
// 3
// 3 5
// 4 9
// 5 8

// Output:
// 2

func MovieFestival(startTime, endTime []int) int {
	c := 1
	skipIndex := make(map[int]bool)
	for i := 0; i < len(endTime); i++ {
		skipIndex[i] = true
		for j := 0; j < len(startTime); j++ {
			if !skipIndex[j] {
				if endTime[i] <= startTime[j] {
					c++
				}
			}
			skipIndex[j] = false
		}
	}
	return c
}

// CSES
// CSES Problem Set
// Sum of Two Values

//     Task Statistics

//     Time limit: 1.00 s Memory limit: 512 MB

// You are given an array of n
// integers, and your task is to find two values (at distinct positions) whose sum is x.

// Input

// The first input line has two integers n and x: the array size and the target sum.

// The second line has n integers a1,a2,…,an: the array values.

// Output

// Print two integers: the positions of the values. If there are several solutions, you may print any of them. If there are no solutions, print IMPOSSIBLE.

// Constraints

//     1≤n≤2⋅105

// 1≤x,ai≤109

// Example

// Input:
// 4 8
// 2 7 5 1

// Output:
// 2 4

func SumOfTwoValues(arr []int, val int) (int, int) {
	for i := 1; i < len(arr)-1; i++ {
		want := val - arr[i]
		for j := i - 1; j >= 0; j-- {
			if arr[j] == want {
				return i + 1, j + 1
			}
		}
		for k := i + 1; k < len(arr); k++ {
			if arr[k] == want {
				return i + 1, k + 1
			}
		}
	}
	return -1, -1
}

// Maximum Subarray Sum

//     Task Statistics

//     Time limit: 1.00 s Memory limit: 512 MB

// Given an array of n
// integers, your task is to find the maximum sum of values in a contiguous, nonempty subarray.

// Input

// The first input line has an integer n: the size of the array.

// The second line has n integers x1,x2,…,xn: the array values.

// Output

// Print one integer: the maximum subarray sum.

// Constraints

//     1≤n≤2⋅105

// −109≤xi≤109

// Example

// Input:
// 8
// -1 3 -2 5 3 -5 2 2

// Output:
// 9

func SubArraySum(ar []int) int {
	maxSum := math.MinInt
	for i := 0; i < len(ar); i++ {
		for j := i; j < len(ar)-1; j++ {
			if maxSum < sum(ar[i:j+1]) {
				maxSum = sum(ar[i : j+1])
			}
		}
	}
	return maxSum

}

func sum(ar []int) int {
	sum := 0
	for _, val := range ar {
		sum += val
	}
	return sum
}

// CSES
// CSES Problem Set
// Stick Lengths

//     Task Statistics

//     Time limit: 1.00 s Memory limit: 512 MB

// There are n
// sticks with some lengths. Your task is to modify the sticks so that each stick has the same length.

// You can either lengthen and shorten each stick. Both operations cost x where x is the difference between the new and original length.

// What is the minimum total cost?

// Input

// The first input line contains an integer n: the number of sticks.

// Then there are n integers: p1,p2,…,pn: the lengths of the sticks.

// Output

// Print one integer: the minimum total cost.

// Constraints

//     1≤n≤2⋅105

// 1≤pi≤109

// Example

// Input:
// 5
// 2 3 1 5 2

// Output:
// 5

func Sticks(sticks []int) int {
	minGlobal := math.MaxInt
	minLocal := 0
	for i := 0; i < len(sticks); i++ {
		minLocal = 0
		for j := 0; j < len(sticks); j++ {
			minLocal += int(math.Abs(float64(sticks[j]) - float64(sticks[i])))
		}
		if minGlobal > minLocal {
			minGlobal = minLocal
		}
	}
	return minGlobal
}

// CSES
// CSES Problem Set
// Playlist
//
// Task Statistics
//
// Time limit: 1.00 s Memory limit: 512 MB
//
// You are given a playlist of a radio station since its establishment. The playlist has a total of n
// songs.
//
// What is the longest sequence of successive songs where each song is unique?
//
// Input
//
// The first input line contains an integer n: the number of songs.
//
// The next line has n integers k1,k2,…,kn: the id number of each song.
//
// Output
//
// Print the length of the longest sequence of unique songs.
//
// Constraints
//
// 1≤n≤2⋅105
//
//
// 1≤ki≤109
//
//
// Example
//
// Input:
// 8
// 1 2 1 3 2 7 4 2
//
// Output:
// 5
//

func Playlist(ar []int) int {
	// y := make([]int, 0)
	c := 0
	occurence := make(map[int]int, 0)
	for i := 0; i < len(ar); i++ {
		if occurence[ar[i]] >= 1 {
			c++
		}
		occurence[ar[i]]++
	}
	fmt.Println(occurence)
	fmt.Println(c)
	return len(ar) - c
}

// CSES
// CSES Problem Set
// Traffic Lights

//     Task Statistics

//     Time limit: 1.00 s Memory limit: 512 MB

// There is a street of length x
// whose positions are numbered 0,1,…,x. Initially there are no traffic lights, but n sets of traffic lights are added to the street one after another.

// Your task is to calculate the length of the longest passage without traffic lights after each addition.

// Input

// The first input line contains two integers x and n: the length of the street and the number of sets of traffic lights.

// Then, the next line contains n integers p1,p2,…,pn: the position of each set of traffic lights. Each position is distinct.

// Output

// Print the length of the longest passage without traffic lights after each addition.

// Constraints

//     1≤x≤109

// 1≤n≤2⋅105

// 0<pi<x

// Example

// Input:
// 8 3
// 3 6 2

// Output:
// 5 3 3

func TrafficLights(ar []int, maxlength int) {

}
