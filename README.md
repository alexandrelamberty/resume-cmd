[![Build](https://github.com/alexandrelamberty/resume-cmd/actions/workflows/build.yml/badge.svg)](https://github.com/alexandrelamberty/resume-cmd/actions/workflows/build.yml)
[![Tests](https://github.com/alexandrelamberty/resume-cmd/actions/workflows/tests.yml/badge.svg)](https://github.com/alexandrelamberty/resume-cmd/actions/workflows/tests.yml)

# Resume Command

Generate a resume in HTML from a YAML file and an HTML template.

## Tasks

- [ ] Single file or directory template
- [ ] Implement cmd flag and usage of Pandoc to generate the PDF

## Requirements

- [Go](https://go.dev/)

## Usage

### Running the Go file with the examples data

```bash
go run ./cmd/resume-cmd/main.go -i ./examples/content/resume-garry-lewis.yml -t ./examples/templates/london/index.gohtml -o resume_garry-lewis.html
```

> Output `-o` not implemented

### Building

```shell
go build -o ./bin/resme-cmd ./cmd/resume-cmd
```

### Installing

```bash
go install ./cmd/resume-cmd
```

Execute the system binary

```bash
resume -i ./examples/content/resume.yml -t ./examples/templates/london/tpl.gohtml -o resume_london.html
```

## Content

The content is written in YAML.

You can add `Picture` key in your YAML it will be parsed and converted into
base64 to be integrated to the HTML page.

Use a relative path to the YAML file for the `Picture` key.

See [`resume.yml`](examples/content/resume.yml)

### Structures

See the `sample` folder.

> TODO: Improve section!

### Template

The template system is based on Go Template.

See
[`sample/templates/london/index.gohtml`](sample/templates/london/index.gohtml)
for a multiple file templates examples and
[`sample/templates/tpl.gohtml`](sample/templates/tpl.gohtml) for a single file
template example.

## Credits

[Photo](https://unsplash.com/photos/2crxTr4jCkc) by
[Connor](https://unsplash.com/@wilks_and_cookies) Wilkins on
[Unsplash](https://unsplash.com)

## References

- <https://pkg.go.dev/html/template>
