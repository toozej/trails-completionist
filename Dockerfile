# setup project and deps
FROM golang:1.21-bullseye AS init

WORKDIR /go/trails-completionist/

COPY go.mod* go.sum* ./
RUN go mod download

COPY . ./

FROM init as vet
RUN go vet ./...

# run tests
FROM init as test
RUN go test -coverprofile c.out -v ./...

# build binary
FROM init as build
ARG LDFLAGS

RUN CGO_ENABLED=0 go build -ldflags="${LDFLAGS}"

# runtime image
FROM scratch
# Copy our static executable.
COPY --from=build /go/trails-completionist/trails-completionist /go/bin/trails-completionist
# Run the binary.
ENTRYPOINT ["/go/bin/trails-completionist"]
