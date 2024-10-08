# Yunroxy
A fully scaleable HTTP/HTTPS/SOCKS4/SOCKS5 Proxy Network, a great tool to manage and scrape proxies.

*The Project consists of two parts:*
- **Proxy Manager (`/api`)**
- **Proxy Scraper (`/updater`)**

This design makes it perfect for personal and commertial uses.

## Features
- Highly scalable
- Easy to setup for multiple users
- Ensures every proxy is working and safe

## Installation
Clone this [repository](https://github.com/vxoid/yunroxy)
Download [Go Compiler](https://go.dev/doc/install) (1.19<=)

Build the project by running `go build -o yunroxy .`

Build the CLI Tools by running `cd cli && go build -o ../yunroxy-cli .`

## Usage
### Create API Key
Running this command will create a new api key.

`yunroxy-cli api-key new`
### Run
After running the program, it will run api on http://0.0.0.0:11555 by default

`./yunroxy`
### API
Create an API Key using the Yunroxy CLI tool

`yunroxy-cli api-key new`

Simply make a GET request to the following endpoint

`http://0.0.0.0:11555/proxy/random&api_key=<API-KEY>`, is the only endpoint available for now.