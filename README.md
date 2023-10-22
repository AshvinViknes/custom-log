# Custom Logger

A straightforward Go (Golang) package for custom logging with additional information such as timestamp, file details, and URL details.

## Overview

This package provides a custom logger with three main log levels: `Info`, `InfoURL`, and `Error`. It enhances traditional logging by including contextual information like timestamp, folder, file, and URL details. This can be particularly useful for debugging and understanding the context in which log messages are generated.

```markdown
## Installation

```bash
go get -u github.com/AshvinViknes/customlogger
```

## Usage

1. Import the package:

    ```go
    import (
        "github.com/AshvinViknes/customlogger"
        // Other imports...
    )
    ```

2. Create a new instance of the logger:

    ```go
    logger := customlogger.NewCustomLogger()
    ```

3. Log messages at different levels:

    ```go
    // Log an informational message with an additional value
    logger.Info("This is an informational message", "SomeValue")

    // Log an informational message with a URL (HTTP method inferred from the URL)
    u, _ := url.Parse("https://example.com")
    logger.InfoURL("This is an informational message with a URL", u)

    // Log an error message with an additional value
    logger.Error("This is an error message", 42)
    ```

## Example Output

```
[2023/10/22 15:04:05]INFO: This is an informational message
	FOLDER: your-folder
	FILE: your-file.go:42
	INFO VALUE: SomeValue

[2023/10/22 15:04:05]INFO: This is an informational message with a URL
	URL: https://example.com
	ADDRESS: example.com
	FOLDER: your-folder
	FILE: your-file.go:42
	SCHEME: HTTPS

[2023/10/22 15:04:05]ERROR: This is an error message
	FOLDER: your-folder
	FILE: your-file.go:42
	ERROR VALUE: 42
```

## Acknowledgments

Special thanks to the Go community and the contributors of the packages used in this project.

---
