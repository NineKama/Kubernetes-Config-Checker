package main

import "log"

func validateServiceAccount(podSecItem podSpec) {
	if podSecItem.Spec.ServiceAccountName == "" && podSecItem.Kind == "Pod" {
		log.Printf("Pod %s should use a service account.", podSecItem.Metadata.Name)
	}
}
