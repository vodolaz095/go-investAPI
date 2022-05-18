FROM centos:8

# Установить зависимости
WORKDIR /etc/yum.repos.d/
RUN sed -i 's/mirrorlist/#mirrorlist/g' /etc/yum.repos.d/CentOS-*
RUN sed -i 's|#baseurl=http://mirror.centos.org|baseurl=http://vault.centos.org|g' /etc/yum.repos.d/CentOS-*
RUN dnf clean all && dnf upgrade -y && dnf install golang wget unzip git -y && dnf clean all

RUN go version
ENV GOBIN=/usr/bin/
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1


# Загружаем protoc и встроенные протоколы
WORKDIR /tmp
RUN wget https://github.com/protocolbuffers/protobuf/releases/download/v3.20.0/protoc-3.20.0-linux-x86_64.zip

# Устанавливаем protoc
RUN unzip /tmp/protoc-3.20.0-linux-x86_64.zip bin/protoc -d /usr/
RUN chown root:root /usr/bin/protoc
RUN chmod 0755 /usr/bin/protoc

# Устанавливаем встроенные протоколы
RUN unzip protoc-3.20.0-linux-x86_64.zip include/* -d /usr/bin/
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
