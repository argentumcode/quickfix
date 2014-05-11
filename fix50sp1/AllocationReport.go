package fix50sp1

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
)

//AllocationReport msg type = AS.
type AllocationReport struct {
	message.Message
}

//AllocationReportBuilder builds AllocationReport messages.
type AllocationReportBuilder struct {
	message.MessageBuilder
}

//CreateAllocationReportBuilder returns an initialized AllocationReportBuilder with specified required fields.
func CreateAllocationReportBuilder(
	allocreportid *field.AllocReportIDField,
	alloctranstype *field.AllocTransTypeField,
	allocreporttype *field.AllocReportTypeField,
	allocstatus *field.AllocStatusField,
	side *field.SideField,
	quantity *field.QuantityField,
	avgpx *field.AvgPxField,
	tradedate *field.TradeDateField) AllocationReportBuilder {
	var builder AllocationReportBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header().Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP1))
	builder.Header().Set(field.NewMsgType("AS"))
	builder.Body().Set(allocreportid)
	builder.Body().Set(alloctranstype)
	builder.Body().Set(allocreporttype)
	builder.Body().Set(allocstatus)
	builder.Body().Set(side)
	builder.Body().Set(quantity)
	builder.Body().Set(avgpx)
	builder.Body().Set(tradedate)
	return builder
}

//AllocReportID is a required field for AllocationReport.
func (m AllocationReport) AllocReportID() (*field.AllocReportIDField, errors.MessageRejectError) {
	f := &field.AllocReportIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocReportID reads a AllocReportID from AllocationReport.
func (m AllocationReport) GetAllocReportID(f *field.AllocReportIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocID is a non-required field for AllocationReport.
func (m AllocationReport) AllocID() (*field.AllocIDField, errors.MessageRejectError) {
	f := &field.AllocIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocID reads a AllocID from AllocationReport.
func (m AllocationReport) GetAllocID(f *field.AllocIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocTransType is a required field for AllocationReport.
func (m AllocationReport) AllocTransType() (*field.AllocTransTypeField, errors.MessageRejectError) {
	f := &field.AllocTransTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocTransType reads a AllocTransType from AllocationReport.
func (m AllocationReport) GetAllocTransType(f *field.AllocTransTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocReportRefID is a non-required field for AllocationReport.
func (m AllocationReport) AllocReportRefID() (*field.AllocReportRefIDField, errors.MessageRejectError) {
	f := &field.AllocReportRefIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocReportRefID reads a AllocReportRefID from AllocationReport.
func (m AllocationReport) GetAllocReportRefID(f *field.AllocReportRefIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocCancReplaceReason is a non-required field for AllocationReport.
func (m AllocationReport) AllocCancReplaceReason() (*field.AllocCancReplaceReasonField, errors.MessageRejectError) {
	f := &field.AllocCancReplaceReasonField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocCancReplaceReason reads a AllocCancReplaceReason from AllocationReport.
func (m AllocationReport) GetAllocCancReplaceReason(f *field.AllocCancReplaceReasonField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryAllocID is a non-required field for AllocationReport.
func (m AllocationReport) SecondaryAllocID() (*field.SecondaryAllocIDField, errors.MessageRejectError) {
	f := &field.SecondaryAllocIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryAllocID reads a SecondaryAllocID from AllocationReport.
func (m AllocationReport) GetSecondaryAllocID(f *field.SecondaryAllocIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocReportType is a required field for AllocationReport.
func (m AllocationReport) AllocReportType() (*field.AllocReportTypeField, errors.MessageRejectError) {
	f := &field.AllocReportTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocReportType reads a AllocReportType from AllocationReport.
func (m AllocationReport) GetAllocReportType(f *field.AllocReportTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocStatus is a required field for AllocationReport.
func (m AllocationReport) AllocStatus() (*field.AllocStatusField, errors.MessageRejectError) {
	f := &field.AllocStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocStatus reads a AllocStatus from AllocationReport.
func (m AllocationReport) GetAllocStatus(f *field.AllocStatusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocRejCode is a non-required field for AllocationReport.
func (m AllocationReport) AllocRejCode() (*field.AllocRejCodeField, errors.MessageRejectError) {
	f := &field.AllocRejCodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocRejCode reads a AllocRejCode from AllocationReport.
func (m AllocationReport) GetAllocRejCode(f *field.AllocRejCodeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RefAllocID is a non-required field for AllocationReport.
func (m AllocationReport) RefAllocID() (*field.RefAllocIDField, errors.MessageRejectError) {
	f := &field.RefAllocIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRefAllocID reads a RefAllocID from AllocationReport.
func (m AllocationReport) GetRefAllocID(f *field.RefAllocIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocIntermedReqType is a non-required field for AllocationReport.
func (m AllocationReport) AllocIntermedReqType() (*field.AllocIntermedReqTypeField, errors.MessageRejectError) {
	f := &field.AllocIntermedReqTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocIntermedReqType reads a AllocIntermedReqType from AllocationReport.
func (m AllocationReport) GetAllocIntermedReqType(f *field.AllocIntermedReqTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocLinkID is a non-required field for AllocationReport.
func (m AllocationReport) AllocLinkID() (*field.AllocLinkIDField, errors.MessageRejectError) {
	f := &field.AllocLinkIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocLinkID reads a AllocLinkID from AllocationReport.
func (m AllocationReport) GetAllocLinkID(f *field.AllocLinkIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocLinkType is a non-required field for AllocationReport.
func (m AllocationReport) AllocLinkType() (*field.AllocLinkTypeField, errors.MessageRejectError) {
	f := &field.AllocLinkTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocLinkType reads a AllocLinkType from AllocationReport.
func (m AllocationReport) GetAllocLinkType(f *field.AllocLinkTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BookingRefID is a non-required field for AllocationReport.
func (m AllocationReport) BookingRefID() (*field.BookingRefIDField, errors.MessageRejectError) {
	f := &field.BookingRefIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBookingRefID reads a BookingRefID from AllocationReport.
func (m AllocationReport) GetBookingRefID(f *field.BookingRefIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocNoOrdersType is a non-required field for AllocationReport.
func (m AllocationReport) AllocNoOrdersType() (*field.AllocNoOrdersTypeField, errors.MessageRejectError) {
	f := &field.AllocNoOrdersTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocNoOrdersType reads a AllocNoOrdersType from AllocationReport.
func (m AllocationReport) GetAllocNoOrdersType(f *field.AllocNoOrdersTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoOrders is a non-required field for AllocationReport.
func (m AllocationReport) NoOrders() (*field.NoOrdersField, errors.MessageRejectError) {
	f := &field.NoOrdersField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoOrders reads a NoOrders from AllocationReport.
func (m AllocationReport) GetNoOrders(f *field.NoOrdersField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoExecs is a non-required field for AllocationReport.
func (m AllocationReport) NoExecs() (*field.NoExecsField, errors.MessageRejectError) {
	f := &field.NoExecsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoExecs reads a NoExecs from AllocationReport.
func (m AllocationReport) GetNoExecs(f *field.NoExecsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PreviouslyReported is a non-required field for AllocationReport.
func (m AllocationReport) PreviouslyReported() (*field.PreviouslyReportedField, errors.MessageRejectError) {
	f := &field.PreviouslyReportedField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPreviouslyReported reads a PreviouslyReported from AllocationReport.
func (m AllocationReport) GetPreviouslyReported(f *field.PreviouslyReportedField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ReversalIndicator is a non-required field for AllocationReport.
func (m AllocationReport) ReversalIndicator() (*field.ReversalIndicatorField, errors.MessageRejectError) {
	f := &field.ReversalIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetReversalIndicator reads a ReversalIndicator from AllocationReport.
func (m AllocationReport) GetReversalIndicator(f *field.ReversalIndicatorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MatchType is a non-required field for AllocationReport.
func (m AllocationReport) MatchType() (*field.MatchTypeField, errors.MessageRejectError) {
	f := &field.MatchTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMatchType reads a MatchType from AllocationReport.
func (m AllocationReport) GetMatchType(f *field.MatchTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Side is a required field for AllocationReport.
func (m AllocationReport) Side() (*field.SideField, errors.MessageRejectError) {
	f := &field.SideField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSide reads a Side from AllocationReport.
func (m AllocationReport) GetSide(f *field.SideField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for AllocationReport.
func (m AllocationReport) Symbol() (*field.SymbolField, errors.MessageRejectError) {
	f := &field.SymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from AllocationReport.
func (m AllocationReport) GetSymbol(f *field.SymbolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for AllocationReport.
func (m AllocationReport) SymbolSfx() (*field.SymbolSfxField, errors.MessageRejectError) {
	f := &field.SymbolSfxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from AllocationReport.
func (m AllocationReport) GetSymbolSfx(f *field.SymbolSfxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for AllocationReport.
func (m AllocationReport) SecurityID() (*field.SecurityIDField, errors.MessageRejectError) {
	f := &field.SecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from AllocationReport.
func (m AllocationReport) GetSecurityID(f *field.SecurityIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for AllocationReport.
func (m AllocationReport) SecurityIDSource() (*field.SecurityIDSourceField, errors.MessageRejectError) {
	f := &field.SecurityIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from AllocationReport.
func (m AllocationReport) GetSecurityIDSource(f *field.SecurityIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for AllocationReport.
func (m AllocationReport) NoSecurityAltID() (*field.NoSecurityAltIDField, errors.MessageRejectError) {
	f := &field.NoSecurityAltIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from AllocationReport.
func (m AllocationReport) GetNoSecurityAltID(f *field.NoSecurityAltIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for AllocationReport.
func (m AllocationReport) Product() (*field.ProductField, errors.MessageRejectError) {
	f := &field.ProductField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from AllocationReport.
func (m AllocationReport) GetProduct(f *field.ProductField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for AllocationReport.
func (m AllocationReport) CFICode() (*field.CFICodeField, errors.MessageRejectError) {
	f := &field.CFICodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from AllocationReport.
func (m AllocationReport) GetCFICode(f *field.CFICodeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for AllocationReport.
func (m AllocationReport) SecurityType() (*field.SecurityTypeField, errors.MessageRejectError) {
	f := &field.SecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from AllocationReport.
func (m AllocationReport) GetSecurityType(f *field.SecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySubType is a non-required field for AllocationReport.
func (m AllocationReport) SecuritySubType() (*field.SecuritySubTypeField, errors.MessageRejectError) {
	f := &field.SecuritySubTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySubType reads a SecuritySubType from AllocationReport.
func (m AllocationReport) GetSecuritySubType(f *field.SecuritySubTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for AllocationReport.
func (m AllocationReport) MaturityMonthYear() (*field.MaturityMonthYearField, errors.MessageRejectError) {
	f := &field.MaturityMonthYearField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from AllocationReport.
func (m AllocationReport) GetMaturityMonthYear(f *field.MaturityMonthYearField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for AllocationReport.
func (m AllocationReport) MaturityDate() (*field.MaturityDateField, errors.MessageRejectError) {
	f := &field.MaturityDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from AllocationReport.
func (m AllocationReport) GetMaturityDate(f *field.MaturityDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for AllocationReport.
func (m AllocationReport) CouponPaymentDate() (*field.CouponPaymentDateField, errors.MessageRejectError) {
	f := &field.CouponPaymentDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from AllocationReport.
func (m AllocationReport) GetCouponPaymentDate(f *field.CouponPaymentDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for AllocationReport.
func (m AllocationReport) IssueDate() (*field.IssueDateField, errors.MessageRejectError) {
	f := &field.IssueDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from AllocationReport.
func (m AllocationReport) GetIssueDate(f *field.IssueDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for AllocationReport.
func (m AllocationReport) RepoCollateralSecurityType() (*field.RepoCollateralSecurityTypeField, errors.MessageRejectError) {
	f := &field.RepoCollateralSecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from AllocationReport.
func (m AllocationReport) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for AllocationReport.
func (m AllocationReport) RepurchaseTerm() (*field.RepurchaseTermField, errors.MessageRejectError) {
	f := &field.RepurchaseTermField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from AllocationReport.
func (m AllocationReport) GetRepurchaseTerm(f *field.RepurchaseTermField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for AllocationReport.
func (m AllocationReport) RepurchaseRate() (*field.RepurchaseRateField, errors.MessageRejectError) {
	f := &field.RepurchaseRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from AllocationReport.
func (m AllocationReport) GetRepurchaseRate(f *field.RepurchaseRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for AllocationReport.
func (m AllocationReport) Factor() (*field.FactorField, errors.MessageRejectError) {
	f := &field.FactorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from AllocationReport.
func (m AllocationReport) GetFactor(f *field.FactorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for AllocationReport.
func (m AllocationReport) CreditRating() (*field.CreditRatingField, errors.MessageRejectError) {
	f := &field.CreditRatingField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from AllocationReport.
func (m AllocationReport) GetCreditRating(f *field.CreditRatingField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for AllocationReport.
func (m AllocationReport) InstrRegistry() (*field.InstrRegistryField, errors.MessageRejectError) {
	f := &field.InstrRegistryField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from AllocationReport.
func (m AllocationReport) GetInstrRegistry(f *field.InstrRegistryField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for AllocationReport.
func (m AllocationReport) CountryOfIssue() (*field.CountryOfIssueField, errors.MessageRejectError) {
	f := &field.CountryOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from AllocationReport.
func (m AllocationReport) GetCountryOfIssue(f *field.CountryOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for AllocationReport.
func (m AllocationReport) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssueField, errors.MessageRejectError) {
	f := &field.StateOrProvinceOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from AllocationReport.
func (m AllocationReport) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for AllocationReport.
func (m AllocationReport) LocaleOfIssue() (*field.LocaleOfIssueField, errors.MessageRejectError) {
	f := &field.LocaleOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from AllocationReport.
func (m AllocationReport) GetLocaleOfIssue(f *field.LocaleOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for AllocationReport.
func (m AllocationReport) RedemptionDate() (*field.RedemptionDateField, errors.MessageRejectError) {
	f := &field.RedemptionDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from AllocationReport.
func (m AllocationReport) GetRedemptionDate(f *field.RedemptionDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for AllocationReport.
func (m AllocationReport) StrikePrice() (*field.StrikePriceField, errors.MessageRejectError) {
	f := &field.StrikePriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from AllocationReport.
func (m AllocationReport) GetStrikePrice(f *field.StrikePriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeCurrency is a non-required field for AllocationReport.
func (m AllocationReport) StrikeCurrency() (*field.StrikeCurrencyField, errors.MessageRejectError) {
	f := &field.StrikeCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeCurrency reads a StrikeCurrency from AllocationReport.
func (m AllocationReport) GetStrikeCurrency(f *field.StrikeCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for AllocationReport.
func (m AllocationReport) OptAttribute() (*field.OptAttributeField, errors.MessageRejectError) {
	f := &field.OptAttributeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from AllocationReport.
func (m AllocationReport) GetOptAttribute(f *field.OptAttributeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for AllocationReport.
func (m AllocationReport) ContractMultiplier() (*field.ContractMultiplierField, errors.MessageRejectError) {
	f := &field.ContractMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from AllocationReport.
func (m AllocationReport) GetContractMultiplier(f *field.ContractMultiplierField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for AllocationReport.
func (m AllocationReport) CouponRate() (*field.CouponRateField, errors.MessageRejectError) {
	f := &field.CouponRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from AllocationReport.
func (m AllocationReport) GetCouponRate(f *field.CouponRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for AllocationReport.
func (m AllocationReport) SecurityExchange() (*field.SecurityExchangeField, errors.MessageRejectError) {
	f := &field.SecurityExchangeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from AllocationReport.
func (m AllocationReport) GetSecurityExchange(f *field.SecurityExchangeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for AllocationReport.
func (m AllocationReport) Issuer() (*field.IssuerField, errors.MessageRejectError) {
	f := &field.IssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from AllocationReport.
func (m AllocationReport) GetIssuer(f *field.IssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for AllocationReport.
func (m AllocationReport) EncodedIssuerLen() (*field.EncodedIssuerLenField, errors.MessageRejectError) {
	f := &field.EncodedIssuerLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from AllocationReport.
func (m AllocationReport) GetEncodedIssuerLen(f *field.EncodedIssuerLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for AllocationReport.
func (m AllocationReport) EncodedIssuer() (*field.EncodedIssuerField, errors.MessageRejectError) {
	f := &field.EncodedIssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from AllocationReport.
func (m AllocationReport) GetEncodedIssuer(f *field.EncodedIssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for AllocationReport.
func (m AllocationReport) SecurityDesc() (*field.SecurityDescField, errors.MessageRejectError) {
	f := &field.SecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from AllocationReport.
func (m AllocationReport) GetSecurityDesc(f *field.SecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for AllocationReport.
func (m AllocationReport) EncodedSecurityDescLen() (*field.EncodedSecurityDescLenField, errors.MessageRejectError) {
	f := &field.EncodedSecurityDescLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from AllocationReport.
func (m AllocationReport) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for AllocationReport.
func (m AllocationReport) EncodedSecurityDesc() (*field.EncodedSecurityDescField, errors.MessageRejectError) {
	f := &field.EncodedSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from AllocationReport.
func (m AllocationReport) GetEncodedSecurityDesc(f *field.EncodedSecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Pool is a non-required field for AllocationReport.
func (m AllocationReport) Pool() (*field.PoolField, errors.MessageRejectError) {
	f := &field.PoolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPool reads a Pool from AllocationReport.
func (m AllocationReport) GetPool(f *field.PoolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractSettlMonth is a non-required field for AllocationReport.
func (m AllocationReport) ContractSettlMonth() (*field.ContractSettlMonthField, errors.MessageRejectError) {
	f := &field.ContractSettlMonthField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractSettlMonth reads a ContractSettlMonth from AllocationReport.
func (m AllocationReport) GetContractSettlMonth(f *field.ContractSettlMonthField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPProgram is a non-required field for AllocationReport.
func (m AllocationReport) CPProgram() (*field.CPProgramField, errors.MessageRejectError) {
	f := &field.CPProgramField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCPProgram reads a CPProgram from AllocationReport.
func (m AllocationReport) GetCPProgram(f *field.CPProgramField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPRegType is a non-required field for AllocationReport.
func (m AllocationReport) CPRegType() (*field.CPRegTypeField, errors.MessageRejectError) {
	f := &field.CPRegTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCPRegType reads a CPRegType from AllocationReport.
func (m AllocationReport) GetCPRegType(f *field.CPRegTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoEvents is a non-required field for AllocationReport.
func (m AllocationReport) NoEvents() (*field.NoEventsField, errors.MessageRejectError) {
	f := &field.NoEventsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoEvents reads a NoEvents from AllocationReport.
func (m AllocationReport) GetNoEvents(f *field.NoEventsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DatedDate is a non-required field for AllocationReport.
func (m AllocationReport) DatedDate() (*field.DatedDateField, errors.MessageRejectError) {
	f := &field.DatedDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDatedDate reads a DatedDate from AllocationReport.
func (m AllocationReport) GetDatedDate(f *field.DatedDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InterestAccrualDate is a non-required field for AllocationReport.
func (m AllocationReport) InterestAccrualDate() (*field.InterestAccrualDateField, errors.MessageRejectError) {
	f := &field.InterestAccrualDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInterestAccrualDate reads a InterestAccrualDate from AllocationReport.
func (m AllocationReport) GetInterestAccrualDate(f *field.InterestAccrualDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityStatus is a non-required field for AllocationReport.
func (m AllocationReport) SecurityStatus() (*field.SecurityStatusField, errors.MessageRejectError) {
	f := &field.SecurityStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityStatus reads a SecurityStatus from AllocationReport.
func (m AllocationReport) GetSecurityStatus(f *field.SecurityStatusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettleOnOpenFlag is a non-required field for AllocationReport.
func (m AllocationReport) SettleOnOpenFlag() (*field.SettleOnOpenFlagField, errors.MessageRejectError) {
	f := &field.SettleOnOpenFlagField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettleOnOpenFlag reads a SettleOnOpenFlag from AllocationReport.
func (m AllocationReport) GetSettleOnOpenFlag(f *field.SettleOnOpenFlagField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrmtAssignmentMethod is a non-required field for AllocationReport.
func (m AllocationReport) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethodField, errors.MessageRejectError) {
	f := &field.InstrmtAssignmentMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInstrmtAssignmentMethod reads a InstrmtAssignmentMethod from AllocationReport.
func (m AllocationReport) GetInstrmtAssignmentMethod(f *field.InstrmtAssignmentMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeMultiplier is a non-required field for AllocationReport.
func (m AllocationReport) StrikeMultiplier() (*field.StrikeMultiplierField, errors.MessageRejectError) {
	f := &field.StrikeMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeMultiplier reads a StrikeMultiplier from AllocationReport.
func (m AllocationReport) GetStrikeMultiplier(f *field.StrikeMultiplierField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeValue is a non-required field for AllocationReport.
func (m AllocationReport) StrikeValue() (*field.StrikeValueField, errors.MessageRejectError) {
	f := &field.StrikeValueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeValue reads a StrikeValue from AllocationReport.
func (m AllocationReport) GetStrikeValue(f *field.StrikeValueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrement is a non-required field for AllocationReport.
func (m AllocationReport) MinPriceIncrement() (*field.MinPriceIncrementField, errors.MessageRejectError) {
	f := &field.MinPriceIncrementField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrement reads a MinPriceIncrement from AllocationReport.
func (m AllocationReport) GetMinPriceIncrement(f *field.MinPriceIncrementField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PositionLimit is a non-required field for AllocationReport.
func (m AllocationReport) PositionLimit() (*field.PositionLimitField, errors.MessageRejectError) {
	f := &field.PositionLimitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPositionLimit reads a PositionLimit from AllocationReport.
func (m AllocationReport) GetPositionLimit(f *field.PositionLimitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NTPositionLimit is a non-required field for AllocationReport.
func (m AllocationReport) NTPositionLimit() (*field.NTPositionLimitField, errors.MessageRejectError) {
	f := &field.NTPositionLimitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNTPositionLimit reads a NTPositionLimit from AllocationReport.
func (m AllocationReport) GetNTPositionLimit(f *field.NTPositionLimitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoInstrumentParties is a non-required field for AllocationReport.
func (m AllocationReport) NoInstrumentParties() (*field.NoInstrumentPartiesField, errors.MessageRejectError) {
	f := &field.NoInstrumentPartiesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoInstrumentParties reads a NoInstrumentParties from AllocationReport.
func (m AllocationReport) GetNoInstrumentParties(f *field.NoInstrumentPartiesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasure is a non-required field for AllocationReport.
func (m AllocationReport) UnitOfMeasure() (*field.UnitOfMeasureField, errors.MessageRejectError) {
	f := &field.UnitOfMeasureField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasure reads a UnitOfMeasure from AllocationReport.
func (m AllocationReport) GetUnitOfMeasure(f *field.UnitOfMeasureField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TimeUnit is a non-required field for AllocationReport.
func (m AllocationReport) TimeUnit() (*field.TimeUnitField, errors.MessageRejectError) {
	f := &field.TimeUnitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTimeUnit reads a TimeUnit from AllocationReport.
func (m AllocationReport) GetTimeUnit(f *field.TimeUnitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityTime is a non-required field for AllocationReport.
func (m AllocationReport) MaturityTime() (*field.MaturityTimeField, errors.MessageRejectError) {
	f := &field.MaturityTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityTime reads a MaturityTime from AllocationReport.
func (m AllocationReport) GetMaturityTime(f *field.MaturityTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityGroup is a non-required field for AllocationReport.
func (m AllocationReport) SecurityGroup() (*field.SecurityGroupField, errors.MessageRejectError) {
	f := &field.SecurityGroupField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityGroup reads a SecurityGroup from AllocationReport.
func (m AllocationReport) GetSecurityGroup(f *field.SecurityGroupField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrementAmount is a non-required field for AllocationReport.
func (m AllocationReport) MinPriceIncrementAmount() (*field.MinPriceIncrementAmountField, errors.MessageRejectError) {
	f := &field.MinPriceIncrementAmountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrementAmount reads a MinPriceIncrementAmount from AllocationReport.
func (m AllocationReport) GetMinPriceIncrementAmount(f *field.MinPriceIncrementAmountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasureQty is a non-required field for AllocationReport.
func (m AllocationReport) UnitOfMeasureQty() (*field.UnitOfMeasureQtyField, errors.MessageRejectError) {
	f := &field.UnitOfMeasureQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasureQty reads a UnitOfMeasureQty from AllocationReport.
func (m AllocationReport) GetUnitOfMeasureQty(f *field.UnitOfMeasureQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXMLLen is a non-required field for AllocationReport.
func (m AllocationReport) SecurityXMLLen() (*field.SecurityXMLLenField, errors.MessageRejectError) {
	f := &field.SecurityXMLLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXMLLen reads a SecurityXMLLen from AllocationReport.
func (m AllocationReport) GetSecurityXMLLen(f *field.SecurityXMLLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXML is a non-required field for AllocationReport.
func (m AllocationReport) SecurityXML() (*field.SecurityXMLField, errors.MessageRejectError) {
	f := &field.SecurityXMLField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXML reads a SecurityXML from AllocationReport.
func (m AllocationReport) GetSecurityXML(f *field.SecurityXMLField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXMLSchema is a non-required field for AllocationReport.
func (m AllocationReport) SecurityXMLSchema() (*field.SecurityXMLSchemaField, errors.MessageRejectError) {
	f := &field.SecurityXMLSchemaField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXMLSchema reads a SecurityXMLSchema from AllocationReport.
func (m AllocationReport) GetSecurityXMLSchema(f *field.SecurityXMLSchemaField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ProductComplex is a non-required field for AllocationReport.
func (m AllocationReport) ProductComplex() (*field.ProductComplexField, errors.MessageRejectError) {
	f := &field.ProductComplexField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProductComplex reads a ProductComplex from AllocationReport.
func (m AllocationReport) GetProductComplex(f *field.ProductComplexField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceUnitOfMeasure is a non-required field for AllocationReport.
func (m AllocationReport) PriceUnitOfMeasure() (*field.PriceUnitOfMeasureField, errors.MessageRejectError) {
	f := &field.PriceUnitOfMeasureField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceUnitOfMeasure reads a PriceUnitOfMeasure from AllocationReport.
func (m AllocationReport) GetPriceUnitOfMeasure(f *field.PriceUnitOfMeasureField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceUnitOfMeasureQty is a non-required field for AllocationReport.
func (m AllocationReport) PriceUnitOfMeasureQty() (*field.PriceUnitOfMeasureQtyField, errors.MessageRejectError) {
	f := &field.PriceUnitOfMeasureQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceUnitOfMeasureQty reads a PriceUnitOfMeasureQty from AllocationReport.
func (m AllocationReport) GetPriceUnitOfMeasureQty(f *field.PriceUnitOfMeasureQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlMethod is a non-required field for AllocationReport.
func (m AllocationReport) SettlMethod() (*field.SettlMethodField, errors.MessageRejectError) {
	f := &field.SettlMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlMethod reads a SettlMethod from AllocationReport.
func (m AllocationReport) GetSettlMethod(f *field.SettlMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExerciseStyle is a non-required field for AllocationReport.
func (m AllocationReport) ExerciseStyle() (*field.ExerciseStyleField, errors.MessageRejectError) {
	f := &field.ExerciseStyleField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExerciseStyle reads a ExerciseStyle from AllocationReport.
func (m AllocationReport) GetExerciseStyle(f *field.ExerciseStyleField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptPayAmount is a non-required field for AllocationReport.
func (m AllocationReport) OptPayAmount() (*field.OptPayAmountField, errors.MessageRejectError) {
	f := &field.OptPayAmountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptPayAmount reads a OptPayAmount from AllocationReport.
func (m AllocationReport) GetOptPayAmount(f *field.OptPayAmountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceQuoteMethod is a non-required field for AllocationReport.
func (m AllocationReport) PriceQuoteMethod() (*field.PriceQuoteMethodField, errors.MessageRejectError) {
	f := &field.PriceQuoteMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceQuoteMethod reads a PriceQuoteMethod from AllocationReport.
func (m AllocationReport) GetPriceQuoteMethod(f *field.PriceQuoteMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListMethod is a non-required field for AllocationReport.
func (m AllocationReport) ListMethod() (*field.ListMethodField, errors.MessageRejectError) {
	f := &field.ListMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetListMethod reads a ListMethod from AllocationReport.
func (m AllocationReport) GetListMethod(f *field.ListMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CapPrice is a non-required field for AllocationReport.
func (m AllocationReport) CapPrice() (*field.CapPriceField, errors.MessageRejectError) {
	f := &field.CapPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCapPrice reads a CapPrice from AllocationReport.
func (m AllocationReport) GetCapPrice(f *field.CapPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FloorPrice is a non-required field for AllocationReport.
func (m AllocationReport) FloorPrice() (*field.FloorPriceField, errors.MessageRejectError) {
	f := &field.FloorPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFloorPrice reads a FloorPrice from AllocationReport.
func (m AllocationReport) GetFloorPrice(f *field.FloorPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PutOrCall is a non-required field for AllocationReport.
func (m AllocationReport) PutOrCall() (*field.PutOrCallField, errors.MessageRejectError) {
	f := &field.PutOrCallField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPutOrCall reads a PutOrCall from AllocationReport.
func (m AllocationReport) GetPutOrCall(f *field.PutOrCallField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlexibleIndicator is a non-required field for AllocationReport.
func (m AllocationReport) FlexibleIndicator() (*field.FlexibleIndicatorField, errors.MessageRejectError) {
	f := &field.FlexibleIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFlexibleIndicator reads a FlexibleIndicator from AllocationReport.
func (m AllocationReport) GetFlexibleIndicator(f *field.FlexibleIndicatorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlexProductEligibilityIndicator is a non-required field for AllocationReport.
func (m AllocationReport) FlexProductEligibilityIndicator() (*field.FlexProductEligibilityIndicatorField, errors.MessageRejectError) {
	f := &field.FlexProductEligibilityIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFlexProductEligibilityIndicator reads a FlexProductEligibilityIndicator from AllocationReport.
func (m AllocationReport) GetFlexProductEligibilityIndicator(f *field.FlexProductEligibilityIndicatorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FuturesValuationMethod is a non-required field for AllocationReport.
func (m AllocationReport) FuturesValuationMethod() (*field.FuturesValuationMethodField, errors.MessageRejectError) {
	f := &field.FuturesValuationMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFuturesValuationMethod reads a FuturesValuationMethod from AllocationReport.
func (m AllocationReport) GetFuturesValuationMethod(f *field.FuturesValuationMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DeliveryForm is a non-required field for AllocationReport.
func (m AllocationReport) DeliveryForm() (*field.DeliveryFormField, errors.MessageRejectError) {
	f := &field.DeliveryFormField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDeliveryForm reads a DeliveryForm from AllocationReport.
func (m AllocationReport) GetDeliveryForm(f *field.DeliveryFormField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PctAtRisk is a non-required field for AllocationReport.
func (m AllocationReport) PctAtRisk() (*field.PctAtRiskField, errors.MessageRejectError) {
	f := &field.PctAtRiskField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPctAtRisk reads a PctAtRisk from AllocationReport.
func (m AllocationReport) GetPctAtRisk(f *field.PctAtRiskField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoInstrAttrib is a non-required field for AllocationReport.
func (m AllocationReport) NoInstrAttrib() (*field.NoInstrAttribField, errors.MessageRejectError) {
	f := &field.NoInstrAttribField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoInstrAttrib reads a NoInstrAttrib from AllocationReport.
func (m AllocationReport) GetNoInstrAttrib(f *field.NoInstrAttribField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementDesc is a non-required field for AllocationReport.
func (m AllocationReport) AgreementDesc() (*field.AgreementDescField, errors.MessageRejectError) {
	f := &field.AgreementDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementDesc reads a AgreementDesc from AllocationReport.
func (m AllocationReport) GetAgreementDesc(f *field.AgreementDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementID is a non-required field for AllocationReport.
func (m AllocationReport) AgreementID() (*field.AgreementIDField, errors.MessageRejectError) {
	f := &field.AgreementIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementID reads a AgreementID from AllocationReport.
func (m AllocationReport) GetAgreementID(f *field.AgreementIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementDate is a non-required field for AllocationReport.
func (m AllocationReport) AgreementDate() (*field.AgreementDateField, errors.MessageRejectError) {
	f := &field.AgreementDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementDate reads a AgreementDate from AllocationReport.
func (m AllocationReport) GetAgreementDate(f *field.AgreementDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementCurrency is a non-required field for AllocationReport.
func (m AllocationReport) AgreementCurrency() (*field.AgreementCurrencyField, errors.MessageRejectError) {
	f := &field.AgreementCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementCurrency reads a AgreementCurrency from AllocationReport.
func (m AllocationReport) GetAgreementCurrency(f *field.AgreementCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TerminationType is a non-required field for AllocationReport.
func (m AllocationReport) TerminationType() (*field.TerminationTypeField, errors.MessageRejectError) {
	f := &field.TerminationTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTerminationType reads a TerminationType from AllocationReport.
func (m AllocationReport) GetTerminationType(f *field.TerminationTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StartDate is a non-required field for AllocationReport.
func (m AllocationReport) StartDate() (*field.StartDateField, errors.MessageRejectError) {
	f := &field.StartDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStartDate reads a StartDate from AllocationReport.
func (m AllocationReport) GetStartDate(f *field.StartDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EndDate is a non-required field for AllocationReport.
func (m AllocationReport) EndDate() (*field.EndDateField, errors.MessageRejectError) {
	f := &field.EndDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEndDate reads a EndDate from AllocationReport.
func (m AllocationReport) GetEndDate(f *field.EndDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DeliveryType is a non-required field for AllocationReport.
func (m AllocationReport) DeliveryType() (*field.DeliveryTypeField, errors.MessageRejectError) {
	f := &field.DeliveryTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDeliveryType reads a DeliveryType from AllocationReport.
func (m AllocationReport) GetDeliveryType(f *field.DeliveryTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarginRatio is a non-required field for AllocationReport.
func (m AllocationReport) MarginRatio() (*field.MarginRatioField, errors.MessageRejectError) {
	f := &field.MarginRatioField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarginRatio reads a MarginRatio from AllocationReport.
func (m AllocationReport) GetMarginRatio(f *field.MarginRatioField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyings is a non-required field for AllocationReport.
func (m AllocationReport) NoUnderlyings() (*field.NoUnderlyingsField, errors.MessageRejectError) {
	f := &field.NoUnderlyingsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyings reads a NoUnderlyings from AllocationReport.
func (m AllocationReport) GetNoUnderlyings(f *field.NoUnderlyingsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoLegs is a non-required field for AllocationReport.
func (m AllocationReport) NoLegs() (*field.NoLegsField, errors.MessageRejectError) {
	f := &field.NoLegsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoLegs reads a NoLegs from AllocationReport.
func (m AllocationReport) GetNoLegs(f *field.NoLegsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Quantity is a required field for AllocationReport.
func (m AllocationReport) Quantity() (*field.QuantityField, errors.MessageRejectError) {
	f := &field.QuantityField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuantity reads a Quantity from AllocationReport.
func (m AllocationReport) GetQuantity(f *field.QuantityField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QtyType is a non-required field for AllocationReport.
func (m AllocationReport) QtyType() (*field.QtyTypeField, errors.MessageRejectError) {
	f := &field.QtyTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQtyType reads a QtyType from AllocationReport.
func (m AllocationReport) GetQtyType(f *field.QtyTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastMkt is a non-required field for AllocationReport.
func (m AllocationReport) LastMkt() (*field.LastMktField, errors.MessageRejectError) {
	f := &field.LastMktField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastMkt reads a LastMkt from AllocationReport.
func (m AllocationReport) GetLastMkt(f *field.LastMktField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeOriginationDate is a non-required field for AllocationReport.
func (m AllocationReport) TradeOriginationDate() (*field.TradeOriginationDateField, errors.MessageRejectError) {
	f := &field.TradeOriginationDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeOriginationDate reads a TradeOriginationDate from AllocationReport.
func (m AllocationReport) GetTradeOriginationDate(f *field.TradeOriginationDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionID is a non-required field for AllocationReport.
func (m AllocationReport) TradingSessionID() (*field.TradingSessionIDField, errors.MessageRejectError) {
	f := &field.TradingSessionIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionID reads a TradingSessionID from AllocationReport.
func (m AllocationReport) GetTradingSessionID(f *field.TradingSessionIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionSubID is a non-required field for AllocationReport.
func (m AllocationReport) TradingSessionSubID() (*field.TradingSessionSubIDField, errors.MessageRejectError) {
	f := &field.TradingSessionSubIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionSubID reads a TradingSessionSubID from AllocationReport.
func (m AllocationReport) GetTradingSessionSubID(f *field.TradingSessionSubIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceType is a non-required field for AllocationReport.
func (m AllocationReport) PriceType() (*field.PriceTypeField, errors.MessageRejectError) {
	f := &field.PriceTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceType reads a PriceType from AllocationReport.
func (m AllocationReport) GetPriceType(f *field.PriceTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AvgPx is a required field for AllocationReport.
func (m AllocationReport) AvgPx() (*field.AvgPxField, errors.MessageRejectError) {
	f := &field.AvgPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAvgPx reads a AvgPx from AllocationReport.
func (m AllocationReport) GetAvgPx(f *field.AvgPxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AvgParPx is a non-required field for AllocationReport.
func (m AllocationReport) AvgParPx() (*field.AvgParPxField, errors.MessageRejectError) {
	f := &field.AvgParPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAvgParPx reads a AvgParPx from AllocationReport.
func (m AllocationReport) GetAvgParPx(f *field.AvgParPxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Spread is a non-required field for AllocationReport.
func (m AllocationReport) Spread() (*field.SpreadField, errors.MessageRejectError) {
	f := &field.SpreadField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSpread reads a Spread from AllocationReport.
func (m AllocationReport) GetSpread(f *field.SpreadField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurveCurrency is a non-required field for AllocationReport.
func (m AllocationReport) BenchmarkCurveCurrency() (*field.BenchmarkCurveCurrencyField, errors.MessageRejectError) {
	f := &field.BenchmarkCurveCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurveCurrency reads a BenchmarkCurveCurrency from AllocationReport.
func (m AllocationReport) GetBenchmarkCurveCurrency(f *field.BenchmarkCurveCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurveName is a non-required field for AllocationReport.
func (m AllocationReport) BenchmarkCurveName() (*field.BenchmarkCurveNameField, errors.MessageRejectError) {
	f := &field.BenchmarkCurveNameField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurveName reads a BenchmarkCurveName from AllocationReport.
func (m AllocationReport) GetBenchmarkCurveName(f *field.BenchmarkCurveNameField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurvePoint is a non-required field for AllocationReport.
func (m AllocationReport) BenchmarkCurvePoint() (*field.BenchmarkCurvePointField, errors.MessageRejectError) {
	f := &field.BenchmarkCurvePointField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurvePoint reads a BenchmarkCurvePoint from AllocationReport.
func (m AllocationReport) GetBenchmarkCurvePoint(f *field.BenchmarkCurvePointField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkPrice is a non-required field for AllocationReport.
func (m AllocationReport) BenchmarkPrice() (*field.BenchmarkPriceField, errors.MessageRejectError) {
	f := &field.BenchmarkPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkPrice reads a BenchmarkPrice from AllocationReport.
func (m AllocationReport) GetBenchmarkPrice(f *field.BenchmarkPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkPriceType is a non-required field for AllocationReport.
func (m AllocationReport) BenchmarkPriceType() (*field.BenchmarkPriceTypeField, errors.MessageRejectError) {
	f := &field.BenchmarkPriceTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkPriceType reads a BenchmarkPriceType from AllocationReport.
func (m AllocationReport) GetBenchmarkPriceType(f *field.BenchmarkPriceTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkSecurityID is a non-required field for AllocationReport.
func (m AllocationReport) BenchmarkSecurityID() (*field.BenchmarkSecurityIDField, errors.MessageRejectError) {
	f := &field.BenchmarkSecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkSecurityID reads a BenchmarkSecurityID from AllocationReport.
func (m AllocationReport) GetBenchmarkSecurityID(f *field.BenchmarkSecurityIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkSecurityIDSource is a non-required field for AllocationReport.
func (m AllocationReport) BenchmarkSecurityIDSource() (*field.BenchmarkSecurityIDSourceField, errors.MessageRejectError) {
	f := &field.BenchmarkSecurityIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkSecurityIDSource reads a BenchmarkSecurityIDSource from AllocationReport.
func (m AllocationReport) GetBenchmarkSecurityIDSource(f *field.BenchmarkSecurityIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for AllocationReport.
func (m AllocationReport) Currency() (*field.CurrencyField, errors.MessageRejectError) {
	f := &field.CurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from AllocationReport.
func (m AllocationReport) GetCurrency(f *field.CurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AvgPxPrecision is a non-required field for AllocationReport.
func (m AllocationReport) AvgPxPrecision() (*field.AvgPxPrecisionField, errors.MessageRejectError) {
	f := &field.AvgPxPrecisionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAvgPxPrecision reads a AvgPxPrecision from AllocationReport.
func (m AllocationReport) GetAvgPxPrecision(f *field.AvgPxPrecisionField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for AllocationReport.
func (m AllocationReport) NoPartyIDs() (*field.NoPartyIDsField, errors.MessageRejectError) {
	f := &field.NoPartyIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from AllocationReport.
func (m AllocationReport) GetNoPartyIDs(f *field.NoPartyIDsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeDate is a required field for AllocationReport.
func (m AllocationReport) TradeDate() (*field.TradeDateField, errors.MessageRejectError) {
	f := &field.TradeDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeDate reads a TradeDate from AllocationReport.
func (m AllocationReport) GetTradeDate(f *field.TradeDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a non-required field for AllocationReport.
func (m AllocationReport) TransactTime() (*field.TransactTimeField, errors.MessageRejectError) {
	f := &field.TransactTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from AllocationReport.
func (m AllocationReport) GetTransactTime(f *field.TransactTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlType is a non-required field for AllocationReport.
func (m AllocationReport) SettlType() (*field.SettlTypeField, errors.MessageRejectError) {
	f := &field.SettlTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlType reads a SettlType from AllocationReport.
func (m AllocationReport) GetSettlType(f *field.SettlTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlDate is a non-required field for AllocationReport.
func (m AllocationReport) SettlDate() (*field.SettlDateField, errors.MessageRejectError) {
	f := &field.SettlDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlDate reads a SettlDate from AllocationReport.
func (m AllocationReport) GetSettlDate(f *field.SettlDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BookingType is a non-required field for AllocationReport.
func (m AllocationReport) BookingType() (*field.BookingTypeField, errors.MessageRejectError) {
	f := &field.BookingTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBookingType reads a BookingType from AllocationReport.
func (m AllocationReport) GetBookingType(f *field.BookingTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//GrossTradeAmt is a non-required field for AllocationReport.
func (m AllocationReport) GrossTradeAmt() (*field.GrossTradeAmtField, errors.MessageRejectError) {
	f := &field.GrossTradeAmtField{}
	err := m.Body.Get(f)
	return f, err
}

//GetGrossTradeAmt reads a GrossTradeAmt from AllocationReport.
func (m AllocationReport) GetGrossTradeAmt(f *field.GrossTradeAmtField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Concession is a non-required field for AllocationReport.
func (m AllocationReport) Concession() (*field.ConcessionField, errors.MessageRejectError) {
	f := &field.ConcessionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetConcession reads a Concession from AllocationReport.
func (m AllocationReport) GetConcession(f *field.ConcessionField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotalTakedown is a non-required field for AllocationReport.
func (m AllocationReport) TotalTakedown() (*field.TotalTakedownField, errors.MessageRejectError) {
	f := &field.TotalTakedownField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTotalTakedown reads a TotalTakedown from AllocationReport.
func (m AllocationReport) GetTotalTakedown(f *field.TotalTakedownField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NetMoney is a non-required field for AllocationReport.
func (m AllocationReport) NetMoney() (*field.NetMoneyField, errors.MessageRejectError) {
	f := &field.NetMoneyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNetMoney reads a NetMoney from AllocationReport.
func (m AllocationReport) GetNetMoney(f *field.NetMoneyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PositionEffect is a non-required field for AllocationReport.
func (m AllocationReport) PositionEffect() (*field.PositionEffectField, errors.MessageRejectError) {
	f := &field.PositionEffectField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPositionEffect reads a PositionEffect from AllocationReport.
func (m AllocationReport) GetPositionEffect(f *field.PositionEffectField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AutoAcceptIndicator is a non-required field for AllocationReport.
func (m AllocationReport) AutoAcceptIndicator() (*field.AutoAcceptIndicatorField, errors.MessageRejectError) {
	f := &field.AutoAcceptIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAutoAcceptIndicator reads a AutoAcceptIndicator from AllocationReport.
func (m AllocationReport) GetAutoAcceptIndicator(f *field.AutoAcceptIndicatorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for AllocationReport.
func (m AllocationReport) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from AllocationReport.
func (m AllocationReport) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for AllocationReport.
func (m AllocationReport) EncodedTextLen() (*field.EncodedTextLenField, errors.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from AllocationReport.
func (m AllocationReport) GetEncodedTextLen(f *field.EncodedTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for AllocationReport.
func (m AllocationReport) EncodedText() (*field.EncodedTextField, errors.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from AllocationReport.
func (m AllocationReport) GetEncodedText(f *field.EncodedTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NumDaysInterest is a non-required field for AllocationReport.
func (m AllocationReport) NumDaysInterest() (*field.NumDaysInterestField, errors.MessageRejectError) {
	f := &field.NumDaysInterestField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNumDaysInterest reads a NumDaysInterest from AllocationReport.
func (m AllocationReport) GetNumDaysInterest(f *field.NumDaysInterestField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AccruedInterestRate is a non-required field for AllocationReport.
func (m AllocationReport) AccruedInterestRate() (*field.AccruedInterestRateField, errors.MessageRejectError) {
	f := &field.AccruedInterestRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccruedInterestRate reads a AccruedInterestRate from AllocationReport.
func (m AllocationReport) GetAccruedInterestRate(f *field.AccruedInterestRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AccruedInterestAmt is a non-required field for AllocationReport.
func (m AllocationReport) AccruedInterestAmt() (*field.AccruedInterestAmtField, errors.MessageRejectError) {
	f := &field.AccruedInterestAmtField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccruedInterestAmt reads a AccruedInterestAmt from AllocationReport.
func (m AllocationReport) GetAccruedInterestAmt(f *field.AccruedInterestAmtField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotalAccruedInterestAmt is a non-required field for AllocationReport.
func (m AllocationReport) TotalAccruedInterestAmt() (*field.TotalAccruedInterestAmtField, errors.MessageRejectError) {
	f := &field.TotalAccruedInterestAmtField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTotalAccruedInterestAmt reads a TotalAccruedInterestAmt from AllocationReport.
func (m AllocationReport) GetTotalAccruedInterestAmt(f *field.TotalAccruedInterestAmtField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InterestAtMaturity is a non-required field for AllocationReport.
func (m AllocationReport) InterestAtMaturity() (*field.InterestAtMaturityField, errors.MessageRejectError) {
	f := &field.InterestAtMaturityField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInterestAtMaturity reads a InterestAtMaturity from AllocationReport.
func (m AllocationReport) GetInterestAtMaturity(f *field.InterestAtMaturityField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EndAccruedInterestAmt is a non-required field for AllocationReport.
func (m AllocationReport) EndAccruedInterestAmt() (*field.EndAccruedInterestAmtField, errors.MessageRejectError) {
	f := &field.EndAccruedInterestAmtField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEndAccruedInterestAmt reads a EndAccruedInterestAmt from AllocationReport.
func (m AllocationReport) GetEndAccruedInterestAmt(f *field.EndAccruedInterestAmtField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StartCash is a non-required field for AllocationReport.
func (m AllocationReport) StartCash() (*field.StartCashField, errors.MessageRejectError) {
	f := &field.StartCashField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStartCash reads a StartCash from AllocationReport.
func (m AllocationReport) GetStartCash(f *field.StartCashField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EndCash is a non-required field for AllocationReport.
func (m AllocationReport) EndCash() (*field.EndCashField, errors.MessageRejectError) {
	f := &field.EndCashField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEndCash reads a EndCash from AllocationReport.
func (m AllocationReport) GetEndCash(f *field.EndCashField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LegalConfirm is a non-required field for AllocationReport.
func (m AllocationReport) LegalConfirm() (*field.LegalConfirmField, errors.MessageRejectError) {
	f := &field.LegalConfirmField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLegalConfirm reads a LegalConfirm from AllocationReport.
func (m AllocationReport) GetLegalConfirm(f *field.LegalConfirmField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoStipulations is a non-required field for AllocationReport.
func (m AllocationReport) NoStipulations() (*field.NoStipulationsField, errors.MessageRejectError) {
	f := &field.NoStipulationsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoStipulations reads a NoStipulations from AllocationReport.
func (m AllocationReport) GetNoStipulations(f *field.NoStipulationsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldType is a non-required field for AllocationReport.
func (m AllocationReport) YieldType() (*field.YieldTypeField, errors.MessageRejectError) {
	f := &field.YieldTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYieldType reads a YieldType from AllocationReport.
func (m AllocationReport) GetYieldType(f *field.YieldTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Yield is a non-required field for AllocationReport.
func (m AllocationReport) Yield() (*field.YieldField, errors.MessageRejectError) {
	f := &field.YieldField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYield reads a Yield from AllocationReport.
func (m AllocationReport) GetYield(f *field.YieldField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldCalcDate is a non-required field for AllocationReport.
func (m AllocationReport) YieldCalcDate() (*field.YieldCalcDateField, errors.MessageRejectError) {
	f := &field.YieldCalcDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYieldCalcDate reads a YieldCalcDate from AllocationReport.
func (m AllocationReport) GetYieldCalcDate(f *field.YieldCalcDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldRedemptionDate is a non-required field for AllocationReport.
func (m AllocationReport) YieldRedemptionDate() (*field.YieldRedemptionDateField, errors.MessageRejectError) {
	f := &field.YieldRedemptionDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYieldRedemptionDate reads a YieldRedemptionDate from AllocationReport.
func (m AllocationReport) GetYieldRedemptionDate(f *field.YieldRedemptionDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldRedemptionPrice is a non-required field for AllocationReport.
func (m AllocationReport) YieldRedemptionPrice() (*field.YieldRedemptionPriceField, errors.MessageRejectError) {
	f := &field.YieldRedemptionPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYieldRedemptionPrice reads a YieldRedemptionPrice from AllocationReport.
func (m AllocationReport) GetYieldRedemptionPrice(f *field.YieldRedemptionPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldRedemptionPriceType is a non-required field for AllocationReport.
func (m AllocationReport) YieldRedemptionPriceType() (*field.YieldRedemptionPriceTypeField, errors.MessageRejectError) {
	f := &field.YieldRedemptionPriceTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYieldRedemptionPriceType reads a YieldRedemptionPriceType from AllocationReport.
func (m AllocationReport) GetYieldRedemptionPriceType(f *field.YieldRedemptionPriceTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotNoAllocs is a non-required field for AllocationReport.
func (m AllocationReport) TotNoAllocs() (*field.TotNoAllocsField, errors.MessageRejectError) {
	f := &field.TotNoAllocsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTotNoAllocs reads a TotNoAllocs from AllocationReport.
func (m AllocationReport) GetTotNoAllocs(f *field.TotNoAllocsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastFragment is a non-required field for AllocationReport.
func (m AllocationReport) LastFragment() (*field.LastFragmentField, errors.MessageRejectError) {
	f := &field.LastFragmentField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastFragment reads a LastFragment from AllocationReport.
func (m AllocationReport) GetLastFragment(f *field.LastFragmentField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoAllocs is a non-required field for AllocationReport.
func (m AllocationReport) NoAllocs() (*field.NoAllocsField, errors.MessageRejectError) {
	f := &field.NoAllocsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoAllocs reads a NoAllocs from AllocationReport.
func (m AllocationReport) GetNoAllocs(f *field.NoAllocsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClearingBusinessDate is a non-required field for AllocationReport.
func (m AllocationReport) ClearingBusinessDate() (*field.ClearingBusinessDateField, errors.MessageRejectError) {
	f := &field.ClearingBusinessDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClearingBusinessDate reads a ClearingBusinessDate from AllocationReport.
func (m AllocationReport) GetClearingBusinessDate(f *field.ClearingBusinessDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TrdType is a non-required field for AllocationReport.
func (m AllocationReport) TrdType() (*field.TrdTypeField, errors.MessageRejectError) {
	f := &field.TrdTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTrdType reads a TrdType from AllocationReport.
func (m AllocationReport) GetTrdType(f *field.TrdTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TrdSubType is a non-required field for AllocationReport.
func (m AllocationReport) TrdSubType() (*field.TrdSubTypeField, errors.MessageRejectError) {
	f := &field.TrdSubTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTrdSubType reads a TrdSubType from AllocationReport.
func (m AllocationReport) GetTrdSubType(f *field.TrdSubTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MultiLegReportingType is a non-required field for AllocationReport.
func (m AllocationReport) MultiLegReportingType() (*field.MultiLegReportingTypeField, errors.MessageRejectError) {
	f := &field.MultiLegReportingTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMultiLegReportingType reads a MultiLegReportingType from AllocationReport.
func (m AllocationReport) GetMultiLegReportingType(f *field.MultiLegReportingTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CustOrderCapacity is a non-required field for AllocationReport.
func (m AllocationReport) CustOrderCapacity() (*field.CustOrderCapacityField, errors.MessageRejectError) {
	f := &field.CustOrderCapacityField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCustOrderCapacity reads a CustOrderCapacity from AllocationReport.
func (m AllocationReport) GetCustOrderCapacity(f *field.CustOrderCapacityField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeInputSource is a non-required field for AllocationReport.
func (m AllocationReport) TradeInputSource() (*field.TradeInputSourceField, errors.MessageRejectError) {
	f := &field.TradeInputSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeInputSource reads a TradeInputSource from AllocationReport.
func (m AllocationReport) GetTradeInputSource(f *field.TradeInputSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RndPx is a non-required field for AllocationReport.
func (m AllocationReport) RndPx() (*field.RndPxField, errors.MessageRejectError) {
	f := &field.RndPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRndPx reads a RndPx from AllocationReport.
func (m AllocationReport) GetRndPx(f *field.RndPxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MessageEventSource is a non-required field for AllocationReport.
func (m AllocationReport) MessageEventSource() (*field.MessageEventSourceField, errors.MessageRejectError) {
	f := &field.MessageEventSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMessageEventSource reads a MessageEventSource from AllocationReport.
func (m AllocationReport) GetMessageEventSource(f *field.MessageEventSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeInputDevice is a non-required field for AllocationReport.
func (m AllocationReport) TradeInputDevice() (*field.TradeInputDeviceField, errors.MessageRejectError) {
	f := &field.TradeInputDeviceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeInputDevice reads a TradeInputDevice from AllocationReport.
func (m AllocationReport) GetTradeInputDevice(f *field.TradeInputDeviceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AvgPxIndicator is a non-required field for AllocationReport.
func (m AllocationReport) AvgPxIndicator() (*field.AvgPxIndicatorField, errors.MessageRejectError) {
	f := &field.AvgPxIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAvgPxIndicator reads a AvgPxIndicator from AllocationReport.
func (m AllocationReport) GetAvgPxIndicator(f *field.AvgPxIndicatorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPosAmt is a non-required field for AllocationReport.
func (m AllocationReport) NoPosAmt() (*field.NoPosAmtField, errors.MessageRejectError) {
	f := &field.NoPosAmtField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPosAmt reads a NoPosAmt from AllocationReport.
func (m AllocationReport) GetNoPosAmt(f *field.NoPosAmtField) errors.MessageRejectError {
	return m.Body.Get(f)
}
