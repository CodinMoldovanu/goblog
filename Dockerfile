FROM golang:latest
RUN export GO111MODULE=on
RUN git clone https://github.com/CodinMoldovanu/goblog.git
RUN pwd
WORKDIR /go/goblog/
RUN ls
RUN GOOS=linux GOARCH=arm GOARM=5 go build

EXPOSE 1333/tcp

# FROM nginx:latest 
# COPY --from=0 /go/src/github.com/codinmoldovanu/goblog/goblog .
CMD ["./goblog"]  
