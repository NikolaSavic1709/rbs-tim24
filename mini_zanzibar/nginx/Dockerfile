FROM nginx:latest

RUN apt-get -y update \
    && apt-get install openssl

RUN mkdir /etc/nginx/ssl

RUN openssl req -x509 -newkey rsa:4096 -sha256 -days 365 -nodes \
        -keyout /etc/nginx/ssl/example.key \
        -out /etc/nginx/ssl/example.crt \
        -subj "/C=RS/ST=Serbia/L=NoviSad/O=42/OU=42/CN=example/UID=example"

COPY ./nginx.conf /etc/nginx/conf.d/default.conf
COPY ./ngconf /etc/nginx/nginx.conf
COPY ./resources /var/www/resources

COPY ./encrypt_logs.sh /usr/local/bin/
RUN chmod +x /usr/local/bin/encrypt_logs.sh
RUN mkdir -p /tmp/nginx_logs
RUN mkdir -p /var/log/nginx
RUN mkdir -p /var/log/nginx/enc
RUN mkfifo /tmp/nginx_logs/access_pipe
RUN mkfifo /var/log/nginx/error_pipe
RUN touch /var/log/nginx/enc/access.log
RUN touch /var/log/nginx/enc/error.log
CMD ["/usr/local/bin/encrypt_logs.sh"]

EXPOSE 80
EXPOSE 443
