/*
Question Link:https://www.hackerearth.com/practice/data-structures/arrays/1-d/practice-problems/algorithm/monk-and-lucky-minimum-3/
Time complexity: O(n)
*/
package main

import(
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	)

var in=bufio.NewReader(os.Stdin)

func main() {
	var test int                    //test:total test cases
	fmt.Fscanf(in,"%d\n",&test)     //Notice "%d\n" for eliminating newline 

	for test>0 {
		var size,setMin int         //size:input array size, setMin:minimum int in the array
		fmt.Fscanf(in,"%d\n",&size)

		setMin,count:=calcMinCount(size,setMin) //count:no. of occurrence of the minimum element

		if count%2==0{
			fmt.Println("Unlucky")
		} else{
			fmt.Println("Lucky")
		}

		test--
	}
}

func calcMinCount(size,setMin int)(int,int){
	lineInput,_ := in.ReadString('\n') //Reads the entire array as a string

	set:= []int(makeIntArray(lineInput)) //convert the string to integer slice
	setMin = set[0]

	count:=1                    //count:no. of occurrence of the minimum element

	for i:=1;i<size;i++{

		if set[i]<setMin{
			setMin=set[i]
			count=1
		}else if set[i]==setMin {
			count++
		}
	}

	return setMin,count
}

func makeIntArray(stream string)[]int{
	// Trims the stream and then splits
  trimmed_stream := strings.TrimSpace(stream)

  split_arr := strings.Split(trimmed_stream, " ")
  // convert strings to integers and store them in a slice

  clean_arr := make([]int, len(split_arr))

  for index,val := range split_arr{
  
  //http://stackoverflow.com/questions/31333353/faster-input-scanning icza's toInt function claims to be faster than strconv
    clean_arr[index], _ = strconv.Atoi(val) 
  }

  return clean_arr
}
