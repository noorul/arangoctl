FROM alpine:3.16.1

COPY arangoctl /usr/local/bin/arangoctl
ENTRYPOINT ["arangoctl"]
