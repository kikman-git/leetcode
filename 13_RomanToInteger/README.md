# 13. Roman to Integer

Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.
`
Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
`

## My dump solution:
* First make a map that store roman symbols and their coresponding number
* The map also contains double symbols like IV =4
* If len of Roman <= 1 then return
* If its > 1 then loop through every char of Roman string:
  * i starts from 0
  * result = 0
  * first check if char[i] and char[i+1] (2 chars next to each other) exists in the map
  * then check if char[i] char[i+1] and char[i+2] are the same -> look up their int values and sum up
  * else look up the char's integer and add to result