package fix40

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//IndicationofInterest msg type = 6.
type IndicationofInterest struct {
	message.Message
}

//IndicationofInterestBuilder builds IndicationofInterest messages.
type IndicationofInterestBuilder struct {
	message.MessageBuilder
}

//CreateIndicationofInterestBuilder returns an initialized IndicationofInterestBuilder with specified required fields.
func CreateIndicationofInterestBuilder(
	ioiid *field.IOIidField,
	ioitranstype *field.IOITransTypeField,
	symbol *field.SymbolField,
	side *field.SideField,
	ioishares *field.IOISharesField) IndicationofInterestBuilder {
	var builder IndicationofInterestBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIX40))
	builder.Header().Set(field.NewMsgType("6"))
	builder.Body().Set(ioiid)
	builder.Body().Set(ioitranstype)
	builder.Body().Set(symbol)
	builder.Body().Set(side)
	builder.Body().Set(ioishares)
	return builder
}

//IOIid is a required field for IndicationofInterest.
func (m IndicationofInterest) IOIid() (*field.IOIidField, errors.MessageRejectError) {
	f := &field.IOIidField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIOIid reads a IOIid from IndicationofInterest.
func (m IndicationofInterest) GetIOIid(f *field.IOIidField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IOITransType is a required field for IndicationofInterest.
func (m IndicationofInterest) IOITransType() (*field.IOITransTypeField, errors.MessageRejectError) {
	f := &field.IOITransTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIOITransType reads a IOITransType from IndicationofInterest.
func (m IndicationofInterest) GetIOITransType(f *field.IOITransTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IOIRefID is a non-required field for IndicationofInterest.
func (m IndicationofInterest) IOIRefID() (*field.IOIRefIDField, errors.MessageRejectError) {
	f := &field.IOIRefIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIOIRefID reads a IOIRefID from IndicationofInterest.
func (m IndicationofInterest) GetIOIRefID(f *field.IOIRefIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a required field for IndicationofInterest.
func (m IndicationofInterest) Symbol() (*field.SymbolField, errors.MessageRejectError) {
	f := &field.SymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from IndicationofInterest.
func (m IndicationofInterest) GetSymbol(f *field.SymbolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for IndicationofInterest.
func (m IndicationofInterest) SymbolSfx() (*field.SymbolSfxField, errors.MessageRejectError) {
	f := &field.SymbolSfxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from IndicationofInterest.
func (m IndicationofInterest) GetSymbolSfx(f *field.SymbolSfxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for IndicationofInterest.
func (m IndicationofInterest) SecurityID() (*field.SecurityIDField, errors.MessageRejectError) {
	f := &field.SecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from IndicationofInterest.
func (m IndicationofInterest) GetSecurityID(f *field.SecurityIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IDSource is a non-required field for IndicationofInterest.
func (m IndicationofInterest) IDSource() (*field.IDSourceField, errors.MessageRejectError) {
	f := &field.IDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIDSource reads a IDSource from IndicationofInterest.
func (m IndicationofInterest) GetIDSource(f *field.IDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for IndicationofInterest.
func (m IndicationofInterest) Issuer() (*field.IssuerField, errors.MessageRejectError) {
	f := &field.IssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from IndicationofInterest.
func (m IndicationofInterest) GetIssuer(f *field.IssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for IndicationofInterest.
func (m IndicationofInterest) SecurityDesc() (*field.SecurityDescField, errors.MessageRejectError) {
	f := &field.SecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from IndicationofInterest.
func (m IndicationofInterest) GetSecurityDesc(f *field.SecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Side is a required field for IndicationofInterest.
func (m IndicationofInterest) Side() (*field.SideField, errors.MessageRejectError) {
	f := &field.SideField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSide reads a Side from IndicationofInterest.
func (m IndicationofInterest) GetSide(f *field.SideField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IOIShares is a required field for IndicationofInterest.
func (m IndicationofInterest) IOIShares() (*field.IOISharesField, errors.MessageRejectError) {
	f := &field.IOISharesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIOIShares reads a IOIShares from IndicationofInterest.
func (m IndicationofInterest) GetIOIShares(f *field.IOISharesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Price is a non-required field for IndicationofInterest.
func (m IndicationofInterest) Price() (*field.PriceField, errors.MessageRejectError) {
	f := &field.PriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPrice reads a Price from IndicationofInterest.
func (m IndicationofInterest) GetPrice(f *field.PriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for IndicationofInterest.
func (m IndicationofInterest) Currency() (*field.CurrencyField, errors.MessageRejectError) {
	f := &field.CurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from IndicationofInterest.
func (m IndicationofInterest) GetCurrency(f *field.CurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ValidUntilTime is a non-required field for IndicationofInterest.
func (m IndicationofInterest) ValidUntilTime() (*field.ValidUntilTimeField, errors.MessageRejectError) {
	f := &field.ValidUntilTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetValidUntilTime reads a ValidUntilTime from IndicationofInterest.
func (m IndicationofInterest) GetValidUntilTime(f *field.ValidUntilTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IOIQltyInd is a non-required field for IndicationofInterest.
func (m IndicationofInterest) IOIQltyInd() (*field.IOIQltyIndField, errors.MessageRejectError) {
	f := &field.IOIQltyIndField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIOIQltyInd reads a IOIQltyInd from IndicationofInterest.
func (m IndicationofInterest) GetIOIQltyInd(f *field.IOIQltyIndField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IOIOthSvc is a non-required field for IndicationofInterest.
func (m IndicationofInterest) IOIOthSvc() (*field.IOIOthSvcField, errors.MessageRejectError) {
	f := &field.IOIOthSvcField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIOIOthSvc reads a IOIOthSvc from IndicationofInterest.
func (m IndicationofInterest) GetIOIOthSvc(f *field.IOIOthSvcField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IOINaturalFlag is a non-required field for IndicationofInterest.
func (m IndicationofInterest) IOINaturalFlag() (*field.IOINaturalFlagField, errors.MessageRejectError) {
	f := &field.IOINaturalFlagField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIOINaturalFlag reads a IOINaturalFlag from IndicationofInterest.
func (m IndicationofInterest) GetIOINaturalFlag(f *field.IOINaturalFlagField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IOIQualifier is a non-required field for IndicationofInterest.
func (m IndicationofInterest) IOIQualifier() (*field.IOIQualifierField, errors.MessageRejectError) {
	f := &field.IOIQualifierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIOIQualifier reads a IOIQualifier from IndicationofInterest.
func (m IndicationofInterest) GetIOIQualifier(f *field.IOIQualifierField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for IndicationofInterest.
func (m IndicationofInterest) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from IndicationofInterest.
func (m IndicationofInterest) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}
