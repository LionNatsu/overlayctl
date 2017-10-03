package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		help()
		return
	}
	switch os.Args[1] {
	case "mount":
		var readOnly bool
		var flags = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
		flags.BoolVar(&readOnly, "ro", false, "mount the overlay filesystem as read-only")
		flags.Parse(os.Args[2:])
		args := flags.Args()
		if len(args) < 3 {
			help()
			os.Exit(127)
		}
		var olfs = Instance{WorkDir: args[len(args)-1], Layers: args[0 : len(args)-1], ReadOnly: readOnly}
		if err := olfs.Mount(); err != nil {
			log.Fatalln(err)
		}
	case "unmount":
		if len(os.Args) < 3 {
			help()
			os.Exit(127)
		}
		var olfs Instance
		if len(os.Args) == 3 {
			olfs = Instance{WorkDir: os.Args[len(os.Args)-1]}
		} else {
			olfs = Instance{WorkDir: os.Args[len(os.Args)-1], Layers: os.Args[2 : len(os.Args)-1]}
		}
		if err := olfs.Unmount(); err != nil {
			log.Fatalln(err)
		}
	case "merge":
		if len(os.Args) < 5 {
			help()
			os.Exit(127)
		}
		var olfs = Instance{WorkDir: "", Layers: os.Args[2 : len(os.Args)-1]}
		if err := olfs.Merge(os.Args[len(os.Args)-1], len(os.Args)-4, len(os.Args)-5); err != nil {
			log.Fatalln(err)
		}
	default:
		help()
		os.Exit(127)
	}
}

func help() {
	fmt.Println(
		`Commands:
	mount [-ro] <bottom0> ...<bottomN> <top> <workdir>;
	unmount [<top>] <workdir>;
	merge <bottom0> ...<bottomN> <dest> <source> <path>
Example:
	Create a simple 2-layer filesystem:
		overlayctl mount test/lower test/upper /mnt/workdir
	Create a 3-layer read-only filesystem:
		overlayctl mount bottom middle top /mnt/workdir2
	Unmount it:
		overlayctl unmount /mnt/workdir2
	Unmount it and delete temporary directory (test/upper.tmp):
		overlayctl unmount test/upper /mnt/workdir
	Merge a directory from middle to bottom1 layer:
		overlayctl merge bottom0 bottom1 middle /file/to/merge`)
}
