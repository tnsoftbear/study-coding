package redis_storage

import (
	"fmt"
	"strconv"
	"task_manager/internal/types"

	"github.com/go-redis/redis"
)

type RedisManager struct {
	rdb *redis.Client
}

func NewRedisManager() *RedisManager {
	return &RedisManager{
		rdb: redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		}),
	}
}

func (rm *RedisManager) SaveTask(task types.Task) error {
	taskId := makeTaskId(task.Id)
	hset := rm.rdb.HSet(taskId, "Id", task.Id)
	if hset.Err() != nil {
		return hset.Err()
	}

	hset = rm.rdb.HSet(taskId, "Name", task.Name)
	if hset.Err() != nil {
		return hset.Err()
	}

	hset = rm.rdb.HSet(taskId, "Description", task.Description)
	if hset.Err() != nil {
		return hset.Err()
	}

	hset = rm.rdb.HSet(taskId, "CreatedAt", task.CreatedAt)
	if hset.Err() != nil {
		return hset.Err()
	}

	z := redis.Z{Score: float64(task.CreatedAt), Member: task.Id}
	zadd := rm.rdb.ZAdd("tasks", z)
	if zadd.Err() != nil {
		return zadd.Err()
	}

	return nil
}

func (rm *RedisManager) LoadTask(id string) (*types.Task, error) {
	hgetAll := rm.rdb.HGetAll(makeTaskId(id))
	if err := hgetAll.Err(); err != nil {
		return nil, err
	}

	ires, err := hgetAll.Result()
	if err != nil {
		return nil, err
	}

	if l := len(ires); l == 0 {
		return nil, nil
	}

	created_at, _ := strconv.ParseInt(ires["CreatedAt"], 10, 64)
	task := types.Task{
		Id:          ires["Id"],
		Name:        ires["Name"],
		Description: ires["Description"],
		CreatedAt:   created_at,
	}
	return &task, nil
}

func (rm *RedisManager) DeleteTask(id string) error {
	if err := rm.rdb.Unlink(makeTaskId(id)).Err(); err != nil {
		return err
	}

	if err := rm.rdb.ZRem("tasks", id).Err(); err != nil {
		return err
	}

	return nil
}

func (rm *RedisManager) LoadTasks() ([]*types.Task, error) {
	var tasks []*types.Task = make([]*types.Task, 0)
	zRange := rm.rdb.ZRange("tasks", 0, -1)
	if err := zRange.Err(); err != nil {
		return nil, err
	}

	ids, err := zRange.Result()
	if err != nil {
		return nil, err
	}

	for _, id := range ids {
		if task, err := rm.LoadTask(id); err != nil {
			return nil, err
		} else {
			tasks = append(tasks, task)
		}
	}
	return tasks, nil
}

func makeTaskId(id string) string {
	return fmt.Sprintf("task:%s", id)
}