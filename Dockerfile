FROM golang:1.23.2

WORKDIR /usr/src/gomall
# 设置代理
ENV GOPROXY=https://atomhub.openatom.cn

COPY app/frontend/go.mod app/frontend/go.sum ./app/frontend/
COPY rpc_gen rpc_gen
RUN cd app/frontend/ && go mod download && go mod verify
COPY app/frontend app/frontend
RUN cd app/frontend/ && go build -v -0 /opt/gomall/frontend/server
COPY app/frontend/conf /opt/gomall/frontend/conf
COPY app/frontend/static /opt/gomall/frontend/static
COPY app/frontend/template /opt/gomall/frontend/template
EXPOSE 8080

CMD ["/opt/gomall/frontend/server"]