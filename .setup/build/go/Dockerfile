FROM public.ecr.aws/docker/library/golang:1.22-bullseye as BUILD

ENV PATH="/go/bin:${PATH}"
ENV WORKDIR="/go/src"

WORKDIR $WORKDIR

COPY . $WORKDIR

RUN apt-get update && \
  go install github.com/githubnemo/CompileDaemon@latest && \
  go mod download

ENTRYPOINT ["CompileDaemon", "-directory=./", "-build=go build -o clean_arch /go/src/cmd/ordersystem/main.go /go/src/cmd/ordersystem/wire_gen.go", "-command=./clean_arch"]