.PHONY: build test run

build:
	cargo build --release

test:
	cargo test

run: build
	cargo run --release -- --source ../mortgagees.txt
