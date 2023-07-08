FROM ubuntu:latest
LABEL authors="pinys"

ENTRYPOINT ["top", "-b"]