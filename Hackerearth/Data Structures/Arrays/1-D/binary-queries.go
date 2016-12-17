/*
Question Link: https://www.hackerearth.com/practice/data-structures/arrays/1-d/practice-problems/algorithm/range-query-2/
Time Complexity: O(max(q,n)) ,correct me if i am wrong
Verdict: writing output to the buffer first in these sort of problems is the right way to go in go.
*/

package main

import(
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"bytes"
	)

var in=bufio.NewReader(os.Stdin) //Check https://github.com/Krashcan/My-GO-Code/projects/2


func main(){
	var n,q int                                   //n:input size,q:no. of queries
	var buffer bytes.Buffer
	fmt.Fscanf(in,"%d %d\n",&n,&q)

	lineInput,_:= in.ReadString('\n')             //lineInput for creating a string of stream
	set := []int(makeIntArray(lineInput))         //set: stores the new integer array from stream
	lineInput,_ = in.ReadString('x')			  //Taking stream input till the end of file

	splitQ := strings.Split(lineInput,"\n")       //converting input to desired format
	for i,_ := range splitQ{
		trimQ := strings.TrimSpace(splitQ[i])
		newQ := strings.Split(trimQ," ")
		qType := toInt(newQ[0])
		if qType==0{
			r := toInt(newQ[2])                          //r= right index value
			if set[r-1]==0{
				buffer.WriteString("EVEN\n")
			}else{
				buffer.WriteString("ODD\n")
			}
		}else{
			k := toInt(newQ[1])                         //k= the provided index value to perform query on
			if set[k-1]==0{
				set[k-1]=1
			}else{
				set[k-1]=0
			}
		}
	}
	fmt.Print(buffer.String())
}

func makeIntArray(stream string)[]int{
	// Trims the stream and then splits
  trimmed_stream := strings.TrimSpace(stream)

  split_arr := strings.Split(trimmed_stream, " ")
  // convert strings to integers and store them in a slice

  clean_arr := make([]int, len(split_arr))

  for index,val := range split_arr{

  clean_arr[index], _ = strconv.Atoi(val) 
  }

  return clean_arr
}
func toInt(buf string) (n int) {
    for i:=0;i<len(buf);i++ {
        n = n*10 + int(buf[i]-'0')
    }
    return
}
