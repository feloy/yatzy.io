package firestore

import (
	"strings"

	"cloud.google.com/go/functions/metadata"
)

// GetResourceIDFromMeta returns the ID of a resource
func GetResourceIDFromMeta(meta metadata.Metadata) string {
	parts := strings.Split(meta.Resource.Name, "/")
	return parts[len(parts)-1]
}

// GetParentIDFromMeta returns the ID of the parent of a resource
func GetParentIDFromMeta(meta metadata.Metadata) string {
	parts := strings.Split(meta.Resource.Name, "/")
	return parts[len(parts)-3]
}
