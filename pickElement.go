package main

import "fmt"

func getSmallest(input []int) int {
    
    smallest := 9999999999999;

    for i := 0; i < len(input); i++ {
        if input[i] <= smallest{
           
            smallest = input[i]
        }

    }

    return smallest;
}



func main(){
    
}
