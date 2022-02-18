FROM ubuntu:latest as builder

RUN apt-get update && apt-get install -y \
    git

WORKDIR /tmp

# Clone zokrates-pycrypto
RUN git clone --depth 1 https://github.com/Zokrates/pycrypto.git

# Create production image
FROM python:3.8-slim-buster

WORKDIR /usr/src/app

COPY requirements.txt requirements.txt
RUN pip3 install -r requirements.txt

# Copy the zokrates-pycrypto package
COPY --from=builder /tmp/pycrypto/zokrates_pycrypto ./zokrates_pycrypto

COPY ./client.py .

ENV FLASK_APP=client.py
ENV FLASK_RUN_PORT=3000
ENV FLASK_RUN_HOST=0.0.0.0
CMD [ "python3", "-m" , "flask", "run"]