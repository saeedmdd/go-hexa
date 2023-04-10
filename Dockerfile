FROM dpage/pgadmin4

RUN mkdir -p /var/lib/pgadmin/sessions
# RUN chown -R 5050:5050 /var/lib/pgadmin/sessions