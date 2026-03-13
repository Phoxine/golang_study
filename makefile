topic=flow_control

new:
	mkdir $(topic)
	touch $(topic)/main.go
	cd $(topic) && go mod init $(topic)