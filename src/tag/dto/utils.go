package tagDTO

import "regexp"

const hexRegexString = "^#(?:[0-9a-fA-F]{3}|[0-9a-fA-F]{6})$"

var hexRegex = regexp.MustCompile(hexRegexString)
