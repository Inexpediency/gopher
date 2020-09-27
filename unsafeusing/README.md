# C code from Go code

## How to build library C code file

`$ cd ./unsafeusing`
`$ gcc -c callClib/*.c`
`$ /usr/bin/ar rs callC.a *.o`
`$ rm callC.o`

# Go code from C code (placed in `./main` folder)

## How to generate from Go code commmon C library 

`$ cd ./unsafeusing/main`
`$ go build -o usedByC.o -buildmode=c-shared usedByC.go`

## Build C code and run

`$ gcc -o willUseGo willUseGo.c ./usedByC.o`
`$ ./willUseGo`
