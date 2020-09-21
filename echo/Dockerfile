FROM alpine:3.12.0

ENV APP=echo-demo
ENV IROOT=/${APP}
ENV PATH="${IROOT}:${PATH}"

RUN mkdir -p ${IROOT}
WORKDIR ${IROOT}

ADD ${APP} ${IROOT}
ADD config.yml ${IROOT}

CMD ["sh", "-c", "$APP"]