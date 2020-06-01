# Protecc - REST Write Blocker

![ForTheBadge powered-by-electricity](http://ForTheBadge.com/images/badges/powered-by-electricity.svg)

![GitHub last commit](https://img.shields.io/github/last-commit/nankeen/sauvignon)
![GitHub](https://img.shields.io/github/license/nankeen/sauvignon)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/nankeen/sauvignon)

---

Protecc is a HTTP reverse proxy that blocks data modifying REST methods like `PUT`, `POST`, and `DELETE`.
It can be used as a write blocker for read only services, e.g. staging databases.
Can be used to maintain a peace of mind tool while working with sensitive services.

## Usage

It's very easy to use, just point it to a HTTP services.
Proxy configuration is not required for transparent mode.

```bash
$ protecc -target 10.42.0.69:1337
```

```
Usage of protecc:
  -listen string
        Where the proxy server listens on (default "localhost:1337")
  -target string
        Where the requests should be proxied to. (default "localhost")
```
