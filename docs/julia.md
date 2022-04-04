## Table
Table containing average time to run on my system for diffrent versions of the Julia program. The average is calculated on the results of three tries.

| Version | Time (sec) |
|:-------:|:----:|
| Original | 14.7 |
| Go Julia | 5.5 |
| Go Julia + Go CreatePng | 4.8 |

## Some conclusions
The program runs each image in different goroutinesm, splits each image in 16 strips that runs in different goroutines and was capped to  4 cores. No further improvement was made when spliting the image further or increasing the max amount of cores.