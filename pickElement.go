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
    input := []int{1,2,3,-5,2}

    fmt.Println(getSmallest(input))
    fmt.Println(getLargest(input))
}



