package docker

import (
	"github.com/urfave/cli"
	"github.com/wenj91/mctl/go-zero/tools/goctl/util"
)

const (
	category           = "docker"
	dockerTemplateFile = "docker.tpl"
	dockerTemplate     = `FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux
{{if .Chinese}}ENV GOPROXY https://goproxy.cn,direct
{{end}}
WORKDIR /build/zero

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
{{if .Argument}}COPY {{.GoRelPath}}/etc /app/etc
{{end}}RUN go build -ldflags="-s -w" -o /app/{{.ExeFile}} {{.GoRelPath}}/{{.GoFile}}


FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/{{.ExeFile}} /app/{{.ExeFile}}{{if .Argument}}
COPY --from=builder /app/etc /app/etc{{end}}
{{if .HasPort}}
EXPOSE {{.Port}}
{{end}}
CMD ["./{{.ExeFile}}"{{.Argument}}]
`
)

func Clean() error {
	return util.Clean(category)
}

func GenTemplates(_ *cli.Context) error {
	return initTemplate()
}

func Category() string {
	return category
}

func RevertTemplate(name string) error {
	return util.CreateTemplate(category, name, dockerTemplate)
}

func Update() error {
	err := Clean()
	if err != nil {
		return err
	}

	return initTemplate()
}

func initTemplate() error {
	return util.InitTemplates(category, map[string]string{
		dockerTemplateFile: dockerTemplate,
	})
}
