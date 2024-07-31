### CloudBees Platform action REST API example
This is an action to execute REST calls from within platform and get the response as an output.

This does not require the calls to be authenticated, but it supports basic auth along with using a Bearer token.

| Input name             | Required? | Description                                                                                                   |
|------------------------|-----------|---------------------------------------------------------------------------------------------------------------|
| url                    | Yes       | The URL of the endpoint that is being hit                                                                     |
| request-type           | Yes       | The type of request. Currently supported: GET, POST, PUT, DELETE                                              |
| payload                | No        | The payload POST/PUT calls                                                                                    |
| bearer-token           | No        | The bearer token to use for authentication (If supplied it will ignore username/password parameters)          |
| username               | No        | The username to use for authentication                                                                        |
| password               | No        | The password to use for authentication (required when using a username)                                       |
| expected-response-code | No        | When provided, the action will return an error if the response code does not match the expected response code |

# Usage example
The below workflow yaml shows a minimum needed to use the action, set the output, and read the output in a later step
```yaml
apiVersion: automation.cloudbees.io/v1alpha1
kind: workflow
name: My workflow
on:
  push:
    branches:
      - "**"
  workflow_dispatch:
jobs:
  build:
    outputs:
      output1: $${{ steps.rest-action.outputs.response }}
    steps:
      - uses: gradinsky-cloudbees/platform-rest@v1.0.0
        name: custom
        id: rest-action
        with:
          url: https://fakeurl.local/api/v1/update/21
          request-type: PUT
          payload: '{"name":"test","salary":"123","age":"23"}'
  job2:
    steps:
      - uses: docker://alpine:3.18
        name: echo
        env:
          OUTPUT1: $${{needs.build.outputs.output1}}
        run: echo "$OUTPUT1"
    needs: build


```
The above example only shows using 1 non-required input parameter. To add more, simply add them to the key/value pairs under the `with` array.