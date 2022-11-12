FROM golang:latest

RUN mkdir /build

WORKDIR /build

RUN export GO111MODULE=ON
RUN go get https://github.com/UniversityOfGdanskProjects/plan_back/main
RUN cd /build && git clone https://github.com/UniversityOfGdanskProjects/plan_back.git

RUN cd /build/plan_back/main && go build

EXPOSE 8080

ENTRYPOINT [ "/build/plan_back/main/main" ]
