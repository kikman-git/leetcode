Given an integer x, return true if x is a palindrome , and false otherwise.

* An integer is a palindrome when it reads the same forward and backward.

For example, 121 is a palindrome while 123 is not.

## Dump solution (my solution):
* Make a temp array to store every digit
* Divide the int by 10 until it is less than 0
* compare last digit with first digit until middle length of array.
* If all equal then it is an palindrome