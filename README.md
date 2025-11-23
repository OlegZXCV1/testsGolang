# Go Test Report

[![Pipeline Status](https://github.com/becheran/go-testreport/actions/workflows/go.yml/badge.svg)](https://github.com/becheran/go-testreport/actions/workflows/go.yml)
[![Go Report Card][go-report-image]][go-report-url]
[![PRs Welcome][pr-welcome-image]][pr-welcome-url]
[![License][license-image]][license-url]
[![GHAction][gh-action-image]][gh-action-url]

[license-url]: https://github.com/becheran/go-testreport/blob/main/LICENSE
[license-image]: https://img.shields.io/badge/License-MIT-brightgreen.svg
[go-report-image]: https://img.shields.io/badge/go%20report-A+-brightgreen.svg?style=flat
[go-report-url]: https://goreportcard.com/report/github.com/becheran/go-testreport
[pr-welcome-image]: https://img.shields.io/badge/PRs-welcome-brightgreen.svg
[pr-welcome-url]: https://github.com/becheran/go-testreport/blob/main/CONTRIBUTING.md
[gh-action-image]: https://img.shields.io/badge/Get-GH_Action-blue
[gh-action-url]: https://github.com/marketplace/actions/golang-test-report

## Install

### Go

Install via the go install command:

``` sh
go install github.com/becheran/go-testreport@latest
```

### Binaries

Or use the pre-compiled binaries for Linux, Windows, and Mac OS from the [github releases page](https://github.com/becheran/go-testreport/releases).

## Usage

Run the following command to get a list of all available command line options:

``` sh
go-testreport -h
```

### Input and Output

When `-input` and `-output` is not set, the stdin stream will be used and return the result will be written to stdout. Will exit with a non zero exit code if at least one test failed:

``` sh
go test ./... -json | go-testreport > result.html
```

Use the `-input` and `-output` file to set files for the input and output. Will always exit with zero also if tests fail:

``` sh
go-testreport -input result.json -output result.html
```

### Running Tests

To run all tests and generate an HTML report, use the following command:

```bash
go test -json ./... | ~/bin/go-testreport > report.html
```

After running the command, you can open `report.html` in your web browser to view the test results.

### Allure Report

To run all tests and generate a unified Allure report, use the following command. This will create an `allure-results` directory.

```bash
go test -json -cover ./... | ~/go/bin/golurectl -o allure-results
```

After the tests have run, you can generate and serve the report with this command:

```bash
allure serve allure-results
```

### Templates

Customize by providing a own [template file](https://pkg.go.dev/text/template). See also the [default markdown template](./src/report/templates/md.tmpl) which is used if the `-template` argument is left empty. With the `vars` options custom dynamic values can be passed to the template from the outside which can be resolved within the template:

``` sh
go test ./... -json | go-testreport -template=./html.tmpl -vars="Title:Test Report Linux" > $GITHUB_STEP_SUMMARY
```

### GitHub Actions

The [Golang Test Report](https://github.com/marketplace/actions/golang-test-report) from the marketplace can be used to integrate the go-testreport tool into an GitHub workflow:

``` yaml
- name: Test
  run: go test ./... -json > report.json
- name: Report
  uses: becheran/go-testreport@main
  with:
    input: report.json
```
