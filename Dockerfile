FROM golang:1.8
MAINTAINER Xavier Bruhiere

# since we will work inside the container from command line
ENV TERM xterm
# needed for ghr to work (+ namespacing of build files)
RUN git config --global user.name hackliff && mkdir /tmp/_build

COPY ./scripts /tmp/_build/scripts
COPY Makefile /tmp/_build/Makefile
RUN cd /tmp/_build && make install-tools

# godoc standard port
EXPOSE 6060

# whatever
CMD ["go"]
