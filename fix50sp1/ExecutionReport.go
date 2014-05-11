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

//ExecutionReport msg type = 8.
type ExecutionReport struct {
	message.Message
}

//ExecutionReportBuilder builds ExecutionReport messages.
type ExecutionReportBuilder struct {
	message.MessageBuilder
}

//CreateExecutionReportBuilder returns an initialized ExecutionReportBuilder with specified required fields.
func CreateExecutionReportBuilder(
	orderid *field.OrderIDField,
	execid *field.ExecIDField,
	exectype *field.ExecTypeField,
	ordstatus *field.OrdStatusField,
	side *field.SideField,
	leavesqty *field.LeavesQtyField,
	cumqty *field.CumQtyField) ExecutionReportBuilder {
	var builder ExecutionReportBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header().Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP1))
	builder.Header().Set(field.NewMsgType("8"))
	builder.Body().Set(orderid)
	builder.Body().Set(execid)
	builder.Body().Set(exectype)
	builder.Body().Set(ordstatus)
	builder.Body().Set(side)
	builder.Body().Set(leavesqty)
	builder.Body().Set(cumqty)
	return builder
}

//OrderID is a required field for ExecutionReport.
func (m ExecutionReport) OrderID() (*field.OrderIDField, errors.MessageRejectError) {
	f := &field.OrderIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderID reads a OrderID from ExecutionReport.
func (m ExecutionReport) GetOrderID(f *field.OrderIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryOrderID is a non-required field for ExecutionReport.
func (m ExecutionReport) SecondaryOrderID() (*field.SecondaryOrderIDField, errors.MessageRejectError) {
	f := &field.SecondaryOrderIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryOrderID reads a SecondaryOrderID from ExecutionReport.
func (m ExecutionReport) GetSecondaryOrderID(f *field.SecondaryOrderIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryClOrdID is a non-required field for ExecutionReport.
func (m ExecutionReport) SecondaryClOrdID() (*field.SecondaryClOrdIDField, errors.MessageRejectError) {
	f := &field.SecondaryClOrdIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryClOrdID reads a SecondaryClOrdID from ExecutionReport.
func (m ExecutionReport) GetSecondaryClOrdID(f *field.SecondaryClOrdIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryExecID is a non-required field for ExecutionReport.
func (m ExecutionReport) SecondaryExecID() (*field.SecondaryExecIDField, errors.MessageRejectError) {
	f := &field.SecondaryExecIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryExecID reads a SecondaryExecID from ExecutionReport.
func (m ExecutionReport) GetSecondaryExecID(f *field.SecondaryExecIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClOrdID is a non-required field for ExecutionReport.
func (m ExecutionReport) ClOrdID() (*field.ClOrdIDField, errors.MessageRejectError) {
	f := &field.ClOrdIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdID reads a ClOrdID from ExecutionReport.
func (m ExecutionReport) GetClOrdID(f *field.ClOrdIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrigClOrdID is a non-required field for ExecutionReport.
func (m ExecutionReport) OrigClOrdID() (*field.OrigClOrdIDField, errors.MessageRejectError) {
	f := &field.OrigClOrdIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrigClOrdID reads a OrigClOrdID from ExecutionReport.
func (m ExecutionReport) GetOrigClOrdID(f *field.OrigClOrdIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClOrdLinkID is a non-required field for ExecutionReport.
func (m ExecutionReport) ClOrdLinkID() (*field.ClOrdLinkIDField, errors.MessageRejectError) {
	f := &field.ClOrdLinkIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdLinkID reads a ClOrdLinkID from ExecutionReport.
func (m ExecutionReport) GetClOrdLinkID(f *field.ClOrdLinkIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteRespID is a non-required field for ExecutionReport.
func (m ExecutionReport) QuoteRespID() (*field.QuoteRespIDField, errors.MessageRejectError) {
	f := &field.QuoteRespIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteRespID reads a QuoteRespID from ExecutionReport.
func (m ExecutionReport) GetQuoteRespID(f *field.QuoteRespIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrdStatusReqID is a non-required field for ExecutionReport.
func (m ExecutionReport) OrdStatusReqID() (*field.OrdStatusReqIDField, errors.MessageRejectError) {
	f := &field.OrdStatusReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrdStatusReqID reads a OrdStatusReqID from ExecutionReport.
func (m ExecutionReport) GetOrdStatusReqID(f *field.OrdStatusReqIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MassStatusReqID is a non-required field for ExecutionReport.
func (m ExecutionReport) MassStatusReqID() (*field.MassStatusReqIDField, errors.MessageRejectError) {
	f := &field.MassStatusReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMassStatusReqID reads a MassStatusReqID from ExecutionReport.
func (m ExecutionReport) GetMassStatusReqID(f *field.MassStatusReqIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotNumReports is a non-required field for ExecutionReport.
func (m ExecutionReport) TotNumReports() (*field.TotNumReportsField, errors.MessageRejectError) {
	f := &field.TotNumReportsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTotNumReports reads a TotNumReports from ExecutionReport.
func (m ExecutionReport) GetTotNumReports(f *field.TotNumReportsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastRptRequested is a non-required field for ExecutionReport.
func (m ExecutionReport) LastRptRequested() (*field.LastRptRequestedField, errors.MessageRejectError) {
	f := &field.LastRptRequestedField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastRptRequested reads a LastRptRequested from ExecutionReport.
func (m ExecutionReport) GetLastRptRequested(f *field.LastRptRequestedField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for ExecutionReport.
func (m ExecutionReport) NoPartyIDs() (*field.NoPartyIDsField, errors.MessageRejectError) {
	f := &field.NoPartyIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from ExecutionReport.
func (m ExecutionReport) GetNoPartyIDs(f *field.NoPartyIDsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeOriginationDate is a non-required field for ExecutionReport.
func (m ExecutionReport) TradeOriginationDate() (*field.TradeOriginationDateField, errors.MessageRejectError) {
	f := &field.TradeOriginationDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeOriginationDate reads a TradeOriginationDate from ExecutionReport.
func (m ExecutionReport) GetTradeOriginationDate(f *field.TradeOriginationDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoContraBrokers is a non-required field for ExecutionReport.
func (m ExecutionReport) NoContraBrokers() (*field.NoContraBrokersField, errors.MessageRejectError) {
	f := &field.NoContraBrokersField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoContraBrokers reads a NoContraBrokers from ExecutionReport.
func (m ExecutionReport) GetNoContraBrokers(f *field.NoContraBrokersField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListID is a non-required field for ExecutionReport.
func (m ExecutionReport) ListID() (*field.ListIDField, errors.MessageRejectError) {
	f := &field.ListIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetListID reads a ListID from ExecutionReport.
func (m ExecutionReport) GetListID(f *field.ListIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CrossID is a non-required field for ExecutionReport.
func (m ExecutionReport) CrossID() (*field.CrossIDField, errors.MessageRejectError) {
	f := &field.CrossIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCrossID reads a CrossID from ExecutionReport.
func (m ExecutionReport) GetCrossID(f *field.CrossIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrigCrossID is a non-required field for ExecutionReport.
func (m ExecutionReport) OrigCrossID() (*field.OrigCrossIDField, errors.MessageRejectError) {
	f := &field.OrigCrossIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrigCrossID reads a OrigCrossID from ExecutionReport.
func (m ExecutionReport) GetOrigCrossID(f *field.OrigCrossIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CrossType is a non-required field for ExecutionReport.
func (m ExecutionReport) CrossType() (*field.CrossTypeField, errors.MessageRejectError) {
	f := &field.CrossTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCrossType reads a CrossType from ExecutionReport.
func (m ExecutionReport) GetCrossType(f *field.CrossTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecID is a required field for ExecutionReport.
func (m ExecutionReport) ExecID() (*field.ExecIDField, errors.MessageRejectError) {
	f := &field.ExecIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExecID reads a ExecID from ExecutionReport.
func (m ExecutionReport) GetExecID(f *field.ExecIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecRefID is a non-required field for ExecutionReport.
func (m ExecutionReport) ExecRefID() (*field.ExecRefIDField, errors.MessageRejectError) {
	f := &field.ExecRefIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExecRefID reads a ExecRefID from ExecutionReport.
func (m ExecutionReport) GetExecRefID(f *field.ExecRefIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecType is a required field for ExecutionReport.
func (m ExecutionReport) ExecType() (*field.ExecTypeField, errors.MessageRejectError) {
	f := &field.ExecTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExecType reads a ExecType from ExecutionReport.
func (m ExecutionReport) GetExecType(f *field.ExecTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrdStatus is a required field for ExecutionReport.
func (m ExecutionReport) OrdStatus() (*field.OrdStatusField, errors.MessageRejectError) {
	f := &field.OrdStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrdStatus reads a OrdStatus from ExecutionReport.
func (m ExecutionReport) GetOrdStatus(f *field.OrdStatusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//WorkingIndicator is a non-required field for ExecutionReport.
func (m ExecutionReport) WorkingIndicator() (*field.WorkingIndicatorField, errors.MessageRejectError) {
	f := &field.WorkingIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetWorkingIndicator reads a WorkingIndicator from ExecutionReport.
func (m ExecutionReport) GetWorkingIndicator(f *field.WorkingIndicatorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrdRejReason is a non-required field for ExecutionReport.
func (m ExecutionReport) OrdRejReason() (*field.OrdRejReasonField, errors.MessageRejectError) {
	f := &field.OrdRejReasonField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrdRejReason reads a OrdRejReason from ExecutionReport.
func (m ExecutionReport) GetOrdRejReason(f *field.OrdRejReasonField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecRestatementReason is a non-required field for ExecutionReport.
func (m ExecutionReport) ExecRestatementReason() (*field.ExecRestatementReasonField, errors.MessageRejectError) {
	f := &field.ExecRestatementReasonField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExecRestatementReason reads a ExecRestatementReason from ExecutionReport.
func (m ExecutionReport) GetExecRestatementReason(f *field.ExecRestatementReasonField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Account is a non-required field for ExecutionReport.
func (m ExecutionReport) Account() (*field.AccountField, errors.MessageRejectError) {
	f := &field.AccountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccount reads a Account from ExecutionReport.
func (m ExecutionReport) GetAccount(f *field.AccountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AcctIDSource is a non-required field for ExecutionReport.
func (m ExecutionReport) AcctIDSource() (*field.AcctIDSourceField, errors.MessageRejectError) {
	f := &field.AcctIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAcctIDSource reads a AcctIDSource from ExecutionReport.
func (m ExecutionReport) GetAcctIDSource(f *field.AcctIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AccountType is a non-required field for ExecutionReport.
func (m ExecutionReport) AccountType() (*field.AccountTypeField, errors.MessageRejectError) {
	f := &field.AccountTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccountType reads a AccountType from ExecutionReport.
func (m ExecutionReport) GetAccountType(f *field.AccountTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DayBookingInst is a non-required field for ExecutionReport.
func (m ExecutionReport) DayBookingInst() (*field.DayBookingInstField, errors.MessageRejectError) {
	f := &field.DayBookingInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDayBookingInst reads a DayBookingInst from ExecutionReport.
func (m ExecutionReport) GetDayBookingInst(f *field.DayBookingInstField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BookingUnit is a non-required field for ExecutionReport.
func (m ExecutionReport) BookingUnit() (*field.BookingUnitField, errors.MessageRejectError) {
	f := &field.BookingUnitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBookingUnit reads a BookingUnit from ExecutionReport.
func (m ExecutionReport) GetBookingUnit(f *field.BookingUnitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PreallocMethod is a non-required field for ExecutionReport.
func (m ExecutionReport) PreallocMethod() (*field.PreallocMethodField, errors.MessageRejectError) {
	f := &field.PreallocMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPreallocMethod reads a PreallocMethod from ExecutionReport.
func (m ExecutionReport) GetPreallocMethod(f *field.PreallocMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlType is a non-required field for ExecutionReport.
func (m ExecutionReport) SettlType() (*field.SettlTypeField, errors.MessageRejectError) {
	f := &field.SettlTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlType reads a SettlType from ExecutionReport.
func (m ExecutionReport) GetSettlType(f *field.SettlTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlDate is a non-required field for ExecutionReport.
func (m ExecutionReport) SettlDate() (*field.SettlDateField, errors.MessageRejectError) {
	f := &field.SettlDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlDate reads a SettlDate from ExecutionReport.
func (m ExecutionReport) GetSettlDate(f *field.SettlDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CashMargin is a non-required field for ExecutionReport.
func (m ExecutionReport) CashMargin() (*field.CashMarginField, errors.MessageRejectError) {
	f := &field.CashMarginField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCashMargin reads a CashMargin from ExecutionReport.
func (m ExecutionReport) GetCashMargin(f *field.CashMarginField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClearingFeeIndicator is a non-required field for ExecutionReport.
func (m ExecutionReport) ClearingFeeIndicator() (*field.ClearingFeeIndicatorField, errors.MessageRejectError) {
	f := &field.ClearingFeeIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClearingFeeIndicator reads a ClearingFeeIndicator from ExecutionReport.
func (m ExecutionReport) GetClearingFeeIndicator(f *field.ClearingFeeIndicatorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for ExecutionReport.
func (m ExecutionReport) Symbol() (*field.SymbolField, errors.MessageRejectError) {
	f := &field.SymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from ExecutionReport.
func (m ExecutionReport) GetSymbol(f *field.SymbolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for ExecutionReport.
func (m ExecutionReport) SymbolSfx() (*field.SymbolSfxField, errors.MessageRejectError) {
	f := &field.SymbolSfxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from ExecutionReport.
func (m ExecutionReport) GetSymbolSfx(f *field.SymbolSfxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for ExecutionReport.
func (m ExecutionReport) SecurityID() (*field.SecurityIDField, errors.MessageRejectError) {
	f := &field.SecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from ExecutionReport.
func (m ExecutionReport) GetSecurityID(f *field.SecurityIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for ExecutionReport.
func (m ExecutionReport) SecurityIDSource() (*field.SecurityIDSourceField, errors.MessageRejectError) {
	f := &field.SecurityIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from ExecutionReport.
func (m ExecutionReport) GetSecurityIDSource(f *field.SecurityIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for ExecutionReport.
func (m ExecutionReport) NoSecurityAltID() (*field.NoSecurityAltIDField, errors.MessageRejectError) {
	f := &field.NoSecurityAltIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from ExecutionReport.
func (m ExecutionReport) GetNoSecurityAltID(f *field.NoSecurityAltIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for ExecutionReport.
func (m ExecutionReport) Product() (*field.ProductField, errors.MessageRejectError) {
	f := &field.ProductField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from ExecutionReport.
func (m ExecutionReport) GetProduct(f *field.ProductField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for ExecutionReport.
func (m ExecutionReport) CFICode() (*field.CFICodeField, errors.MessageRejectError) {
	f := &field.CFICodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from ExecutionReport.
func (m ExecutionReport) GetCFICode(f *field.CFICodeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for ExecutionReport.
func (m ExecutionReport) SecurityType() (*field.SecurityTypeField, errors.MessageRejectError) {
	f := &field.SecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from ExecutionReport.
func (m ExecutionReport) GetSecurityType(f *field.SecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySubType is a non-required field for ExecutionReport.
func (m ExecutionReport) SecuritySubType() (*field.SecuritySubTypeField, errors.MessageRejectError) {
	f := &field.SecuritySubTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySubType reads a SecuritySubType from ExecutionReport.
func (m ExecutionReport) GetSecuritySubType(f *field.SecuritySubTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for ExecutionReport.
func (m ExecutionReport) MaturityMonthYear() (*field.MaturityMonthYearField, errors.MessageRejectError) {
	f := &field.MaturityMonthYearField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from ExecutionReport.
func (m ExecutionReport) GetMaturityMonthYear(f *field.MaturityMonthYearField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for ExecutionReport.
func (m ExecutionReport) MaturityDate() (*field.MaturityDateField, errors.MessageRejectError) {
	f := &field.MaturityDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from ExecutionReport.
func (m ExecutionReport) GetMaturityDate(f *field.MaturityDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for ExecutionReport.
func (m ExecutionReport) CouponPaymentDate() (*field.CouponPaymentDateField, errors.MessageRejectError) {
	f := &field.CouponPaymentDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from ExecutionReport.
func (m ExecutionReport) GetCouponPaymentDate(f *field.CouponPaymentDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for ExecutionReport.
func (m ExecutionReport) IssueDate() (*field.IssueDateField, errors.MessageRejectError) {
	f := &field.IssueDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from ExecutionReport.
func (m ExecutionReport) GetIssueDate(f *field.IssueDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for ExecutionReport.
func (m ExecutionReport) RepoCollateralSecurityType() (*field.RepoCollateralSecurityTypeField, errors.MessageRejectError) {
	f := &field.RepoCollateralSecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from ExecutionReport.
func (m ExecutionReport) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for ExecutionReport.
func (m ExecutionReport) RepurchaseTerm() (*field.RepurchaseTermField, errors.MessageRejectError) {
	f := &field.RepurchaseTermField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from ExecutionReport.
func (m ExecutionReport) GetRepurchaseTerm(f *field.RepurchaseTermField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for ExecutionReport.
func (m ExecutionReport) RepurchaseRate() (*field.RepurchaseRateField, errors.MessageRejectError) {
	f := &field.RepurchaseRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from ExecutionReport.
func (m ExecutionReport) GetRepurchaseRate(f *field.RepurchaseRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for ExecutionReport.
func (m ExecutionReport) Factor() (*field.FactorField, errors.MessageRejectError) {
	f := &field.FactorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from ExecutionReport.
func (m ExecutionReport) GetFactor(f *field.FactorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for ExecutionReport.
func (m ExecutionReport) CreditRating() (*field.CreditRatingField, errors.MessageRejectError) {
	f := &field.CreditRatingField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from ExecutionReport.
func (m ExecutionReport) GetCreditRating(f *field.CreditRatingField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for ExecutionReport.
func (m ExecutionReport) InstrRegistry() (*field.InstrRegistryField, errors.MessageRejectError) {
	f := &field.InstrRegistryField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from ExecutionReport.
func (m ExecutionReport) GetInstrRegistry(f *field.InstrRegistryField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for ExecutionReport.
func (m ExecutionReport) CountryOfIssue() (*field.CountryOfIssueField, errors.MessageRejectError) {
	f := &field.CountryOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from ExecutionReport.
func (m ExecutionReport) GetCountryOfIssue(f *field.CountryOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for ExecutionReport.
func (m ExecutionReport) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssueField, errors.MessageRejectError) {
	f := &field.StateOrProvinceOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from ExecutionReport.
func (m ExecutionReport) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for ExecutionReport.
func (m ExecutionReport) LocaleOfIssue() (*field.LocaleOfIssueField, errors.MessageRejectError) {
	f := &field.LocaleOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from ExecutionReport.
func (m ExecutionReport) GetLocaleOfIssue(f *field.LocaleOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for ExecutionReport.
func (m ExecutionReport) RedemptionDate() (*field.RedemptionDateField, errors.MessageRejectError) {
	f := &field.RedemptionDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from ExecutionReport.
func (m ExecutionReport) GetRedemptionDate(f *field.RedemptionDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for ExecutionReport.
func (m ExecutionReport) StrikePrice() (*field.StrikePriceField, errors.MessageRejectError) {
	f := &field.StrikePriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from ExecutionReport.
func (m ExecutionReport) GetStrikePrice(f *field.StrikePriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeCurrency is a non-required field for ExecutionReport.
func (m ExecutionReport) StrikeCurrency() (*field.StrikeCurrencyField, errors.MessageRejectError) {
	f := &field.StrikeCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeCurrency reads a StrikeCurrency from ExecutionReport.
func (m ExecutionReport) GetStrikeCurrency(f *field.StrikeCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for ExecutionReport.
func (m ExecutionReport) OptAttribute() (*field.OptAttributeField, errors.MessageRejectError) {
	f := &field.OptAttributeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from ExecutionReport.
func (m ExecutionReport) GetOptAttribute(f *field.OptAttributeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for ExecutionReport.
func (m ExecutionReport) ContractMultiplier() (*field.ContractMultiplierField, errors.MessageRejectError) {
	f := &field.ContractMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from ExecutionReport.
func (m ExecutionReport) GetContractMultiplier(f *field.ContractMultiplierField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for ExecutionReport.
func (m ExecutionReport) CouponRate() (*field.CouponRateField, errors.MessageRejectError) {
	f := &field.CouponRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from ExecutionReport.
func (m ExecutionReport) GetCouponRate(f *field.CouponRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for ExecutionReport.
func (m ExecutionReport) SecurityExchange() (*field.SecurityExchangeField, errors.MessageRejectError) {
	f := &field.SecurityExchangeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from ExecutionReport.
func (m ExecutionReport) GetSecurityExchange(f *field.SecurityExchangeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for ExecutionReport.
func (m ExecutionReport) Issuer() (*field.IssuerField, errors.MessageRejectError) {
	f := &field.IssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from ExecutionReport.
func (m ExecutionReport) GetIssuer(f *field.IssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for ExecutionReport.
func (m ExecutionReport) EncodedIssuerLen() (*field.EncodedIssuerLenField, errors.MessageRejectError) {
	f := &field.EncodedIssuerLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from ExecutionReport.
func (m ExecutionReport) GetEncodedIssuerLen(f *field.EncodedIssuerLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for ExecutionReport.
func (m ExecutionReport) EncodedIssuer() (*field.EncodedIssuerField, errors.MessageRejectError) {
	f := &field.EncodedIssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from ExecutionReport.
func (m ExecutionReport) GetEncodedIssuer(f *field.EncodedIssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for ExecutionReport.
func (m ExecutionReport) SecurityDesc() (*field.SecurityDescField, errors.MessageRejectError) {
	f := &field.SecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from ExecutionReport.
func (m ExecutionReport) GetSecurityDesc(f *field.SecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for ExecutionReport.
func (m ExecutionReport) EncodedSecurityDescLen() (*field.EncodedSecurityDescLenField, errors.MessageRejectError) {
	f := &field.EncodedSecurityDescLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from ExecutionReport.
func (m ExecutionReport) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for ExecutionReport.
func (m ExecutionReport) EncodedSecurityDesc() (*field.EncodedSecurityDescField, errors.MessageRejectError) {
	f := &field.EncodedSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from ExecutionReport.
func (m ExecutionReport) GetEncodedSecurityDesc(f *field.EncodedSecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Pool is a non-required field for ExecutionReport.
func (m ExecutionReport) Pool() (*field.PoolField, errors.MessageRejectError) {
	f := &field.PoolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPool reads a Pool from ExecutionReport.
func (m ExecutionReport) GetPool(f *field.PoolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractSettlMonth is a non-required field for ExecutionReport.
func (m ExecutionReport) ContractSettlMonth() (*field.ContractSettlMonthField, errors.MessageRejectError) {
	f := &field.ContractSettlMonthField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractSettlMonth reads a ContractSettlMonth from ExecutionReport.
func (m ExecutionReport) GetContractSettlMonth(f *field.ContractSettlMonthField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPProgram is a non-required field for ExecutionReport.
func (m ExecutionReport) CPProgram() (*field.CPProgramField, errors.MessageRejectError) {
	f := &field.CPProgramField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCPProgram reads a CPProgram from ExecutionReport.
func (m ExecutionReport) GetCPProgram(f *field.CPProgramField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPRegType is a non-required field for ExecutionReport.
func (m ExecutionReport) CPRegType() (*field.CPRegTypeField, errors.MessageRejectError) {
	f := &field.CPRegTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCPRegType reads a CPRegType from ExecutionReport.
func (m ExecutionReport) GetCPRegType(f *field.CPRegTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoEvents is a non-required field for ExecutionReport.
func (m ExecutionReport) NoEvents() (*field.NoEventsField, errors.MessageRejectError) {
	f := &field.NoEventsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoEvents reads a NoEvents from ExecutionReport.
func (m ExecutionReport) GetNoEvents(f *field.NoEventsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DatedDate is a non-required field for ExecutionReport.
func (m ExecutionReport) DatedDate() (*field.DatedDateField, errors.MessageRejectError) {
	f := &field.DatedDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDatedDate reads a DatedDate from ExecutionReport.
func (m ExecutionReport) GetDatedDate(f *field.DatedDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InterestAccrualDate is a non-required field for ExecutionReport.
func (m ExecutionReport) InterestAccrualDate() (*field.InterestAccrualDateField, errors.MessageRejectError) {
	f := &field.InterestAccrualDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInterestAccrualDate reads a InterestAccrualDate from ExecutionReport.
func (m ExecutionReport) GetInterestAccrualDate(f *field.InterestAccrualDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityStatus is a non-required field for ExecutionReport.
func (m ExecutionReport) SecurityStatus() (*field.SecurityStatusField, errors.MessageRejectError) {
	f := &field.SecurityStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityStatus reads a SecurityStatus from ExecutionReport.
func (m ExecutionReport) GetSecurityStatus(f *field.SecurityStatusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettleOnOpenFlag is a non-required field for ExecutionReport.
func (m ExecutionReport) SettleOnOpenFlag() (*field.SettleOnOpenFlagField, errors.MessageRejectError) {
	f := &field.SettleOnOpenFlagField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettleOnOpenFlag reads a SettleOnOpenFlag from ExecutionReport.
func (m ExecutionReport) GetSettleOnOpenFlag(f *field.SettleOnOpenFlagField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrmtAssignmentMethod is a non-required field for ExecutionReport.
func (m ExecutionReport) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethodField, errors.MessageRejectError) {
	f := &field.InstrmtAssignmentMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInstrmtAssignmentMethod reads a InstrmtAssignmentMethod from ExecutionReport.
func (m ExecutionReport) GetInstrmtAssignmentMethod(f *field.InstrmtAssignmentMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeMultiplier is a non-required field for ExecutionReport.
func (m ExecutionReport) StrikeMultiplier() (*field.StrikeMultiplierField, errors.MessageRejectError) {
	f := &field.StrikeMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeMultiplier reads a StrikeMultiplier from ExecutionReport.
func (m ExecutionReport) GetStrikeMultiplier(f *field.StrikeMultiplierField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeValue is a non-required field for ExecutionReport.
func (m ExecutionReport) StrikeValue() (*field.StrikeValueField, errors.MessageRejectError) {
	f := &field.StrikeValueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeValue reads a StrikeValue from ExecutionReport.
func (m ExecutionReport) GetStrikeValue(f *field.StrikeValueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrement is a non-required field for ExecutionReport.
func (m ExecutionReport) MinPriceIncrement() (*field.MinPriceIncrementField, errors.MessageRejectError) {
	f := &field.MinPriceIncrementField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrement reads a MinPriceIncrement from ExecutionReport.
func (m ExecutionReport) GetMinPriceIncrement(f *field.MinPriceIncrementField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PositionLimit is a non-required field for ExecutionReport.
func (m ExecutionReport) PositionLimit() (*field.PositionLimitField, errors.MessageRejectError) {
	f := &field.PositionLimitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPositionLimit reads a PositionLimit from ExecutionReport.
func (m ExecutionReport) GetPositionLimit(f *field.PositionLimitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NTPositionLimit is a non-required field for ExecutionReport.
func (m ExecutionReport) NTPositionLimit() (*field.NTPositionLimitField, errors.MessageRejectError) {
	f := &field.NTPositionLimitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNTPositionLimit reads a NTPositionLimit from ExecutionReport.
func (m ExecutionReport) GetNTPositionLimit(f *field.NTPositionLimitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoInstrumentParties is a non-required field for ExecutionReport.
func (m ExecutionReport) NoInstrumentParties() (*field.NoInstrumentPartiesField, errors.MessageRejectError) {
	f := &field.NoInstrumentPartiesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoInstrumentParties reads a NoInstrumentParties from ExecutionReport.
func (m ExecutionReport) GetNoInstrumentParties(f *field.NoInstrumentPartiesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasure is a non-required field for ExecutionReport.
func (m ExecutionReport) UnitOfMeasure() (*field.UnitOfMeasureField, errors.MessageRejectError) {
	f := &field.UnitOfMeasureField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasure reads a UnitOfMeasure from ExecutionReport.
func (m ExecutionReport) GetUnitOfMeasure(f *field.UnitOfMeasureField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TimeUnit is a non-required field for ExecutionReport.
func (m ExecutionReport) TimeUnit() (*field.TimeUnitField, errors.MessageRejectError) {
	f := &field.TimeUnitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTimeUnit reads a TimeUnit from ExecutionReport.
func (m ExecutionReport) GetTimeUnit(f *field.TimeUnitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityTime is a non-required field for ExecutionReport.
func (m ExecutionReport) MaturityTime() (*field.MaturityTimeField, errors.MessageRejectError) {
	f := &field.MaturityTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityTime reads a MaturityTime from ExecutionReport.
func (m ExecutionReport) GetMaturityTime(f *field.MaturityTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityGroup is a non-required field for ExecutionReport.
func (m ExecutionReport) SecurityGroup() (*field.SecurityGroupField, errors.MessageRejectError) {
	f := &field.SecurityGroupField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityGroup reads a SecurityGroup from ExecutionReport.
func (m ExecutionReport) GetSecurityGroup(f *field.SecurityGroupField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrementAmount is a non-required field for ExecutionReport.
func (m ExecutionReport) MinPriceIncrementAmount() (*field.MinPriceIncrementAmountField, errors.MessageRejectError) {
	f := &field.MinPriceIncrementAmountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrementAmount reads a MinPriceIncrementAmount from ExecutionReport.
func (m ExecutionReport) GetMinPriceIncrementAmount(f *field.MinPriceIncrementAmountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasureQty is a non-required field for ExecutionReport.
func (m ExecutionReport) UnitOfMeasureQty() (*field.UnitOfMeasureQtyField, errors.MessageRejectError) {
	f := &field.UnitOfMeasureQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasureQty reads a UnitOfMeasureQty from ExecutionReport.
func (m ExecutionReport) GetUnitOfMeasureQty(f *field.UnitOfMeasureQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXMLLen is a non-required field for ExecutionReport.
func (m ExecutionReport) SecurityXMLLen() (*field.SecurityXMLLenField, errors.MessageRejectError) {
	f := &field.SecurityXMLLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXMLLen reads a SecurityXMLLen from ExecutionReport.
func (m ExecutionReport) GetSecurityXMLLen(f *field.SecurityXMLLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXML is a non-required field for ExecutionReport.
func (m ExecutionReport) SecurityXML() (*field.SecurityXMLField, errors.MessageRejectError) {
	f := &field.SecurityXMLField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXML reads a SecurityXML from ExecutionReport.
func (m ExecutionReport) GetSecurityXML(f *field.SecurityXMLField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXMLSchema is a non-required field for ExecutionReport.
func (m ExecutionReport) SecurityXMLSchema() (*field.SecurityXMLSchemaField, errors.MessageRejectError) {
	f := &field.SecurityXMLSchemaField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXMLSchema reads a SecurityXMLSchema from ExecutionReport.
func (m ExecutionReport) GetSecurityXMLSchema(f *field.SecurityXMLSchemaField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ProductComplex is a non-required field for ExecutionReport.
func (m ExecutionReport) ProductComplex() (*field.ProductComplexField, errors.MessageRejectError) {
	f := &field.ProductComplexField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProductComplex reads a ProductComplex from ExecutionReport.
func (m ExecutionReport) GetProductComplex(f *field.ProductComplexField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceUnitOfMeasure is a non-required field for ExecutionReport.
func (m ExecutionReport) PriceUnitOfMeasure() (*field.PriceUnitOfMeasureField, errors.MessageRejectError) {
	f := &field.PriceUnitOfMeasureField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceUnitOfMeasure reads a PriceUnitOfMeasure from ExecutionReport.
func (m ExecutionReport) GetPriceUnitOfMeasure(f *field.PriceUnitOfMeasureField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceUnitOfMeasureQty is a non-required field for ExecutionReport.
func (m ExecutionReport) PriceUnitOfMeasureQty() (*field.PriceUnitOfMeasureQtyField, errors.MessageRejectError) {
	f := &field.PriceUnitOfMeasureQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceUnitOfMeasureQty reads a PriceUnitOfMeasureQty from ExecutionReport.
func (m ExecutionReport) GetPriceUnitOfMeasureQty(f *field.PriceUnitOfMeasureQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlMethod is a non-required field for ExecutionReport.
func (m ExecutionReport) SettlMethod() (*field.SettlMethodField, errors.MessageRejectError) {
	f := &field.SettlMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlMethod reads a SettlMethod from ExecutionReport.
func (m ExecutionReport) GetSettlMethod(f *field.SettlMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExerciseStyle is a non-required field for ExecutionReport.
func (m ExecutionReport) ExerciseStyle() (*field.ExerciseStyleField, errors.MessageRejectError) {
	f := &field.ExerciseStyleField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExerciseStyle reads a ExerciseStyle from ExecutionReport.
func (m ExecutionReport) GetExerciseStyle(f *field.ExerciseStyleField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptPayAmount is a non-required field for ExecutionReport.
func (m ExecutionReport) OptPayAmount() (*field.OptPayAmountField, errors.MessageRejectError) {
	f := &field.OptPayAmountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptPayAmount reads a OptPayAmount from ExecutionReport.
func (m ExecutionReport) GetOptPayAmount(f *field.OptPayAmountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceQuoteMethod is a non-required field for ExecutionReport.
func (m ExecutionReport) PriceQuoteMethod() (*field.PriceQuoteMethodField, errors.MessageRejectError) {
	f := &field.PriceQuoteMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceQuoteMethod reads a PriceQuoteMethod from ExecutionReport.
func (m ExecutionReport) GetPriceQuoteMethod(f *field.PriceQuoteMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListMethod is a non-required field for ExecutionReport.
func (m ExecutionReport) ListMethod() (*field.ListMethodField, errors.MessageRejectError) {
	f := &field.ListMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetListMethod reads a ListMethod from ExecutionReport.
func (m ExecutionReport) GetListMethod(f *field.ListMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CapPrice is a non-required field for ExecutionReport.
func (m ExecutionReport) CapPrice() (*field.CapPriceField, errors.MessageRejectError) {
	f := &field.CapPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCapPrice reads a CapPrice from ExecutionReport.
func (m ExecutionReport) GetCapPrice(f *field.CapPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FloorPrice is a non-required field for ExecutionReport.
func (m ExecutionReport) FloorPrice() (*field.FloorPriceField, errors.MessageRejectError) {
	f := &field.FloorPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFloorPrice reads a FloorPrice from ExecutionReport.
func (m ExecutionReport) GetFloorPrice(f *field.FloorPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PutOrCall is a non-required field for ExecutionReport.
func (m ExecutionReport) PutOrCall() (*field.PutOrCallField, errors.MessageRejectError) {
	f := &field.PutOrCallField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPutOrCall reads a PutOrCall from ExecutionReport.
func (m ExecutionReport) GetPutOrCall(f *field.PutOrCallField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlexibleIndicator is a non-required field for ExecutionReport.
func (m ExecutionReport) FlexibleIndicator() (*field.FlexibleIndicatorField, errors.MessageRejectError) {
	f := &field.FlexibleIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFlexibleIndicator reads a FlexibleIndicator from ExecutionReport.
func (m ExecutionReport) GetFlexibleIndicator(f *field.FlexibleIndicatorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlexProductEligibilityIndicator is a non-required field for ExecutionReport.
func (m ExecutionReport) FlexProductEligibilityIndicator() (*field.FlexProductEligibilityIndicatorField, errors.MessageRejectError) {
	f := &field.FlexProductEligibilityIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFlexProductEligibilityIndicator reads a FlexProductEligibilityIndicator from ExecutionReport.
func (m ExecutionReport) GetFlexProductEligibilityIndicator(f *field.FlexProductEligibilityIndicatorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FuturesValuationMethod is a non-required field for ExecutionReport.
func (m ExecutionReport) FuturesValuationMethod() (*field.FuturesValuationMethodField, errors.MessageRejectError) {
	f := &field.FuturesValuationMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFuturesValuationMethod reads a FuturesValuationMethod from ExecutionReport.
func (m ExecutionReport) GetFuturesValuationMethod(f *field.FuturesValuationMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementDesc is a non-required field for ExecutionReport.
func (m ExecutionReport) AgreementDesc() (*field.AgreementDescField, errors.MessageRejectError) {
	f := &field.AgreementDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementDesc reads a AgreementDesc from ExecutionReport.
func (m ExecutionReport) GetAgreementDesc(f *field.AgreementDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementID is a non-required field for ExecutionReport.
func (m ExecutionReport) AgreementID() (*field.AgreementIDField, errors.MessageRejectError) {
	f := &field.AgreementIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementID reads a AgreementID from ExecutionReport.
func (m ExecutionReport) GetAgreementID(f *field.AgreementIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementDate is a non-required field for ExecutionReport.
func (m ExecutionReport) AgreementDate() (*field.AgreementDateField, errors.MessageRejectError) {
	f := &field.AgreementDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementDate reads a AgreementDate from ExecutionReport.
func (m ExecutionReport) GetAgreementDate(f *field.AgreementDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementCurrency is a non-required field for ExecutionReport.
func (m ExecutionReport) AgreementCurrency() (*field.AgreementCurrencyField, errors.MessageRejectError) {
	f := &field.AgreementCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementCurrency reads a AgreementCurrency from ExecutionReport.
func (m ExecutionReport) GetAgreementCurrency(f *field.AgreementCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TerminationType is a non-required field for ExecutionReport.
func (m ExecutionReport) TerminationType() (*field.TerminationTypeField, errors.MessageRejectError) {
	f := &field.TerminationTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTerminationType reads a TerminationType from ExecutionReport.
func (m ExecutionReport) GetTerminationType(f *field.TerminationTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StartDate is a non-required field for ExecutionReport.
func (m ExecutionReport) StartDate() (*field.StartDateField, errors.MessageRejectError) {
	f := &field.StartDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStartDate reads a StartDate from ExecutionReport.
func (m ExecutionReport) GetStartDate(f *field.StartDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EndDate is a non-required field for ExecutionReport.
func (m ExecutionReport) EndDate() (*field.EndDateField, errors.MessageRejectError) {
	f := &field.EndDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEndDate reads a EndDate from ExecutionReport.
func (m ExecutionReport) GetEndDate(f *field.EndDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DeliveryType is a non-required field for ExecutionReport.
func (m ExecutionReport) DeliveryType() (*field.DeliveryTypeField, errors.MessageRejectError) {
	f := &field.DeliveryTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDeliveryType reads a DeliveryType from ExecutionReport.
func (m ExecutionReport) GetDeliveryType(f *field.DeliveryTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarginRatio is a non-required field for ExecutionReport.
func (m ExecutionReport) MarginRatio() (*field.MarginRatioField, errors.MessageRejectError) {
	f := &field.MarginRatioField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarginRatio reads a MarginRatio from ExecutionReport.
func (m ExecutionReport) GetMarginRatio(f *field.MarginRatioField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyings is a non-required field for ExecutionReport.
func (m ExecutionReport) NoUnderlyings() (*field.NoUnderlyingsField, errors.MessageRejectError) {
	f := &field.NoUnderlyingsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyings reads a NoUnderlyings from ExecutionReport.
func (m ExecutionReport) GetNoUnderlyings(f *field.NoUnderlyingsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Side is a required field for ExecutionReport.
func (m ExecutionReport) Side() (*field.SideField, errors.MessageRejectError) {
	f := &field.SideField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSide reads a Side from ExecutionReport.
func (m ExecutionReport) GetSide(f *field.SideField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoStipulations is a non-required field for ExecutionReport.
func (m ExecutionReport) NoStipulations() (*field.NoStipulationsField, errors.MessageRejectError) {
	f := &field.NoStipulationsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoStipulations reads a NoStipulations from ExecutionReport.
func (m ExecutionReport) GetNoStipulations(f *field.NoStipulationsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QtyType is a non-required field for ExecutionReport.
func (m ExecutionReport) QtyType() (*field.QtyTypeField, errors.MessageRejectError) {
	f := &field.QtyTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQtyType reads a QtyType from ExecutionReport.
func (m ExecutionReport) GetQtyType(f *field.QtyTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderQty is a non-required field for ExecutionReport.
func (m ExecutionReport) OrderQty() (*field.OrderQtyField, errors.MessageRejectError) {
	f := &field.OrderQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderQty reads a OrderQty from ExecutionReport.
func (m ExecutionReport) GetOrderQty(f *field.OrderQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CashOrderQty is a non-required field for ExecutionReport.
func (m ExecutionReport) CashOrderQty() (*field.CashOrderQtyField, errors.MessageRejectError) {
	f := &field.CashOrderQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCashOrderQty reads a CashOrderQty from ExecutionReport.
func (m ExecutionReport) GetCashOrderQty(f *field.CashOrderQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderPercent is a non-required field for ExecutionReport.
func (m ExecutionReport) OrderPercent() (*field.OrderPercentField, errors.MessageRejectError) {
	f := &field.OrderPercentField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderPercent reads a OrderPercent from ExecutionReport.
func (m ExecutionReport) GetOrderPercent(f *field.OrderPercentField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RoundingDirection is a non-required field for ExecutionReport.
func (m ExecutionReport) RoundingDirection() (*field.RoundingDirectionField, errors.MessageRejectError) {
	f := &field.RoundingDirectionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRoundingDirection reads a RoundingDirection from ExecutionReport.
func (m ExecutionReport) GetRoundingDirection(f *field.RoundingDirectionField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RoundingModulus is a non-required field for ExecutionReport.
func (m ExecutionReport) RoundingModulus() (*field.RoundingModulusField, errors.MessageRejectError) {
	f := &field.RoundingModulusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRoundingModulus reads a RoundingModulus from ExecutionReport.
func (m ExecutionReport) GetRoundingModulus(f *field.RoundingModulusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrdType is a non-required field for ExecutionReport.
func (m ExecutionReport) OrdType() (*field.OrdTypeField, errors.MessageRejectError) {
	f := &field.OrdTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrdType reads a OrdType from ExecutionReport.
func (m ExecutionReport) GetOrdType(f *field.OrdTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceType is a non-required field for ExecutionReport.
func (m ExecutionReport) PriceType() (*field.PriceTypeField, errors.MessageRejectError) {
	f := &field.PriceTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceType reads a PriceType from ExecutionReport.
func (m ExecutionReport) GetPriceType(f *field.PriceTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Price is a non-required field for ExecutionReport.
func (m ExecutionReport) Price() (*field.PriceField, errors.MessageRejectError) {
	f := &field.PriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPrice reads a Price from ExecutionReport.
func (m ExecutionReport) GetPrice(f *field.PriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StopPx is a non-required field for ExecutionReport.
func (m ExecutionReport) StopPx() (*field.StopPxField, errors.MessageRejectError) {
	f := &field.StopPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStopPx reads a StopPx from ExecutionReport.
func (m ExecutionReport) GetStopPx(f *field.StopPxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegOffsetValue is a non-required field for ExecutionReport.
func (m ExecutionReport) PegOffsetValue() (*field.PegOffsetValueField, errors.MessageRejectError) {
	f := &field.PegOffsetValueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegOffsetValue reads a PegOffsetValue from ExecutionReport.
func (m ExecutionReport) GetPegOffsetValue(f *field.PegOffsetValueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegMoveType is a non-required field for ExecutionReport.
func (m ExecutionReport) PegMoveType() (*field.PegMoveTypeField, errors.MessageRejectError) {
	f := &field.PegMoveTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegMoveType reads a PegMoveType from ExecutionReport.
func (m ExecutionReport) GetPegMoveType(f *field.PegMoveTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegOffsetType is a non-required field for ExecutionReport.
func (m ExecutionReport) PegOffsetType() (*field.PegOffsetTypeField, errors.MessageRejectError) {
	f := &field.PegOffsetTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegOffsetType reads a PegOffsetType from ExecutionReport.
func (m ExecutionReport) GetPegOffsetType(f *field.PegOffsetTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegLimitType is a non-required field for ExecutionReport.
func (m ExecutionReport) PegLimitType() (*field.PegLimitTypeField, errors.MessageRejectError) {
	f := &field.PegLimitTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegLimitType reads a PegLimitType from ExecutionReport.
func (m ExecutionReport) GetPegLimitType(f *field.PegLimitTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegRoundDirection is a non-required field for ExecutionReport.
func (m ExecutionReport) PegRoundDirection() (*field.PegRoundDirectionField, errors.MessageRejectError) {
	f := &field.PegRoundDirectionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegRoundDirection reads a PegRoundDirection from ExecutionReport.
func (m ExecutionReport) GetPegRoundDirection(f *field.PegRoundDirectionField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegScope is a non-required field for ExecutionReport.
func (m ExecutionReport) PegScope() (*field.PegScopeField, errors.MessageRejectError) {
	f := &field.PegScopeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegScope reads a PegScope from ExecutionReport.
func (m ExecutionReport) GetPegScope(f *field.PegScopeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegPriceType is a non-required field for ExecutionReport.
func (m ExecutionReport) PegPriceType() (*field.PegPriceTypeField, errors.MessageRejectError) {
	f := &field.PegPriceTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegPriceType reads a PegPriceType from ExecutionReport.
func (m ExecutionReport) GetPegPriceType(f *field.PegPriceTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegSecurityIDSource is a non-required field for ExecutionReport.
func (m ExecutionReport) PegSecurityIDSource() (*field.PegSecurityIDSourceField, errors.MessageRejectError) {
	f := &field.PegSecurityIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegSecurityIDSource reads a PegSecurityIDSource from ExecutionReport.
func (m ExecutionReport) GetPegSecurityIDSource(f *field.PegSecurityIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegSecurityID is a non-required field for ExecutionReport.
func (m ExecutionReport) PegSecurityID() (*field.PegSecurityIDField, errors.MessageRejectError) {
	f := &field.PegSecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegSecurityID reads a PegSecurityID from ExecutionReport.
func (m ExecutionReport) GetPegSecurityID(f *field.PegSecurityIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegSymbol is a non-required field for ExecutionReport.
func (m ExecutionReport) PegSymbol() (*field.PegSymbolField, errors.MessageRejectError) {
	f := &field.PegSymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegSymbol reads a PegSymbol from ExecutionReport.
func (m ExecutionReport) GetPegSymbol(f *field.PegSymbolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegSecurityDesc is a non-required field for ExecutionReport.
func (m ExecutionReport) PegSecurityDesc() (*field.PegSecurityDescField, errors.MessageRejectError) {
	f := &field.PegSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegSecurityDesc reads a PegSecurityDesc from ExecutionReport.
func (m ExecutionReport) GetPegSecurityDesc(f *field.PegSecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionInst is a non-required field for ExecutionReport.
func (m ExecutionReport) DiscretionInst() (*field.DiscretionInstField, errors.MessageRejectError) {
	f := &field.DiscretionInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionInst reads a DiscretionInst from ExecutionReport.
func (m ExecutionReport) GetDiscretionInst(f *field.DiscretionInstField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionOffsetValue is a non-required field for ExecutionReport.
func (m ExecutionReport) DiscretionOffsetValue() (*field.DiscretionOffsetValueField, errors.MessageRejectError) {
	f := &field.DiscretionOffsetValueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionOffsetValue reads a DiscretionOffsetValue from ExecutionReport.
func (m ExecutionReport) GetDiscretionOffsetValue(f *field.DiscretionOffsetValueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionMoveType is a non-required field for ExecutionReport.
func (m ExecutionReport) DiscretionMoveType() (*field.DiscretionMoveTypeField, errors.MessageRejectError) {
	f := &field.DiscretionMoveTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionMoveType reads a DiscretionMoveType from ExecutionReport.
func (m ExecutionReport) GetDiscretionMoveType(f *field.DiscretionMoveTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionOffsetType is a non-required field for ExecutionReport.
func (m ExecutionReport) DiscretionOffsetType() (*field.DiscretionOffsetTypeField, errors.MessageRejectError) {
	f := &field.DiscretionOffsetTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionOffsetType reads a DiscretionOffsetType from ExecutionReport.
func (m ExecutionReport) GetDiscretionOffsetType(f *field.DiscretionOffsetTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionLimitType is a non-required field for ExecutionReport.
func (m ExecutionReport) DiscretionLimitType() (*field.DiscretionLimitTypeField, errors.MessageRejectError) {
	f := &field.DiscretionLimitTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionLimitType reads a DiscretionLimitType from ExecutionReport.
func (m ExecutionReport) GetDiscretionLimitType(f *field.DiscretionLimitTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionRoundDirection is a non-required field for ExecutionReport.
func (m ExecutionReport) DiscretionRoundDirection() (*field.DiscretionRoundDirectionField, errors.MessageRejectError) {
	f := &field.DiscretionRoundDirectionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionRoundDirection reads a DiscretionRoundDirection from ExecutionReport.
func (m ExecutionReport) GetDiscretionRoundDirection(f *field.DiscretionRoundDirectionField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionScope is a non-required field for ExecutionReport.
func (m ExecutionReport) DiscretionScope() (*field.DiscretionScopeField, errors.MessageRejectError) {
	f := &field.DiscretionScopeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionScope reads a DiscretionScope from ExecutionReport.
func (m ExecutionReport) GetDiscretionScope(f *field.DiscretionScopeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PeggedPrice is a non-required field for ExecutionReport.
func (m ExecutionReport) PeggedPrice() (*field.PeggedPriceField, errors.MessageRejectError) {
	f := &field.PeggedPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPeggedPrice reads a PeggedPrice from ExecutionReport.
func (m ExecutionReport) GetPeggedPrice(f *field.PeggedPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionPrice is a non-required field for ExecutionReport.
func (m ExecutionReport) DiscretionPrice() (*field.DiscretionPriceField, errors.MessageRejectError) {
	f := &field.DiscretionPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionPrice reads a DiscretionPrice from ExecutionReport.
func (m ExecutionReport) GetDiscretionPrice(f *field.DiscretionPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TargetStrategy is a non-required field for ExecutionReport.
func (m ExecutionReport) TargetStrategy() (*field.TargetStrategyField, errors.MessageRejectError) {
	f := &field.TargetStrategyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTargetStrategy reads a TargetStrategy from ExecutionReport.
func (m ExecutionReport) GetTargetStrategy(f *field.TargetStrategyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TargetStrategyParameters is a non-required field for ExecutionReport.
func (m ExecutionReport) TargetStrategyParameters() (*field.TargetStrategyParametersField, errors.MessageRejectError) {
	f := &field.TargetStrategyParametersField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTargetStrategyParameters reads a TargetStrategyParameters from ExecutionReport.
func (m ExecutionReport) GetTargetStrategyParameters(f *field.TargetStrategyParametersField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ParticipationRate is a non-required field for ExecutionReport.
func (m ExecutionReport) ParticipationRate() (*field.ParticipationRateField, errors.MessageRejectError) {
	f := &field.ParticipationRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetParticipationRate reads a ParticipationRate from ExecutionReport.
func (m ExecutionReport) GetParticipationRate(f *field.ParticipationRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TargetStrategyPerformance is a non-required field for ExecutionReport.
func (m ExecutionReport) TargetStrategyPerformance() (*field.TargetStrategyPerformanceField, errors.MessageRejectError) {
	f := &field.TargetStrategyPerformanceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTargetStrategyPerformance reads a TargetStrategyPerformance from ExecutionReport.
func (m ExecutionReport) GetTargetStrategyPerformance(f *field.TargetStrategyPerformanceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for ExecutionReport.
func (m ExecutionReport) Currency() (*field.CurrencyField, errors.MessageRejectError) {
	f := &field.CurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from ExecutionReport.
func (m ExecutionReport) GetCurrency(f *field.CurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ComplianceID is a non-required field for ExecutionReport.
func (m ExecutionReport) ComplianceID() (*field.ComplianceIDField, errors.MessageRejectError) {
	f := &field.ComplianceIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetComplianceID reads a ComplianceID from ExecutionReport.
func (m ExecutionReport) GetComplianceID(f *field.ComplianceIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SolicitedFlag is a non-required field for ExecutionReport.
func (m ExecutionReport) SolicitedFlag() (*field.SolicitedFlagField, errors.MessageRejectError) {
	f := &field.SolicitedFlagField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSolicitedFlag reads a SolicitedFlag from ExecutionReport.
func (m ExecutionReport) GetSolicitedFlag(f *field.SolicitedFlagField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TimeInForce is a non-required field for ExecutionReport.
func (m ExecutionReport) TimeInForce() (*field.TimeInForceField, errors.MessageRejectError) {
	f := &field.TimeInForceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTimeInForce reads a TimeInForce from ExecutionReport.
func (m ExecutionReport) GetTimeInForce(f *field.TimeInForceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EffectiveTime is a non-required field for ExecutionReport.
func (m ExecutionReport) EffectiveTime() (*field.EffectiveTimeField, errors.MessageRejectError) {
	f := &field.EffectiveTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEffectiveTime reads a EffectiveTime from ExecutionReport.
func (m ExecutionReport) GetEffectiveTime(f *field.EffectiveTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExpireDate is a non-required field for ExecutionReport.
func (m ExecutionReport) ExpireDate() (*field.ExpireDateField, errors.MessageRejectError) {
	f := &field.ExpireDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExpireDate reads a ExpireDate from ExecutionReport.
func (m ExecutionReport) GetExpireDate(f *field.ExpireDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExpireTime is a non-required field for ExecutionReport.
func (m ExecutionReport) ExpireTime() (*field.ExpireTimeField, errors.MessageRejectError) {
	f := &field.ExpireTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExpireTime reads a ExpireTime from ExecutionReport.
func (m ExecutionReport) GetExpireTime(f *field.ExpireTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecInst is a non-required field for ExecutionReport.
func (m ExecutionReport) ExecInst() (*field.ExecInstField, errors.MessageRejectError) {
	f := &field.ExecInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExecInst reads a ExecInst from ExecutionReport.
func (m ExecutionReport) GetExecInst(f *field.ExecInstField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderCapacity is a non-required field for ExecutionReport.
func (m ExecutionReport) OrderCapacity() (*field.OrderCapacityField, errors.MessageRejectError) {
	f := &field.OrderCapacityField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderCapacity reads a OrderCapacity from ExecutionReport.
func (m ExecutionReport) GetOrderCapacity(f *field.OrderCapacityField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderRestrictions is a non-required field for ExecutionReport.
func (m ExecutionReport) OrderRestrictions() (*field.OrderRestrictionsField, errors.MessageRejectError) {
	f := &field.OrderRestrictionsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderRestrictions reads a OrderRestrictions from ExecutionReport.
func (m ExecutionReport) GetOrderRestrictions(f *field.OrderRestrictionsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CustOrderCapacity is a non-required field for ExecutionReport.
func (m ExecutionReport) CustOrderCapacity() (*field.CustOrderCapacityField, errors.MessageRejectError) {
	f := &field.CustOrderCapacityField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCustOrderCapacity reads a CustOrderCapacity from ExecutionReport.
func (m ExecutionReport) GetCustOrderCapacity(f *field.CustOrderCapacityField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastQty is a non-required field for ExecutionReport.
func (m ExecutionReport) LastQty() (*field.LastQtyField, errors.MessageRejectError) {
	f := &field.LastQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastQty reads a LastQty from ExecutionReport.
func (m ExecutionReport) GetLastQty(f *field.LastQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingLastQty is a non-required field for ExecutionReport.
func (m ExecutionReport) UnderlyingLastQty() (*field.UnderlyingLastQtyField, errors.MessageRejectError) {
	f := &field.UnderlyingLastQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingLastQty reads a UnderlyingLastQty from ExecutionReport.
func (m ExecutionReport) GetUnderlyingLastQty(f *field.UnderlyingLastQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastPx is a non-required field for ExecutionReport.
func (m ExecutionReport) LastPx() (*field.LastPxField, errors.MessageRejectError) {
	f := &field.LastPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastPx reads a LastPx from ExecutionReport.
func (m ExecutionReport) GetLastPx(f *field.LastPxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingLastPx is a non-required field for ExecutionReport.
func (m ExecutionReport) UnderlyingLastPx() (*field.UnderlyingLastPxField, errors.MessageRejectError) {
	f := &field.UnderlyingLastPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingLastPx reads a UnderlyingLastPx from ExecutionReport.
func (m ExecutionReport) GetUnderlyingLastPx(f *field.UnderlyingLastPxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastParPx is a non-required field for ExecutionReport.
func (m ExecutionReport) LastParPx() (*field.LastParPxField, errors.MessageRejectError) {
	f := &field.LastParPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastParPx reads a LastParPx from ExecutionReport.
func (m ExecutionReport) GetLastParPx(f *field.LastParPxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastSpotRate is a non-required field for ExecutionReport.
func (m ExecutionReport) LastSpotRate() (*field.LastSpotRateField, errors.MessageRejectError) {
	f := &field.LastSpotRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastSpotRate reads a LastSpotRate from ExecutionReport.
func (m ExecutionReport) GetLastSpotRate(f *field.LastSpotRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastForwardPoints is a non-required field for ExecutionReport.
func (m ExecutionReport) LastForwardPoints() (*field.LastForwardPointsField, errors.MessageRejectError) {
	f := &field.LastForwardPointsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastForwardPoints reads a LastForwardPoints from ExecutionReport.
func (m ExecutionReport) GetLastForwardPoints(f *field.LastForwardPointsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastMkt is a non-required field for ExecutionReport.
func (m ExecutionReport) LastMkt() (*field.LastMktField, errors.MessageRejectError) {
	f := &field.LastMktField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastMkt reads a LastMkt from ExecutionReport.
func (m ExecutionReport) GetLastMkt(f *field.LastMktField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionID is a non-required field for ExecutionReport.
func (m ExecutionReport) TradingSessionID() (*field.TradingSessionIDField, errors.MessageRejectError) {
	f := &field.TradingSessionIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionID reads a TradingSessionID from ExecutionReport.
func (m ExecutionReport) GetTradingSessionID(f *field.TradingSessionIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionSubID is a non-required field for ExecutionReport.
func (m ExecutionReport) TradingSessionSubID() (*field.TradingSessionSubIDField, errors.MessageRejectError) {
	f := &field.TradingSessionSubIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionSubID reads a TradingSessionSubID from ExecutionReport.
func (m ExecutionReport) GetTradingSessionSubID(f *field.TradingSessionSubIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TimeBracket is a non-required field for ExecutionReport.
func (m ExecutionReport) TimeBracket() (*field.TimeBracketField, errors.MessageRejectError) {
	f := &field.TimeBracketField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTimeBracket reads a TimeBracket from ExecutionReport.
func (m ExecutionReport) GetTimeBracket(f *field.TimeBracketField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastCapacity is a non-required field for ExecutionReport.
func (m ExecutionReport) LastCapacity() (*field.LastCapacityField, errors.MessageRejectError) {
	f := &field.LastCapacityField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastCapacity reads a LastCapacity from ExecutionReport.
func (m ExecutionReport) GetLastCapacity(f *field.LastCapacityField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LeavesQty is a required field for ExecutionReport.
func (m ExecutionReport) LeavesQty() (*field.LeavesQtyField, errors.MessageRejectError) {
	f := &field.LeavesQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLeavesQty reads a LeavesQty from ExecutionReport.
func (m ExecutionReport) GetLeavesQty(f *field.LeavesQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CumQty is a required field for ExecutionReport.
func (m ExecutionReport) CumQty() (*field.CumQtyField, errors.MessageRejectError) {
	f := &field.CumQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCumQty reads a CumQty from ExecutionReport.
func (m ExecutionReport) GetCumQty(f *field.CumQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AvgPx is a non-required field for ExecutionReport.
func (m ExecutionReport) AvgPx() (*field.AvgPxField, errors.MessageRejectError) {
	f := &field.AvgPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAvgPx reads a AvgPx from ExecutionReport.
func (m ExecutionReport) GetAvgPx(f *field.AvgPxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DayOrderQty is a non-required field for ExecutionReport.
func (m ExecutionReport) DayOrderQty() (*field.DayOrderQtyField, errors.MessageRejectError) {
	f := &field.DayOrderQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDayOrderQty reads a DayOrderQty from ExecutionReport.
func (m ExecutionReport) GetDayOrderQty(f *field.DayOrderQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DayCumQty is a non-required field for ExecutionReport.
func (m ExecutionReport) DayCumQty() (*field.DayCumQtyField, errors.MessageRejectError) {
	f := &field.DayCumQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDayCumQty reads a DayCumQty from ExecutionReport.
func (m ExecutionReport) GetDayCumQty(f *field.DayCumQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DayAvgPx is a non-required field for ExecutionReport.
func (m ExecutionReport) DayAvgPx() (*field.DayAvgPxField, errors.MessageRejectError) {
	f := &field.DayAvgPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDayAvgPx reads a DayAvgPx from ExecutionReport.
func (m ExecutionReport) GetDayAvgPx(f *field.DayAvgPxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//GTBookingInst is a non-required field for ExecutionReport.
func (m ExecutionReport) GTBookingInst() (*field.GTBookingInstField, errors.MessageRejectError) {
	f := &field.GTBookingInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetGTBookingInst reads a GTBookingInst from ExecutionReport.
func (m ExecutionReport) GetGTBookingInst(f *field.GTBookingInstField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeDate is a non-required field for ExecutionReport.
func (m ExecutionReport) TradeDate() (*field.TradeDateField, errors.MessageRejectError) {
	f := &field.TradeDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeDate reads a TradeDate from ExecutionReport.
func (m ExecutionReport) GetTradeDate(f *field.TradeDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a non-required field for ExecutionReport.
func (m ExecutionReport) TransactTime() (*field.TransactTimeField, errors.MessageRejectError) {
	f := &field.TransactTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from ExecutionReport.
func (m ExecutionReport) GetTransactTime(f *field.TransactTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ReportToExch is a non-required field for ExecutionReport.
func (m ExecutionReport) ReportToExch() (*field.ReportToExchField, errors.MessageRejectError) {
	f := &field.ReportToExchField{}
	err := m.Body.Get(f)
	return f, err
}

//GetReportToExch reads a ReportToExch from ExecutionReport.
func (m ExecutionReport) GetReportToExch(f *field.ReportToExchField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Commission is a non-required field for ExecutionReport.
func (m ExecutionReport) Commission() (*field.CommissionField, errors.MessageRejectError) {
	f := &field.CommissionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCommission reads a Commission from ExecutionReport.
func (m ExecutionReport) GetCommission(f *field.CommissionField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CommType is a non-required field for ExecutionReport.
func (m ExecutionReport) CommType() (*field.CommTypeField, errors.MessageRejectError) {
	f := &field.CommTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCommType reads a CommType from ExecutionReport.
func (m ExecutionReport) GetCommType(f *field.CommTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CommCurrency is a non-required field for ExecutionReport.
func (m ExecutionReport) CommCurrency() (*field.CommCurrencyField, errors.MessageRejectError) {
	f := &field.CommCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCommCurrency reads a CommCurrency from ExecutionReport.
func (m ExecutionReport) GetCommCurrency(f *field.CommCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FundRenewWaiv is a non-required field for ExecutionReport.
func (m ExecutionReport) FundRenewWaiv() (*field.FundRenewWaivField, errors.MessageRejectError) {
	f := &field.FundRenewWaivField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFundRenewWaiv reads a FundRenewWaiv from ExecutionReport.
func (m ExecutionReport) GetFundRenewWaiv(f *field.FundRenewWaivField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Spread is a non-required field for ExecutionReport.
func (m ExecutionReport) Spread() (*field.SpreadField, errors.MessageRejectError) {
	f := &field.SpreadField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSpread reads a Spread from ExecutionReport.
func (m ExecutionReport) GetSpread(f *field.SpreadField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurveCurrency is a non-required field for ExecutionReport.
func (m ExecutionReport) BenchmarkCurveCurrency() (*field.BenchmarkCurveCurrencyField, errors.MessageRejectError) {
	f := &field.BenchmarkCurveCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurveCurrency reads a BenchmarkCurveCurrency from ExecutionReport.
func (m ExecutionReport) GetBenchmarkCurveCurrency(f *field.BenchmarkCurveCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurveName is a non-required field for ExecutionReport.
func (m ExecutionReport) BenchmarkCurveName() (*field.BenchmarkCurveNameField, errors.MessageRejectError) {
	f := &field.BenchmarkCurveNameField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurveName reads a BenchmarkCurveName from ExecutionReport.
func (m ExecutionReport) GetBenchmarkCurveName(f *field.BenchmarkCurveNameField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurvePoint is a non-required field for ExecutionReport.
func (m ExecutionReport) BenchmarkCurvePoint() (*field.BenchmarkCurvePointField, errors.MessageRejectError) {
	f := &field.BenchmarkCurvePointField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurvePoint reads a BenchmarkCurvePoint from ExecutionReport.
func (m ExecutionReport) GetBenchmarkCurvePoint(f *field.BenchmarkCurvePointField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkPrice is a non-required field for ExecutionReport.
func (m ExecutionReport) BenchmarkPrice() (*field.BenchmarkPriceField, errors.MessageRejectError) {
	f := &field.BenchmarkPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkPrice reads a BenchmarkPrice from ExecutionReport.
func (m ExecutionReport) GetBenchmarkPrice(f *field.BenchmarkPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkPriceType is a non-required field for ExecutionReport.
func (m ExecutionReport) BenchmarkPriceType() (*field.BenchmarkPriceTypeField, errors.MessageRejectError) {
	f := &field.BenchmarkPriceTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkPriceType reads a BenchmarkPriceType from ExecutionReport.
func (m ExecutionReport) GetBenchmarkPriceType(f *field.BenchmarkPriceTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkSecurityID is a non-required field for ExecutionReport.
func (m ExecutionReport) BenchmarkSecurityID() (*field.BenchmarkSecurityIDField, errors.MessageRejectError) {
	f := &field.BenchmarkSecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkSecurityID reads a BenchmarkSecurityID from ExecutionReport.
func (m ExecutionReport) GetBenchmarkSecurityID(f *field.BenchmarkSecurityIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkSecurityIDSource is a non-required field for ExecutionReport.
func (m ExecutionReport) BenchmarkSecurityIDSource() (*field.BenchmarkSecurityIDSourceField, errors.MessageRejectError) {
	f := &field.BenchmarkSecurityIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkSecurityIDSource reads a BenchmarkSecurityIDSource from ExecutionReport.
func (m ExecutionReport) GetBenchmarkSecurityIDSource(f *field.BenchmarkSecurityIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldType is a non-required field for ExecutionReport.
func (m ExecutionReport) YieldType() (*field.YieldTypeField, errors.MessageRejectError) {
	f := &field.YieldTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYieldType reads a YieldType from ExecutionReport.
func (m ExecutionReport) GetYieldType(f *field.YieldTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Yield is a non-required field for ExecutionReport.
func (m ExecutionReport) Yield() (*field.YieldField, errors.MessageRejectError) {
	f := &field.YieldField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYield reads a Yield from ExecutionReport.
func (m ExecutionReport) GetYield(f *field.YieldField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldCalcDate is a non-required field for ExecutionReport.
func (m ExecutionReport) YieldCalcDate() (*field.YieldCalcDateField, errors.MessageRejectError) {
	f := &field.YieldCalcDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYieldCalcDate reads a YieldCalcDate from ExecutionReport.
func (m ExecutionReport) GetYieldCalcDate(f *field.YieldCalcDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldRedemptionDate is a non-required field for ExecutionReport.
func (m ExecutionReport) YieldRedemptionDate() (*field.YieldRedemptionDateField, errors.MessageRejectError) {
	f := &field.YieldRedemptionDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYieldRedemptionDate reads a YieldRedemptionDate from ExecutionReport.
func (m ExecutionReport) GetYieldRedemptionDate(f *field.YieldRedemptionDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldRedemptionPrice is a non-required field for ExecutionReport.
func (m ExecutionReport) YieldRedemptionPrice() (*field.YieldRedemptionPriceField, errors.MessageRejectError) {
	f := &field.YieldRedemptionPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYieldRedemptionPrice reads a YieldRedemptionPrice from ExecutionReport.
func (m ExecutionReport) GetYieldRedemptionPrice(f *field.YieldRedemptionPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldRedemptionPriceType is a non-required field for ExecutionReport.
func (m ExecutionReport) YieldRedemptionPriceType() (*field.YieldRedemptionPriceTypeField, errors.MessageRejectError) {
	f := &field.YieldRedemptionPriceTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYieldRedemptionPriceType reads a YieldRedemptionPriceType from ExecutionReport.
func (m ExecutionReport) GetYieldRedemptionPriceType(f *field.YieldRedemptionPriceTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//GrossTradeAmt is a non-required field for ExecutionReport.
func (m ExecutionReport) GrossTradeAmt() (*field.GrossTradeAmtField, errors.MessageRejectError) {
	f := &field.GrossTradeAmtField{}
	err := m.Body.Get(f)
	return f, err
}

//GetGrossTradeAmt reads a GrossTradeAmt from ExecutionReport.
func (m ExecutionReport) GetGrossTradeAmt(f *field.GrossTradeAmtField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NumDaysInterest is a non-required field for ExecutionReport.
func (m ExecutionReport) NumDaysInterest() (*field.NumDaysInterestField, errors.MessageRejectError) {
	f := &field.NumDaysInterestField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNumDaysInterest reads a NumDaysInterest from ExecutionReport.
func (m ExecutionReport) GetNumDaysInterest(f *field.NumDaysInterestField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExDate is a non-required field for ExecutionReport.
func (m ExecutionReport) ExDate() (*field.ExDateField, errors.MessageRejectError) {
	f := &field.ExDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExDate reads a ExDate from ExecutionReport.
func (m ExecutionReport) GetExDate(f *field.ExDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AccruedInterestRate is a non-required field for ExecutionReport.
func (m ExecutionReport) AccruedInterestRate() (*field.AccruedInterestRateField, errors.MessageRejectError) {
	f := &field.AccruedInterestRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccruedInterestRate reads a AccruedInterestRate from ExecutionReport.
func (m ExecutionReport) GetAccruedInterestRate(f *field.AccruedInterestRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AccruedInterestAmt is a non-required field for ExecutionReport.
func (m ExecutionReport) AccruedInterestAmt() (*field.AccruedInterestAmtField, errors.MessageRejectError) {
	f := &field.AccruedInterestAmtField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccruedInterestAmt reads a AccruedInterestAmt from ExecutionReport.
func (m ExecutionReport) GetAccruedInterestAmt(f *field.AccruedInterestAmtField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InterestAtMaturity is a non-required field for ExecutionReport.
func (m ExecutionReport) InterestAtMaturity() (*field.InterestAtMaturityField, errors.MessageRejectError) {
	f := &field.InterestAtMaturityField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInterestAtMaturity reads a InterestAtMaturity from ExecutionReport.
func (m ExecutionReport) GetInterestAtMaturity(f *field.InterestAtMaturityField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EndAccruedInterestAmt is a non-required field for ExecutionReport.
func (m ExecutionReport) EndAccruedInterestAmt() (*field.EndAccruedInterestAmtField, errors.MessageRejectError) {
	f := &field.EndAccruedInterestAmtField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEndAccruedInterestAmt reads a EndAccruedInterestAmt from ExecutionReport.
func (m ExecutionReport) GetEndAccruedInterestAmt(f *field.EndAccruedInterestAmtField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StartCash is a non-required field for ExecutionReport.
func (m ExecutionReport) StartCash() (*field.StartCashField, errors.MessageRejectError) {
	f := &field.StartCashField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStartCash reads a StartCash from ExecutionReport.
func (m ExecutionReport) GetStartCash(f *field.StartCashField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EndCash is a non-required field for ExecutionReport.
func (m ExecutionReport) EndCash() (*field.EndCashField, errors.MessageRejectError) {
	f := &field.EndCashField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEndCash reads a EndCash from ExecutionReport.
func (m ExecutionReport) GetEndCash(f *field.EndCashField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradedFlatSwitch is a non-required field for ExecutionReport.
func (m ExecutionReport) TradedFlatSwitch() (*field.TradedFlatSwitchField, errors.MessageRejectError) {
	f := &field.TradedFlatSwitchField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradedFlatSwitch reads a TradedFlatSwitch from ExecutionReport.
func (m ExecutionReport) GetTradedFlatSwitch(f *field.TradedFlatSwitchField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BasisFeatureDate is a non-required field for ExecutionReport.
func (m ExecutionReport) BasisFeatureDate() (*field.BasisFeatureDateField, errors.MessageRejectError) {
	f := &field.BasisFeatureDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBasisFeatureDate reads a BasisFeatureDate from ExecutionReport.
func (m ExecutionReport) GetBasisFeatureDate(f *field.BasisFeatureDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BasisFeaturePrice is a non-required field for ExecutionReport.
func (m ExecutionReport) BasisFeaturePrice() (*field.BasisFeaturePriceField, errors.MessageRejectError) {
	f := &field.BasisFeaturePriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBasisFeaturePrice reads a BasisFeaturePrice from ExecutionReport.
func (m ExecutionReport) GetBasisFeaturePrice(f *field.BasisFeaturePriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Concession is a non-required field for ExecutionReport.
func (m ExecutionReport) Concession() (*field.ConcessionField, errors.MessageRejectError) {
	f := &field.ConcessionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetConcession reads a Concession from ExecutionReport.
func (m ExecutionReport) GetConcession(f *field.ConcessionField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotalTakedown is a non-required field for ExecutionReport.
func (m ExecutionReport) TotalTakedown() (*field.TotalTakedownField, errors.MessageRejectError) {
	f := &field.TotalTakedownField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTotalTakedown reads a TotalTakedown from ExecutionReport.
func (m ExecutionReport) GetTotalTakedown(f *field.TotalTakedownField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NetMoney is a non-required field for ExecutionReport.
func (m ExecutionReport) NetMoney() (*field.NetMoneyField, errors.MessageRejectError) {
	f := &field.NetMoneyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNetMoney reads a NetMoney from ExecutionReport.
func (m ExecutionReport) GetNetMoney(f *field.NetMoneyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlCurrAmt is a non-required field for ExecutionReport.
func (m ExecutionReport) SettlCurrAmt() (*field.SettlCurrAmtField, errors.MessageRejectError) {
	f := &field.SettlCurrAmtField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlCurrAmt reads a SettlCurrAmt from ExecutionReport.
func (m ExecutionReport) GetSettlCurrAmt(f *field.SettlCurrAmtField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlCurrency is a non-required field for ExecutionReport.
func (m ExecutionReport) SettlCurrency() (*field.SettlCurrencyField, errors.MessageRejectError) {
	f := &field.SettlCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlCurrency reads a SettlCurrency from ExecutionReport.
func (m ExecutionReport) GetSettlCurrency(f *field.SettlCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlCurrFxRate is a non-required field for ExecutionReport.
func (m ExecutionReport) SettlCurrFxRate() (*field.SettlCurrFxRateField, errors.MessageRejectError) {
	f := &field.SettlCurrFxRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlCurrFxRate reads a SettlCurrFxRate from ExecutionReport.
func (m ExecutionReport) GetSettlCurrFxRate(f *field.SettlCurrFxRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlCurrFxRateCalc is a non-required field for ExecutionReport.
func (m ExecutionReport) SettlCurrFxRateCalc() (*field.SettlCurrFxRateCalcField, errors.MessageRejectError) {
	f := &field.SettlCurrFxRateCalcField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlCurrFxRateCalc reads a SettlCurrFxRateCalc from ExecutionReport.
func (m ExecutionReport) GetSettlCurrFxRateCalc(f *field.SettlCurrFxRateCalcField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//HandlInst is a non-required field for ExecutionReport.
func (m ExecutionReport) HandlInst() (*field.HandlInstField, errors.MessageRejectError) {
	f := &field.HandlInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetHandlInst reads a HandlInst from ExecutionReport.
func (m ExecutionReport) GetHandlInst(f *field.HandlInstField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinQty is a non-required field for ExecutionReport.
func (m ExecutionReport) MinQty() (*field.MinQtyField, errors.MessageRejectError) {
	f := &field.MinQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMinQty reads a MinQty from ExecutionReport.
func (m ExecutionReport) GetMinQty(f *field.MinQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaxFloor is a non-required field for ExecutionReport.
func (m ExecutionReport) MaxFloor() (*field.MaxFloorField, errors.MessageRejectError) {
	f := &field.MaxFloorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaxFloor reads a MaxFloor from ExecutionReport.
func (m ExecutionReport) GetMaxFloor(f *field.MaxFloorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PositionEffect is a non-required field for ExecutionReport.
func (m ExecutionReport) PositionEffect() (*field.PositionEffectField, errors.MessageRejectError) {
	f := &field.PositionEffectField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPositionEffect reads a PositionEffect from ExecutionReport.
func (m ExecutionReport) GetPositionEffect(f *field.PositionEffectField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaxShow is a non-required field for ExecutionReport.
func (m ExecutionReport) MaxShow() (*field.MaxShowField, errors.MessageRejectError) {
	f := &field.MaxShowField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaxShow reads a MaxShow from ExecutionReport.
func (m ExecutionReport) GetMaxShow(f *field.MaxShowField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BookingType is a non-required field for ExecutionReport.
func (m ExecutionReport) BookingType() (*field.BookingTypeField, errors.MessageRejectError) {
	f := &field.BookingTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBookingType reads a BookingType from ExecutionReport.
func (m ExecutionReport) GetBookingType(f *field.BookingTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for ExecutionReport.
func (m ExecutionReport) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from ExecutionReport.
func (m ExecutionReport) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for ExecutionReport.
func (m ExecutionReport) EncodedTextLen() (*field.EncodedTextLenField, errors.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from ExecutionReport.
func (m ExecutionReport) GetEncodedTextLen(f *field.EncodedTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for ExecutionReport.
func (m ExecutionReport) EncodedText() (*field.EncodedTextField, errors.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from ExecutionReport.
func (m ExecutionReport) GetEncodedText(f *field.EncodedTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlDate2 is a non-required field for ExecutionReport.
func (m ExecutionReport) SettlDate2() (*field.SettlDate2Field, errors.MessageRejectError) {
	f := &field.SettlDate2Field{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlDate2 reads a SettlDate2 from ExecutionReport.
func (m ExecutionReport) GetSettlDate2(f *field.SettlDate2Field) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderQty2 is a non-required field for ExecutionReport.
func (m ExecutionReport) OrderQty2() (*field.OrderQty2Field, errors.MessageRejectError) {
	f := &field.OrderQty2Field{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderQty2 reads a OrderQty2 from ExecutionReport.
func (m ExecutionReport) GetOrderQty2(f *field.OrderQty2Field) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastForwardPoints2 is a non-required field for ExecutionReport.
func (m ExecutionReport) LastForwardPoints2() (*field.LastForwardPoints2Field, errors.MessageRejectError) {
	f := &field.LastForwardPoints2Field{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastForwardPoints2 reads a LastForwardPoints2 from ExecutionReport.
func (m ExecutionReport) GetLastForwardPoints2(f *field.LastForwardPoints2Field) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MultiLegReportingType is a non-required field for ExecutionReport.
func (m ExecutionReport) MultiLegReportingType() (*field.MultiLegReportingTypeField, errors.MessageRejectError) {
	f := &field.MultiLegReportingTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMultiLegReportingType reads a MultiLegReportingType from ExecutionReport.
func (m ExecutionReport) GetMultiLegReportingType(f *field.MultiLegReportingTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CancellationRights is a non-required field for ExecutionReport.
func (m ExecutionReport) CancellationRights() (*field.CancellationRightsField, errors.MessageRejectError) {
	f := &field.CancellationRightsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCancellationRights reads a CancellationRights from ExecutionReport.
func (m ExecutionReport) GetCancellationRights(f *field.CancellationRightsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MoneyLaunderingStatus is a non-required field for ExecutionReport.
func (m ExecutionReport) MoneyLaunderingStatus() (*field.MoneyLaunderingStatusField, errors.MessageRejectError) {
	f := &field.MoneyLaunderingStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMoneyLaunderingStatus reads a MoneyLaunderingStatus from ExecutionReport.
func (m ExecutionReport) GetMoneyLaunderingStatus(f *field.MoneyLaunderingStatusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RegistID is a non-required field for ExecutionReport.
func (m ExecutionReport) RegistID() (*field.RegistIDField, errors.MessageRejectError) {
	f := &field.RegistIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRegistID reads a RegistID from ExecutionReport.
func (m ExecutionReport) GetRegistID(f *field.RegistIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Designation is a non-required field for ExecutionReport.
func (m ExecutionReport) Designation() (*field.DesignationField, errors.MessageRejectError) {
	f := &field.DesignationField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDesignation reads a Designation from ExecutionReport.
func (m ExecutionReport) GetDesignation(f *field.DesignationField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransBkdTime is a non-required field for ExecutionReport.
func (m ExecutionReport) TransBkdTime() (*field.TransBkdTimeField, errors.MessageRejectError) {
	f := &field.TransBkdTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransBkdTime reads a TransBkdTime from ExecutionReport.
func (m ExecutionReport) GetTransBkdTime(f *field.TransBkdTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecValuationPoint is a non-required field for ExecutionReport.
func (m ExecutionReport) ExecValuationPoint() (*field.ExecValuationPointField, errors.MessageRejectError) {
	f := &field.ExecValuationPointField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExecValuationPoint reads a ExecValuationPoint from ExecutionReport.
func (m ExecutionReport) GetExecValuationPoint(f *field.ExecValuationPointField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecPriceType is a non-required field for ExecutionReport.
func (m ExecutionReport) ExecPriceType() (*field.ExecPriceTypeField, errors.MessageRejectError) {
	f := &field.ExecPriceTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExecPriceType reads a ExecPriceType from ExecutionReport.
func (m ExecutionReport) GetExecPriceType(f *field.ExecPriceTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecPriceAdjustment is a non-required field for ExecutionReport.
func (m ExecutionReport) ExecPriceAdjustment() (*field.ExecPriceAdjustmentField, errors.MessageRejectError) {
	f := &field.ExecPriceAdjustmentField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExecPriceAdjustment reads a ExecPriceAdjustment from ExecutionReport.
func (m ExecutionReport) GetExecPriceAdjustment(f *field.ExecPriceAdjustmentField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriorityIndicator is a non-required field for ExecutionReport.
func (m ExecutionReport) PriorityIndicator() (*field.PriorityIndicatorField, errors.MessageRejectError) {
	f := &field.PriorityIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriorityIndicator reads a PriorityIndicator from ExecutionReport.
func (m ExecutionReport) GetPriorityIndicator(f *field.PriorityIndicatorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceImprovement is a non-required field for ExecutionReport.
func (m ExecutionReport) PriceImprovement() (*field.PriceImprovementField, errors.MessageRejectError) {
	f := &field.PriceImprovementField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceImprovement reads a PriceImprovement from ExecutionReport.
func (m ExecutionReport) GetPriceImprovement(f *field.PriceImprovementField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastLiquidityInd is a non-required field for ExecutionReport.
func (m ExecutionReport) LastLiquidityInd() (*field.LastLiquidityIndField, errors.MessageRejectError) {
	f := &field.LastLiquidityIndField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastLiquidityInd reads a LastLiquidityInd from ExecutionReport.
func (m ExecutionReport) GetLastLiquidityInd(f *field.LastLiquidityIndField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoContAmts is a non-required field for ExecutionReport.
func (m ExecutionReport) NoContAmts() (*field.NoContAmtsField, errors.MessageRejectError) {
	f := &field.NoContAmtsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoContAmts reads a NoContAmts from ExecutionReport.
func (m ExecutionReport) GetNoContAmts(f *field.NoContAmtsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoLegs is a non-required field for ExecutionReport.
func (m ExecutionReport) NoLegs() (*field.NoLegsField, errors.MessageRejectError) {
	f := &field.NoLegsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoLegs reads a NoLegs from ExecutionReport.
func (m ExecutionReport) GetNoLegs(f *field.NoLegsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CopyMsgIndicator is a non-required field for ExecutionReport.
func (m ExecutionReport) CopyMsgIndicator() (*field.CopyMsgIndicatorField, errors.MessageRejectError) {
	f := &field.CopyMsgIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCopyMsgIndicator reads a CopyMsgIndicator from ExecutionReport.
func (m ExecutionReport) GetCopyMsgIndicator(f *field.CopyMsgIndicatorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoMiscFees is a non-required field for ExecutionReport.
func (m ExecutionReport) NoMiscFees() (*field.NoMiscFeesField, errors.MessageRejectError) {
	f := &field.NoMiscFeesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoMiscFees reads a NoMiscFees from ExecutionReport.
func (m ExecutionReport) GetNoMiscFees(f *field.NoMiscFeesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoStrategyParameters is a non-required field for ExecutionReport.
func (m ExecutionReport) NoStrategyParameters() (*field.NoStrategyParametersField, errors.MessageRejectError) {
	f := &field.NoStrategyParametersField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoStrategyParameters reads a NoStrategyParameters from ExecutionReport.
func (m ExecutionReport) GetNoStrategyParameters(f *field.NoStrategyParametersField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//HostCrossID is a non-required field for ExecutionReport.
func (m ExecutionReport) HostCrossID() (*field.HostCrossIDField, errors.MessageRejectError) {
	f := &field.HostCrossIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetHostCrossID reads a HostCrossID from ExecutionReport.
func (m ExecutionReport) GetHostCrossID(f *field.HostCrossIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ManualOrderIndicator is a non-required field for ExecutionReport.
func (m ExecutionReport) ManualOrderIndicator() (*field.ManualOrderIndicatorField, errors.MessageRejectError) {
	f := &field.ManualOrderIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetManualOrderIndicator reads a ManualOrderIndicator from ExecutionReport.
func (m ExecutionReport) GetManualOrderIndicator(f *field.ManualOrderIndicatorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CustDirectedOrder is a non-required field for ExecutionReport.
func (m ExecutionReport) CustDirectedOrder() (*field.CustDirectedOrderField, errors.MessageRejectError) {
	f := &field.CustDirectedOrderField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCustDirectedOrder reads a CustDirectedOrder from ExecutionReport.
func (m ExecutionReport) GetCustDirectedOrder(f *field.CustDirectedOrderField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ReceivedDeptID is a non-required field for ExecutionReport.
func (m ExecutionReport) ReceivedDeptID() (*field.ReceivedDeptIDField, errors.MessageRejectError) {
	f := &field.ReceivedDeptIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetReceivedDeptID reads a ReceivedDeptID from ExecutionReport.
func (m ExecutionReport) GetReceivedDeptID(f *field.ReceivedDeptIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CustOrderHandlingInst is a non-required field for ExecutionReport.
func (m ExecutionReport) CustOrderHandlingInst() (*field.CustOrderHandlingInstField, errors.MessageRejectError) {
	f := &field.CustOrderHandlingInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCustOrderHandlingInst reads a CustOrderHandlingInst from ExecutionReport.
func (m ExecutionReport) GetCustOrderHandlingInst(f *field.CustOrderHandlingInstField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderHandlingInstSource is a non-required field for ExecutionReport.
func (m ExecutionReport) OrderHandlingInstSource() (*field.OrderHandlingInstSourceField, errors.MessageRejectError) {
	f := &field.OrderHandlingInstSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderHandlingInstSource reads a OrderHandlingInstSource from ExecutionReport.
func (m ExecutionReport) GetOrderHandlingInstSource(f *field.OrderHandlingInstSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoTrdRegTimestamps is a non-required field for ExecutionReport.
func (m ExecutionReport) NoTrdRegTimestamps() (*field.NoTrdRegTimestampsField, errors.MessageRejectError) {
	f := &field.NoTrdRegTimestampsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoTrdRegTimestamps reads a NoTrdRegTimestamps from ExecutionReport.
func (m ExecutionReport) GetNoTrdRegTimestamps(f *field.NoTrdRegTimestampsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AggressorIndicator is a non-required field for ExecutionReport.
func (m ExecutionReport) AggressorIndicator() (*field.AggressorIndicatorField, errors.MessageRejectError) {
	f := &field.AggressorIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAggressorIndicator reads a AggressorIndicator from ExecutionReport.
func (m ExecutionReport) GetAggressorIndicator(f *field.AggressorIndicatorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CalculatedCcyLastQty is a non-required field for ExecutionReport.
func (m ExecutionReport) CalculatedCcyLastQty() (*field.CalculatedCcyLastQtyField, errors.MessageRejectError) {
	f := &field.CalculatedCcyLastQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCalculatedCcyLastQty reads a CalculatedCcyLastQty from ExecutionReport.
func (m ExecutionReport) GetCalculatedCcyLastQty(f *field.CalculatedCcyLastQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastSwapPoints is a non-required field for ExecutionReport.
func (m ExecutionReport) LastSwapPoints() (*field.LastSwapPointsField, errors.MessageRejectError) {
	f := &field.LastSwapPointsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastSwapPoints reads a LastSwapPoints from ExecutionReport.
func (m ExecutionReport) GetLastSwapPoints(f *field.LastSwapPointsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MatchType is a non-required field for ExecutionReport.
func (m ExecutionReport) MatchType() (*field.MatchTypeField, errors.MessageRejectError) {
	f := &field.MatchTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMatchType reads a MatchType from ExecutionReport.
func (m ExecutionReport) GetMatchType(f *field.MatchTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderCategory is a non-required field for ExecutionReport.
func (m ExecutionReport) OrderCategory() (*field.OrderCategoryField, errors.MessageRejectError) {
	f := &field.OrderCategoryField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderCategory reads a OrderCategory from ExecutionReport.
func (m ExecutionReport) GetOrderCategory(f *field.OrderCategoryField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LotType is a non-required field for ExecutionReport.
func (m ExecutionReport) LotType() (*field.LotTypeField, errors.MessageRejectError) {
	f := &field.LotTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLotType reads a LotType from ExecutionReport.
func (m ExecutionReport) GetLotType(f *field.LotTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceProtectionScope is a non-required field for ExecutionReport.
func (m ExecutionReport) PriceProtectionScope() (*field.PriceProtectionScopeField, errors.MessageRejectError) {
	f := &field.PriceProtectionScopeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceProtectionScope reads a PriceProtectionScope from ExecutionReport.
func (m ExecutionReport) GetPriceProtectionScope(f *field.PriceProtectionScopeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerType is a non-required field for ExecutionReport.
func (m ExecutionReport) TriggerType() (*field.TriggerTypeField, errors.MessageRejectError) {
	f := &field.TriggerTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerType reads a TriggerType from ExecutionReport.
func (m ExecutionReport) GetTriggerType(f *field.TriggerTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerAction is a non-required field for ExecutionReport.
func (m ExecutionReport) TriggerAction() (*field.TriggerActionField, errors.MessageRejectError) {
	f := &field.TriggerActionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerAction reads a TriggerAction from ExecutionReport.
func (m ExecutionReport) GetTriggerAction(f *field.TriggerActionField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerPrice is a non-required field for ExecutionReport.
func (m ExecutionReport) TriggerPrice() (*field.TriggerPriceField, errors.MessageRejectError) {
	f := &field.TriggerPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerPrice reads a TriggerPrice from ExecutionReport.
func (m ExecutionReport) GetTriggerPrice(f *field.TriggerPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerSymbol is a non-required field for ExecutionReport.
func (m ExecutionReport) TriggerSymbol() (*field.TriggerSymbolField, errors.MessageRejectError) {
	f := &field.TriggerSymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerSymbol reads a TriggerSymbol from ExecutionReport.
func (m ExecutionReport) GetTriggerSymbol(f *field.TriggerSymbolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerSecurityID is a non-required field for ExecutionReport.
func (m ExecutionReport) TriggerSecurityID() (*field.TriggerSecurityIDField, errors.MessageRejectError) {
	f := &field.TriggerSecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerSecurityID reads a TriggerSecurityID from ExecutionReport.
func (m ExecutionReport) GetTriggerSecurityID(f *field.TriggerSecurityIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerSecurityIDSource is a non-required field for ExecutionReport.
func (m ExecutionReport) TriggerSecurityIDSource() (*field.TriggerSecurityIDSourceField, errors.MessageRejectError) {
	f := &field.TriggerSecurityIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerSecurityIDSource reads a TriggerSecurityIDSource from ExecutionReport.
func (m ExecutionReport) GetTriggerSecurityIDSource(f *field.TriggerSecurityIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerSecurityDesc is a non-required field for ExecutionReport.
func (m ExecutionReport) TriggerSecurityDesc() (*field.TriggerSecurityDescField, errors.MessageRejectError) {
	f := &field.TriggerSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerSecurityDesc reads a TriggerSecurityDesc from ExecutionReport.
func (m ExecutionReport) GetTriggerSecurityDesc(f *field.TriggerSecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerPriceType is a non-required field for ExecutionReport.
func (m ExecutionReport) TriggerPriceType() (*field.TriggerPriceTypeField, errors.MessageRejectError) {
	f := &field.TriggerPriceTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerPriceType reads a TriggerPriceType from ExecutionReport.
func (m ExecutionReport) GetTriggerPriceType(f *field.TriggerPriceTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerPriceTypeScope is a non-required field for ExecutionReport.
func (m ExecutionReport) TriggerPriceTypeScope() (*field.TriggerPriceTypeScopeField, errors.MessageRejectError) {
	f := &field.TriggerPriceTypeScopeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerPriceTypeScope reads a TriggerPriceTypeScope from ExecutionReport.
func (m ExecutionReport) GetTriggerPriceTypeScope(f *field.TriggerPriceTypeScopeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerPriceDirection is a non-required field for ExecutionReport.
func (m ExecutionReport) TriggerPriceDirection() (*field.TriggerPriceDirectionField, errors.MessageRejectError) {
	f := &field.TriggerPriceDirectionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerPriceDirection reads a TriggerPriceDirection from ExecutionReport.
func (m ExecutionReport) GetTriggerPriceDirection(f *field.TriggerPriceDirectionField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerNewPrice is a non-required field for ExecutionReport.
func (m ExecutionReport) TriggerNewPrice() (*field.TriggerNewPriceField, errors.MessageRejectError) {
	f := &field.TriggerNewPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerNewPrice reads a TriggerNewPrice from ExecutionReport.
func (m ExecutionReport) GetTriggerNewPrice(f *field.TriggerNewPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerOrderType is a non-required field for ExecutionReport.
func (m ExecutionReport) TriggerOrderType() (*field.TriggerOrderTypeField, errors.MessageRejectError) {
	f := &field.TriggerOrderTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerOrderType reads a TriggerOrderType from ExecutionReport.
func (m ExecutionReport) GetTriggerOrderType(f *field.TriggerOrderTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerNewQty is a non-required field for ExecutionReport.
func (m ExecutionReport) TriggerNewQty() (*field.TriggerNewQtyField, errors.MessageRejectError) {
	f := &field.TriggerNewQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerNewQty reads a TriggerNewQty from ExecutionReport.
func (m ExecutionReport) GetTriggerNewQty(f *field.TriggerNewQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerTradingSessionID is a non-required field for ExecutionReport.
func (m ExecutionReport) TriggerTradingSessionID() (*field.TriggerTradingSessionIDField, errors.MessageRejectError) {
	f := &field.TriggerTradingSessionIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerTradingSessionID reads a TriggerTradingSessionID from ExecutionReport.
func (m ExecutionReport) GetTriggerTradingSessionID(f *field.TriggerTradingSessionIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerTradingSessionSubID is a non-required field for ExecutionReport.
func (m ExecutionReport) TriggerTradingSessionSubID() (*field.TriggerTradingSessionSubIDField, errors.MessageRejectError) {
	f := &field.TriggerTradingSessionSubIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerTradingSessionSubID reads a TriggerTradingSessionSubID from ExecutionReport.
func (m ExecutionReport) GetTriggerTradingSessionSubID(f *field.TriggerTradingSessionSubIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PeggedRefPrice is a non-required field for ExecutionReport.
func (m ExecutionReport) PeggedRefPrice() (*field.PeggedRefPriceField, errors.MessageRejectError) {
	f := &field.PeggedRefPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPeggedRefPrice reads a PeggedRefPrice from ExecutionReport.
func (m ExecutionReport) GetPeggedRefPrice(f *field.PeggedRefPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PreTradeAnonymity is a non-required field for ExecutionReport.
func (m ExecutionReport) PreTradeAnonymity() (*field.PreTradeAnonymityField, errors.MessageRejectError) {
	f := &field.PreTradeAnonymityField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPreTradeAnonymity reads a PreTradeAnonymity from ExecutionReport.
func (m ExecutionReport) GetPreTradeAnonymity(f *field.PreTradeAnonymityField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MatchIncrement is a non-required field for ExecutionReport.
func (m ExecutionReport) MatchIncrement() (*field.MatchIncrementField, errors.MessageRejectError) {
	f := &field.MatchIncrementField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMatchIncrement reads a MatchIncrement from ExecutionReport.
func (m ExecutionReport) GetMatchIncrement(f *field.MatchIncrementField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaxPriceLevels is a non-required field for ExecutionReport.
func (m ExecutionReport) MaxPriceLevels() (*field.MaxPriceLevelsField, errors.MessageRejectError) {
	f := &field.MaxPriceLevelsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaxPriceLevels reads a MaxPriceLevels from ExecutionReport.
func (m ExecutionReport) GetMaxPriceLevels(f *field.MaxPriceLevelsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryDisplayQty is a non-required field for ExecutionReport.
func (m ExecutionReport) SecondaryDisplayQty() (*field.SecondaryDisplayQtyField, errors.MessageRejectError) {
	f := &field.SecondaryDisplayQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryDisplayQty reads a SecondaryDisplayQty from ExecutionReport.
func (m ExecutionReport) GetSecondaryDisplayQty(f *field.SecondaryDisplayQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DisplayWhen is a non-required field for ExecutionReport.
func (m ExecutionReport) DisplayWhen() (*field.DisplayWhenField, errors.MessageRejectError) {
	f := &field.DisplayWhenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDisplayWhen reads a DisplayWhen from ExecutionReport.
func (m ExecutionReport) GetDisplayWhen(f *field.DisplayWhenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DisplayMethod is a non-required field for ExecutionReport.
func (m ExecutionReport) DisplayMethod() (*field.DisplayMethodField, errors.MessageRejectError) {
	f := &field.DisplayMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDisplayMethod reads a DisplayMethod from ExecutionReport.
func (m ExecutionReport) GetDisplayMethod(f *field.DisplayMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DisplayLowQty is a non-required field for ExecutionReport.
func (m ExecutionReport) DisplayLowQty() (*field.DisplayLowQtyField, errors.MessageRejectError) {
	f := &field.DisplayLowQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDisplayLowQty reads a DisplayLowQty from ExecutionReport.
func (m ExecutionReport) GetDisplayLowQty(f *field.DisplayLowQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DisplayHighQty is a non-required field for ExecutionReport.
func (m ExecutionReport) DisplayHighQty() (*field.DisplayHighQtyField, errors.MessageRejectError) {
	f := &field.DisplayHighQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDisplayHighQty reads a DisplayHighQty from ExecutionReport.
func (m ExecutionReport) GetDisplayHighQty(f *field.DisplayHighQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DisplayMinIncr is a non-required field for ExecutionReport.
func (m ExecutionReport) DisplayMinIncr() (*field.DisplayMinIncrField, errors.MessageRejectError) {
	f := &field.DisplayMinIncrField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDisplayMinIncr reads a DisplayMinIncr from ExecutionReport.
func (m ExecutionReport) GetDisplayMinIncr(f *field.DisplayMinIncrField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RefreshQty is a non-required field for ExecutionReport.
func (m ExecutionReport) RefreshQty() (*field.RefreshQtyField, errors.MessageRejectError) {
	f := &field.RefreshQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRefreshQty reads a RefreshQty from ExecutionReport.
func (m ExecutionReport) GetRefreshQty(f *field.RefreshQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DisplayQty is a non-required field for ExecutionReport.
func (m ExecutionReport) DisplayQty() (*field.DisplayQtyField, errors.MessageRejectError) {
	f := &field.DisplayQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDisplayQty reads a DisplayQty from ExecutionReport.
func (m ExecutionReport) GetDisplayQty(f *field.DisplayQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Volatility is a non-required field for ExecutionReport.
func (m ExecutionReport) Volatility() (*field.VolatilityField, errors.MessageRejectError) {
	f := &field.VolatilityField{}
	err := m.Body.Get(f)
	return f, err
}

//GetVolatility reads a Volatility from ExecutionReport.
func (m ExecutionReport) GetVolatility(f *field.VolatilityField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TimeToExpiration is a non-required field for ExecutionReport.
func (m ExecutionReport) TimeToExpiration() (*field.TimeToExpirationField, errors.MessageRejectError) {
	f := &field.TimeToExpirationField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTimeToExpiration reads a TimeToExpiration from ExecutionReport.
func (m ExecutionReport) GetTimeToExpiration(f *field.TimeToExpirationField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RiskFreeRate is a non-required field for ExecutionReport.
func (m ExecutionReport) RiskFreeRate() (*field.RiskFreeRateField, errors.MessageRejectError) {
	f := &field.RiskFreeRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRiskFreeRate reads a RiskFreeRate from ExecutionReport.
func (m ExecutionReport) GetRiskFreeRate(f *field.RiskFreeRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceDelta is a non-required field for ExecutionReport.
func (m ExecutionReport) PriceDelta() (*field.PriceDeltaField, errors.MessageRejectError) {
	f := &field.PriceDeltaField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceDelta reads a PriceDelta from ExecutionReport.
func (m ExecutionReport) GetPriceDelta(f *field.PriceDeltaField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TrdMatchID is a non-required field for ExecutionReport.
func (m ExecutionReport) TrdMatchID() (*field.TrdMatchIDField, errors.MessageRejectError) {
	f := &field.TrdMatchIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTrdMatchID reads a TrdMatchID from ExecutionReport.
func (m ExecutionReport) GetTrdMatchID(f *field.TrdMatchIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocID is a non-required field for ExecutionReport.
func (m ExecutionReport) AllocID() (*field.AllocIDField, errors.MessageRejectError) {
	f := &field.AllocIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocID reads a AllocID from ExecutionReport.
func (m ExecutionReport) GetAllocID(f *field.AllocIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoAllocs is a non-required field for ExecutionReport.
func (m ExecutionReport) NoAllocs() (*field.NoAllocsField, errors.MessageRejectError) {
	f := &field.NoAllocsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoAllocs reads a NoAllocs from ExecutionReport.
func (m ExecutionReport) GetNoAllocs(f *field.NoAllocsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotNoFills is a non-required field for ExecutionReport.
func (m ExecutionReport) TotNoFills() (*field.TotNoFillsField, errors.MessageRejectError) {
	f := &field.TotNoFillsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTotNoFills reads a TotNoFills from ExecutionReport.
func (m ExecutionReport) GetTotNoFills(f *field.TotNoFillsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastFragment is a non-required field for ExecutionReport.
func (m ExecutionReport) LastFragment() (*field.LastFragmentField, errors.MessageRejectError) {
	f := &field.LastFragmentField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastFragment reads a LastFragment from ExecutionReport.
func (m ExecutionReport) GetLastFragment(f *field.LastFragmentField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoFills is a non-required field for ExecutionReport.
func (m ExecutionReport) NoFills() (*field.NoFillsField, errors.MessageRejectError) {
	f := &field.NoFillsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoFills reads a NoFills from ExecutionReport.
func (m ExecutionReport) GetNoFills(f *field.NoFillsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DividendYield is a non-required field for ExecutionReport.
func (m ExecutionReport) DividendYield() (*field.DividendYieldField, errors.MessageRejectError) {
	f := &field.DividendYieldField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDividendYield reads a DividendYield from ExecutionReport.
func (m ExecutionReport) GetDividendYield(f *field.DividendYieldField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplID is a non-required field for ExecutionReport.
func (m ExecutionReport) ApplID() (*field.ApplIDField, errors.MessageRejectError) {
	f := &field.ApplIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplID reads a ApplID from ExecutionReport.
func (m ExecutionReport) GetApplID(f *field.ApplIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplSeqNum is a non-required field for ExecutionReport.
func (m ExecutionReport) ApplSeqNum() (*field.ApplSeqNumField, errors.MessageRejectError) {
	f := &field.ApplSeqNumField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplSeqNum reads a ApplSeqNum from ExecutionReport.
func (m ExecutionReport) GetApplSeqNum(f *field.ApplSeqNumField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplLastSeqNum is a non-required field for ExecutionReport.
func (m ExecutionReport) ApplLastSeqNum() (*field.ApplLastSeqNumField, errors.MessageRejectError) {
	f := &field.ApplLastSeqNumField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplLastSeqNum reads a ApplLastSeqNum from ExecutionReport.
func (m ExecutionReport) GetApplLastSeqNum(f *field.ApplLastSeqNumField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplResendFlag is a non-required field for ExecutionReport.
func (m ExecutionReport) ApplResendFlag() (*field.ApplResendFlagField, errors.MessageRejectError) {
	f := &field.ApplResendFlagField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplResendFlag reads a ApplResendFlag from ExecutionReport.
func (m ExecutionReport) GetApplResendFlag(f *field.ApplResendFlagField) errors.MessageRejectError {
	return m.Body.Get(f)
}
