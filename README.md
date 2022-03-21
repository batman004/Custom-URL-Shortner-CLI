# Custom-URL-Shortner-CLI
A lightweight CLI tool written in Go to easily shorten URLs and create custom bookmarks to refer and keep track of  websites you bookmark

## Instructions to Run and Build from Source:

- Pre-requisites
  - Go
- Directions to install

```bash 
git clone https://github.com/batman004/Custom-URL-Shortner-CLI.git
```
CD into the project DIR
```bash 
cd Custom-URL-Shortner-CLI.git
```

Run the build command to create a executable binary of the project
```bash 
go build
```

- Directions to execute

(if you're on linux/MacOS)

```bash
./Custom-URL-Shortner-CLI <COMMAND>
```

(if you're on windows)

```bash
Custom-URL-Shortner-CLI.exe <COMMAND>
```

Or, you can check out the pre-compiled binaries under **Releases**

1. Download the binary for your system (check releases)
2. Create a folder Data and an empty json file with the name `saved_urls.json`
3. Execute the tool by executing the binary according to your OS (see above)

- Usage

```
Usage:
  Custom-URL-Shortner-CLI [command]

Available Commands:
  add        add a new shortened, custom-tagged URL
  get        get bookmarked URL(s)

Flags:
  -custom-tag   help for gtasks
  -description   Help message for toggle
  --all          get all bookmarked URL(s)
  -custom        get bokmarked URL based on custom tag
  -url           add the URL to be shortened
 
```

## Commands

- Creating a new custom-tagged, shortened URL

```bash
 Custom-URL-Shortner-CLI add -custom-tag <tag> -description <desc> -url <url>
```

- Getting all the bookmarked URL(s)
```bash
 Custom-URL-Shortner-CLI get --all
```
- Getting a custom-tagged URL
```bash
 Custom-URL-Shortner-CLI add -custom <tag>
```



