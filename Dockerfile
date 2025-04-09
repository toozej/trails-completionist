# setup project and deps
FROM golang:1.24-bookworm AS init

WORKDIR /go/trails-completionist/

COPY go.mod* go.sum* ./
RUN go mod download

COPY . ./

FROM init AS vet
RUN go vet ./...

# run tests
FROM init AS test
RUN go test -coverprofile c.out -v ./... && \
	echo "Statements missing coverage" && \
	grep -v -e " 1$" c.out

# build binary
FROM init AS build
ARG LDFLAGS

RUN CGO_ENABLED=0 go build -ldflags="${LDFLAGS}"

# runtime image
FROM scratch
# Copy our static executable.
COPY --from=build /go/trails-completionist/trails-completionist /go/bin/trails-completionist
# Run the binary.
ENTRYPOINT ["/go/bin/trails-completionist"]
CMD ["full", "--help"]
