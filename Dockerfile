FROM scratch
ADD ./lengthener /lengthener
ENTRYPOINT ["/lengthener"]
