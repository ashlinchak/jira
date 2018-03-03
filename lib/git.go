package lib

type Git struct{}

func (git *Git) CreateBranch(name string, master string) {
	arguments := []string{"checkout", "-b", name, master}
	ExecuteCommand("git", arguments)
}
