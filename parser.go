package main

import (
    "errors"
    "fmt"
    "log"
    "os"
    "regexp"
    "github.com/antchfx/xmlquery"
)

func fileExists(filepath string) bool {
    _, err := os.Stat(filepath)
    return !errors.Is(err, os.ErrNotExist)
}

func parseVersion(node *xmlquery.Node, globalNode *xmlquery.Node, artifactId *string) string {
    versionTag := xmlquery.FindOne(node, "version")
    var output string
    var versionText string
    versionVariableRegex := regexp.MustCompile(`\${(.*)}`)
    if versionTag != nil {
	versionText = versionTag.InnerText()	
	if isValidVersion(versionText) {
	    return versionText
	} else {
	    matches := versionVariableRegex.FindAllStringSubmatch(versionText, -1)
	    if len(matches) == 0 {
		fmt.Printf("[DEBUG] Version string in artifactId: '%s' is not a variable neither a valid version", *artifactId)
		return ""
	    } else {
		versionVariable := "//" + matches[0][1]
		results := xmlquery.Find(globalNode, versionVariable)
		if len(results) == 0 {
		    fmt.Printf("[DEBUG] Variable does not exist for artifactId: %s", *artifactId)
		    return ""
		}
		output = results[0].InnerText()
		return output
	    }
	}
    }
    return ""
}

func ParsePomFile(filePath string) ([]Dependency) {
    if !fileExists(filePath) {
        log.Fatalf("File %s not found", filePath)
    }
    xmlFile, _ := os.Open(filePath)
    defer xmlFile.Close()
    doc, err := xmlquery.Parse(xmlFile)
    if err != nil {
        log.Fatalf("Error parsing XML: %s", err.Error())
    }
    xmlDeps := xmlquery.Find(doc, "//dependency")
    dependencies := make([]Dependency, len(xmlDeps))
    for idx, dep := range xmlDeps {
	groupId := xmlquery.FindOne(dep, "groupId").InnerText()
	artifactId := xmlquery.FindOne(dep, "artifactId").InnerText()
	dependencies[idx] = Dependency {
	    GroupId: groupId,
	    ArtifactId: artifactId,
	    Version: parseVersion(dep, doc, &artifactId),
	}
    }
    return dependencies
}











