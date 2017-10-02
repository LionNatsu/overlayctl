# overlayctl
An overlayfs controller utils

## Usage
```
Commands:
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
                overlayctl merge bottom0 bottom1 middle /file/to/merge
```
