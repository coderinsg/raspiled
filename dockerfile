FROM arm32v7/python:2.7.13-jessie
MAINTAINER engcheng_lim@dell.com
RUN apt-get update && apt-get install -y python-pip && rm -rf /var/lib/apt/lists/
RUN pip install gpiozero
COPY ./testing.py .
CMD ["python","./testing.py"]
