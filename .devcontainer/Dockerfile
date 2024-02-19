FROM mcr.microsoft.com/devcontainers/base:bullseye

# Install Go and other dependencies
RUN wget https://golang.org/dl/go1.21.6.linux-amd64.tar.gz
RUN tar -xvf go1.21.6.linux-amd64.tar.gz -C /usr/local
RUN rm -rf go1.21.6.linux-amd64.tar.gz    golang
    
ENV GOPATH="/usr/local/go"
ENV PATH="${PATH}:${GOPATH}/bin"

RUN go install sigs.k8s.io/kind@v0.20.0

RUN curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
RUN chmod +x ./kubectl
RUN mv ./kubectl /usr/local/bin