# Ceres

The Comunion business backend service.

## Projects 

We have chosen the [ego](https://github.com/gotomicro/ego) to support the config and other common behaviors of web project.
Fistly, you have to read the documents of Ego seriously.

The user interface API we will support as HTTP APIs, and the eth event will be trigger from other service by gRpc interfaces.

## How to contribute?

Before you start to make PR for Ceres these thing you have to know:

1. Do not use egoctl to generate project file, all code we have to write hand.

As you can see that the whole project business code will be emplaced in the `pkg` directory. And the configuration of Ceres will be placed in `config` directory. But the `metas` directory consistes of many meta informations not only the database schema file and the gRpc definition proto files.

### A simple example 

Here is a simple example which could help you make PR for us.

