FROM golang:1.26-alpine AS build
WORKDIR /src
COPY go.mod main.go monitor.go ./
COPY demo/ ./demo/
RUN CGO_ENABLED=0 go build -o /out/monitor . && \
    CGO_ENABLED=0 go build -o /out/demo ./demo

FROM scratch
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /out/ /
USER 65532:65532
ENTRYPOINT ["/monitor"]
