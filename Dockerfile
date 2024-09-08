FROM golang:1.22.6

# устанавливаем unzip
RUN apt-get update && apt-get install unzip -y

# Задаём версии инструментов
ARG PROTOC_VERSION="25.2"
ARG PROTOC_GEN_GO_VERSION="1.32.0"
ARG PROTOC_GEN_GO_GPRC_VERSION="1.3.0"

# устанавливаем плагины для protoc
ENV GOBIN=/usr/bin/
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v${PROTOC_GEN_GO_VERSION}
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v${PROTOC_GEN_GO_GPRC_VERSION}

# Загружаем protoc и встроенные протоколы
WORKDIR /tmp
RUN wget "https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-x86_64.zip"

# Устанавливаем protoc
RUN unzip /tmp/protoc-${PROTOC_VERSION}-linux-x86_64.zip bin/protoc -d /usr/
RUN chown root:root /usr/bin/protoc
RUN chmod 0755 /usr/bin/protoc

# Устанавливаем встроенные протоколы
RUN unzip protoc-${PROTOC_VERSION}-linux-x86_64.zip include/* -d /usr/bin/
RUN chown root:root /usr/bin/include/google/ -R

# Создаём директории для вывода кода
RUN mkdir -p /opt/client

# Добавляем файлы
ADD investapi/client.go /opt/client
ADD investapi/helpers.go /opt/client
ADD investapi/helpers_test.go /opt/client

# Добавляем скрипт точки входа
ADD entrypoint.sh /usr/bin/entrypoint.sh
RUN chmod a+x /usr/bin/entrypoint.sh

# При каждом запуске образа, будет выполнен скрипт entrypoint.sh
ENTRYPOINT /usr/bin/entrypoint.sh
