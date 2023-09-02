package main

import (
    "bufio"
    "flag"
    "fmt"
    "math/big"
    "os"
    "strconv"
    "strings"
)

func main() {
    var nFlag bool
    var sFlag string
    flag.BoolVar(&nFlag, "n", false, "Print the number of the prime number")
    flag.StringVar(&sFlag, "s", "", "Start from the nth prime number in the file")
    flag.Parse()

    if len(flag.Args()) != 2 {
        fmt.Println("Usage: [-n] [-s file.txt] start and end sequence.")
        os.Exit(1)
    }

    startBig, ok := big.NewInt(0).SetString(flag.Arg(0), 10)
    if !ok {
        fmt.Println("Error reading the start sequence.")
        os.Exit(1)
    }

    endBig, ok := big.NewInt(0).SetString(flag.Arg(1), 10)
    if !ok {
        fmt.Println("Error reading the end sequence.")
        os.Exit(1)
    }

    if startBig.Cmp(endBig) > 0 {
        fmt.Println("The start sequence must be less than the end sequence.")
        os.Exit(1)
    }

    skipCount := big.NewInt(0)
    if sFlag != "" {
        file, err := os.Open(sFlag)
        if err != nil {
            fmt.Printf("Error opening file: %s\n", err.Error())
            os.Exit(1)
        }
        defer file.Close()

        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
            line := scanner.Text()
            if line[0] == '/' && line[1] == '/' {
                continue
            }
            fields := strings.Fields(line)
            if len(fields) > 0 && fields[0][0] != '/' {
                start, _ := new(big.Int).SetString(fields[0], 10)
                if nFlag && len(fields) > 2 {
                    skipCountInt, _ := strconv.Atoi(fields[2])
                    skipCount.SetInt64(int64(skipCountInt))
                } else {
                    skipCount.Add(skipCount, big.NewInt(1))
                }
                startBig.Set(start)
            }
        }

        if err := scanner.Err(); err != nil {
            fmt.Printf("Error reading file: %s\n", err.Error())
            os.Exit(1)
        }

        startBig.Add(startBig, big.NewInt(1))
    }

    one := big.NewInt(1)
    count := new(big.Int).Set(skipCount)
    for i := new(big.Int).Set(startBig); i.Cmp(endBig) <= 0; i.Add(i, one) {
        if i.ProbablyPrime(20) {
            count.Add(count, one)
            if nFlag {
                fmt.Printf("%s : %s\n", i.String(), count.String())
            } else {
                fmt.Println(i)
            }
        }
    }

    fmt.Printf("// Number of prime numbers: %d\n", count.Sub(count, skipCount))
}
