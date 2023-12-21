#!/bin/sh
set -e
rm -rf completions
mkdir completions
for sh in bash zsh fish; do
	go run ./cmd/trails-completionist/ completion "$sh" >"completions/trails-completionist.$sh"
done
