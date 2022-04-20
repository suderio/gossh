# GOSSH - Go Secure Shell

A small decorator over tradicional ssh to easy working with many different hosts.

## Usage

Connecting to a new host:
```bash
> gossh somehost.com
```

Usual arguments:
```bash
> gossh -p 22 -u xpto somehost.com
```

Unusual arguments:
```bash
> gossh --flag some --flag com somehost.com
```

Listing all known hosts:
```bash
> gossh ls
```

Listing some known hosts:
```bash
> gossh ls some
```
