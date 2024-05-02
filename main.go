package main

import "fmt"


func main() {
    pomDependencies := ParsePomFile("pom.xml")
    for _, dep := range pomDependencies {
	fmt.Printf("ArtifactId: %s\nVersion: %s\n\n", dep.ArtifactId, dep.Version)
    }
}

