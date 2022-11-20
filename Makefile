NAME=jrpg-gang
TAG=latest
REGION=europe-central2
IMAGE=${NAME}:${TAG}
SERVICE=${NAME}
PROJECT=${NAME}

build:
	go build -o bin/${NAME} cmd/default/main.go

release:
	go build -ldflags "-s -w" -o bin/${NAME} cmd/default/main.go

image-build:
	docker build . -t ${IMAGE}

gcloud-setup:
	gcloud auth login
	gcloud auth configure-docker

gcloud-image-build:
	docker build . -t gcr.io/${PROJECT}/${IMAGE}

gcloud-image-push:
	docker push gcr.io/${PROJECT}/${IMAGE}

gcloud-image-deploy:
	gcloud run deploy ${SERVICE} --project ${PROJECT} --image gcr.io/${PROJECT}/${IMAGE} --region ${REGION}

gcloud-deploy:
	gcloud run deploy ${SERVICE} --project ${PROJECT} --source . --region ${REGION}

gcloud-regions:
	gcloud compute regions list

item:
	cd ./private/ && node ../tools/make-item-or-unit.mjs
