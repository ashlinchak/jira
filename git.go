package main

type Git struct{}

func (git *Git) CreateBranch(name string, masterBranch string) {
	arguments := []string{"checkout", "-b", name, masterBranch}
	ExecuteCommand("git", arguments)
}
