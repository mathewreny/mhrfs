// "sshftp" implements SFTP version 3. This is the most common version seen in
// the wild, and the version that OpenSSH supports. This is a client only
// implementation. SSH connections are handled by the library
// golang.org/x/crypto/ssh for convienence and interoperability. This client
// library seeks to fully implement the protocol, but will first focus on
// functions and methods relevant to mhrfs. When version 1.0.0 is ready, this
// package will move to its own separate repository.
package sshftp
