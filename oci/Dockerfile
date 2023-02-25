ARG LITESTREAM_VERSION={{ .litestream.version }}
ARG DOCKERHUB=docker.io

ARG UPSTREAM_TAG={{ .images.upstream.tag }}

FROM ${DOCKERHUB}/litestream/litestream:${LITESTREAM_VERSION} as litestream

FROM {{ .images.upstream.name }}:${UPSTREAM_TAG}

COPY --from=litestream /usr/local/bin/litestream /usr/local/bin/litestream
COPY ./entrypoint.sh /usr/local/bin/entrypoint.sh

ARG LITESTREAM_EXEC="{{- .app.exec -}}"
ENV LITESTREAM_EXEC=${LITESTREAM_EXEC}

ENTRYPOINT ["entrypoint.sh"]
