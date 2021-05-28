package models

type StatusType int

const (
	StatusInActive StatusType = iota //禁用
	StatusActive                     //启用
)

var StatusRange = []StatusType{
	StatusInActive,
	StatusActive,
}

func (c StatusType) String() string {
	switch c {
	case StatusInActive:
		return "禁用"
	case StatusActive:
		return "启用"
	default:
		return "启用"
	}
}

type ChannelType int

const (
	BaiWang ChannelType = iota
	PiaoTong
)

func (c ChannelType) String() string {
	switch c {
	case BaiWang:
		return "baiWang"
	case PiaoTong:
		return "piaoTong"
	default:
		return "undefined"
	}
}
