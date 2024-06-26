#!/bin/sh

# Run nginx in the background
nginx &

while true; do
    # Read from the access log named pipe
    if read access_line < /tmp/nginx_logs/access_pipe; then
        echo "$access_line" | openssl enc -aes-256-cbc -salt -pbkdf2 -pass pass:$ENCRYPTION_PASSWORD >> /var/log/nginx/enc/access.log
    fi

    # Read from the error log named pipe
    if read error_line < /var/log/nginx/error.log; then
        echo "$error_line" | openssl enc -aes-256-cbc -salt -pbkdf2 -pass pass:$ENCRYPTION_PASSWORD >> /var/log/nginx/enc/error.log
    fi
done
