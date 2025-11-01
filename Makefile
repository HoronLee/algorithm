.PHONY: clean
clean:
	find . -type f \( -name "*.o" -o -name "*.out" -o -name "main" \) -exec rm -f {} +
