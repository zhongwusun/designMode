package adapter

import "fmt"

// CommonPlug 通用的USB电源插槽
type CommonPlug interface {
	ConnectUSB() string
}

// HuaweiPhonePlugAdapter 华为TypeC充电插槽适配通用USB充电插槽
type HuaweiPhonePlugAdapter struct {
	huaweiPhone HuaweiPlug
}

// NewHuaweiPhonePlugAdapter 创建华为手机适配USB充电插槽适配器
func NewHuaweiPhonePlugAdapter(huaweiPhone HuaweiPlug) *HuaweiPhonePlugAdapter {
	return &HuaweiPhonePlugAdapter{
		huaweiPhone: huaweiPhone,
	}
}

// ConnectUSB 链接USB
func (h *HuaweiPhonePlugAdapter) ConnectUSB() string {
	return fmt.Sprintf("%v adapt to usb ", h.huaweiPhone.ConnectTypeC())
}

// ApplePhonePlugAdapter 苹果Lightning充电插槽适配通用USB充电插槽
type ApplePhonePlugAdapter struct {
	iPhone ApplePlug
}

// NewApplePhonePlugAdapter 创建苹果手机适配USB充电插槽适配器
func NewApplePhonePlugAdapter(iPhone ApplePlug) *ApplePhonePlugAdapter {
	return &ApplePhonePlugAdapter{
		iPhone: iPhone,
	}
}

// ConnectUSB 链接USB
func (a *ApplePhonePlugAdapter) ConnectUSB() string {
	return fmt.Sprintf("%v adapt to usb ", a.iPhone.ConnectLightning())
}

// PowerBank 充电宝
type PowerBank struct {
	brand string
}

// Charge 支持通用USB接口充电
func (p *PowerBank) Charge(plug CommonPlug) string {
	return fmt.Sprintf("%v power bank connect usb plug, start charge for %v", p.brand, plug.ConnectUSB())
}
