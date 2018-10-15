package gce

import "strings"

const (
	prefix = "https://www.googleapis.com/compute/v1/projects"
)

// Ensures machineType is proper URL
// https://www.googleapis.com/compute/v1/projects/<project>/zones/<zone>/machineTypes/<machine>
func machineTypeURL(project, zone, machine string) string {
	if strings.HasPrefix(machine, "https://") {
		return machine
	} else {
		return strings.Join([]string{prefix, project, "zones", zone, "machineTypes", machine}, "/")
	}
}

// Ensures network is proper URL
// https://www.googleapis.com/compute/v1/projects/<gcp-project-id>/global/networks/default
func networkURL(project, network string) string {
	if strings.HasPrefix(network, "https://") {
		return network
	} else {
		return strings.Join([]string{prefix, project, "global/networks", network}, "/")
	}
}
