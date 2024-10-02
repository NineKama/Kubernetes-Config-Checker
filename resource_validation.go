package main

import "log"

func isEmptyResource(c compute) bool {
	return c.Memory == "" && c.CPU == ""
}

func validateContainerResources(podSecItem podSpec) {
	for _, container := range podSecItem.Spec.Container {
		if isEmptyResource(container.Resources.Limits) && isEmptyResource(container.Resources.Requests) {
			log.Printf("Container %s should define CPU/Memory limits and requests.", container.Name)
		}
	}
}
