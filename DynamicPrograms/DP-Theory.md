Algorithm Approaches
A. Divide and Conquer (DAC)
B. Greedy Techniques
C. Dynamic Programming


A. Some Applications/ example of DAC
    * Finding min and max in a list of elements
    * Binary Search
    * Merge Sort
    * Quick Sort
    * Randamize Quick Sort
    * Selection Procedure
    * Strassen's Matrix multiplications
    * Power of an element
    * Karatsuba Algorithm
    * Cooley-Tukey Fast Fourier Transform (FFT) Algorithm 


B. Some Applications/ example of Greedy Techniques
    * Fractional knapsack problem
    * Job sequencing with deadlines
    * Minimum Cost Spanning Tree
        * Kruskal’s algorithm
        * Prim’s algorithm
    * Single Source Shortest Path
        * Dijkstra’s algorithm
        * Bellman-ford algorithm
    
    * Activity search problem
    * Huffman coding
    * Travelling salesperson problem
    * Ford- Fulkerson algorithm
    * Boruvka’s algorithm


C. Dynamic Programming
Dynamic Programming is a method to solve complex problems by breaking them down into smaller and simpler subproblems. By solving each subproblem only once and storing the results, it avoids redundant computations, leading to more efficient solutions for a wide range of problems. It isn’t practical when there aren’t any problems that overlap because it doesn’t make sense to store solutions to the issues that won’t be needed again.

# What is DP?
    * DP is enhanced recursion: DP is an optimization technique for recursive solutions.
    * DP is a technique which is used to solve complex problems by breaking them down into simpler sub-problems. 

# How do you identify that a problem is DP problem.
    if problem has two things:
        1. Choice [if you have choices among items which to choose or pick]
        2. Overlapping subproblems [every function work on very much same input]
        3. Optimal results [like min, max, largest, smallest]

# How to approach DP problems
    * Findout Subproblems
    * Store the results of each subproblems [also know as memoization]
    * Finally, calculate the result of the orignial problem by using result of subproblems [Top-down or Bottom-up]

# Approaches of dynamic programming
    * Top down 
    * Bottom-up

# DP-Patterns Or Useful question sets

    1. 0/1 Knapsack
        * Subset Sum
        * Equal sum partition
        * Count of subset sum
        * Minimum subset sum diff
        * Target sum
        * Number of subset sum and given difference 
    2. Unbounded Kapsack
    3. Fibbonnacci:
        * climbing a staircase
    4. LCS
    5. LIC
    6. Kadan's problem
    7. MCM
    8. DP on trees
    9. DP on Grid
    10. Others
