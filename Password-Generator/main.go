package main

import (
	"bufio"
	cRand "crypto/rand"
	"fmt"
	"math/big"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

var (
    lower = "abcdefghijklmnopqrstuvwxyz"
    upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    numbers = "1234567890"
    special = "!@#$%^&*()_+-={}[]|<>/?~`"
    allChars = lower + upper + numbers + special
)

type PassInfo struct {
    Len int
    Lower int
    Upper int
    Numbers int
    Special int
}

func createPassword() (string, error) {
    passInfo, err := getUserInput()
    if err != nil {
        return "", err
    }

    var password strings.Builder

    for i:=0; i<passInfo.Lower; i++ {
        n, err := randomInt64(len(lower))
        if err != nil {
            return "", err
        }
        password.WriteString(string(lower[n]))
    }

    for i:=0; i<passInfo.Upper; i++ {
        n, err := randomInt64(len(upper))
        if err != nil {
            return "", err
        }
        password.WriteString(string(upper[n]))
    }

    for i:=0; i<passInfo.Numbers; i++ {
        n, err := randomInt64(len(numbers))
        if err != nil {
            return "", err
        }
        password.WriteString(string(numbers[n]))
    }

    for i:=0; i<passInfo.Special; i++ {
        n, err := randomInt64(len(special))
        if err != nil {
            return "", err
        }
        password.WriteString(string(special[n]))
    }

    remaining := passInfo.Len - passInfo.Lower - passInfo.Upper - passInfo.Special - passInfo.Numbers

    for i:=0; i<remaining; i++ {
        n, err := randomInt64(len(allChars))
        if err != nil {
            return "", err
        }
        password.WriteString(string(allChars[n]))
    }

    pass := scrambleString(password.String())
    return pass, nil

}

func randomInt64(len int) (int64, error) {
    n, err := cRand.Int(cRand.Reader, big.NewInt(int64(len)))
    return n.Int64(), err
}



func getUserInput() (PassInfo, error) {
    
    p := PassInfo{}
    fmt.Print("How long should the password be? ")
    str, err := getInput()
    if err != nil {
        return PassInfo{}, err
    }
    p.Len, err = strconv.Atoi(str)
    if err != nil {
        return PassInfo{}, err
    }

    fmt.Print("How many lower case letters? ")
    str, err = getInput()
    if err != nil {
        return PassInfo{}, err
    }
    p.Lower, err = strconv.Atoi(str)
    if err != nil {
        return PassInfo{}, err
    }

    fmt.Print("How many upper case letters? ")
    str, err = getInput()
    if err != nil {
        return PassInfo{}, err
    }
    p.Upper, err = strconv.Atoi(str)
    if err != nil {
        return PassInfo{}, err
    }

    fmt.Print("How many numbers? ")
    str, err = getInput()
    if err != nil {
        return PassInfo{}, err
    }
    p.Numbers, err = strconv.Atoi(str)
    if err != nil {
        return PassInfo{}, err
    }

    fmt.Print("How many special chars? ")
    str, err = getInput()
    if err != nil {
        return PassInfo{}, err
    }
    p.Special, err = strconv.Atoi(str)
    if err != nil {
        return PassInfo{}, err
    }

    return p, nil
}

func getInput() (string, error) {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    err := scanner.Err()
    if err != nil {
        return "", err
    }
    return scanner.Text(), nil    
}

func scrambleString(s string) string {
    r := []rune(s)
    rand.Shuffle(len(r), func(i, j int) {
        r[i], r[j] = r[j], r[i]
    })

    return string(r)

}

func main() {
    pass, err := createPassword()
    if err != nil {
        println(err)
    }
    println(pass)

}