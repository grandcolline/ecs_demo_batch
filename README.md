# ecs_demo_batch

ECSの勉強会で使う簡単なアプリケーション(BATCH)

## Docker Image

[grandcolline/ecs_demo_batch](https://hub.docker.com/r/grandcolline/ecs_demo_batch/)

## RUN

```
$ docker run --rm \
	-e "NAME=ECS_DEMO_BATCH" \
	-e "WEBHOOK_URL=https://hooks.slack.com/services/XXXXX/XXXXX/XXXXX" \
	-e "CHANNEL=#channel_name" \
	-it grandcolline/ecs_demo_batch
```

