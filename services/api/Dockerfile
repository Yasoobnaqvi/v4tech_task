FROM golang:1.20 as development

RUN apt update && apt upgrade -y &&\
    apt install -y git \
    make openssh-client

WORKDIR /go/src/app

COPY . ./

RUN go install github.com/cosmtrek/air@latest

RUN go mod tidy \
    && go mod verify
    

CMD ["air", "-c", ".air.toml"]

# RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
#     && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air

# RUN go install github.com/go-delve/delve/cmd/dlv@latest

ENV environment test
ENV PORT 3000

ENTRYPOINT ["air"]
