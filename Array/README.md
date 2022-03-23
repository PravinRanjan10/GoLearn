#### **********************************
#### Array in Golang
#### **********************************

#### Declaration of array

##### 1. Using keyword var

        Var array_name[length]Type

        or

        var array_name[length]Typle{item1, item2, item3, ...itemN}


#### Example:

         var arr1[5]int32

         var arr1[5]int32{1, 2, 3, 4, 5}

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
