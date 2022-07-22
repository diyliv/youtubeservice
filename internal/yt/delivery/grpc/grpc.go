package grpc

import (
	"context"

	ytservicepb "github.com/diyliv/youtubeservice/proto/ytservice"
	"go.uber.org/zap"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

type ytService struct {
	logger *zap.Logger
	YToken string
}

func NewYTService(logger *zap.Logger, YToken string) *ytService {
	return &ytService{logger: logger, YToken: YToken}
}

func (s *ytService) SearchVideo(ctx context.Context, req *ytservicepb.SearchVideoReq) (*ytservicepb.SearchVideoResp, error) {

	videoName := req.GetVideoName()
	if videoName == "" {
		return &ytservicepb.SearchVideoResp{Resp: map[string]string{"VideoName": "is empty"}}, nil
	}

	yt, err := youtube.NewService(context.Background(), option.WithAPIKey(s.YToken))
	if err != nil {
		s.logger.Error("Error while creating new YT service: " + err.Error())
	}

	search := yt.Search

	list := search.List([]string{"id, snippet"}).Q(videoName)

	resp, err := list.Do()
	if err != nil {
		s.logger.Error("Error while making request: " + err.Error())
	}

	videos := make(map[string]string)
	for _, item := range resp.Items {
		switch item.Id.Kind {
		case "youtube#video":
			videos[item.Id.VideoId] = item.Snippet.Title
		}
	}

	return &ytservicepb.SearchVideoResp{Resp: videos}, nil
}
