create-docker-image:
	docker image build --target prod_image -t andretadeu/go-test-app:$(version) .