package fix44

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//DontKnowTrade msg type = Q.
type DontKnowTrade struct {
	message.Message
}

//DontKnowTradeBuilder builds DontKnowTrade messages.
type DontKnowTradeBuilder struct {
	message.MessageBuilder
}

//CreateDontKnowTradeBuilder returns an initialized DontKnowTradeBuilder with specified required fields.
func CreateDontKnowTradeBuilder(
	orderid *field.OrderIDField,
	execid *field.ExecIDField,
	dkreason *field.DKReasonField,
	side *field.SideField) DontKnowTradeBuilder {
	var builder DontKnowTradeBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIX44))
	builder.Header().Set(field.NewMsgType("Q"))
	builder.Body().Set(orderid)
	builder.Body().Set(execid)
	builder.Body().Set(dkreason)
	builder.Body().Set(side)
	return builder
}

//OrderID is a required field for DontKnowTrade.
func (m DontKnowTrade) OrderID() (*field.OrderIDField, errors.MessageRejectError) {
	f := &field.OrderIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderID reads a OrderID from DontKnowTrade.
func (m DontKnowTrade) GetOrderID(f *field.OrderIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryOrderID is a non-required field for DontKnowTrade.
func (m DontKnowTrade) SecondaryOrderID() (*field.SecondaryOrderIDField, errors.MessageRejectError) {
	f := &field.SecondaryOrderIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryOrderID reads a SecondaryOrderID from DontKnowTrade.
func (m DontKnowTrade) GetSecondaryOrderID(f *field.SecondaryOrderIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecID is a required field for DontKnowTrade.
func (m DontKnowTrade) ExecID() (*field.ExecIDField, errors.MessageRejectError) {
	f := &field.ExecIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExecID reads a ExecID from DontKnowTrade.
func (m DontKnowTrade) GetExecID(f *field.ExecIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DKReason is a required field for DontKnowTrade.
func (m DontKnowTrade) DKReason() (*field.DKReasonField, errors.MessageRejectError) {
	f := &field.DKReasonField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDKReason reads a DKReason from DontKnowTrade.
func (m DontKnowTrade) GetDKReason(f *field.DKReasonField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for DontKnowTrade.
func (m DontKnowTrade) Symbol() (*field.SymbolField, errors.MessageRejectError) {
	f := &field.SymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from DontKnowTrade.
func (m DontKnowTrade) GetSymbol(f *field.SymbolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for DontKnowTrade.
func (m DontKnowTrade) SymbolSfx() (*field.SymbolSfxField, errors.MessageRejectError) {
	f := &field.SymbolSfxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from DontKnowTrade.
func (m DontKnowTrade) GetSymbolSfx(f *field.SymbolSfxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for DontKnowTrade.
func (m DontKnowTrade) SecurityID() (*field.SecurityIDField, errors.MessageRejectError) {
	f := &field.SecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from DontKnowTrade.
func (m DontKnowTrade) GetSecurityID(f *field.SecurityIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for DontKnowTrade.
func (m DontKnowTrade) SecurityIDSource() (*field.SecurityIDSourceField, errors.MessageRejectError) {
	f := &field.SecurityIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from DontKnowTrade.
func (m DontKnowTrade) GetSecurityIDSource(f *field.SecurityIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for DontKnowTrade.
func (m DontKnowTrade) NoSecurityAltID() (*field.NoSecurityAltIDField, errors.MessageRejectError) {
	f := &field.NoSecurityAltIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from DontKnowTrade.
func (m DontKnowTrade) GetNoSecurityAltID(f *field.NoSecurityAltIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for DontKnowTrade.
func (m DontKnowTrade) Product() (*field.ProductField, errors.MessageRejectError) {
	f := &field.ProductField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from DontKnowTrade.
func (m DontKnowTrade) GetProduct(f *field.ProductField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for DontKnowTrade.
func (m DontKnowTrade) CFICode() (*field.CFICodeField, errors.MessageRejectError) {
	f := &field.CFICodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from DontKnowTrade.
func (m DontKnowTrade) GetCFICode(f *field.CFICodeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for DontKnowTrade.
func (m DontKnowTrade) SecurityType() (*field.SecurityTypeField, errors.MessageRejectError) {
	f := &field.SecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from DontKnowTrade.
func (m DontKnowTrade) GetSecurityType(f *field.SecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySubType is a non-required field for DontKnowTrade.
func (m DontKnowTrade) SecuritySubType() (*field.SecuritySubTypeField, errors.MessageRejectError) {
	f := &field.SecuritySubTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySubType reads a SecuritySubType from DontKnowTrade.
func (m DontKnowTrade) GetSecuritySubType(f *field.SecuritySubTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for DontKnowTrade.
func (m DontKnowTrade) MaturityMonthYear() (*field.MaturityMonthYearField, errors.MessageRejectError) {
	f := &field.MaturityMonthYearField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from DontKnowTrade.
func (m DontKnowTrade) GetMaturityMonthYear(f *field.MaturityMonthYearField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for DontKnowTrade.
func (m DontKnowTrade) MaturityDate() (*field.MaturityDateField, errors.MessageRejectError) {
	f := &field.MaturityDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from DontKnowTrade.
func (m DontKnowTrade) GetMaturityDate(f *field.MaturityDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for DontKnowTrade.
func (m DontKnowTrade) CouponPaymentDate() (*field.CouponPaymentDateField, errors.MessageRejectError) {
	f := &field.CouponPaymentDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from DontKnowTrade.
func (m DontKnowTrade) GetCouponPaymentDate(f *field.CouponPaymentDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for DontKnowTrade.
func (m DontKnowTrade) IssueDate() (*field.IssueDateField, errors.MessageRejectError) {
	f := &field.IssueDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from DontKnowTrade.
func (m DontKnowTrade) GetIssueDate(f *field.IssueDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for DontKnowTrade.
func (m DontKnowTrade) RepoCollateralSecurityType() (*field.RepoCollateralSecurityTypeField, errors.MessageRejectError) {
	f := &field.RepoCollateralSecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from DontKnowTrade.
func (m DontKnowTrade) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for DontKnowTrade.
func (m DontKnowTrade) RepurchaseTerm() (*field.RepurchaseTermField, errors.MessageRejectError) {
	f := &field.RepurchaseTermField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from DontKnowTrade.
func (m DontKnowTrade) GetRepurchaseTerm(f *field.RepurchaseTermField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for DontKnowTrade.
func (m DontKnowTrade) RepurchaseRate() (*field.RepurchaseRateField, errors.MessageRejectError) {
	f := &field.RepurchaseRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from DontKnowTrade.
func (m DontKnowTrade) GetRepurchaseRate(f *field.RepurchaseRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for DontKnowTrade.
func (m DontKnowTrade) Factor() (*field.FactorField, errors.MessageRejectError) {
	f := &field.FactorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from DontKnowTrade.
func (m DontKnowTrade) GetFactor(f *field.FactorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for DontKnowTrade.
func (m DontKnowTrade) CreditRating() (*field.CreditRatingField, errors.MessageRejectError) {
	f := &field.CreditRatingField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from DontKnowTrade.
func (m DontKnowTrade) GetCreditRating(f *field.CreditRatingField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for DontKnowTrade.
func (m DontKnowTrade) InstrRegistry() (*field.InstrRegistryField, errors.MessageRejectError) {
	f := &field.InstrRegistryField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from DontKnowTrade.
func (m DontKnowTrade) GetInstrRegistry(f *field.InstrRegistryField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for DontKnowTrade.
func (m DontKnowTrade) CountryOfIssue() (*field.CountryOfIssueField, errors.MessageRejectError) {
	f := &field.CountryOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from DontKnowTrade.
func (m DontKnowTrade) GetCountryOfIssue(f *field.CountryOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for DontKnowTrade.
func (m DontKnowTrade) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssueField, errors.MessageRejectError) {
	f := &field.StateOrProvinceOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from DontKnowTrade.
func (m DontKnowTrade) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for DontKnowTrade.
func (m DontKnowTrade) LocaleOfIssue() (*field.LocaleOfIssueField, errors.MessageRejectError) {
	f := &field.LocaleOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from DontKnowTrade.
func (m DontKnowTrade) GetLocaleOfIssue(f *field.LocaleOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for DontKnowTrade.
func (m DontKnowTrade) RedemptionDate() (*field.RedemptionDateField, errors.MessageRejectError) {
	f := &field.RedemptionDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from DontKnowTrade.
func (m DontKnowTrade) GetRedemptionDate(f *field.RedemptionDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for DontKnowTrade.
func (m DontKnowTrade) StrikePrice() (*field.StrikePriceField, errors.MessageRejectError) {
	f := &field.StrikePriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from DontKnowTrade.
func (m DontKnowTrade) GetStrikePrice(f *field.StrikePriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeCurrency is a non-required field for DontKnowTrade.
func (m DontKnowTrade) StrikeCurrency() (*field.StrikeCurrencyField, errors.MessageRejectError) {
	f := &field.StrikeCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeCurrency reads a StrikeCurrency from DontKnowTrade.
func (m DontKnowTrade) GetStrikeCurrency(f *field.StrikeCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for DontKnowTrade.
func (m DontKnowTrade) OptAttribute() (*field.OptAttributeField, errors.MessageRejectError) {
	f := &field.OptAttributeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from DontKnowTrade.
func (m DontKnowTrade) GetOptAttribute(f *field.OptAttributeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for DontKnowTrade.
func (m DontKnowTrade) ContractMultiplier() (*field.ContractMultiplierField, errors.MessageRejectError) {
	f := &field.ContractMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from DontKnowTrade.
func (m DontKnowTrade) GetContractMultiplier(f *field.ContractMultiplierField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for DontKnowTrade.
func (m DontKnowTrade) CouponRate() (*field.CouponRateField, errors.MessageRejectError) {
	f := &field.CouponRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from DontKnowTrade.
func (m DontKnowTrade) GetCouponRate(f *field.CouponRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for DontKnowTrade.
func (m DontKnowTrade) SecurityExchange() (*field.SecurityExchangeField, errors.MessageRejectError) {
	f := &field.SecurityExchangeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from DontKnowTrade.
func (m DontKnowTrade) GetSecurityExchange(f *field.SecurityExchangeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for DontKnowTrade.
func (m DontKnowTrade) Issuer() (*field.IssuerField, errors.MessageRejectError) {
	f := &field.IssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from DontKnowTrade.
func (m DontKnowTrade) GetIssuer(f *field.IssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for DontKnowTrade.
func (m DontKnowTrade) EncodedIssuerLen() (*field.EncodedIssuerLenField, errors.MessageRejectError) {
	f := &field.EncodedIssuerLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from DontKnowTrade.
func (m DontKnowTrade) GetEncodedIssuerLen(f *field.EncodedIssuerLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for DontKnowTrade.
func (m DontKnowTrade) EncodedIssuer() (*field.EncodedIssuerField, errors.MessageRejectError) {
	f := &field.EncodedIssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from DontKnowTrade.
func (m DontKnowTrade) GetEncodedIssuer(f *field.EncodedIssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for DontKnowTrade.
func (m DontKnowTrade) SecurityDesc() (*field.SecurityDescField, errors.MessageRejectError) {
	f := &field.SecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from DontKnowTrade.
func (m DontKnowTrade) GetSecurityDesc(f *field.SecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for DontKnowTrade.
func (m DontKnowTrade) EncodedSecurityDescLen() (*field.EncodedSecurityDescLenField, errors.MessageRejectError) {
	f := &field.EncodedSecurityDescLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from DontKnowTrade.
func (m DontKnowTrade) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for DontKnowTrade.
func (m DontKnowTrade) EncodedSecurityDesc() (*field.EncodedSecurityDescField, errors.MessageRejectError) {
	f := &field.EncodedSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from DontKnowTrade.
func (m DontKnowTrade) GetEncodedSecurityDesc(f *field.EncodedSecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Pool is a non-required field for DontKnowTrade.
func (m DontKnowTrade) Pool() (*field.PoolField, errors.MessageRejectError) {
	f := &field.PoolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPool reads a Pool from DontKnowTrade.
func (m DontKnowTrade) GetPool(f *field.PoolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractSettlMonth is a non-required field for DontKnowTrade.
func (m DontKnowTrade) ContractSettlMonth() (*field.ContractSettlMonthField, errors.MessageRejectError) {
	f := &field.ContractSettlMonthField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractSettlMonth reads a ContractSettlMonth from DontKnowTrade.
func (m DontKnowTrade) GetContractSettlMonth(f *field.ContractSettlMonthField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPProgram is a non-required field for DontKnowTrade.
func (m DontKnowTrade) CPProgram() (*field.CPProgramField, errors.MessageRejectError) {
	f := &field.CPProgramField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCPProgram reads a CPProgram from DontKnowTrade.
func (m DontKnowTrade) GetCPProgram(f *field.CPProgramField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPRegType is a non-required field for DontKnowTrade.
func (m DontKnowTrade) CPRegType() (*field.CPRegTypeField, errors.MessageRejectError) {
	f := &field.CPRegTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCPRegType reads a CPRegType from DontKnowTrade.
func (m DontKnowTrade) GetCPRegType(f *field.CPRegTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoEvents is a non-required field for DontKnowTrade.
func (m DontKnowTrade) NoEvents() (*field.NoEventsField, errors.MessageRejectError) {
	f := &field.NoEventsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoEvents reads a NoEvents from DontKnowTrade.
func (m DontKnowTrade) GetNoEvents(f *field.NoEventsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DatedDate is a non-required field for DontKnowTrade.
func (m DontKnowTrade) DatedDate() (*field.DatedDateField, errors.MessageRejectError) {
	f := &field.DatedDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDatedDate reads a DatedDate from DontKnowTrade.
func (m DontKnowTrade) GetDatedDate(f *field.DatedDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InterestAccrualDate is a non-required field for DontKnowTrade.
func (m DontKnowTrade) InterestAccrualDate() (*field.InterestAccrualDateField, errors.MessageRejectError) {
	f := &field.InterestAccrualDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInterestAccrualDate reads a InterestAccrualDate from DontKnowTrade.
func (m DontKnowTrade) GetInterestAccrualDate(f *field.InterestAccrualDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyings is a non-required field for DontKnowTrade.
func (m DontKnowTrade) NoUnderlyings() (*field.NoUnderlyingsField, errors.MessageRejectError) {
	f := &field.NoUnderlyingsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyings reads a NoUnderlyings from DontKnowTrade.
func (m DontKnowTrade) GetNoUnderlyings(f *field.NoUnderlyingsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoLegs is a non-required field for DontKnowTrade.
func (m DontKnowTrade) NoLegs() (*field.NoLegsField, errors.MessageRejectError) {
	f := &field.NoLegsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoLegs reads a NoLegs from DontKnowTrade.
func (m DontKnowTrade) GetNoLegs(f *field.NoLegsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Side is a required field for DontKnowTrade.
func (m DontKnowTrade) Side() (*field.SideField, errors.MessageRejectError) {
	f := &field.SideField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSide reads a Side from DontKnowTrade.
func (m DontKnowTrade) GetSide(f *field.SideField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderQty is a non-required field for DontKnowTrade.
func (m DontKnowTrade) OrderQty() (*field.OrderQtyField, errors.MessageRejectError) {
	f := &field.OrderQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderQty reads a OrderQty from DontKnowTrade.
func (m DontKnowTrade) GetOrderQty(f *field.OrderQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CashOrderQty is a non-required field for DontKnowTrade.
func (m DontKnowTrade) CashOrderQty() (*field.CashOrderQtyField, errors.MessageRejectError) {
	f := &field.CashOrderQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCashOrderQty reads a CashOrderQty from DontKnowTrade.
func (m DontKnowTrade) GetCashOrderQty(f *field.CashOrderQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderPercent is a non-required field for DontKnowTrade.
func (m DontKnowTrade) OrderPercent() (*field.OrderPercentField, errors.MessageRejectError) {
	f := &field.OrderPercentField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderPercent reads a OrderPercent from DontKnowTrade.
func (m DontKnowTrade) GetOrderPercent(f *field.OrderPercentField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RoundingDirection is a non-required field for DontKnowTrade.
func (m DontKnowTrade) RoundingDirection() (*field.RoundingDirectionField, errors.MessageRejectError) {
	f := &field.RoundingDirectionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRoundingDirection reads a RoundingDirection from DontKnowTrade.
func (m DontKnowTrade) GetRoundingDirection(f *field.RoundingDirectionField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RoundingModulus is a non-required field for DontKnowTrade.
func (m DontKnowTrade) RoundingModulus() (*field.RoundingModulusField, errors.MessageRejectError) {
	f := &field.RoundingModulusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRoundingModulus reads a RoundingModulus from DontKnowTrade.
func (m DontKnowTrade) GetRoundingModulus(f *field.RoundingModulusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastQty is a non-required field for DontKnowTrade.
func (m DontKnowTrade) LastQty() (*field.LastQtyField, errors.MessageRejectError) {
	f := &field.LastQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastQty reads a LastQty from DontKnowTrade.
func (m DontKnowTrade) GetLastQty(f *field.LastQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastPx is a non-required field for DontKnowTrade.
func (m DontKnowTrade) LastPx() (*field.LastPxField, errors.MessageRejectError) {
	f := &field.LastPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastPx reads a LastPx from DontKnowTrade.
func (m DontKnowTrade) GetLastPx(f *field.LastPxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for DontKnowTrade.
func (m DontKnowTrade) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from DontKnowTrade.
func (m DontKnowTrade) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for DontKnowTrade.
func (m DontKnowTrade) EncodedTextLen() (*field.EncodedTextLenField, errors.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from DontKnowTrade.
func (m DontKnowTrade) GetEncodedTextLen(f *field.EncodedTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for DontKnowTrade.
func (m DontKnowTrade) EncodedText() (*field.EncodedTextField, errors.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from DontKnowTrade.
func (m DontKnowTrade) GetEncodedText(f *field.EncodedTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}
