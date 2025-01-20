.PHONY: init clean

init:
	@for i in {0..6}; do \
		dir="ex0$$i"; \
		if [ -d "$$dir" ]; then \
			cd "$$dir" && \
			rm -f go.mod && \
			go mod init "ex0$$i" && \
			cd ..; \
		fi \
	done

clean:
	@for i in {0..6}; do \
		dir="ex0$$i"; \
		if [ -d "$$dir" ]; then \
			rm -f "$$dir/go.mod"; \
		fi \
	done
