FROM busybox
LABEL authors="boldyrevn"
WORKDIR /go_webapp
COPY ./main .
COPY ./.env .
CMD ./main