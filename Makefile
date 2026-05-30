.PHONY: database-build
database-build:
	make -C database

.PHONY: database-run
database-run:
	make -C database run

.PHONY: database-clean
database-clean:
	make -C database clean