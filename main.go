package main

import (
	"fmt"
	"log"
)

func main() {
    officialVersions, err := fetchDependenciesVersions("3.1.1")
    pomDeps := ParsePomFile("pom.xml")
    if err != nil {
	log.Fatal(err)
    }	
    for _, dep := range pomDeps {
	matchedDep, exists := officialVersions[dep.ArtifactId]
	if exists {
	    fmt.Printf("Dependency verson mismatch found for ArtifactId: %s\nSpring recomended version: %s\nCurrent version: %s", dep.ArtifactId, matchedDep.Version, dep.Version)
	}
    }

}
