package account

import (
	"net/http"
)

type RequestForLoginStatus struct {
}

func (req *RequestForLoginStatus) Path() string {
	return "/login_status"
}

func (req *RequestForLoginStatus) Method() string {
	return http.MethodGet
}

func (req *RequestForLoginStatus) Query() string {
	return ""
}

func (req *RequestForLoginStatus) Payload() []byte {
	return nil
}

type ResponseForLoginStatus struct {
	Account struct {
		BackstopProvider             bool        `json:"backstopProvider"`
		ChargeInterestOnNegativeUsd  bool        `json:"chargeInterestOnNegativeUsd"`
		Collateral                   float64     `json:"collateral"`
		FreeCollateral               float64     `json:"freeCollateral"`
		InitialMarginRequirement     float64     `json:"initialMarginRequirement"`
		Leverage                     float64     `json:"leverage"`
		Liquidating                  bool        `json:"liquidating"`
		MaintenanceMarginRequirement float64     `json:"maintenanceMarginRequirement"`
		MakerFee                     float64     `json:"makerFee"`
		MarginFraction               interface{} `json:"marginFraction"`
		OpenMarginFraction           interface{} `json:"openMarginFraction"`
		PositionLimit                interface{} `json:"positionLimit"`
		PositionLimitUsed            interface{} `json:"positionLimitUsed"`
		Positions                    []struct {
			CollateralUsed               float64     `json:"collateralUsed"`
			Cost                         float64     `json:"cost"`
			EntryPrice                   interface{} `json:"entryPrice"`
			EstimatedLiquidationPrice    interface{} `json:"estimatedLiquidationPrice"`
			Future                       string      `json:"future"`
			InitialMarginRequirement     float64     `json:"initialMarginRequirement"`
			LongOrderSize                float64     `json:"longOrderSize"`
			MaintenanceMarginRequirement float64     `json:"maintenanceMarginRequirement"`
			NetSize                      float64     `json:"netSize"`
			OpenSize                     float64     `json:"openSize"`
			RealizedPnl                  float64     `json:"realizedPnl"`
			ShortOrderSize               float64     `json:"shortOrderSize"`
			Side                         string      `json:"side"`
			Size                         float64     `json:"size"`
			UnrealizedPnl                float64     `json:"unrealizedPnl"`
		} `json:"positions"`
		SpotLendingEnabled bool    `json:"spotLendingEnabled"`
		SpotMarginEnabled  bool    `json:"spotMarginEnabled"`
		TakerFee           float64 `json:"takerFee"`
		TotalAccountValue  float64 `json:"totalAccountValue"`
		TotalPositionSize  float64 `json:"totalPositionSize"`
		UseFttCollateral   bool    `json:"useFttCollateral"`
		Username           string  `json:"username"`
	} `json:"account"`
	Country                  string      `json:"country"`
	InternalTransfersEnabled bool        `json:"internalTransfersEnabled"`
	LoggedIn                 bool        `json:"loggedIn"`
	Mfa                      string      `json:"mfa"`
	MfaRequired              interface{} `json:"mfaRequired"`
	ReadOnly                 bool        `json:"readOnly"`
	RestrictedToSubaccount   bool        `json:"restrictedToSubaccount"`
	Subaccount               interface{} `json:"subaccount"`
	User                     struct {
		CanOtcTradeOptions                    bool        `json:"canOtcTradeOptions"`
		CancelAllOrdersButtonEnabled          bool        `json:"cancelAllOrdersButtonEnabled"`
		ChatApp                               string      `json:"chatApp"`
		ChatHandle                            string      `json:"chatHandle"`
		ChatUserID                            interface{} `json:"chatUserId"`
		ConfirmTrades                         bool        `json:"confirmTrades"`
		CustomRefereeDiscountRate             float64     `json:"customRefereeDiscountRate"`
		CustomReferralRebateRate              float64     `json:"customReferralRebateRate"`
		DailyLeveragedTokenCreationVolume     float64     `json:"dailyLeveragedTokenCreationVolume"`
		DailyLeveragedTokenRedemptionVolume   float64     `json:"dailyLeveragedTokenRedemptionVolume"`
		DailyMakerVolume                      float64     `json:"dailyMakerVolume"`
		DailyVolume                           float64     `json:"dailyVolume"`
		DisplayName                           string      `json:"displayName"`
		Email                                 string      `json:"email"`
		FeeTier                               int         `json:"feeTier"`
		Ftt                                   float64     `json:"ftt"`
		InTradingCompetition                  bool        `json:"inTradingCompetition"`
		Kyc                                   string      `json:"kyc"`
		KycApplicationStatus                  string      `json:"kycApplicationStatus"`
		KycLevel                              int         `json:"kycLevel"`
		KycMessage                            interface{} `json:"kycMessage"`
		KycType                               string      `json:"kycType"`
		Mfa                                   string      `json:"mfa"`
		MmLevel                               int         `json:"mmLevel"`
		MonthlyLeveragedTokenCreationVolume   float64     `json:"monthlyLeveragedTokenCreationVolume"`
		MonthlyLeveragedTokenRedemptionVolume float64     `json:"monthlyLeveragedTokenRedemptionVolume"`
		MonthlyLtVolume                       float64     `json:"monthlyLtVolume"`
		MonthlyMakerVolume                    float64     `json:"monthlyMakerVolume"`
		MonthlyVolume                         float64     `json:"monthlyVolume"`
		OptionsEnabled                        bool        `json:"optionsEnabled"`
		PassedLtQuiz                          bool        `json:"passedLtQuiz"`
		RandomSlug                            string      `json:"randomSlug"`
		ReferralCode                          int         `json:"referralCode"`
		Referred                              bool        `json:"referred"`
		ReferrerID                            int         `json:"referrerId"`
		RequireMfaForWithdrawals              bool        `json:"requireMfaForWithdrawals"`
		RequireWhitelistedWithdrawals         bool        `json:"requireWhitelistedWithdrawals"`
		RequireWithdrawalPassword             bool        `json:"requireWithdrawalPassword"`
		ShowInLeaderboard                     bool        `json:"showInLeaderboard"`
		UseBodPriceChange                     bool        `json:"useBodPriceChange"`
		UseRealNameInLeaderboard              bool        `json:"useRealNameInLeaderboard"`
		Vip                                   int         `json:"vip"`
		WhitelistedAddressDelayDays           interface{} `json:"whitelistedAddressDelayDays"`
	} `json:"user"`
	WithdrawalEnabled bool `json:"withdrawalEnabled"`
}
