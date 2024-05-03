package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)


func fetchDependenciesVersions(springVersion string) (map[string]DependencyInfo, error) {
    if !isValidVersion(springVersion) {
	log.Fatalf("Malformed spring version: %s", springVersion)
    }

    result := make(map[string]DependencyInfo)
    res, err := http.Get(fmt.Sprintf("https://docs.spring.io/spring-boot/docs/%s/reference/html/dependency-versions.html", springVersion))
    if err != nil {
	return nil, err
    }
    defer res.Body.Close()

    doc, err := goquery.NewDocumentFromReader(res.Body)
    if err != nil {
	return nil, err
    }

    doc.Find("#content > div:nth-child(2) > div > table > tbody > tr").Each(func(idx int, s *goquery.Selection) {
	groupId := s.Find("td:nth-of-type(1) > p > code").Text()
	artifactId := s.Find("td:nth-of-type(2) > p > code").Text()
	version := s.Find("td:nth-of-type(3) > p > code").Text() 
	result[groupId + "/" + artifactId] = DependencyInfo{
	    Version: version,
	}
    })

    return result, nil
}
