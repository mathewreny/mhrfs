# What is mhrfs

The *mount home remote filesystem* is the most convienent way to mount remote home directories.  It is an on-demand FUSE filesystem daemon that uses SSH (keys) to securely connect with servers.

## How to use mhrfs

Access remote home directories using the following path scheme.

**`/mhr/<host>[:<port>]/<user>/`**

Setting up public key SSH authentication is the only configuration step required to use mhrfs. This allows the mhrfs daemon to automatically connect to servers when needed.

Examples: `jdoe@reny.io`

    # Copy a file to remote.
    cp /mhr/reny.io/jdoe/foo.txt bar.txt
    
    # Move file to remote.
    mv baz.txt /mhr/reny.io/jdoe/
    
    # Remove the file from remote.
    rm /mhr/reny.io/jdoe/foo.txt
    
    # List files in your remote directory.
    ls /mhr/reny.io/jdoe/
    
    # You can even do remote to remote file transfers.
    # Afterall, these are just paths in your filesystem.
    mv /mhr/10.10.0.5/jdoe/foo.txt /mhr/reny.io/jdoe/


## Security by design.

- Operates entirely within userland.
  - The mhrfs daemon exits when run as `root`.
  - SSH client and server connections are also in userland
- Built using the well established SSH protocol.
  - Connections are encrypted end-to-end.
  - Servers don't require custom daemons.
- Directory traversal programs are hindered by on demand mounting.
  - The `/mhr/` directory does not save history.
  - Tab completion only works once inside the remote home directory.

## Contribute to mhrfs

The mhrfs software architecture is actively being designed.  Please contact me if you would like to help. Breaking changes will be considered up until version 1.0.0 is reached.

## License 

The license will be decided in the future.  Suggestions welcome.
