# What is mhrfs

The *mount home remote filesystem* is a more convienent way to mount remote home directories.  It is an on-demand FUSE filesystem daemon that uses SSH to securely connect with servers. Remote filesystems are mounted when a user, or program, requests information from the filesystem.

## How to use mhrfs

Access remote home directories using the following path scheme.

**`/mhr/[user@]host/`**

Setting up public key SSH authentication is the only configuration step required to use mhrfs. This allows the mhrfs daemon to automatically connect to servers when needed. Mhrfs requires the host to be located in an ssh "config" file.  If the `User` option is present in the config file, then the user is optional in the path.

Hosts must have up-to-date keys in the default ssh "known_hosts" files.

Examples: `jdoe@reny.io`

    # Copy a file to remote.
    cp /mhr/jdoe@reny.io/foo.txt bar.txt
    
    # Move file to remote.
    mv baz.txt /mhr/jdoe@reny.io/
    
    # Remove the file from remote.
    rm /mhr/jdoe@reny.io/foo.txt
    
    # List files in your remote home directory.
    ls /mhr/jdoe@reny.io/
    
    # You can even do remote to remote file transfers.
    # Afterall, these are just paths in your filesystem.
    mv /mhr/jdoe@10.10.0.5/foo.txt /mhr/jdoe@reny.io/


## Security by design.

- Operates entirely within userland.
  - The mhrfs daemon exits when run as `root`.
  - SSH client and server connections are also in userland
- Built using the well established SSH SFTP protocol.
  - Connections are encrypted end-to-end.
  - Servers don't require special daemons.
- Directory traversal programs are hindered by on demand mounting.
  - The `/mhr/` directory does not save history.
  - Tab completion only works once inside the remote home directory.

## Contribute to mhrfs

The mhrfs software architecture is actively being designed.  Please contact me if you would like to help. Breaking changes will be considered up until version 1.0.0 is reached.

## License 

The license will be decided in the future.  Suggestions welcome.
