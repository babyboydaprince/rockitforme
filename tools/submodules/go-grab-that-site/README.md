# Go grab that site

## Overview

GO grab that site is a tool written in Go that downloads an entire website, including HTML pages and assets, while maintaining the original structure.

## Features

- Downloads entire websites
- Extracts and follows links from HTML pages
- Saves all media files and formatting locally
- Preserves directory structure

## Installation

1. Ensure you have **Go** installed on your system.
2. Clone the repository or copy the source code into a local directory.
3. Navigate to the project directory and run:

   ```sh
   go mod init go_grab_that_site
   go mod tidy
   ```

## Usage

Run the tool using the command:

```sh
go run main.go
```

This will start downloading the website specified in `startURL` and save it in the `downloaded_site` directory.

## Functions Explained

### `DownloadFile(url, filePath string) error`

This function fetches a URL and saves its contents to a local file, ensuring directory structure is preserved.

**Parameters:**

- `url`: The URL of the file to download.
- `filePath`: The local path where the file should be saved.

**Returns:**

- `error`: Returns an error if the file cannot be downloaded.

### `ExtractLinks(url string) ([]string, error)`

Extracts all links (`href` and `src` attributes) from an HTML document.

**Parameters:**

- `url`: The webpage URL to extract links from.

**Returns:**

- `[]string`: A list of extracted URLs.
- `error`: Returns an error if the extraction fails.

### `main()`

- Defines the starting URL (`startURL`).
- Calls `ExtractLinks` to gather all internal and external links.
- Calls `DownloadFile` for each extracted link.

## Future Improvements

- Add concurrency for faster downloads.
- Support for handling JavaScript-generated content.
- Implement a crawler depth limit.

## License

This project is licensed under the MIT License.
