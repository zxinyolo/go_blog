// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package service

import (
	"context"
	"go_blog/internal/model"
	"go_blog/internal/model/entity"
)

type IMoment interface {
	CreateMoment(ctx context.Context, in model.CreateMomentInput) (err error)
	ShowMoments(ctx context.Context, page, size int) (out *model.ShowMomentOutput, err error)
	UpdatePublishedStatus(ctx context.Context, momentId int, isPublished bool) (err error)
	MomoentDetil(ctx context.Context, momentId int) (out *model.MomentDetilOutput, err error)
	UpdateMoment(ctx context.Context, in entity.Moment) (err error)
	DeleteMoment(ctx context.Context, momenId int) (err error)
}

var localMoment IMoment

func Moment() IMoment {
	if localMoment == nil {
		panic("implement not found for interface IMoment, forgot register?")
	}
	return localMoment
}

func RegisterMoment(i IMoment) {
	localMoment = i
}