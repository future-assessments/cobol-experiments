.PHONY: build test run csv

build:
	cargo build --release

test:
	cargo test

run: build
	cargo run --release -- --source ../mortgagees.txt

csv: build
	cargo run --release -- --source ../mortgagees.txt --target ./output/mortgagees.csv
