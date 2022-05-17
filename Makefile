target=.

syllables:
	go build -o syllables main.go

.PHONY: clean, install, uninstall

clean:
	rm -f bin syllables

install:
	mkdir -p ${target}
	chmod +x syllables
	cp syllables ${target}/
	cp static ${target}/
	cp templates ${target}/

uninstall:
	rm -rf ${target}
