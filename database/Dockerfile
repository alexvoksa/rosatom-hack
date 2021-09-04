FROM postgres:13.1
HEALTHCHECK --timeout=45s --interval=5m --retries=10 \
    CMD pg_isready -q -d postgres -U postgres
ENV POSTGRES_PASSWORD="postgres"
ENV PGDATA="/var/lib/postgresql/data/"
ENV PGUSER="postgres"
ENV PGPASSWORD="postgres"
ENV PGDATABASE="postgres"
ENV DB_USER="hackathon"
ENV DB_PASS="hackathon"
ENV DB_NAME="hackathon"
COPY ./init/01-init.sh /docker-entrypoint-initdb.d/
COPY constant_files/ /constant_files
EXPOSE "5432"