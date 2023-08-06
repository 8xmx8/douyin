package db

import (
	"context"
	"strconv"
	"sync"
	"time"

	"github.com/Ocyss/douyin/internal/model"
	"github.com/redis/go-redis/v9"
)

var (
	playCountUpdates      = make(chan string, 35)
	videoFavoriteCountKey = make([]byte, 0, 50)
	videoCommentCountKey  = make([]byte, 0, 50)
	videoPlayCountKey     = make([]byte, 0, 50)
	userFollowCountKey    = make([]byte, 0, 50)
	userFollowerCountKey  = make([]byte, 0, 50)
)

func init() {
	videoFavoriteCountKey = append(videoFavoriteCountKey, "video:favorite_count/"...)
	videoCommentCountKey = append(videoCommentCountKey, "video:comment_count/"...)
	videoPlayCountKey = append(videoPlayCountKey, "video:play_count/"...)
	userFollowCountKey = append(userFollowCountKey, "user:follow_count/"...)
	userFollowerCountKey = append(userFollowerCountKey, "user:follower_count/"...)
}

func getKey(id int64, prefix []byte) string {
	s := make([]byte, 0, 50)
	copy(s, prefix)
	s = append(s, strconv.FormatInt(id, 16)...)
	return string(s)
}

func getVideoFavoriteCountKey(id int64) string {
	return getKey(id, videoFavoriteCountKey)
}

func getVideoCommentCountKey(id int64) string {
	return getKey(id, videoCommentCountKey)
}

func getVideoPlayCountKey(id int64) string {
	return getKey(id, videoPlayCountKey)
}

func getUserFollowCountKey(id int64) string {
	return getKey(id, userFollowCountKey)
}

func getUserFollowerCountKey(id int64) string {
	return getKey(id, userFollowerCountKey)
}

// getFavoriteCount 视频的点赞总数
func getFavoriteCount(wg *sync.WaitGroup, vid int64, val *int64) {
	defer wg.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	key := getVideoFavoriteCountKey(vid)
	favoriteCount, err := rdb.Get(ctx, key).Int64()
	if err == redis.Nil {
		db.Table("user_favorite").Where("video_id = ?", vid).Count(&favoriteCount)
		_ = rdb.Set(ctx, key, favoriteCount, 300*time.Second)
	}
	*val = favoriteCount
}

// getCommentCount 视频的评论总数
func getCommentCount(wg *sync.WaitGroup, vid int64, val *int64) {
	defer wg.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	key := getVideoCommentCountKey(vid)
	CommentCount, err := rdb.Get(ctx, key).Int64()
	if err == redis.Nil {
		db.Model(&model.Comment{}).Where("video_id = ?", vid).Count(&CommentCount)
		_ = rdb.Set(ctx, key, CommentCount, 300*time.Second)
	}
	*val = CommentCount
}

// setPlayCount 视频的播放量
func setPlayCount(wg *sync.WaitGroup, ip string, vid int64, val *int64) {
	defer wg.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	key := getVideoPlayCountKey(vid)
	rdb.PFAdd(ctx, key, ip)
	*val, _ = rdb.PFCount(ctx, key).Result()
	go func() {
		playCountUpdates <- key
		if len(playCountUpdates) >= 30 {
			for i := 0; i < 30; i++ {
				v := <-playCountUpdates
				val, _ := rdb.PFCount(ctx, v).Result()
				db.Model(&model.Video{}).Where("id = ?", v[17:]).Update("play_count", val)
			}
		}
	}()
}

// getIsFavorite 视频是否点赞
func getIsFavorite(wg *sync.WaitGroup, uid, vid int64, val *bool) {
	defer wg.Done()
	result := map[string]any{}
	*val = db.Table("user_favorite").Where("user_id = ? AND video_id = ?", uid, vid).Take(&result).RowsAffected == 1
	// data[i].IsFavorite = db.Raw("SELECT * FROM user_favorite WHERE user_id = ? AND video_id = ?", uid, data[i].ID).Scan(&result).RowsAffected == 1
}

func getFollowCount(wg *sync.WaitGroup, uid int64, val *int64) {
	defer wg.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	key := getUserFollowCountKey(uid)
	FollowCount, err := rdb.Get(ctx, key).Int64()
	if err == redis.Nil {
		db.Table("user_follow").Where("user_id = ?", uid).Count(&FollowCount)
		_ = rdb.Set(ctx, key, FollowCount, 300*time.Second)
	}
	*val = FollowCount
}

func getFollowerCount(wg *sync.WaitGroup, uid int64, val *int64) {
	defer wg.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	key := getUserFollowerCountKey(uid)
	FollowerCount, err := rdb.Get(ctx, key).Int64()
	if err == redis.Nil {
		db.Table("user_follower").Where("user_id = ?", uid).Count(&FollowerCount)
		_ = rdb.Set(ctx, key, FollowerCount, 300*time.Second)
	}
	*val = FollowerCount
}
