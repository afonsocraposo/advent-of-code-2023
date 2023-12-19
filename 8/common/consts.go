package common

import (
    "regexp"
)

var R = regexp.MustCompile("[A-Z0-9]{3}")

const START_LABEL = "AAA"
const END_LABEL = "ZZZ"
const START_RUNE = 'A'
const END_RUNE = 'Z'
