coverage.html: cover.out
	go tool cover -o $@ -html $<

cover.out:
	go test -coverprofile $@

.PHONY: \
	cover.out \
