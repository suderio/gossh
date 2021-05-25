# GOSH - Go Secure Shell

A small decorator over tradicional ssh to easy working with many different hosts.

## Usage

Connecting to a new host:
```bash
> gosh open somehost.com
```

Usual arguments:
```bash
> gosh open -p 22 -u xpto somehost.com
```

Unusual arguments:
```bash
> gosh open --flag some --flag com somehost.com
```

Listing all known hosts:
```bash
> gosh ls
```

Listing some known hosts:
```bash
> gosh ls some
```