package public

import (
	"context"
	"suask/internal/consts"
	"suask/internal/dao"
	"suask/internal/model"
	"suask/internal/model/custom"
	"suask/internal/model/do"
	"suask/internal/service"
	"suask/utility"

	"github.com/gogf/gf/v2/util/gconv"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/text/gstr"

	"github.com/gogf/gf/v2/frame/g"
)

type sPublicQuestion struct{}

// var keywordCacheMode = gdb.CacheOption{
// 	Duration: time.Minute * 5,
// 	Name:     "public_keywords_cache",
// 	Force:    false,
// }

// var keywordClearCache = gdb.CacheOption{
// 	Duration: -1,
// 	Name:     "public_keywords_cache",
// 	Force:    false,
// }

func (sPublicQuestion) GetBase(ctx context.Context, input *model.GetBaseInput) (*model.GetBaseOutput, error) {
	md := dao.Questions.Ctx(ctx).WhereNull("dst_user_id")
	if input.Keyword != "" {
		md = md.Where("title LIKE?", "%"+input.Keyword+"%")
	}
	md = md.Page(input.Page, consts.NumOfQuestionsPerPage)
	err := utility.SortByType(&md, input.SortType)
	if err != nil {
		return nil, err
	}
	var q []*custom.Questions
	var remain int
	err = md.ScanAndCount(&q, &remain, true) // 先查不包含favorites的结果
	if err != nil {
		return nil, err
	}
	// 计算剩余页数
	remainNum := remain - consts.NumOfQuestionsPerPage*input.Page
	remain = remainNum / consts.NumOfQuestionsPerPage
	if remainNum%consts.NumOfQuestionsPerPage > 0 {
		remain += 1
	}

	qIDs := make([]int, len(q))                 // 获取问题ID列表 （官方的静态联表也是这么做的）
	pqs := make([]model.PublicQuestion, len(q)) // 用于存放最终结果
	idMap := make(map[int]int)                  // 用于快速查找问题ID对应的索引
	for i, pq := range q {
		idMap[pq.Id] = i
		qIDs[i] = pq.Id
		pqs[i] = model.PublicQuestion{
			ID:        pq.Id,
			Title:     pq.Title,
			Content:   utility.TruncateString(pq.Contents),
			CreatedAt: pq.CreatedAt.TimestampMilli(),
			Views:     pq.Views,
			AnswerNum: pq.ReplyCnt,
		}
	}

	var fav []*custom.MyFavorites
	UserId := 1
	// UserId := gconv.Int(ctx.Value(consts.CtxId))
	md = dao.Favorites.Ctx(ctx).Where("question_id IN (?) AND user_id = ?", qIDs, UserId)
	err = md.Scan(&fav) // 再查favorites
	if err != nil {
		return nil, err
	}

	for _, f := range fav { // 填充IsFavorited字段
		pqs[idMap[f.QuestionId]].IsFavorite = true
	}

	output := model.GetBaseOutput{
		QuestionIDs: qIDs,
		IdMap:       idMap,
		Questions:   pqs,
		RemainPage:  remain,
	}
	return &output, nil
}

func (sPublicQuestion) GetKeyword(ctx context.Context, input *model.GetKeywordsInput) (*model.GetKeywordsOutput, error) {
	// md := dao.Questions.Ctx(ctx).Cache(keywordCacheMode).WhereNull("dst_user_id")
	md := dao.Questions.Ctx(ctx).WhereNull(dao.Questions.Columns().DstUserId)
	// fmt.Println(input.Keyword)
	err := utility.SortByType(&md, input.SortType)
	if err != nil {
		return nil, err
	}
	words := make([]model.Keywords, 8)
	err = md.Where("title LIKE ?", "%"+input.Keyword+"%").Limit(8).Scan(&words)
	if err != nil {
		return nil, err
	}
	output := &model.GetKeywordsOutput{}
	output.Words = words
	return output, nil
}

func (sPublicQuestion) GetAnswers(ctx context.Context, input *model.GetAnswersInput) (*model.GetAnswersOutput, error) {
	db := g.DB()
	sqlStr := `
SELECT al.question_id, u.avatar_file_id
FROM (	SELECT user_id, question_id, ROW_NUMBER() OVER (PARTITION BY question_id) as rowCnt 
		FROM answers WHERE question_id IN (?) ) al, users u
WHERE al.rowCnt <= 5 AND al.user_id = u.id;`
	res, err := db.GetAll(ctx, sqlStr, input.QuestionIDs)
	if err != nil {
		return nil, err
	}
	avatarsMap := make(map[int][]int)
	for _, row := range res {
		if _, ok := avatarsMap[row["question_id"].Int()]; !ok {
			avatarsMap[row["question_id"].Int()] = make([]int, 0, 5)
		}
		avatarsMap[row["question_id"].Int()] = append(avatarsMap[row["question_id"].Int()], row["avatar_file_id"].Int())
	}
	return &model.GetAnswersOutput{
		AvatarsMap: avatarsMap,
	}, nil
}

func (sPublicQuestion) AddQuestion(ctx context.Context, in *model.AddQuestionInput) (out *model.AddQuestionOutput, err error) {
	srcUserId := in.SrcUserID
	if srcUserId == 0 {
		srcUserId = consts.DefaultUserId
	}
	dstUserId := in.DstUserID
	if dstUserId == 0 {
		dstUserId = nil
	}
	question := do.Questions{
		SrcUserId: srcUserId,
		DstUserId: dstUserId,
		Title:     in.Title,
		Contents:  in.Content,
		IsPrivate: in.IsPrivate,
	}
	out = &model.AddQuestionOutput{}
	id, err := dao.Questions.Ctx(ctx).InsertAndGetId(question)
	if err != nil {
		if gstr.Contains(err.Error(), "FOREIGN KEY (`src_user_id`)") {
			return nil, gerror.New("找不到发送者")
		} else if gstr.Contains(err.Error(), "FOREIGN KEY (`dst_user_id`)") {
			return nil, gerror.New("找不到老师")
		}
		return nil, err
	}
	out.ID = int(id)

	// 添加通知
	if in.SrcUserID != nil {
		_, err := service.Notification().Add(ctx, model.AddNotificationInput{UserId: gconv.Int(in.SrcUserID), QuestionId: id})
		if err != nil {
			return nil, err
		}
	}
	return out, nil
}

func init() {
	service.RegisterPublicQuestion(New())
}

func New() *sPublicQuestion {
	return &sPublicQuestion{}
}
