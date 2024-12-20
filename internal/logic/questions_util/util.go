package questions

import (
	"context"
	"suask/internal/dao"
	"suask/internal/model"
	"suask/internal/model/custom"
	"suask/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

type sQuestionUtil struct{}

func (sQuestionUtil) GetImages(ctx context.Context, input *model.GetImagesInput) (*model.GetImagesOutput, error) {
	idList := input.QuestionIDs
	if len(idList) == 0 {
		idList = input.AnswerIDs
	}
	md := dao.Attachments.Ctx(ctx).Where("question_id IN (?)", idList)
	var Images []*custom.Image
	err := md.Scan(&Images)
	if err != nil {
		return nil, err
	}
	imageMap := make(map[int][]int)
	for _, img := range Images {
		if _, ok := imageMap[img.QuestionId]; !ok {
			imageMap[img.QuestionId] = make([]int, 0, 8)
		}
		imageMap[img.QuestionId] = append(imageMap[img.QuestionId], img.FileID)
	}
	output := model.GetImagesOutput{
		ImageMap: imageMap,
	}
	return &output, nil
}

func (sQuestionUtil) Favorite(ctx context.Context, input *model.FavoriteInput) (*model.FavoriteOutput, error) {
	md := dao.Favorites.Ctx(ctx)
	UserId := 1
	// UserId := gconv.Int(ctx.Value(consts.CtxId))
	cnt, err := md.Where("user_id = ? AND question_id = ?", UserId, input.QuestionID).Count()
	if err != nil {
		return nil, err
	}
	if cnt > 0 {
		_, err = md.Delete("user_id = ? AND question_id = ?", UserId, input.QuestionID)
		if err != nil {
			return nil, err
		}
		return &model.FavoriteOutput{
			IsFavorited: false,
		}, nil
	} else {
		_, err = md.Insert(g.Map{
			"user_id":     UserId,
			"question_id": input.QuestionID,
			"package":     "默认收藏夹",
		})
		if err != nil {
			return nil, err
		}
		return &model.FavoriteOutput{
			IsFavorited: true,
		}, nil
	}
}

func init() {
	service.RegisterQuestionUtil(New())
}

func New() *sQuestionUtil {
	return &sQuestionUtil{}
}