#!/bin/bash

# ~ dockerize OTP app. ~

make
docker build . -t ft_otp:latest
docker run \
       -it \
       --rm \
       --name ft_otp \
       ft_otp
