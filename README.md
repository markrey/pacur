[![pacur](https://raw.githubusercontent.com/pacur/pacur-artwork/master/pacur.png)](http://pacur.org)

# pacur: simple packaging

[![Docker Repository](https://img.shields.io/badge/docker-archlinux-1794d1.svg?style=flat "Docker Repository")](https://registry.hub.docker.com/u/pacur/archlinux/)

[![Docker Repository](https://img.shields.io/badge/docker-amazonlinux--2016.09-faaf34.svg?style=flat "Docker Repository")](https://registry.hub.docker.com/u/pacur/amazonlinux-2016.09/)
[![Docker Repository](https://img.shields.io/badge/docker-amazonlinux--2017.03-faaf34.svg?style=flat "Docker Repository")](https://registry.hub.docker.com/u/pacur/amazonlinux-2017.03/)
[![Docker Repository](https://img.shields.io/badge/docker-amazonlinux--2017.09-faaf34.svg?style=flat "Docker Repository")](https://registry.hub.docker.com/u/pacur/amazonlinux-2017.09/)

[![Docker Repository](https://img.shields.io/badge/docker-centos--6-10233f.svg?style=flat "Docker Repository")](https://registry.hub.docker.com/u/pacur/centos-6/)
[![Docker Repository](https://img.shields.io/badge/docker-centos--7-10233f.svg?style=flat "Docker Repository")](https://registry.hub.docker.com/u/pacur/centos-7/)


[![Docker Repository](https://img.shields.io/badge/docker-debian--wheezy-d70a53.svg?style=flat "Docker Repository")](https://registry.hub.docker.com/u/pacur/debian-wheezy/)
[![Docker Repository](https://img.shields.io/badge/docker-debian--jessie-d70a53.svg?style=flat "Docker Repository")](https://registry.hub.docker.com/u/pacur/debian-jessie/)
[![Docker Repository](https://img.shields.io/badge/docker-debian--stretch-d70a53.svg?style=flat "Docker Repository")](https://registry.hub.docker.com/u/pacur/debian-stretch/)


[![Docker Repository](https://img.shields.io/badge/docker-fedora--24-294172.svg?style=flat "Docker Repository")](https://registry.hub.docker.com/u/pacur/fedora-24/)
[![Docker Repository](https://img.shields.io/badge/docker-fedora--25-294172.svg?style=flat "Docker Repository")](https://registry.hub.docker.com/u/pacur/fedora-25/)
[![Docker Repository](https://img.shields.io/badge/docker-fedora--26-294172.svg?style=flat "Docker Repository")](https://registry.hub.docker.com/u/pacur/fedora-26/)
[![Docker Repository](https://img.shields.io/badge/docker-fedora--27-294172.svg?style=flat "Docker Repository")](https://registry.hub.docker.com/u/pacur/fedora-27/)


[![Docker Repository](https://img.shields.io/badge/docker-ubuntu--precise-dd4814.svg?style=flat "Docker Repository")](https://registry.hub.docker.com/u/pacur/ubuntu-precise/)
[![Docker Repository](https://img.shields.io/badge/docker-ubuntu--trusty-dd4814.svg?style=flat "Docker Repository")](https://registry.hub.docker.com/u/pacur/ubuntu-trusty/)
[![Docker Repository](https://img.shields.io/badge/docker-ubuntu--xenial-dd4814.svg?style=flat "Docker Repository")](https://registry.hub.docker.com/u/pacur/ubuntu-xenial/)
[![Docker Repository](https://img.shields.io/badge/docker-ubuntu--yakkety-dd4814.svg?style=flat "Docker Repository")](https://registry.hub.docker.com/u/pacur/ubuntu-yakkety/)
[![Docker Repository](https://img.shields.io/badge/docker-ubuntu--zesty-dd4814.svg?style=flat "Docker Repository")](https://registry.hub.docker.com/u/pacur/ubuntu-zesty/)
[![Docker Repository](https://img.shields.io/badge/docker-ubuntu--artful-dd4814.svg?style=flat "Docker Repository")](https://registry.hub.docker.com/u/pacur/ubuntu-artful/)
[![Docker Repository](https://img.shields.io/badge/docker-ubuntu--bionic-dd4814.svg?style=flat "Docker Repository")](https://registry.hub.docker.com/u/pacur/ubuntu-bionic/)

Pacur allows building packages for multiple linux distributions with a
consistent package spec format. Currently `deb`, `rpm` and `pacman` packages
are available for several linux distributions. Builds are done on Docker
containers without needing to setup any virtual machines or install any
software other then Docker. All packages are built using a simple format that
is similar to [PKGBUILD](https://wiki.archlinux.org/index.php/PKGBUILD) from
ArchLinux. Each distribution is different and will still require different
build instructions but a consistent build process and format can be used for
all builds. Docker only supports 64 bit containers, pacur can't be used to
build packages 32 bit packages. Pacur will also create a `deb`, `rpm` and
`pacman` signed repository that can be used on ArchLinux, CentOS, Fedora,
Debian and Ubuntu to distribute the packages. A
[tutorial](https://medium.com/@zachhuff386/pacur-tutorial-9848b774c84a)
on creating a project is aviaible on medium.

### format

```
key="example string"
key=`example "quoted" string`
key=("list with one element")
key=(
    "list with"
    "multiple elements"
)
key="example ${variable} string"
key:ubuntu="this will apply only to ubuntu builds"
```

### builtin variables

| key | value |
| --- | ----- |
| `${srcdir}` | `Source` directory where all sources are downloaded and extracted |
| `${pkgdir}` | `Package` directory for the root of the package |

### spec

| key | type | value |
| --- | ---- | ----- |
| `targets` | `list` | List of build targets only used for projects. Prefix a `!` to ignore target. |
| `pkgname` | `string` | Package name |
| `pkgver` | `string` | Package version |
| `pkgrel` | `string` | Package release number |
| `pkgdesc` | `string` | Short package description |
| `pkgdesclong` | `list` | List of lines for package description |
| `maintainer` | `string` | Package maintainer |
| `arch` | `string` | Package architecture, can be `all` or `amd64` |
| `license` | `list` | List of licenses for packaged software |
| `section` | `string` | Section for package. Built in sections available:<br> `admin`<br> `localization`<br> `mail`<br> `comm`<br> `math`<br> `database`<br> `misc`<br> `debug`<br> `net`<br> `news`<br> `devel`<br> `doc`<br> `editors`<br> `electronics`<br> `embedded`<br> `fonts`<br> `games`<br> `science`<br> `shells`<br> `sound`<br> `graphics`<br> `text`<br> `httpd`<br> `vcs`<br> `interpreters`<br> `video`<br> `web`<br> `kernel`<br> `x11`<br> `libdevel`<br> `libs` |
| `priority` | `string` | Package priority, only used for debian packages |
| `url` | `string` | Package url |
| `depends` | `list` | List of package dependencies |
| `optdepends` | `list` | List of package optional dependencies |
| `makedepends` | `list` | List of package build dependencies |
| `provides` | `list` | List of packages provided |
| `conflicts` | `list` | List of packages conflicts |
| `sources` | `list` | List of packages sources. Sources can be url or paths that are relative to the PKGBUILD |
| `hashsums` | `list` | List of `md5`/`sha1`/`sha256`/`sha512` hex hashes for sources, hash type is determined by the length of the hash. Use `skip` to ignore hash check |
| `backup` | `list` | List of config files that shouldn't be overwritten on upgrades |
| `build` | `func` | Function to build the source, starts in srcdir |
| `package` | `func` | Function to package the source into the pkgdir, starts in srcdir |
| `preinst` | `func` | Function to run before installing |
| `postinst` | `func` | Function to run after installing |
| `prerm` | `func` | Function to run before removing |
| `postrm` | `func` | Function to run after removing |

### build targets

| target | value |
| ------ | ----- |
| `archlinux` | All archlinux releases |
| `amazonlinux` | All amazonlinux releases |
| `centos` | All centos releases |
| `debian` | All debian releases |
| `fedora` | All fedora releases |
| `ubuntu` | All ubuntu releases |
| `amazonlinux-2016.09` | Amazonlinux 2016.09 |
| `amazonlinux-2017.03` | Amazonlinux 2017.03 |
| `amazonlinux-2017.09` | Amazonlinux 2017.09 |
| `centos-6` | Centos 6 |
| `centos-7` | Centos 7 |
| `debian-wheezy` | Debian wheezy |
| `debian-jessie` | Debian jessie |
| `debian-stretch` | Debian stretch |
| `fedora-24` | Fedora 24 |
| `fedora-25` | Fedora 25 |
| `fedora-26` | Fedora 26 |
| `fedora-27` | Fedora 27 |
| `ubuntu-precise` | Ubuntu precise |
| `ubuntu-trusty` | Ubuntu trusty |
| `ubuntu-xenial` | Ubuntu xenial |
| `ubuntu-yakkety` | Ubuntu yakkety |
| `ubuntu-zesty` | Ubuntu zesty |
| `ubuntu-artful` | Ubuntu artful |
| `ubuntu-bionic` | Ubuntu bionic |

### directives

| directive | value |
| --------- | ----- |
| `apt` | All deb packages |
| `pacman` | All pkg packages |
| `yum` | All rpm packages |
| `archlinux` | All archlinux releases |
| `amazonlinux` | All amazonlinux releases |
| `centos` | All centos releases |
| `debian` | All debian releases |
| `fedora` | All fedora releases |
| `ubuntu` | All ubuntu releases |
| `amazonlinux-2016.09` | Amazonlinux 2016.09 |
| `amazonlinux-2017.03` | Amazonlinux 2017.03 |
| `amazonlinux-2017.09` | Amazonlinux 2017.09 |
| `centos-6` | Centos 6 |
| `centos-7` | Centos 7 |
| `debian-wheezy` | Debian wheezy |
| `debian-jessie` | Debian jessie |
| `debian-stretch` | Debian stretch |
| `fedora-24` | Fedora 24 |
| `fedora-25` | Fedora 25 |
| `fedora-26` | Fedora 26 |
| `fedora-27` | Fedora 27 |
| `ubuntu-precise` | Ubuntu precise |
| `ubuntu-trusty` | Ubuntu trusty |
| `ubuntu-xenial` | Ubuntu xenial |
| `ubuntu-yakkety` | Ubuntu yakkety |
| `ubuntu-zesty` | Ubuntu zesty |
| `ubuntu-artful` | Ubuntu artful |
| `ubuntu-bionic` | Ubuntu bionic |

Directives are used to specify variables that only apply to a limited set of
build targets. All variables can use directives including user definied
variables. To use directives include the directive after a
variable seperated by a colon such as
`pkgdesc:ubuntu="This description will only apply to Ubuntu packages"`
The directives above are sorted lowest to highest priority.

### example

First create a directory for the PKGBUILD file. This directory should only
contain the PKGBUILD file and any other files needed such as patches. Then
create a PKGBUILD the package directory. After creating the PKGBUILD build
the package with docker.

```
$ mkdir httpserver
$ cd httpserver
$ nano PKGBUILD
$ docker run --rm -t -v `pwd`:/pacur pacur/ubuntu-trusty
```

```
targets=(
    "!centos-6"
    "archlinux"
    "centos"
    "debian"
    "fedora"
    "ubuntu"
)
pkgname="httpserver"
pkgver="1.0"
pkgrel="1"
pkgdesc="Http file server written with Go"
pkgdesc:centos="Http file server written with Go for CentOS"
pkgdesc:debian="Http file server written with Go for Debian"
pkgdesc:fedora="Http file server written with Go for Fedora"
pkgdesc:ubuntu="Http file server written with Go for Ubuntu"
pkgdesclong=(
    "Quick http file server written with Go"
    "using directory listing similar to apache"
)
maintainer="Example <example@pacur.org>"
arch="all"
license=("GPLv3")
section="utils"
priority="optional"
url="https://github.com/pacur/${pkgname}"
sources=(
    "${url}/archive/${pkgver}.tar.gz"
)
hashsums=(
    "3548e1263a931b27970e190f04b74623"
)

build() {
    mkdir -p "go/src"
    export GOPATH="${srcdir}/go"
    mv "${pkgname}-${pkgver}" "go/src"
    cd "go/src/${pkgname}-${pkgver}"
    go get
    go build -a
}

package() {
    cd "${srcdir}/go/src/${pkgname}-${pkgver}"
    mkdir -p "${pkgdir}/usr/bin"
    cp ${pkgname}-${pkgver} ${pkgdir}/usr/bin/${pkgname}
}
```

### project example

A project can be created with the cli tools which can be installed using
go get. The packages can be built and added to the repo. An example project is
available in the example directory. The `pull` command should be run before
all builds to update the docker images used for builds.

```
$ go get github.com/pacur/pacur
$ cd example
$ pacur pull
$ pacur project init
$ pacur project build
$ pacur project repo
$ go get github.com/pacur/httpserver
$ cd mirror
$ httpserver --port 80
```

After the repo has been created and is hosted on a server the following
commands can be used to add the repo to the package manager for yum and apt.
For the debian repo the `jessie` should be replaced with the debian/ubuntu
release name. The `pacur` repo name and filenames can be change to suite the
name of your software.

```
$ nano /etc/pacman.conf
[pacur]
Server = http://HTTP_SERVER_IP/arch

$ pacman-key --keyserver hkp://pgp.mit.edu -r KEYID
$ pacman-key --lsign-key KEYID
$ pacman -Sy
$ pacman -S httpserver
```

```
$ nano /etc/yum.repos.d/pacur.repo
[pacur]
name=Pacur Repository
baseurl=http://HTTP_SERVER_IP/yum/centos/7/
gpgcheck=1
enabled=1

$ gpg --keyserver hkp://pgp.mit.edu --recv-keys KEYID
$ gpg --armor --export KEYID > key.tmp; rpm --import key.tmp; rm -f key.tmp
$ yum install httpserver
```

```
$ nano /etc/apt/sources.list.d/pacur.list
deb http://HTTP_SERVER_IP/apt jessie main

$ apt-key adv --keyserver hkp://pgp.mit.edu --recv KEYID
$ apt-get update
$ apt-get install httpserver
```

### signing

Packages in the repository can also be signed by adding a `sign.key` in the
project directory. The signing key cannot use a passphrase. To export a key
first get the key id then export the key with the commands below.

```
$ gpg --list-secret-keys
$ gpg -a --export-secret-keys KEYID > sign.key
```

### generate key

A pacur docker container is available to generate signing keys. A key can be
generated by first pulling the genkey container then running the container in
the project directory. This will generate a `sign.key` and `sign.pub` the
public key will need to be uploaded to [pgp.mit.edu](https://pgp.mit.edu/)

```
$ docker pull pacur/genkey
$ docker run --rm -v `pwd`:/pacur pacur/genkey First Last email@domain.com
```
