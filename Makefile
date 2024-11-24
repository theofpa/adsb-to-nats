release_adsb_to_nats:
	KO_DOCKER_REPO=192.168.1.100:9500/theofpa/adsb-to-nats ~/bin/ko apply -f deploy.yaml