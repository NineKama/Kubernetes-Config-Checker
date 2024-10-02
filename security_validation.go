package main

import "log"

func validateSecurityContext(podSecItem podSpec) {
	if podSecItem.Spec.SecurityContext == (securityContext{}) {
		log.Printf("Pod %s should define security context.", podSecItem.Metadata.Name)
	}
	for _, container := range podSecItem.Spec.Container {
		if container.SecurityContext == (securityContext{}) {
			log.Printf("Container %s should define security context.", container.Name)
			continue
		}
		if container.SecurityContext.Privileged {
			log.Printf("Container %s should disable privileged in security context", container.Name)
		}
		if !container.SecurityContext.ReadOnlyRootFilesystem {
			log.Printf("Container %s should enable readOnlyRootFilesystem for security.", container.Name)
		}
	}
}
