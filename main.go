package main

import "fmt"


func main() {
    dependencies := ParsePomFile("pom.xml")
    for _, dep := range dependencies {
	fmt.Printf("ArtifactId: %s\nVersion: %s\n\n", dep.ArtifactId, dep.Version)
    }
}

