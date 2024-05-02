package main

import "regexp"

func isValidVersion(input string) bool {
    versionRegex := regexp.MustCompile(`^\d{1,3}(?:\.\d{1,3}){0,2}$`)
    return versionRegex.MatchString(input)
}
