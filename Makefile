
prometheus:
	docker run -d -p 9090:9090 -v ./prometheus.yml:/etc/prometheus/prometheus.yml prom/prometheus 

grafana:
	docker run -d -p 3000:3000 --name=grafana grafana/grafana-enterprise