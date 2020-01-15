# gce-metadata

## Usage

```
$ gce-metadata hostname
api-client.c.jbd-deployments.internal
```

Help text:

```
$ gce-metadata help
gce-metadata <cmd> args...
Commands:
- hostname
- external-ip
- internal-ip
- instance-name
- zone
- project-id
- instance-id
- get <attr>
- watch <attr>
```

## Installation

```
$ curl https://storage.googleapis.com/jbd-releases/gce-metadata > gce-metadata && chmod +x ./gce-metadata
```
