package project

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/limes-cloud/kratosx/cmd/kratosx/internal/base"
)

var repoAddIgnores = []string{
	".git", ".github", "api", "README.md", "LICENSE", "go.mod", "go.sum", "third_party", "openapi.yaml", ".gitignore",
}

func (p *Project) Add(ctx context.Context, dir string, layout string, branch string, mod string, pkgPath string) error {
	to := filepath.Join(dir, p.Name)
	if err := p.ensureDir(to); err != nil {
		return err
	}

	fmt.Printf("🚀 Add service %s, layout repo is %s, please wait a moment.\n\n", p.Name, layout)

	pkgPath = fmt.Sprintf("%s/%s", mod, pkgPath)
	repo := base.NewRepo(layout, branch)
	if err := repo.CopyToV2(ctx, to, pkgPath, repoAddIgnores, []string{filepath.Join(p.Path, "api"), "api"}); err != nil {
		return err
	}

	if err := p.renameDefaults(to); err != nil {
		return err
	}

	base.Tree(to, dir)
	p.printSuccess(dir, "Repository creation", "./internal/conf/conf.yaml")
	return nil
}
