package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
)

var (
    outputFile string
    version string
)

func init() {
    flag.StringVar(&outputFile, "output", "deps.json", "Output json file name")
    flag.StringVar(&version, "version", "current", "The spring version you are currently using")
}

type Result struct {
    GroupId string
    ArtifactId string
    CurrentVersion string
    RecomendedVersion string
}

func parseCmdArgs() {
    flag.Parse()

    if outputFile == "deps.json" {
	fmt.Printf("Output not specified defaulting to deps.json\n")
    }

    if !isValidVersion(version) {
	fmt.Printf("Spring version not specified, using: current\n")
    }
}

func main() {
    parseCmdArgs() 
    officialVersions, err := fetchDependenciesVersions(version)
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
    err = os.WriteFile(outputFile, file, 0644)
    if err != nil {
	log.Fatal(err)
    }
}
