# Resume Command

Generate a resume in HTML or PDF from a YAML file.

## Tasks

[ ] Implement cmd flag and usage of Pandoc to generate the PDF

## Usage

```bash
go run cmd/main.go -i "sample/resume.yml" -t "sample/blue-koala" -o "HTML/PDF"
```

## Conventions & Structures

See the `sample` folder.

> TODO: Improve section!


## References

- <https://pkg.go.dev/html/template>
