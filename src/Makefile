
IMAGE?=image_placeholder
NPM?=yarn
PROJECT_PATH=$(shell cd ../;  pwd)

include ../scripts/Makefile

.PHONY:default
default:
	@echo ""
	@echo -e "\033[34mbuilding server... \033[0m"
	@rm -rf "./ui/node_modules"
	@cd  $(SCRIPT_DIR) && bash ./build.sh
	@cd $(SCRIPT_DIR) && python ./generate.py -t '$(BIN_PATH)' -i '${IMAGE}' -o '$(BIN_PATH)/docker'
	@echo -e "\033[34mbuilding web... \033[0m"
	@cd $(SOURCE_ROOT)/ui && bash ./build.sh ${NPM} ${BIN_PATH}

.PHONY:linux
linux:
	@echo ""
	@echo -e "\033[34mbuilding server... \033[0m"
	@rm -rf "./ui/node_modules"
	@cd  $(SCRIPT_DIR) && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 bash ./build.sh
	@cd $(SCRIPT_DIR) && python ./generate.py -t '$(BIN_PATH)' -i '${IMAGE}' -o '$(BIN_PATH)/docker'
	@echo -e "\033[34mbuilding web... \033[0m"
	@cd $(SOURCE_ROOT)/ui && bash ./build.sh ${NPM} ${BIN_PATH}

.PHONY:server
server:
	@echo ""
	@echo -e "\033[34mbuilding server... \033[0m"
	@rm -rf "./ui/node_modules"
	@cd  $(SCRIPT_DIR) &&  bash ./build.sh
	@cd $(SCRIPT_DIR) && python ./generate.py -t '$(BIN_PATH)' -i '${IMAGE}' -o '$(BIN_PATH)/docker'

.PHONY:ui
ui:
	@echo ""
	@echo -e "\033[34mbuilding web... \033[0m"
	@rm -rf "./ui/node_modules"
	@cd $(SOURCE_ROOT)/ui && bash ./build.sh ${NPM} ${BIN_PATH}

.PHONY:debug
debug:
	@echo ""
	@echo -e "\033[34mbuilding debug... \033[0m"
	@rm -rf "./ui/node_modules"
	@cd $(SCRIPT_DIR) && bash ./build.sh debug
	@cd $(SCRIPT_DIR) && python ./generate.py -t '$(BIN_PATH)' -i '${IMAGE}' -o '$(BIN_PATH)/docker'

.PHONY:package
package:
	rm -rf $(SOURCE_ROOT)/bin/pub/cmdb
	mkdir -p $(SOURCE_ROOT)/bin/pub && cp -R $(BIN_PATH) ./bin/pub/cmdb
	cd $(SOURCE_ROOT)/bin/pub && tar -zvcf cmdb.tar.gz cmdb
	@echo -e "\033[34mpackage success and store in $(SOURCE_ROOT)/bin/pub/cmdb.tar.gz\033[0m"

.PHONY:enterprise
enterprise:
	rm -rf $(SOURCE_ROOT)/bin/enterprise/cmdb
	mkdir -p $(SOURCE_ROOT)/bin/enterprise/cmdb/server/bin/
	mkdir -p $(SOURCE_ROOT)/bin/enterprise/cmdb/errors/
	mkdir -p $(SOURCE_ROOT)/bin/enterprise/cmdb/language/
	mkdir -p $(SOURCE_ROOT)/bin/enterprise/cmdb/web/
	mkdir -p $(SOURCE_ROOT)/bin/enterprise/cmdb/monstache/etc
	mkdir -p $(SOURCE_ROOT)/bin/enterprise/cmdb/support-files/

	cp -R ${RESOURCE_DIR}/errors/*  $(SOURCE_ROOT)/bin/enterprise/cmdb/errors
	cp -R ${RESOURCE_DIR}/language/*  $(SOURCE_ROOT)/bin/enterprise/cmdb/language
	cp -R ${BIN_PATH}/monstache/  $(SOURCE_ROOT)/bin/enterprise/cmdb/monstache
	cp -R ${BIN_PATH}/web/*  $(SOURCE_ROOT)/bin/enterprise/cmdb/web/
	cp  ${BIN_PATH}/cmdb_*/cmdb_*  $(SOURCE_ROOT)/bin/enterprise/cmdb/server/bin/
	cp  ${BIN_PATH}/tool_*/tool_*  $(SOURCE_ROOT)/bin/enterprise/cmdb/server/bin/
	cp ${DOCS_DIR}/support-file/config/projects.yaml $(SOURCE_ROOT)/bin/enterprise/cmdb/projects.yaml
	cp -R ${DOCS_DIR}/support-file/config/*  $(SOURCE_ROOT)/bin/enterprise/cmdb/support-files
	cp -R ${DOCS_DIR}/support-file/changelog/*  $(SOURCE_ROOT)/bin/enterprise/cmdb

.PHONY:clean
clean:
	rm -rf ${BIN_PATH}

.PHONY:cleanall
cleanall:
	rm -rf $(SOURCE_ROOT)/bin


release:
	rm -Rf /opt/release/cmdb
	mkdir /opt/release/cmdb
	cp -Rf ../resources/errors  /opt/release/cmdb
	cp -Rf ../resources/language  /opt/release/cmdb
	cp -Rf ./bin/build/weops/monstache /opt/release/cmdb
	cp -Rf ./projects.yaml /opt/release/cmdb
	cp -Rf ./VERSION /opt/release/cmdb
	cp ../readme.md /opt/release/cmdb
	mkdir -p /opt/release/cmdb/support-files
	cp -Rf   ../docs/support-file/config/templates /opt/release/cmdb/support-files/
	mkdir -p /opt/release/cmdb/server/bin
	cp -Rf ./bin/build/weops/cmdb_adminserver/cmdb_adminserver /opt/release/cmdb/server/bin/
	cp -Rf ./bin/build/weops/cmdb_apiserver/cmdb_apiserver  /opt/release/cmdb/server/bin/
	cp -Rf ./bin/build/weops/cmdb_authserver/cmdb_authserver  /opt/release/cmdb/server/bin/
	cp -Rf ./bin/build/weops/cmdb_cacheservice/cmdb_cacheservice  /opt/release/cmdb/server/bin/
	cp -Rf ./bin/build/weops/cmdb_cloudserver/cmdb_cloudserver  /opt/release/cmdb/server/bin/
	cp -Rf ./bin/build/weops/cmdb_coreservice/cmdb_coreservice  /opt/release/cmdb/server/bin/
	cp -Rf ./bin/build/weops/cmdb_datacollection/cmdb_datacollection  /opt/release/cmdb/server/bin/
	cp -Rf ./bin/build/weops/cmdb_eventserver/cmdb_eventserver  /opt/release/cmdb/server/bin/
	cp -Rf ./bin/build/weops/cmdb_hostserver/cmdb_hostserver  /opt/release/cmdb/server/bin/
	cp -Rf ./bin/build/weops/cmdb_operationserver/cmdb_operationserver  /opt/release/cmdb/server/bin/
	cp -Rf ./bin/build/weops/cmdb_procserver/cmdb_procserver  /opt/release/cmdb/server/bin/
	cp -Rf ./bin/build/weops/cmdb_synchronizeserver/cmdb_synchronizeserver  /opt/release/cmdb/server/bin/
	cp -Rf ./bin/build/weops/cmdb_taskserver/cmdb_taskserver  /opt/release/cmdb/server/bin/
	cp -Rf ./bin/build/weops/cmdb_toposerver/cmdb_toposerver  /opt/release/cmdb/server/bin/
	cp -Rf ./bin/build/weops/cmdb_webserver/cmdb_webserver  /opt/release/cmdb/server/bin/
	cp -Rf ./bin/build/weops/tool_ctl/tool_ctl  /opt/release/cmdb/server/bin/
	cp -Rf ./bin/build/weops/web  /opt/release/cmdb/
	cd /opt/release