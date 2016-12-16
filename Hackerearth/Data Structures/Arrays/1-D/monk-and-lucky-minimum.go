/*
Question Link:https://www.hackerearth.com/practice/data-structures/arrays/1-d/practice-problems/algorithm/monk-and-lucky-minimum-3/
Time complexity: O(n)
Verdict: Time Limit exceeded. Problem lies with the scanning process. Apparently fmt.Scanf is too slow.
*/

package main

import "fmt"

func main() {
	var test int                    //test:total test cases
	fmt.Scanf("%d",&test)
	for test>0 {
		var size,setMin int         //size:input array size, setMin:minimum int in the array
		fmt.Scanf("%d",&size)
		count:=calcMinCount(size,setMin) //count:no. of occurrence of the minimum element
		if count%2==0{
			fmt.Println("UNLUCKY")
		} else{
			fmt.Println("LUCKY")
		}
		test--
	}
}

func calcMinCount(size,setMin int)(int){
	fmt.Scanf("%d",&setMin)
		count:=1                    //count:no. of occurrence of the minimum element
		for i:=1;i<size;i++{
			var x int               //x:array elements input stored here
			fmt.Scanf("%d",&x)
			if x<setMin{
				setMin=x
				count=1
			} else if x==setMin {
				count++
			}
		}
		return count
}
