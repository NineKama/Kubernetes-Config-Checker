package main

import "log"

func validateNameSpace(podSecItem podSpec) {
	if podSecItem.Metadata.Namespace == "" || podSecItem.Metadata.Namespace == "default" {
		log.Printf("Pod %s should use a separate namespace (not default).", podSecItem.Metadata.Name)
	}
}
