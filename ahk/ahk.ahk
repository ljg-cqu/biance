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

	FromCurrencyValueFieldX := 428
	FromCurrencyValueFieldY := 478
	FromCurrencyEntryX := 577
	FromCurrencyEntryY := 472
	FromCurrencyFieldX := 460
	FromCurrencyFieldY := 425
	FromCurrency1X := 462
	FromCurrency1Y := 486
	FromCurrency2X := 458
	FromCurrency2Y := 552
	FromCurrency3X := 467
	FromCurrency3Y := 617
	FromCurrency4X := 470
	FromCurrency4Y := 690
	FromCurrency5X := 457
	FromCurrency5Y := 760

	ToCurrencyValueFieldX := 400
	ToCurrencyValueFieldY := 678
	ToCurrencyEntryX := 586
	ToCurrencyEntryY := 668
	ToCurrencyFieldX := 468
	ToCurrencyFieldY := 425
	ToCurrency1X := 467
	ToCurrency1Y := 486
	ToCurrency2X := 467
	ToCurrency2Y := 549
	ToCurrency3X := 464
	ToCurrency3Y := 626
	ToCurrency4X := 454
	ToCurrency4Y := 690
	ToCurrency5X := 452
	ToCurrency5Y := 756

	PreviewConversionButtonX := 471
	PreviewConversionButtonY := 753

	PreviewConversionButtonXRisk := 479
	PreviewConversionButtonYRisk := 839

	ConvertButtonX := 464
	ConvertButtonY := 768

	ConvertButtonXVolatile := 474
	ConvertButtonYVolatile := 808

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