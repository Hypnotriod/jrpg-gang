NAME=jrpg-gang
TAG=latest
REGION=europe-central2
IMAGE=${NAME}:${TAG}
SERVICE=${NAME}
PROJECT=${NAME}
REPOSITORY=${NAME}
ARTIFACT_REGISTRY_IMAGE=${REGION}-docker.pkg.dev/${PROJECT}/${REPOSITORY}/${IMAGE}

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
	docker build . -t ${ARTIFACT_REGISTRY_IMAGE}

gcloud-image-push:
	docker push ${ARTIFACT_REGISTRY_IMAGE}

gcloud-image-deploy:
	gcloud run deploy ${SERVICE} --project ${PROJECT} --image ${ARTIFACT_REGISTRY_IMAGE} --platform managed --region ${REGION}

gcloud-deploy:
	gcloud run deploy ${SERVICE} --project ${PROJECT} --source . --region ${REGION}

gcloud-regions:
	gcloud compute regions list

item:
	cd ./private/ && node ../tools/make-item-or-unit.mjs

certificate:
	openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout key.pem -out cert.pem
