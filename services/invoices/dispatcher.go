package invoices

import (
	"errors"
	"github.com/spf13/viper"
	baiWang2 "invoiceOcr/services/invoices/baiWang"
	piaoTong2 "invoiceOcr/services/invoices/piaoTong"
)

var (
	ErrorChannelNotAllowed   = errors.New("暂不支持该渠道调用！")
	ErrorMethodNotAllowed    = errors.New("该渠道暂不支持该操作！")
	ErrorConfigEmpty         = errors.New("配置不全，接口无法调用！")
	ErrorServiceNotAvailable = errors.New("服务暂时不可用！")
	ErrorServiceUnAuthorized = errors.New("服务未被授权！")
)
var channelQueue map[string]interface{}

type channelType int

const (
	BaiWang channelType = iota
	PiaoTong
)

func (c channelType) String() string {
	switch c {
	case BaiWang:
		return "baiWang"
	case PiaoTong:
		return "piaoTong"
	default:
		return "undefined"
	}
}

type Context struct {
	patcher *Patcher
	Invoice interface{}
}

type Patcher struct {
	Channel     string
	PatcherType channelType
	Config      interface{}
}

type InvoicePatcher interface {
	// CheckConfig 检验配置
	CheckConfig() (bool, error)
	// MakeInvoice 开票
	MakeInvoice() (bool, error)
	// CheckInvoice 查询发票状态
	CheckInvoice() (bool, error)
	// NullifyInvoice 作废发票
	NullifyInvoice() (bool, error)
	// PrintInvoice 打印发票
	PrintInvoice() (interface{}, error)
}

type BehaviorInterface interface {
	// CheckParams 其他参数校验(不同活动类型实现不同)
	CheckParams(p *Context) error
	// InvoiceByVoucher 开票流程
	InvoiceByVoucher(p *Context) error
}

type InvoiceChannel struct {
	ConcreteBehavior BehaviorInterface
}

func (i *InvoiceChannel) MakeInvoice(ctx *Context) {

}

func (i *InvoiceChannel) CheckAuth(ctx *Context) {

}

func (i *InvoiceChannel) CheckData(ctx *Context) {

}

func (i *InvoiceChannel) CheckConfig(ctx *Context) {

}

func init() {
	config := viper.New()
	config.AddConfigPath("./config") //设置读取的文件路径
	config.SetConfigName("invoice")  //设置读取的文件名
	config.SetConfigType("yaml")     //设置文件的类型
	//尝试进行配置读取
	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}
	channelQueue = config.GetStringMap("invoiceQueue")
}

func Run() error {
	for channel, conf := range channelQueue {
		ctx := &Context{
			patcher: &Patcher{
				Channel: channel,
				Config:  conf,
			},
		}

		switch ctx.patcher.PatcherType {
		case PiaoTong: // 票通
			instance := &baiWang2.InputData{}
			instance.ConcreteBehavior = instance
			return instance.Run(ctx)
		case BaiWang: // 百望
			instance := &piaoTong2.Invoice{}
			instance.ConcreteBehavior = instance
			instance.Run(ctx)
		default:
			// 报错
			return ErrorChannelNotAllowed
		}
	}
	return nil
}
