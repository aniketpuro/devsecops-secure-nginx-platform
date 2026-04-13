FROM nginxinc/nginx-unprivileged:1.27-alpine-slim

LABEL org.opencontainers.image.title="devsecops-secure-nginx-platform" \
      org.opencontainers.image.description="Production-grade hardened NGINX static application" \
      org.opencontainers.image.source="https://github.com/acme-platform/devsecops-secure-nginx-platform" \
      org.opencontainers.image.licenses="Apache-2.0"

USER root

RUN mkdir -p /var/cache/nginx /var/run /tmp /usr/share/nginx/html \
    && chown -R 101:101 /var/cache/nginx /var/run /tmp /usr/share/nginx/html

COPY --chown=101:101 app/nginx.conf /etc/nginx/nginx.conf
COPY --chown=101:101 app/index.html /usr/share/nginx/html/index.html

USER 101:101

EXPOSE 8080

STOPSIGNAL SIGQUIT

HEALTHCHECK --interval=30s --timeout=3s --start-period=10s --retries=3 \
  CMD wget -q -O /dev/null http://127.0.0.1:8080/healthz || exit 1

CMD ["nginx", "-g", "daemon off;"]
