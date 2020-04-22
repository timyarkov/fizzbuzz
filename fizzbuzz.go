package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Get num of iterations to do
	fmt.Println("Input number of iterations to do:")
	reader := bufio.NewReader(os.Stdin) //:= only once for this; after, type doesn't need further inferring
	numstr, _ := reader.ReadString('\n')
	num, err := strconv.ParseInt(strings.TrimSpace(numstr), 10, 64) //Converts to an int, base 10, bitsize 64.
	errorCheck(err)                                                 //Trimspace needed or else \n wont let int be parsed
	if num <= 0 {
		panic("You can't do that :(")
	}

	//Get num of factors to get
	fmt.Println("Input number of factors:")
	reader = bufio.NewReader(os.Stdin)
	divnumstr, _ := reader.ReadString('\n')
	divnum, err := strconv.ParseInt(strings.TrimSpace(divnumstr), 10, 64) //Converts to an int, base 10, bitsize 64
	errorCheck(err)
	if num <= 0 {
		panic("You can't do that :(")
	}

	//Get factors and associated words
	divlsnum := make([]int64, 0, 2)
	divlsword := make([]string, 0, 2)
	var i int64 = 0 //Declare i as int64, or else will have a type mismatch between that and parsed ints
	var newdivnum int64 = 0
	var newdivnumstr string = ""
	var newdivword string = ""
	for i < divnum {
		//Get Factor
		fmt.Printf("Input factor %v / %v:\n", (i + 1), divnum)
		reader = bufio.NewReader(os.Stdin)
		newdivnumstr, _ = reader.ReadString('\n')
		newdivnum, err = strconv.ParseInt(strings.TrimSpace(newdivnumstr), 10, 64)
		errorCheck(err)
		divlsnum = append(divlsnum, newdivnum)
		//Get Word
		fmt.Printf("Input word for factor %v, %v / %v:\n", newdivnum, (i + 1), divnum)
		reader = bufio.NewReader(os.Stdin)
		newdivword, _ = reader.ReadString('\n')
		divlsword = append(divlsword, newdivword)
		i++ //Dont forget to iterate silly
	}

	//Every factor of 3 is Fizz, every factor of 5 is Buzz, factor of 3 and 5 is FizzBuzz
	i = 0
	var j int64 = 0
	var result string = ""
	for i < num {
		result = ""
		j = 0
		for j < divnum {
			if (i+1)%divlsnum[j] == 0 {
				result += strings.TrimSpace(divlsword[j])
			}
			j++
		}
		if result == "" {
			fmt.Println(i + 1)
		} else {
			fmt.Println(result)
		}
		i++
	}

}

//Error Check function that can be called easily
func errorCheck(err error) {
	if err != nil {
		panic(err)
	}
}
