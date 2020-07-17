# Copyright 2018 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
FROM gcr.io/gcp-runtimes/go1-builder:1.11 as builder
ENV GO111MODULE=auto
WORKDIR /go/src/github.com/duniapay/api/
RUN ls -al
COPY ./ ./
RUN ls -al
COPY go.mod .
COPY go.sum .

RUN apt-get update -y
RUN apt-get install git -y
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 /usr/local/go/bin/go build -o app api-service/main.go
RUN ls -al

FROM alpine
LABEL author="Fiifi Baidoo"
WORKDIR /go/src/github.com/duniapay/api/api-service
RUN ls -al
COPY --from=builder /go/src/github.com/duniapay/api/app /go/src/github.com/duniapay/api/api-service
COPY --from=builder /go/src/github.com/duniapay/api/app /go/src/github.com/duniapay/api/api-service/
RUN ls -al

EXPOSE 80
RUN apk add --no-cache ca-certificates
ENTRYPOINT ["./app"]
