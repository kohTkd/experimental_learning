.PHONY: mysql.console
mysql.console:
	docker compose exec database-auth sh -c 'mysql -u $${DB_WRITE_USER} -p$${DB_WRITE_PASSWORD} $${DB_NAME}'
