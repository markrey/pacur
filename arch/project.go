package arch

import (
	"github.com/pacur/pacur/constants"
	"github.com/pacur/pacur/utils"
	"path/filepath"
)

type ArchProject struct {
	Name       string
	Root       string
	MirrorRoot string
	BuildRoot  string
	Path       string
	Distro     string
	Release    string
}

func (p *ArchProject) getBuildDir() (path string, err error) {
	path = filepath.Join(p.BuildRoot, p.Distro)

	err = utils.MkdirAll(path)
	if err != nil {
		return
	}

	return
}

func (p *ArchProject) Prep() (err error) {
	buildDir, err := p.getBuildDir()
	if err != nil {
		return
	}

	keyPath := filepath.Join(p.Path, "..", "..", "sign.key")
	exists, err := utils.Exists(keyPath)
	if err != nil {
		return
	}

	if exists {
		err = utils.CopyFile("", keyPath, buildDir, true)
		if err != nil {
			return
		}
	}

	err = utils.RsyncExt(p.Path, buildDir, ".pkg.tar.xz")
	if err != nil {
		return
	}

	return
}

func (p *ArchProject) Create() (err error) {
	buildDir, err := p.getBuildDir()
	if err != nil {
		return
	}

	err = utils.Exec("", "docker", "run", "--rm", "-t", "-v",
		buildDir+":/pacur", constants.DockerOrg+p.Distro, "create",
		p.Distro, p.Name)
	if err != nil {
		return
	}

	path := filepath.Join(p.MirrorRoot, "arch")

	_ = utils.RemoveAll(path)

	err = utils.Rsync(filepath.Join(buildDir, "arch"), path)
	if err != nil {
		return
	}

	return
}
