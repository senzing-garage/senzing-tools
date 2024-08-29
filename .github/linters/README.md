# Linters

## .checkov.yaml

- [.checkov.yaml]
- Used by [lint-workflows.yaml]
- [checkov]
  - [checkov configuration]

## .golangci.yaml

- [.golangci.yaml]
- Used by [golangci-lint.yaml]
- [golangci-lint]
  - [golangci-lint configuration]

## .jscpd.json

- [.jscpd.json]
- Used by [lint-workflows.yaml]
- [jscpd]
  - [jscpd configuration]
  - Example:

      ```json
      {
      "ignore": [
          "**/*.go,**/go-test*.yaml"
      ],
      "threshold": 10
      }
      ```

## .yaml-lint.yml

- [.yaml-lint.yml]
- Used by [lint-workflows.yaml]
- [yaml-lint]
  - [yaml-lint configuration]

[.checkov.yaml]:  .checkov.yaml
[.golangci.yaml]: .golangci.yaml
[.jscpd.json]: .jscpd.json
[.yaml-lint.yml]: .yaml-lint.yml
[checkov configuration]: https://www.checkov.io/2.Basics/CLI%20Command%20Reference.html
[checkov]: https://www.checkov.io/
[golangci-lint configuration]: https://golangci-lint.run/usage/configuration/
[golangci-lint.yaml]: ../workflows/README.md#golangci-lintyaml
[golangci-lint]: https://golangci-lint.run/
[jscpd configuration]: https://github.com/kucherenko/jscpd/tree/master/apps/jscpd#options
[jscpd]: https://github.com/kucherenko/jscpd
[lint-workflows.yaml]: ../workflows/README.md#lint-workflowsyaml
[yaml-lint configuration]: https://yamllint.readthedocs.io/en/stable/configuration.html
[yaml-lint]: https://github.com/adrienverge/yamllint
