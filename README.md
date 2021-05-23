# Problem statement
Write a program that given a directory as an argument, finds all duplicate files (with same contents) there and prints their names. No need to enter sub-directories - the first level is fine.
For example, it should print
- fileA fileB
- fileS fileT fileV
- If A,B are identical between themselves and S,T,V are identical between themselves.
- It should not print "single" files without duplicates.
The directory may contain dozens of thousands of files with different sizes from very small to very big. Two files are identical if their contents are the same using bytewise comparison.