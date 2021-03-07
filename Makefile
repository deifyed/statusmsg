BUILD_DIR=./build
INSTALL_DIR=~/.local/bin

${BUID_DIR}/status:
	mkdir -p ${BUILD_DIR}
	go build -o ${BUILD_DIR}/status

build: ${BUID_DIR}/status

install: build
	mkdir -p ${INSTALL_DIR}
	cp ${BUILD_DIR}/status ${INSTALL_DIR}

uninstall:
	rm ${INSTALL_DIR}/status

clean:
	rm -r ${BUILD_DIR}
