# Resume Command

Generate a resume in HTML or PDF from a YAML file.

## Tasks

- [ ] Single file or directory template
- [ ] Implement cmd flag and usage of Pandoc to generate the PDF

## Usage

Running Go file

```bash
go run cmd/main.go -i sample/content/resume.yml -t
sample/templates/london/tpl.gohtml -o resume_london.html
```

Installing

```bash
go install -o ~/.local/share/go/bin/resume
```

Executing

```bash
resume -i sample/content/resume.yml -t sample/templates/london/tpl.gohtml -o
resume_london.html
```

## Conventions & Structures

See the `sample` folder.

> TODO: Improve section!

## Content

The content is written in YAML.

You can add `Picture` key in your YAML it will be parsed and converted into
base64 to be integrated to the HTML page.

Use a relative path to the YAML file for the `Picture` key.

See [`resume.yml`](sample/content/resume.yml)

## Template

The templating system is based on Go Template.

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
