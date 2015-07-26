package redhat

import (
	"fmt"
	"github.com/dropbox/godropbox/errors"
	"github.com/pacur/pacur/pack"
	"github.com/pacur/pacur/utils"
	"os"
	"path/filepath"
)

type Redhat struct {
	Distro       string
	Release      string
	Pack         *pack.Pack
	redhatDir    string
	buildDir     string
	buildRootDir string
	rpmsDir      string
	sourcesDir   string
	specsDir     string
	srpmsDir     string
}

func (r *Redhat) createSpec() (err error) {
	path := filepath.Join(r.specsDir, r.Pack.PkgName+".spec")

	file, err := os.Create(path)
	if err != nil {
		err = &WriteError{
			errors.Wrapf(err,
				"redhat: Failed to create redhat spec at '%s'", path),
		}
		return
	}
	defer file.Close()

	data := ""

	data += fmt.Sprintf("Name: %s\n", r.Pack.PkgName)
	data += fmt.Sprintf("Summary: %s\n", r.Pack.PkgDesc)
	data += fmt.Sprintf("Version: %s\n", r.Pack.PkgVer)
	data += fmt.Sprintf("Release: %s", r.Pack.PkgName) + "%{?dist}\n"
	data += fmt.Sprintf("Group: %s\n", r.Pack.Section)
	data += fmt.Sprintf("URL: %s\n", r.Pack.Url)
	data += fmt.Sprintf("License: %s\n", r.Pack.License)
	data += fmt.Sprintf("Packager: %s\n", r.Pack.Maintainer)

	for _, pkg := range r.Pack.Provides {
		data += fmt.Sprintf("Provides: %s\n", pkg)
	}

	for _, pkg := range r.Pack.Conflicts {
		data += fmt.Sprintf("Conflicts: %s\n", pkg)
	}

	for _, pkg := range r.Pack.Depends {
		data += fmt.Sprintf("Requires: %s\n", pkg)
	}

	for _, pkg := range r.Pack.MakeDepends {
		data += fmt.Sprintf("BuildRequires: %s\n", pkg)
	}

	data += "\n"

	if len(r.Pack.PkgDescLong) > 0 {
		data += "%description\n"
		for _, line := range r.Pack.PkgDescLong {
			data += line + "\n"
		}
		data += "\n"
	}

	data += "%install\n"
	data += fmt.Sprintf("mv -f %s/.[!.]* $RPM_BUILD_ROOT || true\n",
		r.Pack.PackageDir)
	data += fmt.Sprintf("mv -f %s/* $RPM_BUILD_ROOT || true\n",
		r.Pack.PackageDir)
	data += "\n"

	data += "%files\n"
	data += "/\n"
	for _, file := range r.Pack.Backup {
		data += "%config " + file + "\n"
	}
	data += "\n"

	if len(r.Pack.PreInst) > 0 {
		data += "%pre\n"
		for _, line := range r.Pack.PreInst {
			data += line + "\n"
		}
		data += "\n"
	}

	if len(r.Pack.PostInst) > 0 {
		data += "%post\n"
		for _, line := range r.Pack.PostInst {
			data += line + "\n"
		}
		data += "\n"
	}

	if len(r.Pack.PreRm) > 0 {
		data += "%preun\n"
		for _, line := range r.Pack.PreRm {
			data += line + "\n"
		}
		data += "\n"
	}

	if len(r.Pack.PostRm) > 0 {
		data += "%postun\n"
		for _, line := range r.Pack.PostRm {
			data += line + "\n"
		}
	}

	_, err = file.WriteString(data)
	if err != nil {
		err = &WriteError{
			errors.Wrapf(err,
				"redhat: Failed to write redhat spec at '%s'", path),
		}
		return
	}

	return
}

func (r *Redhat) Prep() (err error) {
	return
}

func (r *Redhat) Build() (err error) {
	r.redhatDir = filepath.Join(r.Pack.Root, "redhat")
	r.buildDir = filepath.Join(r.redhatDir, "BUILD")
	r.buildRootDir = filepath.Join(r.redhatDir, "BUILDROOT")
	r.rpmsDir = filepath.Join(r.redhatDir, "RPMS")
	r.sourcesDir = filepath.Join(r.redhatDir, "SOURCES")
	r.specsDir = filepath.Join(r.redhatDir, "SPECS")
	r.srpmsDir = filepath.Join(r.redhatDir, "SRPMS")

	for _, path := range []string{
		r.redhatDir,
		r.buildDir,
		r.buildRootDir,
		r.rpmsDir,
		r.sourcesDir,
		r.specsDir,
		r.srpmsDir,
	} {
		err = utils.ExistsMakeDir(path)
		if err != nil {
			return
		}
	}
	//defer os.RemoveAll(r.redhatDir)

	err = r.createSpec()
	if err != nil {
		return
	}

	return
}