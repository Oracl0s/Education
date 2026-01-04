package payments

type PaymentMethod interface {
	Pay(usd int) int
	Cancel(id int)
}

type PaymentModule struct {
	paymentsInfo map[int]PaymentInfo

	paymentMethod PaymentMethod
}

func NewPaymentModule(paymentMethod PaymentMethod) *PaymentModule {
	return &PaymentModule{
		paymentsInfo:  make(map[int]PaymentInfo),
		paymentMethod: paymentMethod,
	}
}

// Метод Pay()
// Принемает:
// 1.Описание проводимой оплаты
// 2.Сумма оплаты
// Возврощает:
// 1.IDпроведенной операции
func (p *PaymentModule) Pay(descrition string, usd int) int {
	//
	//
	id := p.paymentMethod.Pay(usd)
	info := PaymentInfo{
		Description: descrition,
		Usd:         usd,
		Cancelled:   false,
	}
	//1.Проводить оплату
	//2.получать id проведенной оплаты
	//3.сохронять информацию о проведенной операции
	//- описание операций
	//- сколько было потрачено
	//- отмененная ли операция
	p.paymentsInfo[id] = info
	//4.возврощать id проведенной оплаты
	return id
}

// Метод Cacel()
// Принемает:
// 1.ID операции
// Возврощает:
// -неченго не возврощает
func (p *PaymentModule) Cancel(id int) {
	p.paymentMethod.Cancel(id)
	info, ok := p.paymentsInfo[id]
	if !ok {
		return
	}
	info.Cancelled = true
	p.paymentsInfo[id] = info
}

// Метод Info()
// Принемает:
// 1.ID операции
// Возврощает:
// Информации о проведеной операции
func (p *PaymentModule) Info(id int) PaymentInfo {
	info, ok := p.paymentsInfo[id]
	if !ok {
		return PaymentInfo{}
	}
	return info
}

// Метод AllInfo()
// Принемает:
// 1.ID операции
// Возврощает:
// 1. Информацию о всех проведенных операциях
func (p *PaymentModule) AllInfo() map[int]PaymentInfo {
	tempMap := make(map[int]PaymentInfo, len(p.paymentsInfo))
	for k, v := range p.paymentsInfo {
		tempMap[k] = v
	}
	return tempMap
}
