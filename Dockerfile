FROM golang:1.12-alpine as builder

RUN apk add git
RUN git clone https://github.com/boodyvo/tendermint.git $GOPATH/src/github.com/boodyvo/tendermint
WORKDIR $GOPATH/src/github.com/boodyvo/tendermint
RUN go install .
RUN cp ./input.txt /input.txt

FROM alpine as final

COPY --from=builder /go/bin/tendermint /bin/tendermint
COPY --from=builder /input.txt /input.txt

CMD tendermint -n=10 -input_path=/input.txt