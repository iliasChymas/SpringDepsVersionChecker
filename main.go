package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Result struct {
    GroupId string
    ArtifactId string
    CurrentVersion string
    RecomendedVersion string
}

func main() {
    officialVersions, err := fetchDependenciesVersions("3.1.1")
    pomDeps := ParsePomFile("pom.xml")
    var output []Result
    if err != nil {
	log.Fatal(err)
    }	
    for _, dep := range pomDeps {
	matchedDep, exists := officialVersions[dep.GroupId + "/" + dep.ArtifactId]
	if exists && dep.Version != "" && dep.Version != matchedDep.Version {
	    output = append(output, Result{
		GroupId: dep.GroupId,
		ArtifactId: dep.ArtifactId,
		CurrentVersion: dep.Version,
		RecomendedVersion: matchedDep.Version,
	    }) 
	}
    }
    
    if len(output) < 0 {
	fmt.Println("No recomendations")
	os.Exit(0)
    }	

    file, _ := json.MarshalIndent(output, "", " ")
    err = os.WriteFile("pomVersions.json", file, 0644)
    if err != nil {
	log.Fatal(err)
    }
}
