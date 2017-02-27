package util

import (
	"errors"
	"time"

	"github.com/spf13/viper"
)

/**
队列方式的频率控制, 精准的控制指定时间内允许访问的次数
适用场景是业务相关的频率控制,比如卖座卡使用频率控制,领奖频率控制等等
访问限制次数过大时, 比如接口频率控制, ip频率控制,不建议使用此方式, 建议使用令牌桶方式进行频率控制
*/
func RateLimit(businessName, uniqueKey string, count, second int64, needAdd bool) (bool, error) {
	key := viper.GetString("rateLimit.key") + "_" + businessName + "_" + uniqueKey
	if key == "" {
		return false, errors.New("取不到reids中ratelimit")
	}
	//获取当前队列长度
	len := Redis.LLen(key).Val()
	//如果请求次数已经超过当前最大的请求次数
	if len >= count {
		firstTime, err := Redis.RPop(key).Int64()
		if err != nil {
			return false, errors.New("取不到reids中ratelimitTime")
		}
		//判断当前时间是否已经超过限制时间,恢复次数
		if time.Now().Unix() > firstTime+second {
			//队列push 将当前请求时间放入队列中
			if needAdd {
				Redis.LPush(key, time.Now().Unix())
				//设置过期时间
				Redis.Expire(key, time.Duration(second*int64(time.Second)))
			}
			return true, nil
		} else {
			Redis.RPush(key, firstTime)
			return false, nil
		}

	}

	//队列push 将当前请求时间放入队列中
	if needAdd {
		Redis.LPush(key, time.Now().Unix())
		Redis.Expire(key, time.Duration(second*int64(time.Second)))
	}

	return true, nil
}
