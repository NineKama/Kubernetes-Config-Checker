package main

import "log"

func validateImmutable(s podSpec) {
	if (s.Kind == "ConfigMap" || s.Kind == "Secret") && !s.Immutable {
		log.Printf("%s %s should be immutable.", s.Kind, s.Metadata.Name)
	}
}
