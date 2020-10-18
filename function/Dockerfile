FROM openfaas/of-watchdog:0.8.1 as watchdog

COPY ./dist/handler ./home/app/handler

FROM alpine:3.12

RUN mkdir -p /home/app

COPY --from=watchdog /fwatchdog /usr/bin/fwatchdog
RUN chmod +x /usr/bin/fwatchdog

COPY --from=watchdog /home/app/handler /home/app/handler

RUN apk add --no-cache libc6-compat

# Add non root user
RUN addgroup -S app && adduser app -S -G app
RUN chown app /home/app

WORKDIR /home/app

USER app

ENV fprocess="./handler"
ENV mode="http"
ENV upstream_url="http://127.0.0.1:8201"
ENV write_debug="true"

HEALTHCHECK --interval=3s CMD [ -e /tmp/.lock ] || exit 1

CMD ["fwatchdog"]