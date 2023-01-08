# 192. Word Frequency

Write a bash script to calculate the
frequency
of each word in a text file words.txt.

For simplicity sake, you may assume:

    words.txt contains only lowercase characters and space ' ' characters.
    Each word must consist of lowercase characters only.
    Words are separated by one or more whitespace characters.


[preference](https://stackoverflow.com/questions/10552803/how-to-create-a-frequency-list-of-every-word-in-a-file)

```sh
cat words.txt | tr -s ' ' '\n' | sort| uniq -c | sort -r | awk '{print $2" "$1}' 
```