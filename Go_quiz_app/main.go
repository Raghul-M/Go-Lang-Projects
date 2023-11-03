package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func problemPuller(fileName string) ([]problem, error) {
	// read all the problems from the quiz.csv
	//1. open the file
	if fObj, err := os.Open(fileName); err == nil {
		//2. we will create a new reader
		csvR := csv.NewReader(fObj)
		//3. it will need to read the file
		if cLines, err := csvR.ReadAll(); err == nil {
			//4. call the parseProblem function
			return parseProblem(cLines), nil
		} else {
			return nil, fmt.Errorf("Error in reading data in csv"+"format from %s file; %s", fileName, err.Error())
		}
	} else {
		return nil, fmt.Errorf("Error in opening the %s file; %s", fileName, err.Error())
	}

}

func main() {
	//1.input the name of the file
	fName := flag.String("f", "quiz.csv", "path of csv file")
	//2.set the duration of the timer
	timer := flag.Int("t", 30, "timer of the quiz")
	flag.Parse()
	//3.pull the problems from the file (calling our problem puller func)
	problems, err := problemPuller(*fName)
	//4. handle the error
	if err != nil {
		exit(fmt.Sprintf("something went wrong: %s", err.Error()))
	}
	//5.create a variable to count our correct answers
	correctAns := 0
	//6. using the duration of the timer, we want to initialize the time
	tObj := time.NewTimer(time.Duration(*timer) * time.Second)
	ansC := make(chan string)
	//7. loop through problems, print the questions, we will accept the answers

problemLoop:
	for i, p := range problems {
		var answer string
		fmt.Printf("Problem %d: %s=", i+1, p.q)

		go func() {
			fmt.Scanf("%s", &answer)
			ansC <- answer
		}()
		select {
		case <-tObj.C:
			fmt.Println()
			break problemLoop
		case iAns := <-ansC:
			if iAns == p.a {
				correctAns++
				
			}
			if i == len(problems)-1 {
				close(ansC)
			}
			
		}
		fmt.Println(correctAns)
	}

	//8. we will calculate and print out the result
	fmt.Printf("Your result is %d out of %d \n", correctAns, len(problems))
	fmt.Printf("Press enter to exit")
	<-ansC

}

func parseProblem(lines [][]string) []problem {
	//go over the lines and parse them, with problem struct
	r := make([]problem, len(lines))
	for i := 0; i < len(lines); i++ {
		r[i] = problem{q: lines[i][0], a: lines[i][1]}
	}
	return r
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func noUse() {
	fmt.Println("no use")
}