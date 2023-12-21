package main

import (
	"os"
	"strconv"
	"strings"
)

func parseInstructions(rules map[string][]string, instr []string) expr {
	if len(instr) == 0 {
		panic("empty list")
	}
	head := instr[0]
	if len(instr) == 1 {
		if head == "A" {
			return accepted{}
		}
		if head == "R" {
			return rejected{}
		}

		return parseInstructions(rules, rules[head])
	}
	subject := head[0]
	cond := head[1]
	value, _ := strconv.Atoi(head[2:])
	then := instr[1]
	tail := instr[2:]

	if cond == '<' {
		return ifThenElse{
			cond:  lt{subject, value},
			then:  parseInstructions(rules, rules[then]),
			else_: parseInstructions(rules, tail),
		}
	}
	if cond == '>' {
		return ifThenElse{
			cond:  gt{subject, value},
			then:  parseInstructions(rules, rules[then]),
			else_: parseInstructions(rules, tail),
		}
	}
	panic("not implemented")
}

func parseRules(lines []string, start string) expr {
	var rules = make(map[string][]string)
	for _, line := range lines {
		name, after, _ := strings.Cut(line, "{")
		fields := strings.FieldsFunc(after, func(c rune) bool {
			return c == ',' || c == ':' || c == '}'
		})
		rules[name] = fields
	}
	rules["A"] = []string{"A"}
	rules["R"] = []string{"R"}
	return parseInstructions(rules, rules[start])
}

func part2(inputFile string) int {
	contents, err := os.ReadFile(inputFile)
	if err != nil {
		panic("could't read the input")
	}

	input := strings.TrimSpace(string(contents))
	parts := strings.Split(input, "\n\n")
	lines := strings.Split(parts[0], "\n")

	var expr = parseRules(lines, "in")

	start := interval{1, 4000}
	var c = constraint{'x': start, 'm': start, 'a': start, 's': start}
	return expr.propagate(c)
}
