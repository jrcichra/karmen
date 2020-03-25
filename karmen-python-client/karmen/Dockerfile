FROM python:3.8.0-alpine
RUN apk add --no-cache --virtual .build-deps gcc musl-dev wireless-tools
ADD requirements.txt /
RUN pip install -r /requirements.txt
ADD karmen.py common.py /
CMD python karmen.py