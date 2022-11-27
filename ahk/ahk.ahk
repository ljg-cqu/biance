#p::Pause  ; Win+P
^n::
; run https://www.binance.com/en/convert?fromCoin=BUSD
{
; token icon path
    TokenIconPath := "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook"
; mouse move speed
    MouseMoveSpeed := 0
    SleepTimeAfterRefreshPage := 3000
	SleepTimeAfterMouseMove := 50

	SleepTimeForAlternativeClick := 10
	SleepTimeAfterClickActionButton := 1000
	SleepTimeAfterClickInputField := 10
    SleepTimeAfterClickCurrencyEntry := 1500
    SleepTimeAfterClickCurrencyOption := 1000

    SleepTimeAfterSendTxt := 10

; login time control
	LoggedInTimeCount := 0
	MaxLoggedInTimeCount := 300 ;18000
	SleepTimeToMeasureLoggedInTime := 100
; Login positions
    MenueX := 907
    MenueY := 144

    LoginX := 774
    LoginY := 197

    ContinueWithGoogleX := 216
    ContinueWithGoogleY := 575

    ChooseAccountX := 891
    ChooseAccountY := 570

; xy_asus_vivobook common positions https://www.binance.com/en/convert?fromCoin=BUSD
	CurrencyGapPixelY := 66

	SelectCurrencyPanX := 951
	SelectCurrencyPanY := 1127

	FromCurrencyValueFieldX := 661
	FromCurrencyValueFieldY := 600
	FromCurrencyEntryX := 868
	FromCurrencyEntryY := 592
	FromCurrencyFieldX := 700
	FromCurrencyFieldY := 664
	FromCurrency1X := 711
	FromCurrency1Y := 740
	FromCurrency2X := 711
	FromCurrency2Y := 831
	FromCurrency3X := 704
	FromCurrency3Y := 920
	FromCurrency4X := 692
	FromCurrency4Y := 1005
	FromCurrency5X := 688
	FromCurrency5Y := 1081

	ToCurrencyValueFieldX := 605
	ToCurrencyValueFieldY := 846
	ToCurrencyEntryX := 864
	ToCurrencyEntryY := 838
	ToCurrencyFieldX := 687
	ToCurrencyFieldY := 668
	ToCurrency1X := 711
	ToCurrency1Y := 745
	ToCurrency2X := 699
	ToCurrency2Y := 828
	ToCurrency3X := 692
	ToCurrency3Y := 916
	ToCurrency4X := 712
	ToCurrency4Y := 997
	ToCurrency5X := 686
	ToCurrency5Y := 1080

	PreviewConversionButtonX := 710
	PreviewConversionButtonY := 940

	PreviewConversionButtonXRisk := 727
	PreviewConversionButtonYRisk := 1029

	ConvertButtonX := 728
	ConvertButtonY := 1095

	ConvertButtonXVolatile := 728
	ConvertButtonYVolatile := 1135

loop
    {
        LableProcessGainConvert:
		; read value data
			FileGainConvertFrom := "gainConvertFrom.txt"
			FileGainConvertTo := "gainConvertTo.txt"
			FileGainConvertValue := "gainConvertValue.txt"

			GainConvertFrom := ""
			GainConvertTo := ""
			GainConvertValue := ""
			if FileExist(FileGainConvertFrom) && FileExist(FileGainConvertTo) && FileExist(FileGainConvertValue) {
				GainConvertFrom := FileRead(FileGainConvertFrom)
				GainConvertTo := FileRead(FileGainConvertTo)
				GainConvertValue := FileRead(FileGainConvertValue)
			}

		if (GainConvertFrom != "" && GainConvertTo != "" && GainConvertValue != "")
		{
			; set from token
				MouseMove FromCurrencyEntryX, FromCurrencyEntryY, MouseMoveSpeed
				Sleep SleepTimeAfterMouseMove
				Click
				Sleep SleepTimeAfterClickCurrencyEntry
				MouseMove FromCurrencyFieldX, FromCurrencyFieldY, MouseMoveSpeed
				Sleep SleepTimeAfterMouseMove
				Click
				Sleep SleepTimeAfterClickInputField
				Send GainConvertFrom
				Sleep SleepTimeAfterSendTxt
				MouseMove FromCurrency1X, FromCurrency1Y, MouseMoveSpeed

                ; TODO: goto
                ; TODO: not foud abnormal case
				if (GainConvertFrom = "ACA")
				{
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\ACA.png")
                    {        MouseMove FoundX, FoundY, MouseMoveSpeed
                            ;Goto LabelAfterGainConvertFromTokenLocated
                    } ;} else
                            ;Goto LableProcessLossConvert
				}
				if (GainConvertFrom = "ANT")
				{
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\ANT.png")
                    {
                            MouseMove FoundX, FoundY, MouseMoveSpeed
                            ;Goto LabelAfterGainConvertFromTokenLocated
                    } ;} else
                             ;Goto LableProcessLossConvert
				}
				if (GainConvertFrom = "AR")
				{
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\AR.png")
                    {
                            MouseMove FoundX, FoundY, MouseMoveSpeed
                            ;Goto LabelAfterGainConvertFromTokenLocated
                    } ;} else
                             ;Goto LableProcessLossConvert
				}
				if (GainConvertFrom = "ATA")
				{
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\ATA.png")
                    {
                            MouseMove FoundX, FoundY, MouseMoveSpeed
                            ;Goto LabelAfterGainConvertFromTokenLocated
                    } ;} else
                             ;Goto LableProcessLossConvert
				}
				if (GainConvertFrom = "AUD")
				{
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\AUD.png")
                    {
                            MouseMove FoundX, FoundY, MouseMoveSpeed
                            ;Goto LabelAfterGainConvertFromTokenLocated
                    } ;} else
                             ;Goto LableProcessLossConvert
				}
				if (GainConvertFrom = "AVA")
				{
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\AVA.png")
                    {
                            MouseMove FoundX, FoundY, MouseMoveSpeed
                            ;Goto LabelAfterGainConvertFromTokenLocated
                    } ;} else
                             ;Goto LableProcessLossConvert
				}
				if (GainConvertFrom = "BAR")
				{
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\BAR.png")
                    {
                            MouseMove FoundX, FoundY, MouseMoveSpeed
                            ;Goto LabelAfterGainConvertFromTokenLocated
                    } ;} else
                             ;Goto LableProcessLossConvert
				}
				if (GainConvertFrom = "BNB")
				{
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\BNB.png")
                    {
                            MouseMove FoundX, FoundY, MouseMoveSpeed
                            ;Goto LabelAfterGainConvertFromTokenLocated
                    } ;} else
                             ;Goto LableProcessLossConvert
				}
				if (GainConvertFrom = "BTC")
				{
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\BTC.png")
                    {
                            MouseMove FoundX, FoundY, MouseMoveSpeed
                            ;Goto LabelAfterGainConvertFromTokenLocated
                    } ;} else
                             ;Goto LableProcessLossConvert
				}
				if (GainConvertFrom = "COS")
				{
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\COS.png")
                    {
                            MouseMove FoundX, FoundY, MouseMoveSpeed
                            ;Goto LabelAfterGainConvertFromTokenLocated
                    } ;} else
                             ;Goto LableProcessLossConvert
				}
				if (GainConvertFrom = "ETH")
				{
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\ETH.png")
                    {
                            MouseMove FoundX, FoundY, MouseMoveSpeed
                            ;Goto LabelAfterGainConvertFromTokenLocated
                    } ;} else
                             ;Goto LableProcessLossConvert
				}
				if (GainConvertFrom = "FOR")
				{
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\FOR.png")
                    {
                            MouseMove FoundX, FoundY, MouseMoveSpeed
                            ;Goto LabelAfterGainConvertFromTokenLocated
                    } ;} else
                             ;Goto LableProcessLossConvert
				}
				if (GainConvertFrom = "GAL")
				{
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\GAL.png")
                    {
                            MouseMove FoundX, FoundY, MouseMoveSpeed
                            ;Goto LabelAfterGainConvertFromTokenLocated
                    } ;} else
                             ;Goto LableProcessLossConvert
				}
				if (GainConvertFrom = "GBP")
				{
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\GBP.png")
                    {
                            MouseMove FoundX, FoundY, MouseMoveSpeed
                            ;Goto LabelAfterGainConvertFromTokenLocated
                    } ;} else
                             ;Goto LableProcessLossConvert
				}
				if (GainConvertFrom = "GLM")
				{
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\GLM.png")
                    {
                            MouseMove FoundX, FoundY, MouseMoveSpeed
                            ;Goto LabelAfterGainConvertFromTokenLocated
                    } ;} else
                             ;Goto LableProcessLossConvert
				}
				if (GainConvertFrom = "OG")
				{
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\OG.png")
                    {
                            MouseMove FoundX, FoundY, MouseMoveSpeed
                            ;Goto LabelAfterGainConvertFromTokenLocated
                    } ;} else
                             ;Goto LableProcessLossConvert
				}
				if (GainConvertFrom = "OM")
				{
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\OM.png")
                    {
                            MouseMove FoundX, FoundY, MouseMoveSpeed
                            ;Goto LabelAfterGainConvertFromTokenLocated
                    } ;} else
                             ;Goto LableProcessLossConvert
				}
				if (GainConvertFrom = "ONT")
				{
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\ONT.png")
                    {
                            MouseMove FoundX, FoundY, MouseMoveSpeed
                            ;Goto LabelAfterGainConvertFromTokenLocated
                    } ;} else
                             ;Goto LableProcessLossConvert
				}
				if (GainConvertFrom = "OP")
				{
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\OP.png")
                    {
                            MouseMove FoundX, FoundY, MouseMoveSpeed
                            ;Goto LabelAfterGainConvertFromTokenLocated
                    } ;} else
                             ;Goto LableProcessLossConvert
				}
				if (GainConvertFrom = "ORN")
				{
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\ORN.png")
                    {
                            MouseMove FoundX, FoundY, MouseMoveSpeed
                            ;Goto LabelAfterGainConvertFromTokenLocated
                    } ;} else
                             ;Goto LableProcessLossConvert
				}
				if (GainConvertFrom = "PHA")
				{
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\PHA.png")
                    {
                            MouseMove FoundX, FoundY, MouseMoveSpeed
                            ;Goto LabelAfterGainConvertFromTokenLocated
                    } ;} else
                             ;Goto LableProcessLossConvert
				}
				if (GainConvertFrom = "REP")
				{
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\REP.png")
                    {
                            MouseMove FoundX, FoundY, MouseMoveSpeed
                            ;Goto LabelAfterGainConvertFromTokenLocated
                    } ;} else
                             ;Goto LableProcessLossConvert
				}
				if (GainConvertFrom = "SC")
				{
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\SC.png")
                    {
                            MouseMove FoundX, FoundY, MouseMoveSpeed
                            ;Goto LabelAfterGainConvertFromTokenLocated
                    } ;} else
                             ;Goto LableProcessLossConvert
				}
				if (GainConvertFrom = "T")
				{
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\T.png")
                    {
                            MouseMove FoundX, FoundY, MouseMoveSpeed
                            ;Goto LabelAfterGainConvertFromTokenLocated
                    } ;} else
                             ;Goto LableProcessLossConvert
				}
				if (GainConvertFrom = "WIN")
				{
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\WIN.png")
                    {
                            MouseMove FoundX, FoundY, MouseMoveSpeed
                            ;Goto LabelAfterGainConvertFromTokenLocated
                    } ;} else
                             ;Goto LableProcessLossConvert
				}
				if (GainConvertFrom = "YFI")
				{
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\YFI.png")
                    {
                            MouseMove FoundX, FoundY, MouseMoveSpeed
                            ;Goto LabelAfterGainConvertFromTokenLocated
                    } ;} else
                             ;Goto LableProcessLossConvert
				}


                LabelAfterGainConvertFromTokenLocated:
				Sleep SleepTimeAfterMouseMove
				Click
				Sleep SleepTimeAfterClickCurrencyOption

			; set to token: BUSD, or USDT
				MouseMove ToCurrencyEntryX, ToCurrencyEntryY, MouseMoveSpeed
				Sleep SleepTimeAfterMouseMove
				Click
				Sleep SleepTimeAfterClickCurrencyEntry
				MouseMove ToCurrencyFieldX, ToCurrencyFieldY, MouseMoveSpeed
				Sleep SleepTimeAfterMouseMove
				Click
				Sleep SleepTimeAfterClickInputField
				Send GainConvertTo
				Sleep SleepTimeAfterSendTxt
				MouseMove ToCurrency1X, ToCurrency1Y, MouseMoveSpeed
				Sleep SleepTimeAfterMouseMove
				Click
				Sleep SleepTimeAfterClickCurrencyOption

			; set convert value
				MouseMove FromCurrencyValueFieldX, FromCurrencyValueFieldY, MouseMoveSpeed
				Sleep SleepTimeAfterMouseMove
				Click
				Sleep SleepTimeAfterClickInputField
				Send GainConvertValue
				Sleep SleepTimeAfterSendTxt

			; preview conversion
				MouseMove PreviewConversionButtonX, PreviewConversionButtonY, MouseMoveSpeed
				Sleep SleepTimeAfterMouseMove
				Click
				Sleep SleepTimeForAlternativeClick
				MouseMove PreviewConversionButtonXRisk, PreviewConversionButtonYRisk, MouseMoveSpeed
				Sleep SleepTimeAfterMouseMove
				Click
				Sleep SleepTimeAfterClickActionButton

			; convert
				MouseMove ConvertButtonX, ConvertButtonY, MouseMoveSpeed
				Sleep SleepTimeAfterMouseMove
				Click
				MouseMove ConvertButtonXVolatile, ConvertButtonYVolatile, MouseMoveSpeed
				Sleep SleepTimeAfterMouseMove
				Click
				Sleep SleepTimeAfterClickActionButton

			; refresh web page
				Send "{F5}" ; TODO: remove it after balance check by server side
				Sleep SleepTimeAfterRefreshPage ;todo: window wait
				Send "^{Tab}"
		}

        LableProcessLossConvert:
		FileLossConvertFrom := "lossConvertFrom.txt"
		FileLossConvertTo := "lossConvertTo.txt"
		FileLossConvertValue := "lossConvertValue.txt"

		LossConvertFrom := ""
		LossConvertTo := ""
		LossConvertValue := ""
		if FileExist(FileLossConvertFrom) && FileExist(FileLossConvertTo) && FileExist(FileLossConvertValue) {
			LossConvertFrom := FileRead(FileLossConvertFrom)
			LossConvertTo := FileRead(FileLossConvertTo)
			LossConvertValue := FileRead(FileLossConvertValue)
		}

		if (LossConvertFrom != "" && LossConvertTo != "" && LossConvertValue != "")
		{
			; set from tokenMOVR, limit to: BUSD, USDT
				MouseMove FromCurrencyEntryX, FromCurrencyEntryY, MouseMoveSpeed
				Sleep SleepTimeAfterMouseMove
				Click
				Sleep SleepTimeAfterClickCurrencyEntry
				MouseMove FromCurrencyFieldX, FromCurrencyFieldY, MouseMoveSpeed
				Sleep SleepTimeAfterMouseMove
				Click
				Sleep SleepTimeAfterClickInputField
				Send LossConvertFrom
				Sleep SleepTimeAfterSendTxt
				MouseMove FromCurrency1X, FromCurrency1Y, MouseMoveSpeed
				Sleep SleepTimeAfterMouseMove
				Click
				Sleep SleepTimeAfterClickCurrencyOption

			; set to token
				MouseMove ToCurrencyEntryX, ToCurrencyEntryY, MouseMoveSpeed
				Sleep SleepTimeAfterMouseMove
				Click
				Sleep SleepTimeAfterClickCurrencyEntry
				MouseMove ToCurrencyFieldX, ToCurrencyFieldY, MouseMoveSpeed
				Sleep SleepTimeAfterMouseMove
				Click
				Sleep SleepTimeAfterClickInputField
				Send LossConvertTo
				Sleep SleepTimeAfterSendTxt
				MouseMove ToCurrency1X, ToCurrency1Y, MouseMoveSpeed
				if (LossConvertTo = "ACA")
				{
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\ACA.png")
                    {        MouseMove FoundX, FoundY, MouseMoveSpeed
                            ;Goto LabelAfterLossConvertToTokenLocated
                    } ;} else
                            ;Goto LableProcessGainConvert
				}
				if (LossConvertTo = "ANT")
				{
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\ANT.png")
                    {
                            MouseMove FoundX, FoundY, MouseMoveSpeed
                            ;Goto LabelAfterLossConvertToTokenLocated
                    } ;} else
                             ;Goto LableProcessGainConvert
				}
				if (LossConvertTo = "AR")
				{
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\AR.png")
                    {
                            MouseMove FoundX, FoundY, MouseMoveSpeed
                            ;Goto LabelAfterLossConvertToTokenLocated
                    } ;} else
                             ;Goto LableProcessGainConvert
				}
				if (LossConvertTo = "ATA")
				{
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\ATA.png")
                    {
                            MouseMove FoundX, FoundY, MouseMoveSpeed
                            ;Goto LabelAfterLossConvertToTokenLocated
                    } ;} else
                             ;Goto LableProcessGainConvert
				}
				if (LossConvertTo = "AUD")
				{
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\AUD.png")
                    {
                            MouseMove FoundX, FoundY, MouseMoveSpeed
                            ;Goto LabelAfterLossConvertToTokenLocated
                    } ;} else
                             ;Goto LableProcessGainConvert
				}
				if (LossConvertTo = "AVA")
				{
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\AVA.png")
                    {
                            MouseMove FoundX, FoundY, MouseMoveSpeed
                            ;Goto LabelAfterLossConvertToTokenLocated
                    } ;} else
                             ;Goto LableProcessGainConvert
				}
				if (LossConvertTo = "BAR")
				{
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\BAR.png")
                    {
                            MouseMove FoundX, FoundY, MouseMoveSpeed
                            ;Goto LabelAfterLossConvertToTokenLocated
                    } ;} else
                             ;Goto LableProcessGainConvert
				}
				if (LossConvertTo = "BNB")
				{
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\BNB.png")
                    {
                            MouseMove FoundX, FoundY, MouseMoveSpeed
                            ;Goto LabelAfterLossConvertToTokenLocated
                    } ;} else
                             ;Goto LableProcessGainConvert
				}
				if (LossConvertTo = "BTC")
				{
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\BTC.png")
                    {
                            MouseMove FoundX, FoundY, MouseMoveSpeed
                            ;Goto LabelAfterLossConvertToTokenLocated
                    } ;} else
                             ;Goto LableProcessGainConvert
				}
				if (LossConvertTo = "COS")
				{
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\COS.png")
                    {
                            MouseMove FoundX, FoundY, MouseMoveSpeed
                            ;Goto LabelAfterLossConvertToTokenLocated
                    } ;} else
                             ;Goto LableProcessGainConvert
				}
				if (LossConvertTo = "ETH")
				{
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\ETH.png")
                    {
                            MouseMove FoundX, FoundY, MouseMoveSpeed
                            ;Goto LabelAfterLossConvertToTokenLocated
                    } ;} else
                             ;Goto LableProcessGainConvert
				}
				if (LossConvertTo = "FOR")
				{
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\FOR.png")
                    {
                            MouseMove FoundX, FoundY, MouseMoveSpeed
                            ;Goto LabelAfterLossConvertToTokenLocated
                    } ;} else
                             ;Goto LableProcessGainConvert
				}
				if (LossConvertTo = "GAL")
				{
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\GAL.png")
                    {
                            MouseMove FoundX, FoundY, MouseMoveSpeed
                            ;Goto LabelAfterLossConvertToTokenLocated
                    } ;} else
                             ;Goto LableProcessGainConvert
				}
				if (LossConvertTo = "GBP")
				{
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\GBP.png")
                    {
                            MouseMove FoundX, FoundY, MouseMoveSpeed
                            ;Goto LabelAfterLossConvertToTokenLocated
                    } ;} else
                             ;Goto LableProcessGainConvert
				}
				if (LossConvertTo = "GLM")
				{
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\GLM.png")
                    {
                            MouseMove FoundX, FoundY, MouseMoveSpeed
                            ;Goto LabelAfterLossConvertToTokenLocated
                    } ;} else
                             ;Goto LableProcessGainConvert
				}
				if (LossConvertTo = "OG")
				{
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\OG.png")
                    {
                            MouseMove FoundX, FoundY, MouseMoveSpeed
                            ;Goto LabelAfterLossConvertToTokenLocated
                    } ;} else
                             ;Goto LableProcessGainConvert
				}
				if (LossConvertTo = "OM")
				{
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\OM.png")
                    {
                            MouseMove FoundX, FoundY, MouseMoveSpeed
                            ;Goto LabelAfterLossConvertToTokenLocated
                    } ;} else
                             ;Goto LableProcessGainConvert
				}
				if (LossConvertTo = "ONT")
				{
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\ONT.png")
                    {
                            MouseMove FoundX, FoundY, MouseMoveSpeed
                            ;Goto LabelAfterLossConvertToTokenLocated
                    } ;} else
                             ;Goto LableProcessGainConvert
				}
				if (LossConvertTo = "OP")
				{
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\OP.png")
                    {
                            MouseMove FoundX, FoundY, MouseMoveSpeed
                            ;Goto LabelAfterLossConvertToTokenLocated
                    } ;} else
                             ;Goto LableProcessGainConvert
				}
				if (LossConvertTo = "ORN")
				{
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\ORN.png")
                    {
                            MouseMove FoundX, FoundY, MouseMoveSpeed
                            ;Goto LabelAfterLossConvertToTokenLocated
                    } ;} else
                             ;Goto LableProcessGainConvert
				}
				if (LossConvertTo = "PHA")
				{
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\PHA.png")
                    {
                            MouseMove FoundX, FoundY, MouseMoveSpeed
                            ;Goto LabelAfterLossConvertToTokenLocated
                    } ;} else
                             ;Goto LableProcessGainConvert
				}
				if (LossConvertTo = "REP")
				{
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\REP.png")
                    {
                            MouseMove FoundX, FoundY, MouseMoveSpeed
                            ;Goto LabelAfterLossConvertToTokenLocated
                    } ;} else
                             ;Goto LableProcessGainConvert
				}
				if (LossConvertTo = "SC")
				{
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\SC.png")
                    {
                            MouseMove FoundX, FoundY, MouseMoveSpeed
                            ;Goto LabelAfterLossConvertToTokenLocated
                    } ;} else
                             ;Goto LableProcessGainConvert
				}
				if (LossConvertTo = "T")
				{
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\T.png")
                    {
                            MouseMove FoundX, FoundY, MouseMoveSpeed
                            ;Goto LabelAfterLossConvertToTokenLocated
                    } ;} else
                             ;Goto LableProcessGainConvert
				}
				if (LossConvertTo = "WIN")
				{
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\WIN.png")
                    {
                            MouseMove FoundX, FoundY, MouseMoveSpeed
                            ;Goto LabelAfterLossConvertToTokenLocated
                    } ;} else
                             ;Goto LableProcessGainConvert
				}
				if (LossConvertTo = "YFI")
				{
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\YFI.png")
                    {
                            MouseMove FoundX, FoundY, MouseMoveSpeed
                            ;Goto LabelAfterLossConvertToTokenLocated
                    } ;} else
                             ;Goto LableProcessGainConvert
				}

				LabelAfterLossConvertToTokenLocated:
				Sleep SleepTimeAfterMouseMove
				Click
				Sleep SleepTimeAfterClickCurrencyEntry

			; set convert value
				MouseMove FromCurrencyValueFieldX, FromCurrencyValueFieldY, MouseMoveSpeed
				Sleep SleepTimeAfterMouseMove
				Click
				Sleep SleepTimeAfterClickInputField
				Send LossConvertValue
				Sleep SleepTimeAfterSendTxt

			; preview conversion
				MouseMove PreviewConversionButtonX, PreviewConversionButtonY, MouseMoveSpeed
				Sleep SleepTimeAfterMouseMove
				Click
				Sleep SleepTimeForAlternativeClick
				MouseMove PreviewConversionButtonXRisk, PreviewConversionButtonYRisk, MouseMoveSpeed
				Sleep SleepTimeAfterMouseMove
				Click
				Sleep SleepTimeAfterClickActionButton

			; convert
				MouseMove ConvertButtonX, ConvertButtonY, MouseMoveSpeed
				Sleep SleepTimeAfterMouseMove
				Click
				Sleep SleepTimeForAlternativeClick
				MouseMove ConvertButtonXVolatile, ConvertButtonYVolatile, MouseMoveSpeed
				Sleep SleepTimeAfterMouseMove
				Click
				Sleep SleepTimeAfterClickActionButton

			; refresh web page
				Send "{F5}" ; TODO: remove it after balance check by server side
				Sleep SleepTimeAfterRefreshPage
				Send "^{Tab}"
		}

;		LoggedInTimeCount += 1
;		Sleep SleepTimeToMeasureLoggedInTime
;
;		if (LoggedInTimeCount > MaxLoggedInTimeCount )
;		{
;			MouseMove MenueX, MenueY, MouseMoveSpeed
;            Sleep SleepTimeAfterMouseMove
;            Click
;            Sleep 5000
;            MouseMove LoginX, LoginY, MouseMoveSpeed
;            Sleep SleepTimeAfterMouseMove
;            Click
;            Sleep 5000
;            MouseMove ContinueWithGoogleX, ContinueWithGoogleY, MouseMoveSpeed
;            Sleep SleepTimeAfterMouseMove
;            Click
;            Sleep 5000
;            MouseMove ChooseAccountX, ChooseAccountY, MouseMoveSpeed
;            Sleep SleepTimeAfterMouseMove
;            Click
;            Sleep 5000
;			LoggedInTimeCount := 0
;		; TODO: extract function
;	    }
    }
}