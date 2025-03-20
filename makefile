init:
	hugo mod get && hugo mod tidy && git clone https://github.com/CaiJimmy/hugo-theme-stack.git themes/hugo-theme-stack 
	cd themes/hugo-theme-stack && git checkout v3.30.0 && cd ../..

run_main:
	go run main.go

build: run_main
	rm -rf resources
	hugo --config ./config.yaml

server: build
	hugo server --bind=0.0.0.0 --port=1313 --disableFastRender