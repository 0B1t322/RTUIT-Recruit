rebuild_all_docker:
	cd service.factory && make build_docker_image && make push_docker_image
	cd service.shops && make build_docker_image && make push_docker_image
	cd service.purchases && make build_docker_image && make push_docker_image