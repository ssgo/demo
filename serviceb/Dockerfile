FROM alpine
ADD server /opt/
ADD env.json /opt/
ENTRYPOINT /opt/server
HEALTHCHECK --interval=10s --timeout=3s CMD /opt/server check
