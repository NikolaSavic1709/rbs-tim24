FROM nginx:latest

RUN apt-get -y update \
    && apt-get install openssl

RUN mkdir /etc/nginx/ssl

RUN openssl req -x509 -newkey rsa:4096 -sha256 -days 365 -nodes \
        -keyout /etc/nginx/ssl/example.key \
        -out /etc/nginx/ssl/example.crt \
        -subj "/C=RS/ST=Serbia/L=NoviSad/O=42/OU=42/CN=example/UID=example"

COPY ./nginx.conf /etc/nginx/conf.d/default.conf

COPY ./resources /var/www/resources

EXPOSE 80
EXPOSE 443
