lint/be:
	qodana scan -d ./src/go  --yaml-name .\cfg\lint\qodana-backend.yaml -o ./cfg/lint/report/back # -c
lint/fe:
	qodana scan -d ./src/js --yaml-name .\cfg\lint\qodana-frontend.yaml -o ./cfg/lint/report/front # -c
lint: lint/be lint/fe