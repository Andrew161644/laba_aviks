FROM python:3.8 AS builder
COPY requirements.txt .

RUN pip install --user -r requirements.txt

FROM python:3.8-slim
WORKDIR /code

COPY --from=builder /root/.local /root/.local
COPY ./src .

#RUN apt update
#RUN apt install -y git
#RUN pip install git+https://github.com/MicroPyramid/forex-python.git

ENV PATH=/root/.local:$PATH

CMD [ "python", "./server.py" ]