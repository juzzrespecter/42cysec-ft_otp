FROM debian:bullseye-slim

WORKDIR /app

COPY ./ft_otp /app/

# ~  install oathtool for testing purposes ~ #
RUN apt-get update && \
    apt-get install ca-certificates oathtool python3 -y

RUN /bin/sh
