default:
	@python3 -m nose --with-doctest 1-4-24.py
	@mypy --strict 1-4-24.py
