FROM dpage/pgadmin4

RUN mkdir -p /var/lib/pgadmin/sessions
RUN chmod -R 777 /var/lib/