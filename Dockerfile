FROM jrcichra/smartcar_python_base
ARG commit
EXPOSE 8080
COPY requirements.txt /
RUN pip install -r /requirements.txt
COPY . /
RUN echo -n $commit > /commit.txt
CMD python -u controller.py