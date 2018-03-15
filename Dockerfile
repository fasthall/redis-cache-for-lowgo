FROM redis:3.2

RUN apt update
RUN apt install python python-requests python-yaml -y
COPY docker-entrypoint.sh /usr/local/bin/
