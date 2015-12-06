package main

type Result struct {
	ActiveProcesses int
	TotalProcesses  int
	AcceptedConn    int
	Available       bool
}

func NewResult(pools []Pool) *Result {
	result := &Result{0, 0, 0, true}
	for _, pool := range pools {
		if !pool.Available {
			result.Available = false
			continue
		}
		result.ActiveProcesses += pool.Status.ActiveProcesses
		result.TotalProcesses += pool.Status.TotalProcesses
		result.AcceptedConn += pool.Status.AcceptedConn
	}
	return result
}
