## Synopsis

This is a basic program in GO that grabs the input from a file (chars.txt) and then chooses a random line from that file and outputs it to a new file (output.txt). 

If the random line exists in output.txt already, the code with shuffle the slice and select again. It will do this until it finds one that is not present in output.txt.

If input (chars.txt) is the same length as the output, the code assumes all options have been exhausted and exits.

## Motivation

I basically made this to randomly choose things for me, but also keep track of what has been chosen in the past.

## To-Do

Clean up the code, add some comments maybe.
~~Make it read the output file to do "something"~~
Allow the manipulation of input (chars.txt). Perhaps to add/remove before continuing.

## Contributors

Just me
