# MOVING AVERAGES EXERCISE
This exercises evaluates critical thinking and problem solving skills, specifically coding and algorithms.
Use any tools or resources that you normally would for programming tasks.


## James L Note:
I took the pdf and extracted text and put in this document to keep things together.  My notes/comments will be inline with initials starting the line -JL

### INSTRUCTIONS
Write an implementation of the function specification below. Use whatever mainstream programming
language you are most comfortable with.  The code should be efficient and robust. Consider edge cases; what might happen with different inputs (e.g. various array lengths, data values, window sizes)? Describe any potential problems even if you can’t think of a solution for a particular case. A driver program is nice, but not required.

### -JL Discussian
Solution was created in Golang and notes are for V 1.15+ (one added testing dependency - managed in Go Modules).  


### SUBMISSIONS SHOULD CONTAIN THE FOLLOWING:
- Code that solves the problem above
- Instructions for how to compile and run your code
-- Two ways to run
1.  Download code
2.  #go get
3.  #go run mmodal.go "(3, [0, 1, 2, 3])" -jl quotes are required around input
4.  #go test -v ./
- A description of your solution and any assumptions it makes
- jl With requirements listed - I decided to create a solution that was a decent solution that could be implemented quickly and be fairly resilient. With that - problems could be run into problems with extremely large data input.  The program would probably max out with available system memory. For the integer window - a golang int was used and for the values - float64 was chosen. The solution loops through the values - adding new value and possible subtracting the max value - before average is calculated.  Somekind of streaming solution would be one way to get around this limitation.  Golang - testing at moment is basic and ideally would utilize a harness that exercise larger data sets.

- Analysis of the time complexity of your solution, preferably in terms of big-O notation
```
##########################################################
JL NOtes
Time complexity analysis
O(1).
Space complexity analysis
O(size).
##########################################################
```
### SPECIFICATION
A function that accepts inputs for an integer window size and an array of double-precision floating point
values, and returns an array of double-precision floating point values. The nth value of the output array
should be the average after processing the first n elements of the input array. The average should be
computed as a cumulative average until the number of elements reaches the window size, and as a
simple moving average using the window size afterward.
### EXAMPLES
- compute(3, [0, 1, 2, 3]) => [0, 0.5, 1, 2]
- compute(5, [0, 1, -2, 3, -4, 5, -6, 7, -8, 9]) => [0, 0.5, -0.333333333333333, 0.5, -0.4, 0.6, -0.8, 1, -1.2, 1.4]
