default:
	@python3 -m nose --with-doctest 1424.py
	@mypy --strict 1424.py
