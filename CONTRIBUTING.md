# Getting Involved
 
Thanks for your interest in the project, we'd love to have you involved!
 
## Opening an Issue
 
We always welcome issues, if you've seen something that isn't quite right or you have a suggestion for a new feature, please go ahead and open an issue in this project. Include as much information as you have, it really helps.
 
## Making a Code Change

**Please open an issue for discussion before creating a pull request on this project** - this library is a work in progress and we don't want to waste/duplicate anyone's efforts.

Please note that only the code in the top level of the project, and the `examples/` directory is available for pull requests. The other directories contain code generated from [our OpenAPI specs](https://github.com/Nexmo/api-specification/tree/master/definitions) using [OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator).

When you're ready to start coding, fork this repository to your own GitHub account and make your changes in a new branch. Once you're happy (see below for information on how to run the tests), open a pull request and explain what the change is and why you think we should include it in our project.

### Generate Code from OpenAPI Specs

We're using [OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator) to generate code from the OpenAPI specifications. This output is **never edited**, but wrapped by user-facing code. It may be regenerated and completely replaced when the API spec updates.

For the docker setup of the code generator tool, try a command like this one:

```
docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli generate -i https://raw.githubusercontent.com/Nexmo/api-specification/master/definitions/sms.yml -g go --package-name sms -o /local/out/golang
```

Then copy the resulting `*.go` files into a directory named to match the `package-name` value given in your command and matching the API spec used.
 
### Running Tests

We run the tests when we build the project, including when you open a pull request. To run the test suite locally, do:

```
go test
```
 
**Pro tip:** if you get an error about a missing module or a newer version of Go being required, try setting your `GO111MODULE` environment variable to `"on"`.

