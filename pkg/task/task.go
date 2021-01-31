package task

// ITask defines a task contract
type ITask interface {
	// Returns the task ID
	Identity() int

	// Starts the execution of the task, returns an error if present
	Exec() error
}
