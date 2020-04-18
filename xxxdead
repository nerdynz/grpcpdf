FROM partlab/ubuntu-golang

RUN apt-get update -y && apt-get install -qq -y  xvfb libfontconfig wkhtmltopdf

ADD . /opt/go/src/jeevesrpcpdf
ENTRYPOINT /opt/go/bin/jeevesrpcpdf

RUN go get -d

ENV GRPC_PORT=5532
ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor
# ENV APP_USER app
# ENV APP_HOME /go/src/mathapp


# USER $APP_USER
EXPOSE 5532
CMD ["go", "run", "main.go"]