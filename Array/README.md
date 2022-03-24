#### **********************************
#### Array in Golang
#### **********************************

#### Declaration of array

##### 1. Using keyword var

        Var array_name[length]Type

#### Example:

         var arr1[5]int32

#### 2. Using shorthand declaration:

        array_name:= [length]Type{item1, item2, item3,...itemN}

        Example:

        arr1 := [5]int32{1, 2, 3, 4, 5}


#### Iterate the array

        Example:

        for i:=0 ; i <= len(array); i++ {
            fmt.Println(array[i])
        }




#### 2D Array:

        arr:= [3][3]int32{{1, 2, 3}, 
                           {10, 20, 30},
                            {200, 200, 300},}


        fmt.Println("Elements of Array 1")
        for x:= 0; x < 3; x++{
        for y:= 0; y < 3; y++{
        fmt.Println(arr[x][y])
        }
        }

#### Some Important Facts:
        myarray:= [...]int32{1, 2, 3, 4, 5}

        fmt.Println("Elements of the array: ", myarray) // {1, 2, 3, 4, 5}
        fmt.Println("Length of the array is:", len(myarray)) // 5

#### ****************************************
#### Slices in Golang
#### ****************************************
A slice is declared just like an array, but it doesn’t contain the size of the slice. So it can grow or shrink according to the requirement. 

        1. Syntax:  

        var my_slice[]int
        var my_slice[]int{}
        var my_slice[]int{1, 2, 3} 


        2. Using make() function:

        func make([]int, length, capacity)

        example:
        var my_slice_1 = make([]int, 4, 7)

        var my_slice_2 = make([]int, 7)


#### Iterate the slice:

    1.  myslice := []int{1, 2, 3, 4 5, 6, 7}
    
        // Iterate using for loop
        for e := 0; e < len(myslice); e++ {
            fmt.Println(myslice[e])
        }

    2. for index, ele := range myslice {
            fmt.Printf("Index = %d and element = %s\n", index+3, ele)
        }

    3.  for _, ele := range myslice {
            fmt.Printf("Element = %s\n", ele)
        }


#### Important facts of slices:

    1. var my_slice = []int{1, 2, 3, 4, 5}

        println(my_slice[2:]) // {4, 5}
