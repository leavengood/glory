# glory
A package and command-line tool for updating long-running processes

## Overview

Your server uses this package to provide an endpoint over HTTP for being notified
of updates. You use the tool of your choice to upload an update to a file server
you control, and then the glory command-line tool can be used to notify you server
that an update is available. The glory code in your server will then download the
update and verify its checksum before replacing the existing server binary with
the new version.

At that point your server can exit and your chosen supervisor system will restart
it (using the newly updated binary), or you can make use of graceful restart
libraries, such as https://github.com/fvbock/endless.

## Using the Package

Add something like the following to your HTTP server

```go
glory.Glorify("secret", func() {
  // Imagine this was a graceful restart, or assume this is managed by a supervisor
  os.Exit(0)
})
```

This will expose an endpoint at /glory/update/available which can accept requests
signed with the given secret. The given function is called when an update occurs.

Currently when an update succeeds the existing executable will be moved to
executable_name_old and the update will be moved to the name of the existing
executable.

## Commands

Currently the command-line tool `glory` only sends an update request with the server URL,
updated file URL, updated file SHA1 and shared secret provided on the command-line.
