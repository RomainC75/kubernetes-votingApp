package bootstrap

import redis_client "front-server/api/redis"

func Bootstrap() {
	redis_client.SetRedis()

}
