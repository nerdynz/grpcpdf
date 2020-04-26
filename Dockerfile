FROM golang:latest

RUN apt-get update && apt-get upgrade -y --no-install-recommends \
	&& apt-get install -y ca-certificates wget xz-utils libssl-dev libxrender1 libxt6 libxtst6 fontconfig \
	&& wget https://downloads.wkhtmltopdf.org/0.12/0.12.5/wkhtmltox_0.12.5-1.buster_amd64.deb \
	&& apt install -y ./wkhtmltox_0.12.5-1.buster_amd64.deb \
	&& apt-get purge -y --auto-remove wget xz-utils \
	&& rm -rf /var/lib/apt/lists/*

RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]
