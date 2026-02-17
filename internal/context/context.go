package context

import (
	"os"
	"os/exec"
)

type ProjectType string

const (
	ProjectNode   ProjectType = "node"
	ProjectPython ProjectType = "python"
	ProjectGo     ProjectType = "go"
	ProjectRust   ProjectType = "rust"
	ProjectNone   ProjectType = "generic"
)

type Context struct {
	Cwd         string
	ProjectType ProjectType
	Files       []string
	Tools       []string
}

// Get gathers context about the current environment
func Get(cwd string) Context {
	if cwd == "" {
		cwd, _ = os.Getwd()
	}

	ctx := Context{
		Cwd:         cwd,
		ProjectType: ProjectNone,
		Tools:       detectTools(),
	}

	files, err := os.ReadDir(cwd)
	if err == nil {
		for _, f := range files {
			ctx.Files = append(ctx.Files, f.Name())
		}
	}

	ctx.ProjectType = detectProject(ctx.Files)
	return ctx
}

func detectProject(files []string) ProjectType {
	for _, f := range files {
		switch f {
		case "package.json":
			return ProjectNode
		case "requirements.txt", "pyproject.toml", "Pipfile", "setup.py":
			return ProjectPython
		case "go.mod":
			return ProjectGo
		case "Cargo.toml":
			return ProjectRust
		}
	}
	return ProjectNone
}

func detectTools() []string {
	// Common tools to check for
	candidates := []string{
		"git", "docker", "npm", "pnpm", "yarn", "python3", "python", "pip", "go", "cargo", "node",
	}

	var found []string
	for _, tool := range candidates {
		if _, err := exec.LookPath(tool); err == nil {
			found = append(found, tool)
		}
	}
	return found
}
