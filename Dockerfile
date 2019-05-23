FROM scratch
EXPOSE 8080
ENTRYPOINT ["/devpod-demo"]
COPY ./bin/ /