package main 

type Dependency struct {
    GroupId string `xml:"groupId"`
    ArtifactId string `xml:"artifactId"`
    Version string `xml:"version,omitempty"`
}
