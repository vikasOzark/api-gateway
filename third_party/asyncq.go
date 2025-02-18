package third_party

import (
	helper "gateway/helpers"
	"gateway/helpers/constant"
	"os"

	"github.com/hibiken/asynq"
)

type Taskq struct {
	Asynq *asynq.Server
}

func (t *Taskq) AsyncServer() *asynq.Server {

	REDIS_ADDRESS := os.Getenv(constant.ENV_QUEUE_REDIS_ADDRESS)

	CRITICAL_WORKERS, _ := helper.ConvertEnvInt(constant.ENV_QUEUE_CRITICAL_WORKERS)
	DEFAULT_WORKERS, _ := helper.ConvertEnvInt(constant.ENV_QUEUE_DEFAULT_WORKERS)
	CONCURRENT_WORKERS, _ := helper.ConvertEnvInt(constant.ENV_QUEUE_CONCURRENT_WORKERS)
	LOW_WORKERS, _ := helper.ConvertEnvInt(constant.ENV_QUEUE_LOW_WORKERS)

	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: REDIS_ADDRESS},
		asynq.Config{
			Concurrency: int(CONCURRENT_WORKERS),
			Queues: map[string]int{
				"critical": CRITICAL_WORKERS,
				"default":  DEFAULT_WORKERS,
				"low":      LOW_WORKERS,
			},
		},
	)
	t.Asynq = srv
	return srv
}
