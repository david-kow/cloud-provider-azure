FROM mcr.microsoft.com/devcontainers/base:bullseye

ARG KIND_VERSION=0.20.0
ARG GO_VERSION=1.21.6

# Install Go and other dependencies
RUN wget https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz
RUN tar -xvf go${GO_VERSION}.linux-amd64.tar.gz -C /usr/local
RUN rm -rf go${GO_VERSION}.linux-amd64.tar.gz    golang
    
ENV GOPATH="/usr/local/go"
ENV PATH="${PATH}:${GOPATH}/bin"

RUN go install sigs.k8s.io/kind@v${KIND_VERSION}

RUN curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
RUN chmod +x ./kubectl
RUN mv ./kubectl /usr/local/bin

RUN curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3
RUN chmod 700 get_helm.sh
RUN ./get_helm.sh