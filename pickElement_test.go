package main 


import "testing"

func TestGetSmallest(t *testing.T){

    input := []int{-1,1,2,3,-4,5}
    got := getSmallest(input)
    want := -4
    
    if got != want {
        t.Errorf(" got %q, wanted %q",got, want);
    }

}


func TestGetLargest(t *testing.T){

    input := []int{-1,1,2,3,-4,5}
    got := getLargest(input)
    want := 5 
    
    if got != want {
        t.Errorf(" got %q, wanted %q",got, want);
    }

}


