.PHONY: database-build
database-build:
	make -C database

.PHONY: database-run
database-run:
	make -C database run

.PHONY: database-clean
database-clean:
	make -C database clean

.PHONY: backend-build
backend-build:
	make -C database

.PHONY: backend-run
backend-run:
	make -C backend run

.PHONY: backend-clean
backend-clean:
	make -C backend clean