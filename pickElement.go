package main

import "fmt"

func getSmallest() int {

    return -1;
}

func getLargest(input []int) int {
    largest := -1;
    
    for i := 0; i < len(input); i++ {
        if input[i] >= largest{
            largest = input[i];
        }
    }
    
    return largest;
}


func main(){
    
}
