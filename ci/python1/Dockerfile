FROM python:3.8.0-alpine
WORKDIR /app
ARG GITHUB_SHA
ADD python1.py requirements.txt https://raw.githubusercontent.com/jrcichra/karmen/${GITHUB_SHA}/karmen-python-client/karmen/karmen.py /app/
RUN pip install -r requirements.txt
CMD python3 python1.py
