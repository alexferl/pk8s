// Code generated by pk8s; DO NOT EDIT.

package k8s

// ExecActionV1 ExecAction describes a "run in container" action.
type ExecActionV1 struct {
	// Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.
	Command []string `json:"command,omitempty"`
}
