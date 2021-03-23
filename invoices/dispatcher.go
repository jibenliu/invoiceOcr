package invoices

import (
	"github.com/spf13/viper"
	"invoiceOcr/invoices/baiWang"
	"invoiceOcr/invoices/piaoTong"
)

var channelQueue [interface]interface{}

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
	PatcherType channelType
}

type BehaviorInterface interface {
	// 其他参数校验(不同活动类型实现不同)
	CheckParams(p *Context) error
	//开票流程
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
	config.SetConfigType("yml")      //设置文件的类型
	//尝试进行配置读取
	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}
	channelQueue = config.Get("invoiceQueue")
}
func Run() {
	for channel := range channelQueue{
		ctx := &Context{
			patcher: &Patcher{
				PatcherType: channel,
			},
		}

		switch ctx.patcher.PatcherType {
		case PiaoTong: // 票通
			instance := &baiWang.Invoice{}
			instance.ConcreteBehavior = instance
			instance.Run(ctx)
		case BaiWang: // 百望
			instance := &piaoTong.Invoice{}
			instance.ConcreteBehavior = instance
			instance.Run(ctx)
		default:
			// 报错
			return
		}
	}

}
