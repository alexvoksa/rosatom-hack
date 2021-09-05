FROM python:3.9-buster
RUN pip3 install cython SQLAlchemy==1.3.22 jsonschema==3.2.0 meilisearch psycopg2-binary==2.8.6
COPY update_index.sh ./
RUN chmod +x update_index.sh
CMD ./update_index.sh