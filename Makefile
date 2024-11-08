run:
	cd migrations && go run migrations.go droptables
	cd migrations && go run migrations.go migrate
	cd migrations && go run migrations.go seed

	@echo "database migrated sucessfully"
