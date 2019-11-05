FROM alphine

ADD ./search-service /app/search-service

WORKDIR /app

CMD ["/app/search-service"]