package main

import (
	"flag"
	"fmt"
	"github.com/Piszmog/adventofcode/utils"
	"log"
	"regexp"
	"strconv"
)

type operationType int

const (
	accumulator operationType = iota
	jump
	noOp
)

const (
	accumulatorField = "acc"
	jumpField        = "jmp"
	noOpField        = "nop"
)

func (o operationType) String() string {
	return [...]string{accumulatorField, jumpField, noOpField}[o]
}

func toOperationType(s string) operationType {
	switch s {
	case accumulatorField:
		return accumulator
	case jumpField:
		return jump
	case noOpField:
		return noOp
	}
	return -99
}

type operation struct {
	operationType operationType
	argument      int
}

func main() {
	inputFile := flag.String("f", "../input.txt", "File containing the input")

	ops, err := getOperations(*inputFile)
	if err != nil {
		log.Fatalln(err)
	}

	// part 1
	accTotal := getAccumulator(ops)
	fmt.Printf("Part 1: Accumulator value: %d\n", accTotal)

	// part 2
	actualAccTotal := getFixedAccumulator(ops)
	fmt.Printf("Part 2: Accumulator value: %d\n", actualAccTotal)
}

func getOperations(inputFile string) ([]operation, error) {
	var operations []operation
	operationRegex := regexp.MustCompile("([a-z]+)\\s([+|-])([0-9]+)")
	err := utils.ReadTextFile(inputFile, func(line string) error {
		matches := operationRegex.FindStringSubmatch(line)
		opType := toOperationType(matches[1])
		argumentSign := matches[2]
		argument, _ := strconv.Atoi(matches[3])
		if argumentSign == "-" {
			argument *= -1
		}
		operations = append(operations, operation{operationType: opType, argument: argument})
		return nil
	})
	return operations, err
}

func getAccumulator(ops []operation) int {
	acc := 0
	i := 0
	opsSize := len(ops)
	visitCounts := make(map[int]int)
	for {
		// prevent us from causing an issue
		if i >= opsSize {
			break
		}
		if visitCounts[i] == 1 {
			break
		}
		visitCounts[i] += 1
		op := ops[i]
		switch op.operationType {
		case accumulator:
			acc += op.argument
			i++
		case jump:
			i += op.argument
		case noOp:
			i++
		}
	}
	return acc
}

func getFixedAccumulator(ops []operation) int {
	hasFlipped := false
	acc := 0
	i := 0
	opsSize := len(ops)
	visitCounts := make(map[int]int)
	for {
		// prevent us from causing an issue
		if i >= opsSize {
			break
		}
		visitCounts[i] += 1
		op := ops[i]
		switch op.operationType {
		case accumulator:
			acc += op.argument
			i++
		case jump:
			if !hasFlipped && !isInfinite(ops, i, operation{operationType: noOp, argument: op.argument}) {
				i++
				hasFlipped = true
			} else {
				i += op.argument
			}
		case noOp:
			if !hasFlipped && !isInfinite(ops, i, operation{operationType: jump, argument: op.argument}) {
				i += op.argument
				hasFlipped = true
			} else {
				i++
			}
		}
	}
	return acc
}

func isInfinite(ops []operation, index int, op operation) bool {
	infinite := false
	i := index
	currentOp := op
	opsSize := len(ops)
	visitCounts := make(map[int]int)
	for {
		if visitCounts[i] == 1 {
			infinite = true
			break
		}
		visitCounts[i] += 1
		switch currentOp.operationType {
		case jump:
			i += currentOp.argument
		case accumulator:
			fallthrough
		case noOp:
			i++
		}
		if i >= opsSize {
			break
		}
		currentOp = ops[i]
	}
	return infinite
}
