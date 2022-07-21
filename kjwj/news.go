package kjwj

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-redis/redis/v9"
	"github.com/joewongex/tasks/g"
)

func CheckLatestNews(ctx context.Context, wg sync.WaitGroup) {
	wg.Add(1)
	g.Log.Info("【科技玩家】获取热讯开始执行")
	err := latestNews(ctx)
	if err != nil {
		g.Log.Errorf("【科技玩家】获取最新热讯出错：%v", err)
	}
	g.Log.Info("【科技玩家】获取热讯执行结束")
	wg.Done()
}

func latestNews(ctx context.Context) (err error) {
	const url string = "https://www.kejiwanjia.com/newsflashes"
	resp, err := g.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("【科技玩家】请求%s，返回状态码：%d", url, resp.StatusCode)
		return
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return
	}

	newsItem := doc.Find(".news-list-box .news-item:nth-child(1) ul li:nth-child(1)")
	if newsItem.Length() == 0 {
		err = fmt.Errorf("【科技玩家】打不到热讯信息")
		return
	}

	date := doc.Find(".news-list-box .news-item:nth-child(1) .news-item-date").Text()
	time := newsItem.Find(".news-item-header span:nth-child(1)").Text()
	title := newsItem.Find(".news-item-content .anhover").Text()
	content := newsItem.Find(".news-item-content .b2-hover").Text()

	if strings.TrimSpace(date) == "" || strings.TrimSpace(time) == "" {
		err = fmt.Errorf("【科技玩家】找不到热讯的发布时间")
		return
	}

	const key = "tasks:kjwj:latest_news"
	createdAt, err := g.Redis.HGet(ctx, key, "created_at").Result()
	if err != nil && err != redis.Nil || createdAt == date+" "+time {
		return
	}

	err = g.Redis.HSet(ctx, "tasks:kjwj:latest_news", map[string]interface{}{
		"created_at": date + " " + time,
		"title":      title,
		"content":    content,
	}).Err()
	if err != nil {
		return
	}

	msg := fmt.Sprintf("有新的热讯：\n【标题】%s\n【内容】%s\n【时间】%s", title, content, date+" "+time)
	return g.BarkNotify("科技玩家", msg)
}
