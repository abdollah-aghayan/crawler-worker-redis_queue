package logic

import (
	"context"
	"crawler/domain"
	"crawler/queue"
	"crawler/utils/errorh"
	"log"
)

//CrawlerInterface interface
type CrawlerInterface interface {
	FetchUrls(ctx context.Context, req domain.Request) error
}

//UserLogic user logic
type CrawlerLogic struct {
	// Self CrawlerInterface
}

func New() CrawlerLogic {
	return CrawlerLogic{}
}

//FetchUrls get user by id
func (u *CrawlerLogic) FetchUrls(ctx context.Context, req domain.Request) *errorh.Errorh {

	// get queue
	q := queue.GetQueue()

	for _, url := range req.Urls {
		_, err := q.Push(url)

		if err != nil {
			log.Fatal(err)
			return errorh.InternalError("Something went wrong please try later")
		}
	}

	return nil
}
