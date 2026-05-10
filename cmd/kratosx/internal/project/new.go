package project

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"

	"github.com/limes-cloud/kratosx/cmd/kratosx/internal/base"
)

// Project is a project template.
type Project struct {
	Name string
	Path string
}

func (p *Project) ensureDir(to string) error {
	if _, err := os.Stat(to); !os.IsNotExist(err) {
		fmt.Printf("🚫 %s already exists\n", p.Name)
		var override bool
		prompt := &survey.Confirm{
			Message: "📂 Do you want to override the folder ?",
			Help:    "Delete the existing folder and create the project.",
		}
		e := survey.AskOne(prompt, &override)
		if e != nil {
			return e
		}
		if !override {
			return err
		}
		if err := os.RemoveAll(to); err != nil {
			return err
		}
	}
	return nil
}

func (p *Project) renameDefaults(to string) error {
	if err := os.Rename(
		filepath.Join(to, "cmd", "server"),
		filepath.Join(to, "cmd", p.Name),
	); err != nil {
		return err
	}
	return os.Rename(
		filepath.Join(to, "api", "layout"),
		filepath.Join(to, "api", p.Name),
	)
}

func (p *Project) printSuccess(dir, action, confPath string) {
	fmt.Printf("\n🍺 %s succeeded %s\n", action, color.GreenString(p.Name))
	fmt.Print("💻 Use the following command to start the project 👇:\n\n")
	fmt.Println(color.WhiteString("$ cd %s", p.Name))
	fmt.Println(color.WhiteString("$ go generate ./..."))
	fmt.Println(color.WhiteString("$ go build -o ./bin/ ./... "))
	fmt.Println(color.WhiteString("$ ./bin/%s -conf %s\n", p.Name, confPath))
	fmt.Println("	🤝 Thanks for using Kratosx")
	fmt.Println("	📚 Tutorial: http://docs.qlime.cn")
}

// New new a project from remote repo.
func (p *Project) New(ctx context.Context, dir string, layout string, branch string) error {
	to := filepath.Join(dir, p.Name)
	if err := p.ensureDir(to); err != nil {
		return err
	}
	fmt.Printf("🚀 Creating service %s, layout repo is %s, please wait a moment.\n\n", p.Name, layout)
	repo := base.NewRepo(layout, branch)
	if err := repo.CopyTo(ctx, to, p.Name, []string{".git", ".github"}); err != nil {
		return err
	}
	if err := p.renameDefaults(to); err != nil {
		return err
	}
	base.Tree(to, dir)
	p.printSuccess(dir, "Project creation", "./configs")
	return nil
}
