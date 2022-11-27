^n::
; run https://www.binance.com/en/convert?fromCoin=BUSD
{
; mouse move speed
    MouseMoveSpeed := 0
; common sleep time
    SleepTimeAfterRefreshPage := 3000
	SleepTimeAfterMouseMove := 100

	SleepTimeForAlternativeClick := 100
	SleepTimeAfterClickActionButton := 1500
	SleepTimeAfterClickInputField := 100
    SleepTimeAfterClickCurrencyEntry := 2500
    SleepTimeAfterClickCurrencyOption := 1500

    SleepTimeAfterSendTxt := 100

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

; common positions
	CurrencyGapPixelY := 66

	SelectCurrencyPanX := 618
	SelectCurrencyPanY := 798

	FromCurrencyValueFieldX := 617
	FromCurrencyValueFieldY := 597
	FromCurrencyEntryX := 863
	FromCurrencyEntryY := 600
	FromCurrencyFieldX := 684
	FromCurrencyFieldY := 668
	FromCurrency1X := 684
	FromCurrency1Y := 743
	FromCurrency2X := 714
	FromCurrency2Y := 831
	FromCurrency3X := 704
	FromCurrency3Y := 914
	FromCurrency4X := 692
	FromCurrency4Y := 994
	FromCurrency5X := 688
	FromCurrency5Y := 1078

	ToCurrencyValueFieldX := 605
	ToCurrencyValueFieldY := 844
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
	PreviewConversionButtonY := 994

	PreviewConversionButtonXRisk := 727
	PreviewConversionButtonYRisk := 1029

	ConvertButtonX := 728
	ConvertButtonY := 1095

	ConvertButtonXVolatile := 728
	ConvertButtonYVolatile := 1135

loop
    {
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

				; T, AST
				if (GainConvertFrom = "ATA") ; TODO: add more token check ; TODO: reuse, image s.
				{
				MouseMove FromCurrency2X, FromCurrency2Y, MouseMoveSpeed
				}
				if (GainConvertFrom = "PHA")
				{
				MouseMove FromCurrency2X, FromCurrency2Y, MouseMoveSpeed
				}
				if (GainConvertFrom = "WIN")
				{
				MouseMove FromCurrency2X, FromCurrency2Y, MouseMoveSpeed
				}
				if (GainConvertFrom = "OP")
				{
				MouseMove FromCurrency2X, FromCurrency2Y, MouseMoveSpeed
				}
				if (GainConvertFrom = "GLM")
				{
				MouseMove FromCurrency2X, FromCurrency2Y, MouseMoveSpeed
				}
				if (GainConvertFrom = "AUD") ; TODO: add more token check
				{
					MouseMove FromCurrency2X, FromCurrency2Y, MouseMoveSpeed
				}
				if (GainConvertFrom = "OG")
                {
                	MouseMove FromCurrency2X, FromCurrency2Y, MouseMoveSpeed
                }
				if (GainConvertFrom = "BAR")
				{
					MouseMove FromCurrency2X, FromCurrency2Y, MouseMoveSpeed
				}
				if (GainConvertFrom = "COS") ;
				{
					MouseMove FromCurrency2X, FromCurrency2Y, MouseMoveSpeed
				}
				if (GainConvertFrom = "ETH") ;
				{
					MouseMove FromCurrency2X, FromCurrency2Y, MouseMoveSpeed
				}
				if (GainConvertFrom = "AVA")
				{
					MouseMove FromCurrency3X, FromCurrency3Y, MouseMoveSpeed
				}
				if (GainConvertFrom = "OM")
				{
					MouseMove FromCurrency5X, FromCurrency5Y, MouseMoveSpeed
				}
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
				; OM, AST, T
				if (LossConvertTo = "ONT") ;
				{
					MouseMove ToCurrency2X, ToCurrency2Y, MouseMoveSpeed
				}
				if (LossConvertTo = "OG") ;
				{
					MouseMove ToCurrency2X, ToCurrency2Y, MouseMoveSpeed
				}
				if (LossConvertTo = "GAL") ;
				{
					MouseMove ToCurrency2X, ToCurrency2Y, MouseMoveSpeed
				}
				if (LossConvertTo = "PHA") ;
				{
					MouseMove ToCurrency2X, ToCurrency2Y, MouseMoveSpeed
				}
				if (LossConvertTo = "BAR") ;
				{
					MouseMove ToCurrency2X, ToCurrency2Y, MouseMoveSpeed
				}
				if (LossConvertTo = "COS") ;
				{
					MouseMove ToCurrency2X, ToCurrency2Y, MouseMoveSpeed
				}
				if (LossConvertTo = "OM") ;
				{
					MouseMove ToCurrency3X, ToCurrency3Y, MouseMoveSpeed
				}
				if (LossConvertTo = "AVA") ;
				{
					MouseMove ToCurrency2X, ToCurrency2Y, MouseMoveSpeed
				}
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