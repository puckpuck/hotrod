FROM scratch
EXPOSE 8080
COPY /bin/hotrod /hotrod
ENTRYPOINT ["/hotrod"]
CMD ["all"]
