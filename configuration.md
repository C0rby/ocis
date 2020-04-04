---
title: "Configuration"
date: "2020-04-03T21:39:56"
weight: 20
geekdocRepo: https://github.com/owncloud/ocis
geekdocEditPath: edit/master/docs
geekdocFilePath: configuration.md
---

{{< toc >}}

## Configuration

## Configuration of extensions

oCIS Single Binary is not responsible for configuring extensions. Instead, each extension could either be configured by environment variables, cli flags or config files.

Each extension has its dedicated documentation page (e.g. https://owncloud.github.io/extensions/ocis_proxy/configuration) which lists all possible configurations. Config files and environment variables are picked up if you use the `./bin/ocis server` command within the oCIS single binary. Command line flags must be set explicitly on the extensions subcommands.

## Configuration using config files

Out of the box extensions will attempt to read configuration details from:

```console
/etc/ocis
$HOME/.ocis
./config
```

For this configuration to be picked up, have a look at your extension `root` command and look for which default config name it has assigned. *i.e: ocis-proxy reads `proxy.json | yaml | toml ...`*.

### Configuration file

So far we support the file formats `JSON` and `YAML`, if you want to get a full example configuration just take a look at [our repository](https://github.com/owncloud/ocis/tree/master/config), there you can always see the latest configuration format. These example configurations include all available options and the default values. The configuration file will be automatically loaded if it's placed at `/etc/ocis/ocis.yml`, `${HOME}/.ocis/ocis.yml` or `$(pwd)/config/ocis.yml`.

## Envrionment variables

If you prefer to configure the service with environment variables you can see the available variables below.

### Server

OCIS_TRACING_ENABLED
: Enable sending traces.

OCIS_TRACING_TYPE
: Tracing backend type. Default: `jaeger`.

OCIS_TRACING_ENDPOINT
: Endpoint for the agent.

OCIS_TRACING_COLLECTOR
: Endpoint for the collector.

OCIS_TRACING_SERVICE
: Service name for tracing. Default: `ocis`.

OCIS_DEBUG_ADDR
: Address to bind debug server. Default: `0.0.0.0:9010`.

OCIS_DEBUG_TOKEN
: Token to grant metrics access.

OCIS_DEBUG_PPROF
: Enable pprof debugging.

OCIS_DEBUG_ZPAGES
: Enable zpages debugging.

OCIS_HTTP_ADDR
: Address to bind http server. Default: `0.0.0.0:9000`.

OCIS_HTTP_ROOT
: Root path of http server. Default: `/`.

OCIS_GRPC_ADDR
: Address to bind grpc server. Default: `0.0.0.0:9001`.

### Root Command

OCIS_CONFIG_FILE
: Path to config file.

OCIS_LOG_LEVEL
: Set logging level. Default: `info`.

OCIS_LOG_PRETTY
: Enable pretty logging. Default: `true`.

OCIS_LOG_COLOR
: Enable colored logging. Default: `true`.

### Health

OCIS_DEBUG_ADDR
: Address to debug endpoint. Default: `0.0.0.0:9010`.

## Commandline flags

If you prefer to configure the service with commandline flags you can see the available variables below.

### Server

--tracing-enabled
: Enable sending traces.

--tracing-type
: Tracing backend type. Default: `jaeger`.

--tracing-endpoint
: Endpoint for the agent.

--tracing-collector
: Endpoint for the collector.

--tracing-service
: Service name for tracing. Default: `ocis`.

--debug-addr
: Address to bind debug server. Default: `0.0.0.0:9010`.

--debug-token
: Token to grant metrics access.

--debug-pprof
: Enable pprof debugging.

--debug-zpages
: Enable zpages debugging.

--http-addr
: Address to bind http server. Default: `0.0.0.0:9000`.

--http-root
: Root path of http server. Default: `/`.

--grpc-addr
: Address to bind grpc server. Default: `0.0.0.0:9001`.

### Root Command

--config-file
: Path to config file.

--log-level
: Set logging level. Default: `info`.

--log-pretty
: Enable pretty logging. Default: `true`.

--log-color
: Enable colored logging. Default: `true`.

### Health

--debug-addr
: Address to debug endpoint. Default: `0.0.0.0:9010`.
