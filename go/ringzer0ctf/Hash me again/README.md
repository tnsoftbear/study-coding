# Ringzer0cfg "Hash me again" challenge

Author: Igor Kot

Solution for [Coding Challenge #14](https://ringzer0ctf.com/challenges/14)

This solution splits zero-one input to 8 character slices, because we want to represent it as array of binary strings, where each character represents a bit and 8 characters represent a string of binary byte.  
We transform this string of binary byte to integer and use it to represent ASCII character.  
We implode all detected ASCII characters to string and use it as input for SHA-512 hashing function.  
Then we send this hashed value to respective url and get desired flag.

Use `github.com/gocolly/colly` library for:

* loading challenge page with initial input and scrapping it for parsing;
* sending the determined hash result to respective challenge url;
* scrapping the flag from the result page with success response.
