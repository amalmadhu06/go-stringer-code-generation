## Generating Code Using Stringer CLI
To generate code using the Stringer CLI tool, follow these simple steps:

**Step 1: Installation**<br>
Make sure you have Go installed on your system. If you don't have it installed, you can download and install it from the official Go website: [https://golang.org/dl/](https://golang.org/dl/).

**Step 2: Install `stringer` tool**<br>


```go
go install  "golang.org/x/tools/cmd/stringer"
```

**Step 3: Add the //go:generate Directive**<br>
Place the `//go:generate` directive comment above the type declaration you want to generate code for. In your example, it should be placed above the StatusCode type declaration:
```go
//go:generate stringer -type=StatusCode
type StatusCode int

const (
    // Your constant definitions here
)
```

**Step 4: Run the Generate Command**<br>
Open a terminal or command prompt, navigate to the root of your project, and execute the following command:
```shell
go generate ./...
```
This command will automatically invoke the `go:generate` directive and execute the Stringer tool. It will analyze your code and generate the necessary stringer methods.

**Step 5: Generated Code**<br>
After running the go generate command, you will find a new file named statuscode_string.go (in my example, it would be http_string.go) in the same package directory. This file will contain the generated code for the String() method based on the StatusCode type.<br>

<i>Note: Make sure to re-run the go generate command whenever you make changes to your StatusCode type or its constants to ensure that the generated code stays up to date.