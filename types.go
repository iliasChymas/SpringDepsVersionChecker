package main 

type Dependency struct {
    GroupId string `xml:"groupId"`
    ArtifactId string `xml:"artifactId"`
    Version string `xml:"version,omitempty"`
}

type DependencyInfo struct {
    GroupId string `xml:"groupId"`
    Version string `xml:"version,omitempty"`
}
