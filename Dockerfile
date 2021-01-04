FROM golang:alpine as golang   

RUN mkdir -p /root/vedioUploadPath
RUN mkdir -p /root/vedioUploadPath/output
COPY ./ /root/vedioUploadPath

WORKDIR /root/vedioUploadPath/output
RUN go build /root/vedioUploadPath/main.go


FROM alpine as alpine

RUN mkdir -p /root/vedioUploadPath
COPY --from=golang --chown=root:root /root/vedioUploadPath/output /root/vedioUploadPath

CMD /root/vedioUploadPath/main

