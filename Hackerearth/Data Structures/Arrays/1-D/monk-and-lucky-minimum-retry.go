/*
Question Link:https://www.hackerearth.com/practice/data-structures/arrays/1-d/practice-problems/algorithm/monk-and-lucky-minimum-3/
Time complexity: O(n)
Verdict:Accepted. Time-taken : 8s over all inputs. C++ programmers got it done in around 1s.
*/
package main

import(
	"fmt"
	"bufio"
	"os"
	)
var in=bufio.NewReader(os.Stdin)
func main() {
	var test int                    //test:total test cases
	fmt.Fscanf(in,"%d\n",&test)     //Notice "%d\n" for eliminating newline 
	for test>0 {
		var size,setMin int         //size:input array size, setMin:minimum int in the array
		fmt.Fscanf(in,"%d\n",&size)
		setMin,count:=calcMinCount(size,setMin) //count:no. of occurrence of the minimum element
		fmt.Fscanf(in,"\n")
		if count%2==0{
			fmt.Println("Unlucky")
		} else{
			fmt.Println("Lucky")
		}
		test--
	}
}

func calcMinCount(size,setMin int)(int,int){
	fmt.Fscanf(in,"%d",&setMin)
		count:=1                    //count:no. of occurrence of the minimum element
		for i:=1;i<size;i++{
			var x int               //x:array elements input stored here
			fmt.Fscanf(in,"%d",&x)
			if x<setMin{
				setMin=x
				count=1
			} else if x==setMin {
				count++
			}
		}
		return setMin,count
}
