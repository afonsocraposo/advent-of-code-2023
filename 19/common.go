package main

import (
	"strings"

	. "github.com/afonsocraposo/advent-of-code-2023/utils"
)

type Workflow struct {
	dest       string
	conditions []Condition
}

func ParseWorkfowLine(line string) (string, Workflow) {
	label, text, _ := strings.Cut(line, "{")
	text = text[:len(text)-1]
	workflow := ParseWorkfow(text)
	return label, workflow
}

func ParseWorkfow(text string) Workflow {
	c := strings.Split(text, ",")
	conditions := make([]Condition, len(c)-1)
	for i, slice := range c[:len(c)-1] {
		conditions[i] = ParseCondition(slice)
	}

	dest := c[len(c)-1]
	workflow := Workflow{}
	workflow.dest = dest
	workflow.conditions = conditions

	return workflow
}

func (w *Workflow) goTo(part map[string]int) string {
    for _, condition := range w.conditions {
        if condition.test(part) {
            return condition.dest
        }
    }
    return w.dest
}

type Condition struct {
	label   string
	greater bool
	value   int
	dest    string
}

func (c *Condition) test(part map[string]int) bool {
	pv := part[c.label]
	g := pv > c.value
	return g == c.greater
}

func ParseCondition(text string) Condition {
	var label, after, value, dest string
	greater := false

	label, after, greater = strings.Cut(text, ">")
	if !greater {
		label, after, _ = strings.Cut(text, "<")
	}
	value, dest, _ = strings.Cut(after, ":")

	condition := Condition{}
	condition.label = label
	condition.greater = greater
	condition.value = ParseInt(value)
	condition.dest = dest

	return condition
}

func ParsePart(text string) map[string]int {
	text = text[1 : len(text)-1]
	parts := strings.Split(text, ",")

	part := map[string]int{}
	for _, p := range parts {
		label, value, _ := strings.Cut(p, "=")
		part[label] = ParseInt(value)
	}
	return part
}

const (
	ACCEPT = "A"
	REJECT = "R"
)
